package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/taskcluster/mozilla-history/workerpool"
	tcclient "github.com/taskcluster/taskcluster/v47/clients/client-go"
	"github.com/taskcluster/taskcluster/v47/clients/client-go/tcauth"
	"github.com/taskcluster/taskcluster/v47/clients/client-go/tchooks"
	"github.com/taskcluster/taskcluster/v47/clients/client-go/tcsecrets"
	"github.com/taskcluster/taskcluster/v47/clients/client-go/tcworkermanager"
)

var (
	auth     *tcauth.Auth   = tcauth.NewFromEnv()
	hooks    *tchooks.Hooks = tchooks.NewFromEnv()
	fakeDate time.Time
)

func init() {
	var err error
	fakeDate, err = time.Parse("Jan 2, 2006 at 3:04pm -0700", "Aug 19, 1977 at 5:00pm +0100")
	if err != nil {
		panic(err)
	}
}

func FilenameEscape(raw string) (escaped string) {
	return strings.Replace(strings.Replace(raw, "*", "★", -1), "/", "⁄", -1)
}

func main() {
	wp := workerpool.New(100)
	hookGroupsPool := workerpool.New(10)
	hookGroupsPool.AddWork(FetchHookGroups)
	hookGroupsPool.Done()
	hookGroupsPool.OnComplete(func(result workerpool.Result) {
		wp.AddWork(result.(func(*workerpool.SubmitterContext)))
	})
	wp.AddWork(FetchRoles)
	wp.AddWork(FetchClients)
	wp.AddWork(FetchWorkerPools)
	// https://bugzil.la/1572135 - ** don't ** log hashed secrets
	// Note, secrets are json blobs rather than plain strings, so rainbow table
	// attacks[1] are likely to be ineffective, but better to be safe than
	// sorry.
	// ---
	// [1] https://blogs.getcertifiedgetahead.com/rainbow-table-attacks/
	// wp.AddWork(FetchSecrets)
	wp.Done()
	wp.OnComplete(func(result workerpool.Result) {
		log.Printf("%s", result)
	})
}

func EmptyDirectory(dir string) {
	err := os.RemoveAll(dir)
	if err != nil {
		panic(err)
	}
}

func WriteEntityToFileAsJSON(entity interface{}, path string, result workerpool.Result) workerpool.Work {
	return func(workerId int) workerpool.Result {
		var file *os.File
		err := os.MkdirAll(filepath.Dir(path), 0755)
		if err != nil {
			panic(err)
		}
		file, err = os.Create(path)
		if err != nil {
			panic(err)
		}
		defer func() {
			if err == nil {
				err = file.Close()
			} else {
				file.Close()
			}
		}()
		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "  ")
		encoder.Encode(entity)
		return result
	}
}

func FetchRoles(context *workerpool.SubmitterContext) {
	EmptyDirectory("Roles")
	continuationToken := ""
	for {
		roles, err := auth.ListRoles2(continuationToken, "")
		if err != nil {
			panic(err)
		}
		for _, role := range roles.Roles {
			context.RequestChannel <- WriteEntityToFileAsJSON(
				role,
				filepath.Join("Roles", FilenameEscape(role.RoleID)),
				"Fetched role "+role.RoleID,
			)
		}
		continuationToken = roles.ContinuationToken
		if continuationToken == "" {
			break
		}
	}
}

func FetchClients(context *workerpool.SubmitterContext) {
	EmptyDirectory("Clients")
	continuationToken := ""
	for {
		clients, err := auth.ListClients(continuationToken, "", "")
		if err != nil {
			panic(err)
		}
		for _, client := range clients.Clients {
			client.LastDateUsed = tcclient.Time(fakeDate)
			context.RequestChannel <- WriteEntityToFileAsJSON(
				client,
				filepath.Join("Clients", FilenameEscape(client.ClientID)),
				"Fetched client "+client.ClientID,
			)
		}
		continuationToken = clients.ContinuationToken
		if continuationToken == "" {
			break
		}
	}
}

func FetchHookGroups(context *workerpool.SubmitterContext) {
	EmptyDirectory("Hooks")
	allHookGroups, err := hooks.ListHookGroups()
	if err != nil {
		panic(err)
	}
	for _, hookGroup := range allHookGroups.Groups {
		context.RequestChannel <- func(hookGroup string) workerpool.Work {
			return func(workerId int) workerpool.Result {
				return func(context *workerpool.SubmitterContext) {
					hookList, err := hooks.ListHooks(hookGroup)
					if err != nil {
						panic(err)
					}
					for _, hook := range hookList.Hooks {
						context.RequestChannel <- WriteEntityToFileAsJSON(
							hook,
							filepath.Join("Hooks", hookGroup, FilenameEscape(hook.HookID)),
							"Fetched hook "+hookGroup+"/"+hook.HookID,
						)
					}
				}
			}
		}(hookGroup)
	}
}

// FetchSecrets is currently not called in production, but is useful if e.g.
// publishing to a private repo, or updating a repo just to say that a secret
// has changed, without publishing the hash of the secret. Therefore leaving
// the uncalled function in place so it can be adapted in the future.
//
// See https://bugzil.la/1572135 for more context.
func FetchSecrets(context *workerpool.SubmitterContext) {
	EmptyDirectory("Secrets")
	ss := tcsecrets.NewFromEnv()
	continuationToken := ""
	for {
		secretsList, err := ss.List(continuationToken, "")
		if err != nil {
			panic(err)
		}
		for _, secretName := range secretsList.Secrets {
			context.RequestChannel <- func(secretName string) workerpool.Work {
				return func(workerId int) workerpool.Result {
					secret, err := ss.Get(secretName)
					if err != nil {
						panic(err)
					}
					secret.Secret = json.RawMessage(fmt.Sprintf(`"%x"`, sha256.Sum256([]byte(secret.Secret))))
					return WriteEntityToFileAsJSON(
						secret,
						filepath.Join("Secrets", FilenameEscape(secretName)),
						"Fetched secret "+secretName,
					)(workerId)
				}
			}(secretName)
		}
		continuationToken = secretsList.ContinuationToken
		if continuationToken == "" {
			break
		}
	}
}

func FetchWorkerPools(context *workerpool.SubmitterContext) {
	EmptyDirectory("WorkerPools")
	workermanager := tcworkermanager.NewFromEnv()
	continuationToken := ""
	for {
		workerPools, err := workermanager.ListWorkerPools(continuationToken, "")
		if err != nil {
			panic(err)
		}
		for _, workerPool := range workerPools.WorkerPools {
  		// Strip out fields that change frequently
  		workerPool.RequestedCapacity = 0
  		workerPool.RequestedCount = 0
  		workerPool.RunningCapacity = 0
  		workerPool.RunningCount = 0
  		workerPool.StoppedCapacity = 0
  		workerPool.StoppedCount = 0
  		workerPool.StoppingCapacity = 0
  		workerPool.StoppingCount = 0

			context.RequestChannel <- WriteEntityToFileAsJSON(
				workerPool,
				filepath.Join("WorkerPools", FilenameEscape(workerPool.WorkerPoolID)),
				"Fetched worker pool "+workerPool.WorkerPoolID,
			)
		}
		continuationToken = workerPools.ContinuationToken
		if continuationToken == "" {
			break
		}
	}
}
