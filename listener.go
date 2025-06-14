package main

import (
	"strconv"

	"github.com/coreruleset/seclang_parser/parser"
)

type ParserResult struct {
	comments         []string
	variables        []string
	collections      []string
	collectionArgs   []string
	directiveList    []string
	directiveValues  []string
	rangeEvents      []string
	rangeStartEvents []int
	rangeEndEvents   []int
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

func (l *TreeShapeListener) EnterComment(ctx *parser.CommentContext) {
	l.results.comments = append(l.results.comments, ctx.GetText())
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
