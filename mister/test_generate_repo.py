import hashlib
import tempfile
import unittest
from pathlib import Path

import generate_repo


class GenerateRepoTests(unittest.TestCase):
    def test_files_ref_and_updater_tag_are_independent(self) -> None:
        with tempfile.TemporaryDirectory() as temp:
            root = Path(temp)
            mapping = root / "input_v3.map"
            updater = root / "reflex_updater.sh"
            mapping.write_bytes(b"mapping")
            updater.write_bytes(b"updater")

            database = generate_repo.create_repo_db(
                [mapping],
                tag="v2.0",
                files_ref="main",
                updater_path=updater,
            )

        self.assertEqual(
            database["base_files_url"],
            "https://raw.githubusercontent.com/misteraddons/Reflex-Adapt-Legacy/main/mister/",
        )
        self.assertEqual(
            database["files"]["Scripts/reflex_updater.sh"]["url"],
            "https://github.com/misteraddons/Reflex-Adapt-Legacy/releases/download/v2.0/reflex_updater.sh",
        )
        self.assertEqual(
            database["files"]["config/inputs/input_v3.map"]["hash"],
            hashlib.md5(b"mapping", usedforsecurity=False).hexdigest(),
        )


if __name__ == "__main__":
    unittest.main()