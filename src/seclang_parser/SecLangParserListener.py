# Generated from SecLangParser.g4 by ANTLR 4.13.2
from antlr4 import *
if "." in __name__:
    from .SecLangParser import SecLangParser
else:
    from SecLangParser import SecLangParser

# This class defines a complete listener for a parse tree produced by SecLangParser.
class SecLangParserListener(ParseTreeListener):

    # Enter a parse tree produced by SecLangParser#configuration.
    def enterConfiguration(self, ctx:SecLangParser.ConfigurationContext):
        pass

    # Exit a parse tree produced by SecLangParser#configuration.
    def exitConfiguration(self, ctx:SecLangParser.ConfigurationContext):
        pass


    # Enter a parse tree produced by SecLangParser#stmt.
    def enterStmt(self, ctx:SecLangParser.StmtContext):
        pass

    # Exit a parse tree produced by SecLangParser#stmt.
    def exitStmt(self, ctx:SecLangParser.StmtContext):
        pass


    # Enter a parse tree produced by SecLangParser#comment_block.
    def enterComment_block(self, ctx:SecLangParser.Comment_blockContext):
        pass

    # Exit a parse tree produced by SecLangParser#comment_block.
    def exitComment_block(self, ctx:SecLangParser.Comment_blockContext):
        pass


    # Enter a parse tree produced by SecLangParser#comment.
    def enterComment(self, ctx:SecLangParser.CommentContext):
        pass

    # Exit a parse tree produced by SecLangParser#comment.
    def exitComment(self, ctx:SecLangParser.CommentContext):
        pass


    # Enter a parse tree produced by SecLangParser#engine_config_rule_directive.
    def enterEngine_config_rule_directive(self, ctx:SecLangParser.Engine_config_rule_directiveContext):
        pass

    # Exit a parse tree produced by SecLangParser#engine_config_rule_directive.
    def exitEngine_config_rule_directive(self, ctx:SecLangParser.Engine_config_rule_directiveContext):
        pass


    # Enter a parse tree produced by SecLangParser#engine_config_directive.
    def enterEngine_config_directive(self, ctx:SecLangParser.Engine_config_directiveContext):
        pass

    # Exit a parse tree produced by SecLangParser#engine_config_directive.
    def exitEngine_config_directive(self, ctx:SecLangParser.Engine_config_directiveContext):
        pass


    # Enter a parse tree produced by SecLangParser#string_engine_config_directive.
    def enterString_engine_config_directive(self, ctx:SecLangParser.String_engine_config_directiveContext):
        pass

    # Exit a parse tree produced by SecLangParser#string_engine_config_directive.
    def exitString_engine_config_directive(self, ctx:SecLangParser.String_engine_config_directiveContext):
        pass


    # Enter a parse tree produced by SecLangParser#sec_marker_directive.
    def enterSec_marker_directive(self, ctx:SecLangParser.Sec_marker_directiveContext):
        pass

    # Exit a parse tree produced by SecLangParser#sec_marker_directive.
    def exitSec_marker_directive(self, ctx:SecLangParser.Sec_marker_directiveContext):
        pass


    # Enter a parse tree produced by SecLangParser#engine_config_directive_with_param.
    def enterEngine_config_directive_with_param(self, ctx:SecLangParser.Engine_config_directive_with_paramContext):
        pass

    # Exit a parse tree produced by SecLangParser#engine_config_directive_with_param.
    def exitEngine_config_directive_with_param(self, ctx:SecLangParser.Engine_config_directive_with_paramContext):
        pass


    # Enter a parse tree produced by SecLangParser#rule_script_directive.
    def enterRule_script_directive(self, ctx:SecLangParser.Rule_script_directiveContext):
        pass

    # Exit a parse tree produced by SecLangParser#rule_script_directive.
    def exitRule_script_directive(self, ctx:SecLangParser.Rule_script_directiveContext):
        pass


    # Enter a parse tree produced by SecLangParser#file_path.
    def enterFile_path(self, ctx:SecLangParser.File_pathContext):
        pass

    # Exit a parse tree produced by SecLangParser#file_path.
    def exitFile_path(self, ctx:SecLangParser.File_pathContext):
        pass


    # Enter a parse tree produced by SecLangParser#remove_rule_by_id.
    def enterRemove_rule_by_id(self, ctx:SecLangParser.Remove_rule_by_idContext):
        pass

    # Exit a parse tree produced by SecLangParser#remove_rule_by_id.
    def exitRemove_rule_by_id(self, ctx:SecLangParser.Remove_rule_by_idContext):
        pass


    # Enter a parse tree produced by SecLangParser#remove_rule_by_id_int.
    def enterRemove_rule_by_id_int(self, ctx:SecLangParser.Remove_rule_by_id_intContext):
        pass

    # Exit a parse tree produced by SecLangParser#remove_rule_by_id_int.
    def exitRemove_rule_by_id_int(self, ctx:SecLangParser.Remove_rule_by_id_intContext):
        pass


    # Enter a parse tree produced by SecLangParser#remove_rule_by_id_int_range.
    def enterRemove_rule_by_id_int_range(self, ctx:SecLangParser.Remove_rule_by_id_int_rangeContext):
        pass

    # Exit a parse tree produced by SecLangParser#remove_rule_by_id_int_range.
    def exitRemove_rule_by_id_int_range(self, ctx:SecLangParser.Remove_rule_by_id_int_rangeContext):
        pass


    # Enter a parse tree produced by SecLangParser#operator_int_range.
    def enterOperator_int_range(self, ctx:SecLangParser.Operator_int_rangeContext):
        pass

    # Exit a parse tree produced by SecLangParser#operator_int_range.
    def exitOperator_int_range(self, ctx:SecLangParser.Operator_int_rangeContext):
        pass


    # Enter a parse tree produced by SecLangParser#int_range.
    def enterInt_range(self, ctx:SecLangParser.Int_rangeContext):
        pass

    # Exit a parse tree produced by SecLangParser#int_range.
    def exitInt_range(self, ctx:SecLangParser.Int_rangeContext):
        pass


    # Enter a parse tree produced by SecLangParser#range_start.
    def enterRange_start(self, ctx:SecLangParser.Range_startContext):
        pass

    # Exit a parse tree produced by SecLangParser#range_start.
    def exitRange_start(self, ctx:SecLangParser.Range_startContext):
        pass


    # Enter a parse tree produced by SecLangParser#range_end.
    def enterRange_end(self, ctx:SecLangParser.Range_endContext):
        pass

    # Exit a parse tree produced by SecLangParser#range_end.
    def exitRange_end(self, ctx:SecLangParser.Range_endContext):
        pass


    # Enter a parse tree produced by SecLangParser#remove_rule_by_msg.
    def enterRemove_rule_by_msg(self, ctx:SecLangParser.Remove_rule_by_msgContext):
        pass

    # Exit a parse tree produced by SecLangParser#remove_rule_by_msg.
    def exitRemove_rule_by_msg(self, ctx:SecLangParser.Remove_rule_by_msgContext):
        pass


    # Enter a parse tree produced by SecLangParser#remove_rule_by_tag.
    def enterRemove_rule_by_tag(self, ctx:SecLangParser.Remove_rule_by_tagContext):
        pass

    # Exit a parse tree produced by SecLangParser#remove_rule_by_tag.
    def exitRemove_rule_by_tag(self, ctx:SecLangParser.Remove_rule_by_tagContext):
        pass


    # Enter a parse tree produced by SecLangParser#string_remove_rules_values.
    def enterString_remove_rules_values(self, ctx:SecLangParser.String_remove_rules_valuesContext):
        pass

    # Exit a parse tree produced by SecLangParser#string_remove_rules_values.
    def exitString_remove_rules_values(self, ctx:SecLangParser.String_remove_rules_valuesContext):
        pass


    # Enter a parse tree produced by SecLangParser#update_target_by_id.
    def enterUpdate_target_by_id(self, ctx:SecLangParser.Update_target_by_idContext):
        pass

    # Exit a parse tree produced by SecLangParser#update_target_by_id.
    def exitUpdate_target_by_id(self, ctx:SecLangParser.Update_target_by_idContext):
        pass


    # Enter a parse tree produced by SecLangParser#update_target_by_msg.
    def enterUpdate_target_by_msg(self, ctx:SecLangParser.Update_target_by_msgContext):
        pass

    # Exit a parse tree produced by SecLangParser#update_target_by_msg.
    def exitUpdate_target_by_msg(self, ctx:SecLangParser.Update_target_by_msgContext):
        pass


    # Enter a parse tree produced by SecLangParser#update_target_by_tag.
    def enterUpdate_target_by_tag(self, ctx:SecLangParser.Update_target_by_tagContext):
        pass

    # Exit a parse tree produced by SecLangParser#update_target_by_tag.
    def exitUpdate_target_by_tag(self, ctx:SecLangParser.Update_target_by_tagContext):
        pass


    # Enter a parse tree produced by SecLangParser#update_action_rule.
    def enterUpdate_action_rule(self, ctx:SecLangParser.Update_action_ruleContext):
        pass

    # Exit a parse tree produced by SecLangParser#update_action_rule.
    def exitUpdate_action_rule(self, ctx:SecLangParser.Update_action_ruleContext):
        pass


    # Enter a parse tree produced by SecLangParser#id.
    def enterId(self, ctx:SecLangParser.IdContext):
        pass

    # Exit a parse tree produced by SecLangParser#id.
    def exitId(self, ctx:SecLangParser.IdContext):
        pass


    # Enter a parse tree produced by SecLangParser#engine_config_sec_cache_transformations.
    def enterEngine_config_sec_cache_transformations(self, ctx:SecLangParser.Engine_config_sec_cache_transformationsContext):
        pass

    # Exit a parse tree produced by SecLangParser#engine_config_sec_cache_transformations.
    def exitEngine_config_sec_cache_transformations(self, ctx:SecLangParser.Engine_config_sec_cache_transformationsContext):
        pass


    # Enter a parse tree produced by SecLangParser#option_list.
    def enterOption_list(self, ctx:SecLangParser.Option_listContext):
        pass

    # Exit a parse tree produced by SecLangParser#option_list.
    def exitOption_list(self, ctx:SecLangParser.Option_listContext):
        pass


    # Enter a parse tree produced by SecLangParser#option.
    def enterOption(self, ctx:SecLangParser.OptionContext):
        pass

    # Exit a parse tree produced by SecLangParser#option.
    def exitOption(self, ctx:SecLangParser.OptionContext):
        pass


    # Enter a parse tree produced by SecLangParser#option_name.
    def enterOption_name(self, ctx:SecLangParser.Option_nameContext):
        pass

    # Exit a parse tree produced by SecLangParser#option_name.
    def exitOption_name(self, ctx:SecLangParser.Option_nameContext):
        pass


    # Enter a parse tree produced by SecLangParser#config_dir_sec_action.
    def enterConfig_dir_sec_action(self, ctx:SecLangParser.Config_dir_sec_actionContext):
        pass

    # Exit a parse tree produced by SecLangParser#config_dir_sec_action.
    def exitConfig_dir_sec_action(self, ctx:SecLangParser.Config_dir_sec_actionContext):
        pass


    # Enter a parse tree produced by SecLangParser#config_dir_sec_default_action.
    def enterConfig_dir_sec_default_action(self, ctx:SecLangParser.Config_dir_sec_default_actionContext):
        pass

    # Exit a parse tree produced by SecLangParser#config_dir_sec_default_action.
    def exitConfig_dir_sec_default_action(self, ctx:SecLangParser.Config_dir_sec_default_actionContext):
        pass


    # Enter a parse tree produced by SecLangParser#stmt_audit_log.
    def enterStmt_audit_log(self, ctx:SecLangParser.Stmt_audit_logContext):
        pass

    # Exit a parse tree produced by SecLangParser#stmt_audit_log.
    def exitStmt_audit_log(self, ctx:SecLangParser.Stmt_audit_logContext):
        pass


    # Enter a parse tree produced by SecLangParser#values.
    def enterValues(self, ctx:SecLangParser.ValuesContext):
        pass

    # Exit a parse tree produced by SecLangParser#values.
    def exitValues(self, ctx:SecLangParser.ValuesContext):
        pass


    # Enter a parse tree produced by SecLangParser#action_ctl_target_value.
    def enterAction_ctl_target_value(self, ctx:SecLangParser.Action_ctl_target_valueContext):
        pass

    # Exit a parse tree produced by SecLangParser#action_ctl_target_value.
    def exitAction_ctl_target_value(self, ctx:SecLangParser.Action_ctl_target_valueContext):
        pass


    # Enter a parse tree produced by SecLangParser#update_target_rules_values.
    def enterUpdate_target_rules_values(self, ctx:SecLangParser.Update_target_rules_valuesContext):
        pass

    # Exit a parse tree produced by SecLangParser#update_target_rules_values.
    def exitUpdate_target_rules_values(self, ctx:SecLangParser.Update_target_rules_valuesContext):
        pass


    # Enter a parse tree produced by SecLangParser#operator_not.
    def enterOperator_not(self, ctx:SecLangParser.Operator_notContext):
        pass

    # Exit a parse tree produced by SecLangParser#operator_not.
    def exitOperator_not(self, ctx:SecLangParser.Operator_notContext):
        pass


    # Enter a parse tree produced by SecLangParser#operator.
    def enterOperator(self, ctx:SecLangParser.OperatorContext):
        pass

    # Exit a parse tree produced by SecLangParser#operator.
    def exitOperator(self, ctx:SecLangParser.OperatorContext):
        pass


    # Enter a parse tree produced by SecLangParser#operator_name.
    def enterOperator_name(self, ctx:SecLangParser.Operator_nameContext):
        pass

    # Exit a parse tree produced by SecLangParser#operator_name.
    def exitOperator_name(self, ctx:SecLangParser.Operator_nameContext):
        pass


    # Enter a parse tree produced by SecLangParser#operator_value.
    def enterOperator_value(self, ctx:SecLangParser.Operator_valueContext):
        pass

    # Exit a parse tree produced by SecLangParser#operator_value.
    def exitOperator_value(self, ctx:SecLangParser.Operator_valueContext):
        pass


    # Enter a parse tree produced by SecLangParser#var_not.
    def enterVar_not(self, ctx:SecLangParser.Var_notContext):
        pass

    # Exit a parse tree produced by SecLangParser#var_not.
    def exitVar_not(self, ctx:SecLangParser.Var_notContext):
        pass


    # Enter a parse tree produced by SecLangParser#var_count.
    def enterVar_count(self, ctx:SecLangParser.Var_countContext):
        pass

    # Exit a parse tree produced by SecLangParser#var_count.
    def exitVar_count(self, ctx:SecLangParser.Var_countContext):
        pass


    # Enter a parse tree produced by SecLangParser#variables.
    def enterVariables(self, ctx:SecLangParser.VariablesContext):
        pass

    # Exit a parse tree produced by SecLangParser#variables.
    def exitVariables(self, ctx:SecLangParser.VariablesContext):
        pass


    # Enter a parse tree produced by SecLangParser#update_variables.
    def enterUpdate_variables(self, ctx:SecLangParser.Update_variablesContext):
        pass

    # Exit a parse tree produced by SecLangParser#update_variables.
    def exitUpdate_variables(self, ctx:SecLangParser.Update_variablesContext):
        pass


    # Enter a parse tree produced by SecLangParser#new_target.
    def enterNew_target(self, ctx:SecLangParser.New_targetContext):
        pass

    # Exit a parse tree produced by SecLangParser#new_target.
    def exitNew_target(self, ctx:SecLangParser.New_targetContext):
        pass


    # Enter a parse tree produced by SecLangParser#var_stmt.
    def enterVar_stmt(self, ctx:SecLangParser.Var_stmtContext):
        pass

    # Exit a parse tree produced by SecLangParser#var_stmt.
    def exitVar_stmt(self, ctx:SecLangParser.Var_stmtContext):
        pass


    # Enter a parse tree produced by SecLangParser#variable_enum.
    def enterVariable_enum(self, ctx:SecLangParser.Variable_enumContext):
        pass

    # Exit a parse tree produced by SecLangParser#variable_enum.
    def exitVariable_enum(self, ctx:SecLangParser.Variable_enumContext):
        pass


    # Enter a parse tree produced by SecLangParser#collection_enum.
    def enterCollection_enum(self, ctx:SecLangParser.Collection_enumContext):
        pass

    # Exit a parse tree produced by SecLangParser#collection_enum.
    def exitCollection_enum(self, ctx:SecLangParser.Collection_enumContext):
        pass


    # Enter a parse tree produced by SecLangParser#actions.
    def enterActions(self, ctx:SecLangParser.ActionsContext):
        pass

    # Exit a parse tree produced by SecLangParser#actions.
    def exitActions(self, ctx:SecLangParser.ActionsContext):
        pass


    # Enter a parse tree produced by SecLangParser#action.
    def enterAction(self, ctx:SecLangParser.ActionContext):
        pass

    # Exit a parse tree produced by SecLangParser#action.
    def exitAction(self, ctx:SecLangParser.ActionContext):
        pass


    # Enter a parse tree produced by SecLangParser#action_only.
    def enterAction_only(self, ctx:SecLangParser.Action_onlyContext):
        pass

    # Exit a parse tree produced by SecLangParser#action_only.
    def exitAction_only(self, ctx:SecLangParser.Action_onlyContext):
        pass


    # Enter a parse tree produced by SecLangParser#disruptive_action_only.
    def enterDisruptive_action_only(self, ctx:SecLangParser.Disruptive_action_onlyContext):
        pass

    # Exit a parse tree produced by SecLangParser#disruptive_action_only.
    def exitDisruptive_action_only(self, ctx:SecLangParser.Disruptive_action_onlyContext):
        pass


    # Enter a parse tree produced by SecLangParser#non_disruptive_action_only.
    def enterNon_disruptive_action_only(self, ctx:SecLangParser.Non_disruptive_action_onlyContext):
        pass

    # Exit a parse tree produced by SecLangParser#non_disruptive_action_only.
    def exitNon_disruptive_action_only(self, ctx:SecLangParser.Non_disruptive_action_onlyContext):
        pass


    # Enter a parse tree produced by SecLangParser#flow_action_only.
    def enterFlow_action_only(self, ctx:SecLangParser.Flow_action_onlyContext):
        pass

    # Exit a parse tree produced by SecLangParser#flow_action_only.
    def exitFlow_action_only(self, ctx:SecLangParser.Flow_action_onlyContext):
        pass


    # Enter a parse tree produced by SecLangParser#action_with_params.
    def enterAction_with_params(self, ctx:SecLangParser.Action_with_paramsContext):
        pass

    # Exit a parse tree produced by SecLangParser#action_with_params.
    def exitAction_with_params(self, ctx:SecLangParser.Action_with_paramsContext):
        pass


    # Enter a parse tree produced by SecLangParser#ACTION_PHASE.
    def enterACTION_PHASE(self, ctx:SecLangParser.ACTION_PHASEContext):
        pass

    # Exit a parse tree produced by SecLangParser#ACTION_PHASE.
    def exitACTION_PHASE(self, ctx:SecLangParser.ACTION_PHASEContext):
        pass


    # Enter a parse tree produced by SecLangParser#ACTION_ID.
    def enterACTION_ID(self, ctx:SecLangParser.ACTION_IDContext):
        pass

    # Exit a parse tree produced by SecLangParser#ACTION_ID.
    def exitACTION_ID(self, ctx:SecLangParser.ACTION_IDContext):
        pass


    # Enter a parse tree produced by SecLangParser#ACTION_MATURITY.
    def enterACTION_MATURITY(self, ctx:SecLangParser.ACTION_MATURITYContext):
        pass

    # Exit a parse tree produced by SecLangParser#ACTION_MATURITY.
    def exitACTION_MATURITY(self, ctx:SecLangParser.ACTION_MATURITYContext):
        pass


    # Enter a parse tree produced by SecLangParser#ACTION_MSG.
    def enterACTION_MSG(self, ctx:SecLangParser.ACTION_MSGContext):
        pass

    # Exit a parse tree produced by SecLangParser#ACTION_MSG.
    def exitACTION_MSG(self, ctx:SecLangParser.ACTION_MSGContext):
        pass


    # Enter a parse tree produced by SecLangParser#ACTION_REV.
    def enterACTION_REV(self, ctx:SecLangParser.ACTION_REVContext):
        pass

    # Exit a parse tree produced by SecLangParser#ACTION_REV.
    def exitACTION_REV(self, ctx:SecLangParser.ACTION_REVContext):
        pass


    # Enter a parse tree produced by SecLangParser#ACTION_SEVERITY.
    def enterACTION_SEVERITY(self, ctx:SecLangParser.ACTION_SEVERITYContext):
        pass

    # Exit a parse tree produced by SecLangParser#ACTION_SEVERITY.
    def exitACTION_SEVERITY(self, ctx:SecLangParser.ACTION_SEVERITYContext):
        pass


    # Enter a parse tree produced by SecLangParser#ACTION_TAG.
    def enterACTION_TAG(self, ctx:SecLangParser.ACTION_TAGContext):
        pass

    # Exit a parse tree produced by SecLangParser#ACTION_TAG.
    def exitACTION_TAG(self, ctx:SecLangParser.ACTION_TAGContext):
        pass


    # Enter a parse tree produced by SecLangParser#ACTION_VER.
    def enterACTION_VER(self, ctx:SecLangParser.ACTION_VERContext):
        pass

    # Exit a parse tree produced by SecLangParser#ACTION_VER.
    def exitACTION_VER(self, ctx:SecLangParser.ACTION_VERContext):
        pass


    # Enter a parse tree produced by SecLangParser#disruptive_action_with_params.
    def enterDisruptive_action_with_params(self, ctx:SecLangParser.Disruptive_action_with_paramsContext):
        pass

    # Exit a parse tree produced by SecLangParser#disruptive_action_with_params.
    def exitDisruptive_action_with_params(self, ctx:SecLangParser.Disruptive_action_with_paramsContext):
        pass


    # Enter a parse tree produced by SecLangParser#non_disruptive_action_with_params.
    def enterNon_disruptive_action_with_params(self, ctx:SecLangParser.Non_disruptive_action_with_paramsContext):
        pass

    # Exit a parse tree produced by SecLangParser#non_disruptive_action_with_params.
    def exitNon_disruptive_action_with_params(self, ctx:SecLangParser.Non_disruptive_action_with_paramsContext):
        pass


    # Enter a parse tree produced by SecLangParser#data_action_with_params.
    def enterData_action_with_params(self, ctx:SecLangParser.Data_action_with_paramsContext):
        pass

    # Exit a parse tree produced by SecLangParser#data_action_with_params.
    def exitData_action_with_params(self, ctx:SecLangParser.Data_action_with_paramsContext):
        pass


    # Enter a parse tree produced by SecLangParser#flow_action_with_params.
    def enterFlow_action_with_params(self, ctx:SecLangParser.Flow_action_with_paramsContext):
        pass

    # Exit a parse tree produced by SecLangParser#flow_action_with_params.
    def exitFlow_action_with_params(self, ctx:SecLangParser.Flow_action_with_paramsContext):
        pass


    # Enter a parse tree produced by SecLangParser#action_value.
    def enterAction_value(self, ctx:SecLangParser.Action_valueContext):
        pass

    # Exit a parse tree produced by SecLangParser#action_value.
    def exitAction_value(self, ctx:SecLangParser.Action_valueContext):
        pass


    # Enter a parse tree produced by SecLangParser#action_value_types.
    def enterAction_value_types(self, ctx:SecLangParser.Action_value_typesContext):
        pass

    # Exit a parse tree produced by SecLangParser#action_value_types.
    def exitAction_value_types(self, ctx:SecLangParser.Action_value_typesContext):
        pass


    # Enter a parse tree produced by SecLangParser#string_literal.
    def enterString_literal(self, ctx:SecLangParser.String_literalContext):
        pass

    # Exit a parse tree produced by SecLangParser#string_literal.
    def exitString_literal(self, ctx:SecLangParser.String_literalContext):
        pass


    # Enter a parse tree produced by SecLangParser#ctl_action.
    def enterCtl_action(self, ctx:SecLangParser.Ctl_actionContext):
        pass

    # Exit a parse tree produced by SecLangParser#ctl_action.
    def exitCtl_action(self, ctx:SecLangParser.Ctl_actionContext):
        pass


    # Enter a parse tree produced by SecLangParser#transformation_action_value.
    def enterTransformation_action_value(self, ctx:SecLangParser.Transformation_action_valueContext):
        pass

    # Exit a parse tree produced by SecLangParser#transformation_action_value.
    def exitTransformation_action_value(self, ctx:SecLangParser.Transformation_action_valueContext):
        pass


    # Enter a parse tree produced by SecLangParser#collection_value.
    def enterCollection_value(self, ctx:SecLangParser.Collection_valueContext):
        pass

    # Exit a parse tree produced by SecLangParser#collection_value.
    def exitCollection_value(self, ctx:SecLangParser.Collection_valueContext):
        pass


    # Enter a parse tree produced by SecLangParser#setvar_action.
    def enterSetvar_action(self, ctx:SecLangParser.Setvar_actionContext):
        pass

    # Exit a parse tree produced by SecLangParser#setvar_action.
    def exitSetvar_action(self, ctx:SecLangParser.Setvar_actionContext):
        pass


    # Enter a parse tree produced by SecLangParser#col_name.
    def enterCol_name(self, ctx:SecLangParser.Col_nameContext):
        pass

    # Exit a parse tree produced by SecLangParser#col_name.
    def exitCol_name(self, ctx:SecLangParser.Col_nameContext):
        pass


    # Enter a parse tree produced by SecLangParser#setvar_stmt.
    def enterSetvar_stmt(self, ctx:SecLangParser.Setvar_stmtContext):
        pass

    # Exit a parse tree produced by SecLangParser#setvar_stmt.
    def exitSetvar_stmt(self, ctx:SecLangParser.Setvar_stmtContext):
        pass


    # Enter a parse tree produced by SecLangParser#assignment.
    def enterAssignment(self, ctx:SecLangParser.AssignmentContext):
        pass

    # Exit a parse tree produced by SecLangParser#assignment.
    def exitAssignment(self, ctx:SecLangParser.AssignmentContext):
        pass


    # Enter a parse tree produced by SecLangParser#var_assignment.
    def enterVar_assignment(self, ctx:SecLangParser.Var_assignmentContext):
        pass

    # Exit a parse tree produced by SecLangParser#var_assignment.
    def exitVar_assignment(self, ctx:SecLangParser.Var_assignmentContext):
        pass


    # Enter a parse tree produced by SecLangParser#ctl_id.
    def enterCtl_id(self, ctx:SecLangParser.Ctl_idContext):
        pass

    # Exit a parse tree produced by SecLangParser#ctl_id.
    def exitCtl_id(self, ctx:SecLangParser.Ctl_idContext):
        pass



del SecLangParser