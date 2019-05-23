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
	tcclient "github.com/taskcluster/taskcluster-client-go"
	"github.com/taskcluster/taskcluster-client-go/tcauth"
	"github.com/taskcluster/taskcluster-client-go/tcawsprovisioner"
	"github.com/taskcluster/taskcluster-client-go/tchooks"
	"github.com/taskcluster/taskcluster-client-go/tcsecrets"
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
	wp.AddWork(FetchWorkerTypes)
	wp.AddWork(FetchRoles)
	wp.AddWork(FetchClients)
	wp.AddWork(FetchSecrets)
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

func FetchWorkerTypes(context *workerpool.SubmitterContext) {
	EmptyDirectory("AWSWorkerTypes")
	prov := tcawsprovisioner.NewFromEnv()
	allWorkerTypes, err := prov.ListWorkerTypes()
	if err != nil {
		panic(err)
	}
	for _, workerType := range *allWorkerTypes {
		context.RequestChannel <- func(workerType string) workerpool.Work {
			return func(workerId int) workerpool.Result {
				wt, err := prov.WorkerType(workerType)
				if err != nil {
					panic(err)
				}
				return WriteEntityToFileAsJSON(
					wt,
					filepath.Join("AWSWorkerTypes", FilenameEscape(workerType)),
					"Fetched workerType "+workerType,
				)(workerId)
			}
		}(workerType)
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
