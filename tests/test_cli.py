# Copyright 2025 OWASP CRS Project
# SPDX-License-Identifier: Apache-2.0

import sys

from seclang_parser.cli import run


def test_cli(monkeypatch):
    monkeypatch.setattr(
        sys,
        "argv",
        [
            "seclang_parser",
            "-f",
            "testdata/test_01_comment.conf",
        ],
    )

    assert run() is True
