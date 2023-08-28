// Copyright 2023 Felipe Zipitria
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"testing"

	"github.com/antlr4-go/antlr/v4"

	"github.com/fzipi/seclang_parser/parsing"
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

var testfiles = []string{
	"testdata/test1.conf",
	"testdata/test2.conf",
	"testdata/test3.conf",
	"testdata/test4.conf",
	"testdata/test5.conf",
	"testdata/crs-setup.conf",
	"testdata/REQUEST-901-INITIALIZATION.conf",
}

func TestSampleFiles(t *testing.T) {
	for _, file := range testfiles {
		input, err := antlr.NewFileStream(file)
		if err != nil {
			t.Errorf("Error reading file %s", file)
			continue
		}
		lexer := parsing.NewSecLangLexer(input)
		for {
			token := lexer.NextToken()
			if token.GetTokenType() == antlr.TokenEOF {
				break
			}
			fmt.Printf("%s (%q)\n",
				lexer.SymbolicNames[token.GetTokenType()], token.GetText())
		}
		stream := antlr.NewCommonTokenStream(lexer, 0)
		p := parsing.NewSecLangParser(stream)
		p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
		p.BuildParseTrees = true
		tree := p.Configuration()
		antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
	}
}
