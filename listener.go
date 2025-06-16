package main

import (
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/coreruleset/seclang_parser/parser"
)

type ParserResult struct {
	variables             []string
	negatedVarCount       int
	collectionLengthCount int
	collections           []string
	collectionArgs        []string
	operatorList          []string
	operatorValueList     []string
	negatedOperatorCount  int
	directiveList         []string
	directiveValues       []string
	rangeEvents           []string
	rangeStartEvents      []int
	rangeEndEvents        []int
}

type TreeShapeListener struct {
	*parser.BaseSecLangParserListener
	results ParserResult
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

type CustomErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []error
}

func NewCustomErrorListener() *CustomErrorListener {
	return &CustomErrorListener{antlr.NewDefaultErrorListener(), make([]error, 0)}
}

func (c *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	var err error
	if offendingSymbol == nil {
		err = fmt.Errorf("Recognition error at line %d, column %d: %s", line, column, msg)
	} else {
		err = fmt.Errorf("Syntax error at line %d, column %d: %v", line, column, offendingSymbol)
	}
	c.Errors = append(c.Errors, err)
}

func (t *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	// if you need to debug, enable this one below
	// fmt.Println(ctx.GetText())
}

func (l *TreeShapeListener) EnterInt_range(ctx *parser.Int_rangeContext) {
	l.results.rangeEvents = append(l.results.rangeEvents, ctx.GetText())
}

func (l *TreeShapeListener) EnterRange_start(ctx *parser.Range_startContext) {
	i, err := strconv.Atoi(ctx.GetText())
	if err != nil {
		panic(err)
	}
	l.results.rangeStartEvents = append(l.results.rangeStartEvents, i)
}

func (l *TreeShapeListener) EnterRange_end(ctx *parser.Range_endContext) {
	i, err := strconv.Atoi(ctx.GetText())
	if err != nil {
		panic(err)
	}
	l.results.rangeEndEvents = append(l.results.rangeEndEvents, i)
}

func (l *TreeShapeListener) EnterRemove_rule_by_id(ctx *parser.Remove_rule_by_idContext) {
	l.results.directiveList = append(l.results.directiveList, ctx.GetText())
}

func (l *TreeShapeListener) EnterRemove_rule_by_id_int(ctx *parser.Remove_rule_by_id_intContext) {
	l.results.directiveValues = append(l.results.directiveValues, ctx.GetText())
}

func (l *TreeShapeListener) EnterRemove_rule_by_id_int_range(ctx *parser.Remove_rule_by_id_int_rangeContext) {
	l.results.directiveValues = append(l.results.directiveValues, ctx.GetText())
}

func (l *TreeShapeListener) EnterRemove_rule_by_msg(ctx *parser.Remove_rule_by_msgContext) {
	l.results.directiveList = append(l.results.directiveList, ctx.GetText())
}

func (l *TreeShapeListener) EnterRemove_rule_by_tag(ctx *parser.Remove_rule_by_tagContext) {
	l.results.directiveList = append(l.results.directiveList, ctx.GetText())
}

func (l *TreeShapeListener) EnterValues(ctx *parser.ValuesContext) {
	l.results.directiveValues = append(l.results.directiveValues, ctx.GetText())
}

func (l *TreeShapeListener) EnterUpdate_target_by_id(ctx *parser.Update_target_by_idContext) {
	l.results.directiveList = append(l.results.directiveList, ctx.GetText())
}

func (l *TreeShapeListener) EnterUpdate_target_by_msg(ctx *parser.Update_target_by_msgContext) {
	l.results.directiveList = append(l.results.directiveList, ctx.GetText())
}

func (l *TreeShapeListener) EnterUpdate_target_by_tag(ctx *parser.Update_target_by_tagContext) {
	l.results.directiveList = append(l.results.directiveList, ctx.GetText())
}

func (l *TreeShapeListener) EnterUpdate_target_rules_values(ctx *parser.Update_target_rules_valuesContext) {
	l.results.directiveValues = append(l.results.directiveValues, ctx.GetText())
}

func (l *TreeShapeListener) EnterVariable_enum(ctx *parser.Variable_enumContext) {
	l.results.variables = append(l.results.variables, ctx.GetText())
}

func (l *TreeShapeListener) EnterCollection_enum(ctx *parser.Collection_enumContext) {
	l.results.collections = append(l.results.collections, ctx.GetText())
}

func (l *TreeShapeListener) EnterCollection_value(ctx *parser.Collection_valueContext) {
	l.results.collectionArgs = append(l.results.collectionArgs, ctx.GetText())
}

func (l *TreeShapeListener) EnterVar_not(ctx *parser.Var_notContext) {
	l.results.negatedVarCount++
}

func (l *TreeShapeListener) EnterVar_count(ctx *parser.Var_countContext) {
	l.results.collectionLengthCount++
}

func (l *TreeShapeListener) EnterRules_directive(ctx *parser.Rules_directiveContext) {
	l.results.directiveList = append(l.results.directiveList, ctx.GetText())
}

func (l *TreeShapeListener) EnterOperator_name(ctx *parser.Operator_nameContext) {
	l.results.operatorList = append(l.results.operatorList, ctx.GetText())
}

func (l *TreeShapeListener) EnterOperator_value(ctx *parser.Operator_valueContext) {
	l.results.operatorValueList = append(l.results.operatorValueList, ctx.GetText())
}

func (l *TreeShapeListener) EnterOperator_not(ctx *parser.Operator_notContext) {
	l.results.negatedOperatorCount++
}
