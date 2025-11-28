// Copyright 2025 OWASP CRS Project
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"

	"github.com/coreruleset/seclang_parser/parser"
)

type ParserResult struct {
	Comments              []string `yaml:"comments"`
	CommentsBlocks        int      `yaml:"comment_blocks"`
	Variables             []string `yaml:"variables"`
	NegatedVarCount       int      `yaml:"negated_var_count"`
	CollectionLengthCount int      `yaml:"collection_length_count"`
	Collections           []string `yaml:"collections"`
	CollectionArgs        []string `yaml:"collection_args"`
	OperatorList          []string `yaml:"operator_list"`
	OperatorValueList     []string `yaml:"operator_value_list"`
	NegatedOperatorCount  int      `yaml:"negated_operator_count"`
	DirectiveList         []string `yaml:"directive_list"`
	DirectiveValues       []string `yaml:"directive_values"`
	RangeEvents           []string `yaml:"range_events"`
	RangeStartEvents      []int    `yaml:"range_start_events"`
	RangeEndEvents        []int    `yaml:"range_end_events"`
	SetvarCollections     []string `yaml:"setvar_collections"`
	SetvarNames           []string `yaml:"setvar_names"`
	SetvarOperations      []string `yaml:"setvar_operations"`
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
	l.results.RangeEvents = append(l.results.RangeEvents, ctx.GetText())
}

func (l *TreeShapeListener) EnterRange_start(ctx *parser.Range_startContext) {
	i, err := strconv.Atoi(ctx.GetText())
	if err != nil {
		panic(err)
	}
	l.results.RangeStartEvents = append(l.results.RangeStartEvents, i)
}

func (l *TreeShapeListener) EnterRange_end(ctx *parser.Range_endContext) {
	i, err := strconv.Atoi(ctx.GetText())
	if err != nil {
		panic(err)
	}
	l.results.RangeEndEvents = append(l.results.RangeEndEvents, i)
}

func (l *TreeShapeListener) EnterOperator_int_range(ctx *parser.Operator_int_rangeContext) {
	l.results.RangeEvents = append(l.results.RangeEvents, ctx.GetText())
	i, err := strconv.Atoi(ctx.INT_RANGE_VALUE(0).GetText())
	if err != nil {
		panic(err)
	}
	l.results.RangeStartEvents = append(l.results.RangeStartEvents, i)
	i, err = strconv.Atoi(ctx.INT_RANGE_VALUE(1).GetText())
	if err != nil {
		panic(err)
	}
	l.results.RangeEndEvents = append(l.results.RangeEndEvents, i)
}

func (l *TreeShapeListener) EnterRemove_rule_by_id(ctx *parser.Remove_rule_by_idContext) {
	l.results.DirectiveList = append(l.results.DirectiveList, ctx.GetText())
}

func (l *TreeShapeListener) EnterRemove_rule_by_id_int(ctx *parser.Remove_rule_by_id_intContext) {
	l.results.DirectiveValues = append(l.results.DirectiveValues, ctx.GetText())
}

func (l *TreeShapeListener) EnterRemove_rule_by_id_int_range(ctx *parser.Remove_rule_by_id_int_rangeContext) {
	l.results.DirectiveValues = append(l.results.DirectiveValues, ctx.GetText())
}

func (l *TreeShapeListener) EnterRemove_rule_by_msg(ctx *parser.Remove_rule_by_msgContext) {
	l.results.DirectiveList = append(l.results.DirectiveList, ctx.GetText())
}

func (l *TreeShapeListener) EnterRemove_rule_by_tag(ctx *parser.Remove_rule_by_tagContext) {
	l.results.DirectiveList = append(l.results.DirectiveList, ctx.GetText())
}

func (l *TreeShapeListener) EnterString_remove_rules_values(ctx *parser.String_remove_rules_valuesContext) {
	l.results.DirectiveValues = append(l.results.DirectiveValues, ctx.GetText())
}

func (l *TreeShapeListener) EnterUpdate_target_by_id(ctx *parser.Update_target_by_idContext) {
	l.results.DirectiveList = append(l.results.DirectiveList, ctx.GetText())
}

func (l *TreeShapeListener) EnterUpdate_target_by_msg(ctx *parser.Update_target_by_msgContext) {
	l.results.DirectiveList = append(l.results.DirectiveList, ctx.GetText())
}

