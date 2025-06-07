package main

import (
	"strconv"

	"github.com/coreruleset/seclang_parser/parsing"
)

type ParserResult struct {
	directiveList    []string
	directiveValues  []string
	rangeEvents      []string
	rangeStartEvents []int
	rangeEndEvents   []int
}

func (l *TreeShapeListener) EnterInt_range(ctx *parsing.Int_rangeContext) {
	l.results.rangeEvents = append(l.results.rangeEvents, ctx.GetText())
}

func (l *TreeShapeListener) EnterRange_start(ctx *parsing.Range_startContext) {
	i, err := strconv.Atoi(ctx.GetText())
	if err != nil {
		panic(err)
	}
	l.results.rangeStartEvents = append(l.results.rangeStartEvents, i)
}

func (l *TreeShapeListener) EnterRange_end(ctx *parsing.Range_endContext) {
	i, err := strconv.Atoi(ctx.GetText())
	if err != nil {
		panic(err)
	}
	l.results.rangeEndEvents = append(l.results.rangeEndEvents, i)
}

func (l *TreeShapeListener) EnterRemove_rule_by_id(ctx *parsing.Remove_rule_by_idContext) {
	l.results.directiveList = append(l.results.directiveList, ctx.GetText())
}

func (l *TreeShapeListener) EnterRemove_rule_by_id_int(ctx *parsing.Remove_rule_by_id_intContext) {
	l.results.directiveValues = append(l.results.directiveValues, ctx.GetText())
}

func (l *TreeShapeListener) EnterRemove_rule_by_id_int_range(ctx *parsing.Remove_rule_by_id_int_rangeContext) {
	l.results.directiveValues = append(l.results.directiveValues, ctx.GetText())
}

func (l *TreeShapeListener) EnterRemove_rule_by_msg(ctx *parsing.Remove_rule_by_msgContext) {
	l.results.directiveList = append(l.results.directiveList, ctx.GetText())
}

func (l *TreeShapeListener) EnterRemove_rule_by_tag(ctx *parsing.Remove_rule_by_tagContext) {
	l.results.directiveList = append(l.results.directiveList, ctx.GetText())
}

func (l *TreeShapeListener) EnterValues(ctx *parsing.ValuesContext) {
	l.results.directiveValues = append(l.results.directiveValues, ctx.GetText())
}
