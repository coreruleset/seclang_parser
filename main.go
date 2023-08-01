package main

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/fzipi/seclang_parser/parsing"
	"os"
)

type TreeShapeListener struct {
	*parsing.BaseSecLangListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parsing.NewSecLangLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parsing.NewSecLangParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Config()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}