func (l *TreeShapeListener) EnterUpdate_target_by_tag(ctx *parser.Update_target_by_tagContext) {
	l.results.DirectiveList = append(l.results.DirectiveList, ctx.GetText())
}

func (l *TreeShapeListener) EnterUpdate_target_rules_values(ctx *parser.Update_target_rules_valuesContext) {
	l.results.DirectiveValues = append(l.results.DirectiveValues, ctx.GetText())
}

func (l *TreeShapeListener) EnterVariable_enum(ctx *parser.Variable_enumContext) {
	l.results.Variables = append(l.results.Variables, ctx.GetText())
}

func (l *TreeShapeListener) EnterCollection_enum(ctx *parser.Collection_enumContext) {
	l.results.Collections = append(l.results.Collections, ctx.GetText())
}

func (l *TreeShapeListener) EnterCollection_value(ctx *parser.Collection_valueContext) {
	l.results.CollectionArgs = append(l.results.CollectionArgs, ctx.GetText())
}

func (l *TreeShapeListener) EnterVar_not(ctx *parser.Var_notContext) {
	l.results.NegatedVarCount++
}

func (l *TreeShapeListener) EnterVar_count(ctx *parser.Var_countContext) {
	l.results.CollectionLengthCount++
}

func (l *TreeShapeListener) EnterEngine_config_rule_directive(ctx *parser.Engine_config_rule_directiveContext) {
	l.results.DirectiveList = append(l.results.DirectiveList, ctx.GetText())
}

func (l *TreeShapeListener) EnterConfig_dir_sec_action(ctx *parser.Config_dir_sec_actionContext) {
	l.results.DirectiveList = append(l.results.DirectiveList, ctx.GetText())
}

func (l *TreeShapeListener) EnterOperator_name(ctx *parser.Operator_nameContext) {
	l.results.OperatorList = append(l.results.OperatorList, ctx.GetText())
}

func (l *TreeShapeListener) EnterOperator_value(ctx *parser.Operator_valueContext) {
	l.results.OperatorValueList = append(l.results.OperatorValueList, ctx.GetText())
}

func (l *TreeShapeListener) EnterOperator_not(ctx *parser.Operator_notContext) {
	l.results.NegatedOperatorCount++
}

func (l *TreeShapeListener) EnterCol_name(ctx *parser.Col_nameContext) {
	l.results.SetvarCollections = append(l.results.SetvarCollections, ctx.GetText())
}

func (l *TreeShapeListener) EnterSetvar_stmt(ctx *parser.Setvar_stmtContext) {
	l.results.SetvarNames = append(l.results.SetvarNames, ctx.GetText())
}

func (l *TreeShapeListener) EnterAssignment(ctx *parser.AssignmentContext) {
	l.results.SetvarOperations = append(l.results.SetvarOperations, ctx.GetText())
}

func (l *TreeShapeListener) EnterVar_assignment(ctx *parser.Var_assignmentContext) {
	l.results.DirectiveValues = append(l.results.DirectiveValues, ctx.GetText())
}

func (l *TreeShapeListener) EnterCtl_action(ctx *parser.Ctl_actionContext) {
	l.results.DirectiveValues = append(l.results.DirectiveValues, ctx.GetText())
}

func (l *TreeShapeListener) EnterCtl_id(ctx *parser.Ctl_idContext) {
	l.results.DirectiveValues = append(l.results.DirectiveValues, ctx.GetText())
}

func (l *TreeShapeListener) EnterComment(ctx *parser.CommentContext) {
	// ctx.COMMENT() can be nil if there is only a HASH without comment text
	if ctx.COMMENT() != nil {
		l.results.Comments = append(l.results.Comments, ctx.COMMENT().GetText())
	} else {
		l.results.Comments = append(l.results.Comments, "")
	}
}

func (l *TreeShapeListener) EnterTransformation_action_value(ctx *parser.Transformation_action_valueContext) {
	l.results.DirectiveValues = append(l.results.DirectiveValues, ctx.GetText())
}

func (l *TreeShapeListener) EnterComment_block(ctx *parser.Comment_blockContext) {
	l.results.CommentsBlocks++
}
