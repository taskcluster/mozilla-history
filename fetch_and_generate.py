#!/usr/bin/env python3
"""Refresh all local Taskcluster snapshots and generated report data."""

from __future__ import annotations

import argparse
import json
import os
import re
import subprocess
import sys
import tempfile
import time
from pathlib import Path

REPOSITORY = Path(__file__).resolve().parent
TASK_GROUP_PATTERN = re.compile(r"^Task Group ID: (\S+)$")
WAITING_EXIT_CODE = 3
TOKEN_FILE = Path.home() / ".tc_token"


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description=(
            "Refresh local Taskcluster configuration snapshots, probe worker versions, "
            "wait for every probe to resolve, and generate report data."
        )
    )
    parser.add_argument(
        "--poll-interval",
        type=positive_number,
        default=60.0,
        metavar="SECONDS",
        help="seconds between task-group status checks (default: 60)",
    )
    parser.add_argument(
        "--task-group-id",
        help="resume polling and collecting an existing probe group instead of creating one",
    )
    return parser.parse_args()


def positive_number(raw: str) -> float:
    value = float(raw)
    if value <= 0:
        raise argparse.ArgumentTypeError("must be greater than zero")
    return value


def run(command: list[str], *, env: dict[str, str]) -> None:
    print(f"\n+ {' '.join(command)}", flush=True)
    subprocess.run(command, cwd=REPOSITORY, env=env, check=True)


def create_probe_group(auditor: str, *, env: dict[str, str]) -> str:
    print(f"\n+ {auditor}", flush=True)
    process = subprocess.Popen(
        [auditor],
        cwd=REPOSITORY,
        env=env,
        stdout=subprocess.PIPE,
        text=True,
        bufsize=1,
    )
    task_group_id = None
    assert process.stdout is not None
    for line in process.stdout:
        print(line, end="", flush=True)
        match = TASK_GROUP_PATTERN.match(line.rstrip("\n"))
        if match:
            task_group_id = match.group(1)

    return_code = process.wait()
    if return_code != 0:
        raise subprocess.CalledProcessError(return_code, [auditor])
    if task_group_id is None:
        raise RuntimeError("auditor did not report a task group ID")
    return task_group_id


def wait_for_probe_group(
    auditor: str,
    task_group_id: str,
    poll_interval: float,
    *,
    env: dict[str, str],
) -> None:
    while True:
        result = subprocess.run(
            [auditor, "status", task_group_id],
            cwd=REPOSITORY,
            env=env,
            check=False,
        )
        if result.returncode == 0:
            return
        if result.returncode != WAITING_EXIT_CODE:
            raise subprocess.CalledProcessError(
                result.returncode, [auditor, "status", task_group_id]
            )
        time.sleep(poll_interval)


def load_taskcluster_token(env: dict[str, str], token_file: Path = TOKEN_FILE) -> bool:
    credential_keys = {
        "TASKCLUSTER_CLIENT_ID": "clientId",
        "TASKCLUSTER_ACCESS_TOKEN": "accessToken",
    }
    if not token_file.exists():
        return False

    try:
        token = json.loads(token_file.read_text(encoding="utf-8"))
    except (OSError, UnicodeDecodeError, json.JSONDecodeError) as error:
        raise RuntimeError(
            f"could not read Taskcluster token file {token_file}: {error}"
        ) from error
    if not isinstance(token, dict):
        raise TypeError(
            f"Taskcluster token file {token_file} must contain a JSON object"
        )

    invalid = [
        token_name
        for token_name in credential_keys.values()
        if not isinstance(token.get(token_name), str) or not token[token_name]
    ]
    if invalid:
        names = ", ".join(invalid)
        raise RuntimeError(
            f"Taskcluster token file {token_file} has missing or invalid field(s): {names}"
        )

    for environment_name, token_name in credential_keys.items():
        env[environment_name] = token[token_name]
    return True


def taskcluster_environment() -> dict[str, str]:
    env = os.environ.copy()
    if load_taskcluster_token(env):
        print(f"Using Taskcluster credentials from {TOKEN_FILE}", flush=True)
    env.setdefault(
        "TASKCLUSTER_ROOT_URL", "https://firefox-ci-tc.services.mozilla.com/"
    )
    env.setdefault("REPORT_SCHEDULER_ID", "smoketest")
    env.setdefault(
        "REPORT_PREFIX",
        "https://github.com/taskcluster/mozilla-history/blob/master/WorkerVersions/",
    )
    return env


def require_probe_credentials(env: dict[str, str]) -> None:
    missing = [
        name
        for name in ("TASKCLUSTER_CLIENT_ID", "TASKCLUSTER_ACCESS_TOKEN")
        if not env.get(name)
    ]
    if missing:
        names = ", ".join(missing)
        raise RuntimeError(f"missing required environment variable(s): {names}")


def main() -> int:
    args = parse_args()
    env = taskcluster_environment()
    if args.task_group_id is None:
        require_probe_credentials(env)

    with tempfile.TemporaryDirectory(prefix="mozilla-history-") as temp_dir:
        snapshot_tool = str(Path(temp_dir) / "mozilla-history")
        auditor = str(Path(temp_dir) / "audit-worker-versions")

        run(
            ["go", "build", "-buildvcs=false", "-o", snapshot_tool, "."],
            env=env,
        )
        run(
            [
                "go",
                "build",
                "-buildvcs=false",
                "-o",
                auditor,
                "./audit-worker-versions",
            ],
            env=env,
        )

        run([snapshot_tool], env=env)

        task_group_id = args.task_group_id
        if task_group_id is None:
            task_group_id = create_probe_group(auditor, env=env)
        print(f"\nPolling Taskcluster task group {task_group_id}...", flush=True)
        wait_for_probe_group(auditor, task_group_id, args.poll_interval, env=env)
        run([auditor, task_group_id], env=env)

    print(
        "\nLocal refresh complete. Review the generated files before committing them.",
        flush=True,
    )
    return 0


if __name__ == "__main__":
    try:
        sys.exit(main())
    except KeyboardInterrupt:
        print("\nInterrupted.", file=sys.stderr)
        sys.exit(130)
    except (OSError, RuntimeError, TypeError, subprocess.CalledProcessError) as error:
        print(f"\nError: {error}", file=sys.stderr)
        sys.exit(1)
