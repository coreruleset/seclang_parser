package main

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/fzipi/seclang_parser/parsing"
	"os"
)

type TreeShapeListener struct {
	*parsing.BaseSecLangParserListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (t *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parsing.NewSecLangLexer(input)
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%s (%q)\n",
			lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parsing.NewSecLangParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Configuration()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}
