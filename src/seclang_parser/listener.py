# Copyright 2025 OWASP CRS Project
# SPDX-License-Identifier: Apache-2.0

"""Test listener for collecting parser results."""

from dataclasses import dataclass, field

from seclang_parser.SecLangParserListener import SecLangParserListener
from seclang_parser.SecLangParser import SecLangParser


@dataclass
class ParserResult:
    """Stores results from parsing for test validation."""

    comments: list[str] = field(default_factory=list)
    comment_blocks: int = 0
    variables: list[str] = field(default_factory=list)
    negated_var_count: int = 0
    collection_length_count: int = 0
    collections: list[str] = field(default_factory=list)
    collection_args: list[str] = field(default_factory=list)
    operator_list: list[str] = field(default_factory=list)
    operator_value_list: list[str] = field(default_factory=list)
    negated_operator_count: int = 0
    directive_list: list[str] = field(default_factory=list)
    directive_values: list[str] = field(default_factory=list)
    range_events: list[str] = field(default_factory=list)
    range_start_events: list[int] = field(default_factory=list)
    range_end_events: list[int] = field(default_factory=list)
    setvar_collections: list[str] = field(default_factory=list)
    setvar_names: list[str] = field(default_factory=list)
    setvar_operations: list[str] = field(default_factory=list)


class TreeShapeListener(SecLangParserListener):
    """Listener that collects parsing results for test validation."""

    def __init__(self):
        super().__init__()
        self.results = ParserResult()

    def enterInt_range(self, ctx: SecLangParser.Int_rangeContext):
        self.results.range_events.append(ctx.getText())

    def enterRange_start(self, ctx: SecLangParser.Range_startContext):
        self.results.range_start_events.append(int(ctx.getText()))

    def enterRange_end(self, ctx: SecLangParser.Range_endContext):
        self.results.range_end_events.append(int(ctx.getText()))

    def enterRemove_rule_by_id(self, ctx: SecLangParser.Remove_rule_by_idContext):
        self.results.directive_list.append(ctx.getText())

    def enterRemove_rule_by_id_int(self, ctx: SecLangParser.Remove_rule_by_id_intContext):
        self.results.directive_values.append(ctx.getText())

    def enterRemove_rule_by_id_int_range(
        self, ctx: SecLangParser.Remove_rule_by_id_int_rangeContext
    ):
        self.results.directive_values.append(ctx.getText())

    def enterRemove_rule_by_msg(self, ctx: SecLangParser.Remove_rule_by_msgContext):
        self.results.directive_list.append(ctx.getText())

    def enterRemove_rule_by_tag(self, ctx: SecLangParser.Remove_rule_by_tagContext):
        self.results.directive_list.append(ctx.getText())

    def enterString_remove_rules_values(self, ctx: SecLangParser.String_remove_rules_valuesContext):
        self.results.directive_values.append(ctx.getText())

    def enterUpdate_target_by_id(self, ctx: SecLangParser.Update_target_by_idContext):
        self.results.directive_list.append(ctx.getText())

    def enterUpdate_target_by_msg(self, ctx: SecLangParser.Update_target_by_msgContext):
        self.results.directive_list.append(ctx.getText())

    def enterUpdate_target_by_tag(self, ctx: SecLangParser.Update_target_by_tagContext):
        self.results.directive_list.append(ctx.getText())

    def enterUpdate_target_rules_values(
        self, ctx: SecLangParser.Update_target_rules_valuesContext
    ):
        self.results.directive_values.append(ctx.getText())

    def enterVariable_enum(self, ctx: SecLangParser.Variable_enumContext):
        self.results.variables.append(ctx.getText())

    def enterCollection_enum(self, ctx: SecLangParser.Collection_enumContext):
        self.results.collections.append(ctx.getText())

    def enterCollection_value(self, ctx: SecLangParser.Collection_valueContext):
        self.results.collection_args.append(ctx.getText())

    def enterVar_not(self, ctx: SecLangParser.Var_notContext):
        self.results.negated_var_count += 1

    def enterVar_count(self, ctx: SecLangParser.Var_countContext):
        self.results.collection_length_count += 1

    def enterEngine_config_rule_directive(self, ctx: SecLangParser.Engine_config_rule_directiveContext):
        self.results.directive_list.append(ctx.getText())

    def enterOperator_name(self, ctx: SecLangParser.Operator_nameContext):
        self.results.operator_list.append(ctx.getText())

    def enterOperator_value(self, ctx: SecLangParser.Operator_valueContext):
        self.results.operator_value_list.append(ctx.getText())

    def enterOperator_not(self, ctx: SecLangParser.Operator_notContext):
        self.results.negated_operator_count += 1

    def enterCol_name(self, ctx: SecLangParser.Col_nameContext):
        self.results.setvar_collections.append(ctx.getText())

    def enterSetvar_stmt(self, ctx: SecLangParser.Setvar_stmtContext):
        self.results.setvar_names.append(ctx.getText())

    def enterAssignment(self, ctx: SecLangParser.AssignmentContext):
        self.results.setvar_operations.append(ctx.getText())

    def enterVar_assignment(self, ctx: SecLangParser.Var_assignmentContext):
        self.results.directive_values.append(ctx.getText())

    def enterCtl_action(self, ctx: SecLangParser.Ctl_actionContext):
        self.results.directive_values.append(ctx.getText())

    def enterCtl_id(self, ctx: SecLangParser.Ctl_idContext):
        self.results.directive_values.append(ctx.getText())

    def enterComment(self, ctx: SecLangParser.CommentContext):
        # ctx.COMMENT() can be None if there is only a HASH without comment text
        if ctx.COMMENT() is not None:
            self.results.comments.append(ctx.COMMENT().getText())
        else:
            self.results.comments.append("")

    def enterTransformation_action_value(self, ctx: SecLangParser.Transformation_action_valueContext):
        self.results.directive_values.append(ctx.getText())

    def enterComment_block(self, ctx: SecLangParser.Comment_blockContext):
        self.results.comment_blocks += 1