// Copyright 2023 Felipe Zipitria
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"os"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"

	"github.com/coreruleset/seclang_parser/parser"
)

// TestData represents the structure of the shared test YAML file
type TestData struct {
	CRSTestFiles      []string                       `yaml:"crs_test_files"`
	PluginsTestFiles  []string                       `yaml:"plugins_test_files"`
	GenericTests      map[string]GenericTestCase     `yaml:"generic_tests"`
	CheckOutputTests  map[string]CheckOutputTestCase `yaml:"check_output_tests"`
}

type GenericTestCase struct {
	ErrorCount int    `yaml:"error_count"`
	Comment    string `yaml:"comment"`
}

type CheckOutputTestCase struct {
	ErrorCount     int          `yaml:"error_count"`
	Comment        string       `yaml:"comment"`
	ExpectedResult ParserResult `yaml:"expected_result"`
}

var (
	crsTestFiles     []string
	pluginsTestFiles []string
	genericTests     map[string]GenericTestCase
	checkOutputTests map[string]CheckOutputTestCase
)

func init() {
	// Load test data from shared YAML file
	data, err := os.ReadFile("testdata/tests.yaml")
	if err != nil {
		panic("Failed to read test data file: " + err.Error())
	}

	var testData TestData
	if err := yaml.Unmarshal(data, &testData); err != nil {
		panic("Failed to parse test data file: " + err.Error())
	}

	crsTestFiles = testData.CRSTestFiles
	pluginsTestFiles = testData.PluginsTestFiles
	genericTests = testData.GenericTests
	checkOutputTests = testData.CheckOutputTests
}

func TestSecLang(t *testing.T) {
	for file, data := range genericTests {
		t.Logf("Testing file %s", file)
		input, err := antlr.NewFileStream(file)
		if err != nil {
			t.Errorf("Error reading file %s", file)
			continue
		}
		lexer := parser.NewSecLangLexer(input)

		lexerErrors := NewCustomErrorListener()
		lexer.RemoveErrorListeners()
		lexer.AddErrorListener(lexerErrors)

		parserErrors := NewCustomErrorListener()
		stream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewSecLangParser(stream)
		p.RemoveErrorListeners()
		p.AddErrorListener(parserErrors)

		p.BuildParseTrees = true
		tree := p.Configuration()

		antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)

		if len(lexerErrors.Errors) > 0 && data.ErrorCount != (len(lexerErrors.Errors)+len(parserErrors.Errors)) {
			t.Logf("Lexer %d errors found\n", len(lexerErrors.Errors))
			t.Logf("First error: %v\n", lexerErrors.Errors[0])
		}
		if len(parserErrors.Errors) > 0 && data.ErrorCount != (len(lexerErrors.Errors)+len(parserErrors.Errors)) {
			t.Logf("Parser %d errors found\n", len(parserErrors.Errors))
			t.Logf("First error: %v\n", parserErrors.Errors[0])
		}
		require.Equalf(t, data.ErrorCount, (len(lexerErrors.Errors) + len(parserErrors.Errors)), "Error count mismatch for file %s -> we want: %s\n", file, data.Comment)
	}
}

func TestSeclangOutput(t *testing.T) {
	for file, data := range checkOutputTests {
		t.Logf("Testing file %s", file)
		input, err := antlr.NewFileStream(file)
		if err != nil {
			t.Errorf("Error reading file %s", file)
			continue
		}
		lexer := parser.NewSecLangLexer(input)

		lexerErrors := NewCustomErrorListener()
		lexer.RemoveErrorListeners()
		lexer.AddErrorListener(lexerErrors)

		parserErrors := NewCustomErrorListener()
		stream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewSecLangParser(stream)
		p.RemoveErrorListeners()
		p.AddErrorListener(parserErrors)

		p.BuildParseTrees = true
		tree := p.Configuration()

		listener := NewTreeShapeListener()

		antlr.ParseTreeWalkerDefault.Walk(listener, tree)

		if len(lexerErrors.Errors) > 0 && data.ErrorCount != (len(lexerErrors.Errors)+len(parserErrors.Errors)) {
			t.Logf("Lexer %d errors found\n", len(lexerErrors.Errors))
			t.Logf("First error: %v\n", lexerErrors.Errors[0])
		}
		if len(parserErrors.Errors) > 0 && data.ErrorCount != (len(lexerErrors.Errors)+len(parserErrors.Errors)) {
			t.Logf("Parser %d errors found\n", len(parserErrors.Errors))
			t.Logf("First error: %v\n", parserErrors.Errors[0])
		}
		require.Equalf(t, data.ErrorCount, (len(lexerErrors.Errors) + len(parserErrors.Errors)), "Error count mismatch for file %s -> we want: %s\n", file, data.Comment)
		require.Equalf(t, data.ExpectedResult, listener.results, "Expected result mismatch for file %s -> we want: %v\n", file, data.ExpectedResult)
	}
}

func TestCRSLang(t *testing.T) {
	for _, file := range crsTestFiles {
		t.Logf("Testing file %s", file)
		input, err := antlr.NewFileStream(file)
		if err != nil {
			t.Fatalf("Error reading file %s", file)
		}
		lexer := parser.NewSecLangLexer(input)

		lexerErrors := NewCustomErrorListener()
		lexer.RemoveErrorListeners()
		lexer.AddErrorListener(lexerErrors)

		parserErrors := NewCustomErrorListener()
		stream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewSecLangParser(stream)
		p.RemoveErrorListeners()
		p.AddErrorListener(parserErrors)
		p.BuildParseTrees = true
		tree := p.Configuration()

		antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)

		if len(lexerErrors.Errors) > 0 {
			t.Logf("Lexer %d errors found\n", len(lexerErrors.Errors))
			t.Logf("First error: %v\n", lexerErrors.Errors[0])
		}
		if len(parserErrors.Errors) > 0 {
			t.Logf("Parser %d errors found\n", len(parserErrors.Errors))
			t.Logf("First error: %v\n", parserErrors.Errors[0])
		}
		require.Equalf(t, 0, (len(lexerErrors.Errors) + len(parserErrors.Errors)), "Error count mismatch for file %s -> we want no errors\n", file)
	}
}

func TestPlugins(t *testing.T) {
	for _, file := range pluginsTestFiles {
		t.Logf("Testing file %s", file)
		input, err := antlr.NewFileStream(file)
		if err != nil {
			t.Fatalf("Error reading file %s", file)
		}
		lexer := parser.NewSecLangLexer(input)

		lexerErrors := NewCustomErrorListener()
		lexer.RemoveErrorListeners()
		lexer.AddErrorListener(lexerErrors)

		parserErrors := NewCustomErrorListener()
		stream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewSecLangParser(stream)
		p.RemoveErrorListeners()
		p.AddErrorListener(parserErrors)
		p.BuildParseTrees = true
		tree := p.Configuration()

		antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)

		if len(lexerErrors.Errors) > 0 {
			t.Logf("Lexer %d errors found\n", len(lexerErrors.Errors))
			t.Logf("First error: %v\n", lexerErrors.Errors[0])
		}
		if len(parserErrors.Errors) > 0 {
			t.Logf("Parser %d errors found\n", len(parserErrors.Errors))
			t.Logf("First error: %v\n", parserErrors.Errors[0])
		}
		require.Equalf(t, 0, (len(lexerErrors.Errors) + len(parserErrors.Errors)), "Error count mismatch for file %s -> we want no errors\n", file)
	}
}
