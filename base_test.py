# Copyright 2025 Felipe Zipitria
# SPDX-License-Identifier: Apache-2.0

import sys
from antlr4 import *
from parsing.SecLangLexer import SecLangLexer
from parsing.SecLangParser import SecLangParser
from parsing.SecLangParserListener import SecLangParserListener

class Rule:
    def __init__(self):
        pass

class SecLangListener(SecLangParserListener):
    def enterStmt(self, ctx):
        print("Entering statement:", ctx.getText())
        Rule(variables, operators, actions)

    def enterAction(self, ctx:SecLangParser.ActionContext):
        print("Entering action:", ctx.getText())

    def enterConfig_value_types(self, ctx:SecLangParser.Config_value_typesContext):
        print("Entering config value:", ctx.getText())


def main(argv):
    if len(sys.argv) > 1:
        input = FileStream(sys.argv[1])
    else:
        input = InputStream(sys.stdin.readline())
    lexer = SecLangLexer(input)
    stream = CommonTokenStream(lexer)
    parser = SecLangParser(stream)
    tree = parser.configuration()
    # walk and print
    printer = SecLangListener()
    walker = ParseTreeWalker()
    walker.walk(printer, tree)


if __name__ == '__main__':
    main(sys.argv)
