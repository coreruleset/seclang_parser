#!/bin/sh
# Copyright 2023 Felipe Zipitria
# SPDX-License-Identifier: Apache-2.0

# This script is used to generate the parser files for the seclang DSL.
# It is used by the pre-commit hook to ensure that the parser files are up to date.

# Check if java is installed
if ! command -v java >/dev/null 2>&1; then
    echo "Java is not installed. Please install Java and try again."
    exit 1
fi

# Find g4 files and change directory
g4_files=$(find . -name "SecLangLexer.g4")
g4_dir=$(dirname "$g4_files")

# Change directory to g4
cd "$g4_dir"

alias antlr4='java -Xmx500M -cp "../lib/antlr-4.13.2-complete.jar:$CLASSPATH" org.antlr.v4.Tool'
antlr4 -Dlanguage=Go -no-visitor -package parser -o ../parser *.g4
antlr4 -Dlanguage=Python3 -no-visitor -package parser -o ../src/seclang_parser *.g4
