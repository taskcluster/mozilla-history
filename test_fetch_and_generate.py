import json
import tempfile
import unittest
from pathlib import Path

from fetch_and_generate import load_taskcluster_token


class LoadTaskclusterTokenTests(unittest.TestCase):
    def test_loads_missing_credentials_from_json(self) -> None:
        with tempfile.TemporaryDirectory() as temp_dir:
            token_file = Path(temp_dir) / ".tc_token"
            token_file.write_text(
                json.dumps({"clientId": "file-client", "accessToken": "file-token"}),
                encoding="utf-8",
            )
            env: dict[str, str] = {}

            loaded = load_taskcluster_token(env, token_file)

        self.assertTrue(loaded)
        self.assertEqual(env["TASKCLUSTER_CLIENT_ID"], "file-client")
        self.assertEqual(env["TASKCLUSTER_ACCESS_TOKEN"], "file-token")

    def test_token_file_takes_precedence_over_environment(self) -> None:
        with tempfile.TemporaryDirectory() as temp_dir:
            token_file = Path(temp_dir) / ".tc_token"
            token_file.write_text(
                json.dumps({"clientId": "file-client", "accessToken": "file-token"}),
                encoding="utf-8",
            )
            env = {
                "TASKCLUSTER_CLIENT_ID": "environment-client",
                "TASKCLUSTER_ACCESS_TOKEN": "environment-token",
            }

            loaded = load_taskcluster_token(env, token_file)

        self.assertTrue(loaded)
        self.assertEqual(env["TASKCLUSTER_CLIENT_ID"], "file-client")
        self.assertEqual(env["TASKCLUSTER_ACCESS_TOKEN"], "file-token")

    def test_environment_is_unchanged_when_token_file_is_absent(self) -> None:
        env = {
            "TASKCLUSTER_CLIENT_ID": "environment-client",
            "TASKCLUSTER_ACCESS_TOKEN": "environment-token",
        }

        loaded = load_taskcluster_token(env, Path("/does/not/exist/.tc_token"))

        self.assertFalse(loaded)
        self.assertEqual(env["TASKCLUSTER_CLIENT_ID"], "environment-client")
        self.assertEqual(env["TASKCLUSTER_ACCESS_TOKEN"], "environment-token")

    def test_rejects_malformed_json_when_credentials_are_needed(self) -> None:
        with tempfile.TemporaryDirectory() as temp_dir:
            token_file = Path(temp_dir) / ".tc_token"
            token_file.write_text("not json", encoding="utf-8")

            with self.assertRaisesRegex(
                RuntimeError, "could not read Taskcluster token file"
            ):
                load_taskcluster_token({}, token_file)


if __name__ == "__main__":
    unittest.main()
