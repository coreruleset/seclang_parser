#!/bin/sh
# Copyright 2025 Felipe Zipitria
# SPDX-License-Identifier: Apache-2.0

# This script is used to generate the parser files for the seclang DSL.
# It is used by the pre-commit hook to ensure that the parser files are up to date.

# Check if java is installed
if ! command -v npx >/dev/null 2>&1; then
    echo "Node is not installed. Please install nodejs and try again."
    exit 1
fi

# Find g4 files and change directory
g4_files=$(find . -name "SecLangLexer.g4")
g4_dir=$(dirname "$g4_files")

npx antlr-ng -Dlanguage=Go -p parser -o parser g4/*.g4
npx antlr-ng -Dlanguage=Python3 -o src/seclang_parser g4/*.g4
