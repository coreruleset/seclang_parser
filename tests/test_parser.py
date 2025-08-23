import sys

from antlr4 import *

from seclang_parser.SecLangLexer import SecLangLexer
from seclang_parser.SecLangParser import SecLangParser
from seclang_parser.SecLangParserListener import SecLangParserListener

class SecLangListener(SecLangParserListener):
    def enterStmt(self, ctx):
        print("Entering statement:", ctx.getText())

    def enterAction(self, ctx: SecLangParser.ActionContext):
        print("Entering action:", ctx.getText())

    def enterComment(self, ctx: SecLangParser.CommentContext):
        print("Entering comment:", ctx.getText())


def test_listener():
    assert True
