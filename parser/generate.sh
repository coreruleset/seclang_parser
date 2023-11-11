#!/bin/sh
# Copyright 2023 Felipe Zipitria
# SPDX-License-Identifier: Apache-2.0


alias antlr4='java -Xmx500M -cp "../lib/antlr-4.13.0-complete.jar:$CLASSPATH" org.antlr.v4.Tool'
antlr4 -Dlanguage=Go -no-visitor -package parsing -o ../parsing *.g4
antlr4 -Dlanguage=Python3 -no-visitor -package parsing -o ../parsing *.g4
