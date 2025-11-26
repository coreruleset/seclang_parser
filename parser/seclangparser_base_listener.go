// Code generated from SecLangParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // SecLangParser
import "github.com/antlr4-go/antlr/v4"

// BaseSecLangParserListener is a complete listener for a parse tree produced by SecLangParser.
type BaseSecLangParserListener struct{}

var _ SecLangParserListener = &BaseSecLangParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSecLangParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSecLangParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSecLangParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSecLangParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterConfiguration is called when production configuration is entered.
func (s *BaseSecLangParserListener) EnterConfiguration(ctx *ConfigurationContext) {}

// ExitConfiguration is called when production configuration is exited.
func (s *BaseSecLangParserListener) ExitConfiguration(ctx *ConfigurationContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseSecLangParserListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseSecLangParserListener) ExitStmt(ctx *StmtContext) {}

// EnterComment_block is called when production comment_block is entered.
func (s *BaseSecLangParserListener) EnterComment_block(ctx *Comment_blockContext) {}

// ExitComment_block is called when production comment_block is exited.
func (s *BaseSecLangParserListener) ExitComment_block(ctx *Comment_blockContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseSecLangParserListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseSecLangParserListener) ExitComment(ctx *CommentContext) {}

// EnterEngine_config_rule_directive is called when production engine_config_rule_directive is entered.
func (s *BaseSecLangParserListener) EnterEngine_config_rule_directive(ctx *Engine_config_rule_directiveContext) {
}

// ExitEngine_config_rule_directive is called when production engine_config_rule_directive is exited.
func (s *BaseSecLangParserListener) ExitEngine_config_rule_directive(ctx *Engine_config_rule_directiveContext) {
}

// EnterEngine_config_directive is called when production engine_config_directive is entered.
func (s *BaseSecLangParserListener) EnterEngine_config_directive(ctx *Engine_config_directiveContext) {
}

// ExitEngine_config_directive is called when production engine_config_directive is exited.
func (s *BaseSecLangParserListener) ExitEngine_config_directive(ctx *Engine_config_directiveContext) {
}

// EnterString_engine_config_directive is called when production string_engine_config_directive is entered.
func (s *BaseSecLangParserListener) EnterString_engine_config_directive(ctx *String_engine_config_directiveContext) {
}

// ExitString_engine_config_directive is called when production string_engine_config_directive is exited.
func (s *BaseSecLangParserListener) ExitString_engine_config_directive(ctx *String_engine_config_directiveContext) {
}

// EnterSec_marker_directive is called when production sec_marker_directive is entered.
func (s *BaseSecLangParserListener) EnterSec_marker_directive(ctx *Sec_marker_directiveContext) {}

// ExitSec_marker_directive is called when production sec_marker_directive is exited.
func (s *BaseSecLangParserListener) ExitSec_marker_directive(ctx *Sec_marker_directiveContext) {}

// EnterEngine_config_directive_with_param is called when production engine_config_directive_with_param is entered.
func (s *BaseSecLangParserListener) EnterEngine_config_directive_with_param(ctx *Engine_config_directive_with_paramContext) {
}

// ExitEngine_config_directive_with_param is called when production engine_config_directive_with_param is exited.
func (s *BaseSecLangParserListener) ExitEngine_config_directive_with_param(ctx *Engine_config_directive_with_paramContext) {
}

// EnterRule_script_directive is called when production rule_script_directive is entered.
func (s *BaseSecLangParserListener) EnterRule_script_directive(ctx *Rule_script_directiveContext) {}

// ExitRule_script_directive is called when production rule_script_directive is exited.
func (s *BaseSecLangParserListener) ExitRule_script_directive(ctx *Rule_script_directiveContext) {}

// EnterFile_path is called when production file_path is entered.
func (s *BaseSecLangParserListener) EnterFile_path(ctx *File_pathContext) {}

// ExitFile_path is called when production file_path is exited.
func (s *BaseSecLangParserListener) ExitFile_path(ctx *File_pathContext) {}

// EnterRemove_rule_by_id is called when production remove_rule_by_id is entered.
func (s *BaseSecLangParserListener) EnterRemove_rule_by_id(ctx *Remove_rule_by_idContext) {}

// ExitRemove_rule_by_id is called when production remove_rule_by_id is exited.
func (s *BaseSecLangParserListener) ExitRemove_rule_by_id(ctx *Remove_rule_by_idContext) {}

// EnterRemove_rule_by_id_int is called when production remove_rule_by_id_int is entered.
func (s *BaseSecLangParserListener) EnterRemove_rule_by_id_int(ctx *Remove_rule_by_id_intContext) {}

// ExitRemove_rule_by_id_int is called when production remove_rule_by_id_int is exited.
func (s *BaseSecLangParserListener) ExitRemove_rule_by_id_int(ctx *Remove_rule_by_id_intContext) {}

// EnterRemove_rule_by_id_int_range is called when production remove_rule_by_id_int_range is entered.
func (s *BaseSecLangParserListener) EnterRemove_rule_by_id_int_range(ctx *Remove_rule_by_id_int_rangeContext) {
}

// ExitRemove_rule_by_id_int_range is called when production remove_rule_by_id_int_range is exited.
func (s *BaseSecLangParserListener) ExitRemove_rule_by_id_int_range(ctx *Remove_rule_by_id_int_rangeContext) {
}

// EnterInt_range is called when production int_range is entered.
func (s *BaseSecLangParserListener) EnterInt_range(ctx *Int_rangeContext) {}

// ExitInt_range is called when production int_range is exited.
func (s *BaseSecLangParserListener) ExitInt_range(ctx *Int_rangeContext) {}

// EnterRange_start is called when production range_start is entered.
func (s *BaseSecLangParserListener) EnterRange_start(ctx *Range_startContext) {}

// ExitRange_start is called when production range_start is exited.
func (s *BaseSecLangParserListener) ExitRange_start(ctx *Range_startContext) {}

// EnterRange_end is called when production range_end is entered.
func (s *BaseSecLangParserListener) EnterRange_end(ctx *Range_endContext) {}

// ExitRange_end is called when production range_end is exited.
func (s *BaseSecLangParserListener) ExitRange_end(ctx *Range_endContext) {}

// EnterRemove_rule_by_msg is called when production remove_rule_by_msg is entered.
func (s *BaseSecLangParserListener) EnterRemove_rule_by_msg(ctx *Remove_rule_by_msgContext) {}

// ExitRemove_rule_by_msg is called when production remove_rule_by_msg is exited.
func (s *BaseSecLangParserListener) ExitRemove_rule_by_msg(ctx *Remove_rule_by_msgContext) {}

// EnterRemove_rule_by_tag is called when production remove_rule_by_tag is entered.
func (s *BaseSecLangParserListener) EnterRemove_rule_by_tag(ctx *Remove_rule_by_tagContext) {}

// ExitRemove_rule_by_tag is called when production remove_rule_by_tag is exited.
func (s *BaseSecLangParserListener) ExitRemove_rule_by_tag(ctx *Remove_rule_by_tagContext) {}

// EnterString_remove_rules_values is called when production string_remove_rules_values is entered.
func (s *BaseSecLangParserListener) EnterString_remove_rules_values(ctx *String_remove_rules_valuesContext) {
}

// ExitString_remove_rules_values is called when production string_remove_rules_values is exited.
func (s *BaseSecLangParserListener) ExitString_remove_rules_values(ctx *String_remove_rules_valuesContext) {
}

// EnterUpdate_target_by_id is called when production update_target_by_id is entered.
func (s *BaseSecLangParserListener) EnterUpdate_target_by_id(ctx *Update_target_by_idContext) {}

// ExitUpdate_target_by_id is called when production update_target_by_id is exited.
func (s *BaseSecLangParserListener) ExitUpdate_target_by_id(ctx *Update_target_by_idContext) {}

// EnterUpdate_target_by_msg is called when production update_target_by_msg is entered.
func (s *BaseSecLangParserListener) EnterUpdate_target_by_msg(ctx *Update_target_by_msgContext) {}

// ExitUpdate_target_by_msg is called when production update_target_by_msg is exited.
func (s *BaseSecLangParserListener) ExitUpdate_target_by_msg(ctx *Update_target_by_msgContext) {}

// EnterUpdate_target_by_tag is called when production update_target_by_tag is entered.
func (s *BaseSecLangParserListener) EnterUpdate_target_by_tag(ctx *Update_target_by_tagContext) {}

// ExitUpdate_target_by_tag is called when production update_target_by_tag is exited.
func (s *BaseSecLangParserListener) ExitUpdate_target_by_tag(ctx *Update_target_by_tagContext) {}

// EnterUpdate_action_rule is called when production update_action_rule is entered.
func (s *BaseSecLangParserListener) EnterUpdate_action_rule(ctx *Update_action_ruleContext) {}

// ExitUpdate_action_rule is called when production update_action_rule is exited.
func (s *BaseSecLangParserListener) ExitUpdate_action_rule(ctx *Update_action_ruleContext) {}

// EnterId is called when production id is entered.
func (s *BaseSecLangParserListener) EnterId(ctx *IdContext) {}

// ExitId is called when production id is exited.
func (s *BaseSecLangParserListener) ExitId(ctx *IdContext) {}

// EnterEngine_config_sec_cache_transformations is called when production engine_config_sec_cache_transformations is entered.
func (s *BaseSecLangParserListener) EnterEngine_config_sec_cache_transformations(ctx *Engine_config_sec_cache_transformationsContext) {
}

// ExitEngine_config_sec_cache_transformations is called when production engine_config_sec_cache_transformations is exited.
func (s *BaseSecLangParserListener) ExitEngine_config_sec_cache_transformations(ctx *Engine_config_sec_cache_transformationsContext) {
}

// EnterOption_list is called when production option_list is entered.
func (s *BaseSecLangParserListener) EnterOption_list(ctx *Option_listContext) {}

// ExitOption_list is called when production option_list is exited.
func (s *BaseSecLangParserListener) ExitOption_list(ctx *Option_listContext) {}

// EnterOption is called when production option is entered.
func (s *BaseSecLangParserListener) EnterOption(ctx *OptionContext) {}

// ExitOption is called when production option is exited.
func (s *BaseSecLangParserListener) ExitOption(ctx *OptionContext) {}

// EnterOption_name is called when production option_name is entered.
func (s *BaseSecLangParserListener) EnterOption_name(ctx *Option_nameContext) {}

// ExitOption_name is called when production option_name is exited.
func (s *BaseSecLangParserListener) ExitOption_name(ctx *Option_nameContext) {}

// EnterConfig_dir_sec_action is called when production config_dir_sec_action is entered.
func (s *BaseSecLangParserListener) EnterConfig_dir_sec_action(ctx *Config_dir_sec_actionContext) {}

// ExitConfig_dir_sec_action is called when production config_dir_sec_action is exited.
func (s *BaseSecLangParserListener) ExitConfig_dir_sec_action(ctx *Config_dir_sec_actionContext) {}

// EnterConfig_dir_sec_default_action is called when production config_dir_sec_default_action is entered.
func (s *BaseSecLangParserListener) EnterConfig_dir_sec_default_action(ctx *Config_dir_sec_default_actionContext) {
}

// ExitConfig_dir_sec_default_action is called when production config_dir_sec_default_action is exited.
func (s *BaseSecLangParserListener) ExitConfig_dir_sec_default_action(ctx *Config_dir_sec_default_actionContext) {
}

// EnterStmt_audit_log is called when production stmt_audit_log is entered.
func (s *BaseSecLangParserListener) EnterStmt_audit_log(ctx *Stmt_audit_logContext) {}

// ExitStmt_audit_log is called when production stmt_audit_log is exited.
func (s *BaseSecLangParserListener) ExitStmt_audit_log(ctx *Stmt_audit_logContext) {}

// EnterValues is called when production values is entered.
func (s *BaseSecLangParserListener) EnterValues(ctx *ValuesContext) {}

// ExitValues is called when production values is exited.
func (s *BaseSecLangParserListener) ExitValues(ctx *ValuesContext) {}

// EnterAction_ctl_target_value is called when production action_ctl_target_value is entered.
func (s *BaseSecLangParserListener) EnterAction_ctl_target_value(ctx *Action_ctl_target_valueContext) {
}

// ExitAction_ctl_target_value is called when production action_ctl_target_value is exited.
func (s *BaseSecLangParserListener) ExitAction_ctl_target_value(ctx *Action_ctl_target_valueContext) {
}

// EnterUpdate_target_rules_values is called when production update_target_rules_values is entered.
func (s *BaseSecLangParserListener) EnterUpdate_target_rules_values(ctx *Update_target_rules_valuesContext) {
}

// ExitUpdate_target_rules_values is called when production update_target_rules_values is exited.
func (s *BaseSecLangParserListener) ExitUpdate_target_rules_values(ctx *Update_target_rules_valuesContext) {
}

// EnterOperator_not is called when production operator_not is entered.
func (s *BaseSecLangParserListener) EnterOperator_not(ctx *Operator_notContext) {}

// ExitOperator_not is called when production operator_not is exited.
func (s *BaseSecLangParserListener) ExitOperator_not(ctx *Operator_notContext) {}

// EnterOperator is called when production operator is entered.
func (s *BaseSecLangParserListener) EnterOperator(ctx *OperatorContext) {}

// ExitOperator is called when production operator is exited.
func (s *BaseSecLangParserListener) ExitOperator(ctx *OperatorContext) {}

// EnterOperator_name is called when production operator_name is entered.
func (s *BaseSecLangParserListener) EnterOperator_name(ctx *Operator_nameContext) {}

// ExitOperator_name is called when production operator_name is exited.
func (s *BaseSecLangParserListener) ExitOperator_name(ctx *Operator_nameContext) {}

// EnterOperator_value is called when production operator_value is entered.
func (s *BaseSecLangParserListener) EnterOperator_value(ctx *Operator_valueContext) {}

// ExitOperator_value is called when production operator_value is exited.
func (s *BaseSecLangParserListener) ExitOperator_value(ctx *Operator_valueContext) {}

// EnterVar_not is called when production var_not is entered.
func (s *BaseSecLangParserListener) EnterVar_not(ctx *Var_notContext) {}

// ExitVar_not is called when production var_not is exited.
func (s *BaseSecLangParserListener) ExitVar_not(ctx *Var_notContext) {}

// EnterVar_count is called when production var_count is entered.
func (s *BaseSecLangParserListener) EnterVar_count(ctx *Var_countContext) {}

// ExitVar_count is called when production var_count is exited.
func (s *BaseSecLangParserListener) ExitVar_count(ctx *Var_countContext) {}

// EnterVariables is called when production variables is entered.
func (s *BaseSecLangParserListener) EnterVariables(ctx *VariablesContext) {}

// ExitVariables is called when production variables is exited.
func (s *BaseSecLangParserListener) ExitVariables(ctx *VariablesContext) {}

// EnterUpdate_variables is called when production update_variables is entered.
func (s *BaseSecLangParserListener) EnterUpdate_variables(ctx *Update_variablesContext) {}

// ExitUpdate_variables is called when production update_variables is exited.
func (s *BaseSecLangParserListener) ExitUpdate_variables(ctx *Update_variablesContext) {}

// EnterNew_target is called when production new_target is entered.
func (s *BaseSecLangParserListener) EnterNew_target(ctx *New_targetContext) {}

// ExitNew_target is called when production new_target is exited.
func (s *BaseSecLangParserListener) ExitNew_target(ctx *New_targetContext) {}

// EnterVar_stmt is called when production var_stmt is entered.
func (s *BaseSecLangParserListener) EnterVar_stmt(ctx *Var_stmtContext) {}

// ExitVar_stmt is called when production var_stmt is exited.
func (s *BaseSecLangParserListener) ExitVar_stmt(ctx *Var_stmtContext) {}

// EnterVariable_enum is called when production variable_enum is entered.
func (s *BaseSecLangParserListener) EnterVariable_enum(ctx *Variable_enumContext) {}

// ExitVariable_enum is called when production variable_enum is exited.
func (s *BaseSecLangParserListener) ExitVariable_enum(ctx *Variable_enumContext) {}

// EnterCollection_enum is called when production collection_enum is entered.
func (s *BaseSecLangParserListener) EnterCollection_enum(ctx *Collection_enumContext) {}

// ExitCollection_enum is called when production collection_enum is exited.
func (s *BaseSecLangParserListener) ExitCollection_enum(ctx *Collection_enumContext) {}

// EnterActions is called when production actions is entered.
func (s *BaseSecLangParserListener) EnterActions(ctx *ActionsContext) {}

// ExitActions is called when production actions is exited.
func (s *BaseSecLangParserListener) ExitActions(ctx *ActionsContext) {}

// EnterAction is called when production action is entered.
func (s *BaseSecLangParserListener) EnterAction(ctx *ActionContext) {}

// ExitAction is called when production action is exited.
func (s *BaseSecLangParserListener) ExitAction(ctx *ActionContext) {}

// EnterAction_only is called when production action_only is entered.
func (s *BaseSecLangParserListener) EnterAction_only(ctx *Action_onlyContext) {}

// ExitAction_only is called when production action_only is exited.
func (s *BaseSecLangParserListener) ExitAction_only(ctx *Action_onlyContext) {}

// EnterDisruptive_action_only is called when production disruptive_action_only is entered.
func (s *BaseSecLangParserListener) EnterDisruptive_action_only(ctx *Disruptive_action_onlyContext) {}

// ExitDisruptive_action_only is called when production disruptive_action_only is exited.
func (s *BaseSecLangParserListener) ExitDisruptive_action_only(ctx *Disruptive_action_onlyContext) {}

// EnterNon_disruptive_action_only is called when production non_disruptive_action_only is entered.
func (s *BaseSecLangParserListener) EnterNon_disruptive_action_only(ctx *Non_disruptive_action_onlyContext) {
}

// ExitNon_disruptive_action_only is called when production non_disruptive_action_only is exited.
func (s *BaseSecLangParserListener) ExitNon_disruptive_action_only(ctx *Non_disruptive_action_onlyContext) {
}

// EnterFlow_action_only is called when production flow_action_only is entered.
func (s *BaseSecLangParserListener) EnterFlow_action_only(ctx *Flow_action_onlyContext) {}

// ExitFlow_action_only is called when production flow_action_only is exited.
func (s *BaseSecLangParserListener) ExitFlow_action_only(ctx *Flow_action_onlyContext) {}

// EnterAction_with_params is called when production action_with_params is entered.
func (s *BaseSecLangParserListener) EnterAction_with_params(ctx *Action_with_paramsContext) {}

// ExitAction_with_params is called when production action_with_params is exited.
func (s *BaseSecLangParserListener) ExitAction_with_params(ctx *Action_with_paramsContext) {}

// EnterACTION_PHASE is called when production ACTION_PHASE is entered.
func (s *BaseSecLangParserListener) EnterACTION_PHASE(ctx *ACTION_PHASEContext) {}

// ExitACTION_PHASE is called when production ACTION_PHASE is exited.
func (s *BaseSecLangParserListener) ExitACTION_PHASE(ctx *ACTION_PHASEContext) {}

// EnterACTION_ID is called when production ACTION_ID is entered.
func (s *BaseSecLangParserListener) EnterACTION_ID(ctx *ACTION_IDContext) {}

// ExitACTION_ID is called when production ACTION_ID is exited.
func (s *BaseSecLangParserListener) ExitACTION_ID(ctx *ACTION_IDContext) {}

// EnterACTION_MATURITY is called when production ACTION_MATURITY is entered.
func (s *BaseSecLangParserListener) EnterACTION_MATURITY(ctx *ACTION_MATURITYContext) {}

// ExitACTION_MATURITY is called when production ACTION_MATURITY is exited.
func (s *BaseSecLangParserListener) ExitACTION_MATURITY(ctx *ACTION_MATURITYContext) {}

// EnterACTION_MSG is called when production ACTION_MSG is entered.
func (s *BaseSecLangParserListener) EnterACTION_MSG(ctx *ACTION_MSGContext) {}

// ExitACTION_MSG is called when production ACTION_MSG is exited.
func (s *BaseSecLangParserListener) ExitACTION_MSG(ctx *ACTION_MSGContext) {}

// EnterACTION_REV is called when production ACTION_REV is entered.
func (s *BaseSecLangParserListener) EnterACTION_REV(ctx *ACTION_REVContext) {}

// ExitACTION_REV is called when production ACTION_REV is exited.
func (s *BaseSecLangParserListener) ExitACTION_REV(ctx *ACTION_REVContext) {}

// EnterACTION_SEVERITY is called when production ACTION_SEVERITY is entered.
func (s *BaseSecLangParserListener) EnterACTION_SEVERITY(ctx *ACTION_SEVERITYContext) {}

// ExitACTION_SEVERITY is called when production ACTION_SEVERITY is exited.
func (s *BaseSecLangParserListener) ExitACTION_SEVERITY(ctx *ACTION_SEVERITYContext) {}

// EnterACTION_TAG is called when production ACTION_TAG is entered.
func (s *BaseSecLangParserListener) EnterACTION_TAG(ctx *ACTION_TAGContext) {}

// ExitACTION_TAG is called when production ACTION_TAG is exited.
func (s *BaseSecLangParserListener) ExitACTION_TAG(ctx *ACTION_TAGContext) {}

// EnterACTION_VER is called when production ACTION_VER is entered.
func (s *BaseSecLangParserListener) EnterACTION_VER(ctx *ACTION_VERContext) {}

// ExitACTION_VER is called when production ACTION_VER is exited.
func (s *BaseSecLangParserListener) ExitACTION_VER(ctx *ACTION_VERContext) {}

// EnterDisruptive_action_with_params is called when production disruptive_action_with_params is entered.
func (s *BaseSecLangParserListener) EnterDisruptive_action_with_params(ctx *Disruptive_action_with_paramsContext) {
}

// ExitDisruptive_action_with_params is called when production disruptive_action_with_params is exited.
func (s *BaseSecLangParserListener) ExitDisruptive_action_with_params(ctx *Disruptive_action_with_paramsContext) {
}

// EnterNon_disruptive_action_with_params is called when production non_disruptive_action_with_params is entered.
func (s *BaseSecLangParserListener) EnterNon_disruptive_action_with_params(ctx *Non_disruptive_action_with_paramsContext) {
}

// ExitNon_disruptive_action_with_params is called when production non_disruptive_action_with_params is exited.
func (s *BaseSecLangParserListener) ExitNon_disruptive_action_with_params(ctx *Non_disruptive_action_with_paramsContext) {
}

// EnterData_action_with_params is called when production data_action_with_params is entered.
func (s *BaseSecLangParserListener) EnterData_action_with_params(ctx *Data_action_with_paramsContext) {
}

// ExitData_action_with_params is called when production data_action_with_params is exited.
func (s *BaseSecLangParserListener) ExitData_action_with_params(ctx *Data_action_with_paramsContext) {
}

// EnterFlow_action_with_params is called when production flow_action_with_params is entered.
func (s *BaseSecLangParserListener) EnterFlow_action_with_params(ctx *Flow_action_with_paramsContext) {
}

// ExitFlow_action_with_params is called when production flow_action_with_params is exited.
func (s *BaseSecLangParserListener) ExitFlow_action_with_params(ctx *Flow_action_with_paramsContext) {
}

// EnterAction_value is called when production action_value is entered.
func (s *BaseSecLangParserListener) EnterAction_value(ctx *Action_valueContext) {}

// ExitAction_value is called when production action_value is exited.
func (s *BaseSecLangParserListener) ExitAction_value(ctx *Action_valueContext) {}

// EnterAction_value_types is called when production action_value_types is entered.
func (s *BaseSecLangParserListener) EnterAction_value_types(ctx *Action_value_typesContext) {}

// ExitAction_value_types is called when production action_value_types is exited.
func (s *BaseSecLangParserListener) ExitAction_value_types(ctx *Action_value_typesContext) {}

// EnterString_literal is called when production string_literal is entered.
func (s *BaseSecLangParserListener) EnterString_literal(ctx *String_literalContext) {}

// ExitString_literal is called when production string_literal is exited.
func (s *BaseSecLangParserListener) ExitString_literal(ctx *String_literalContext) {}

// EnterCtl_action is called when production ctl_action is entered.
func (s *BaseSecLangParserListener) EnterCtl_action(ctx *Ctl_actionContext) {}

// ExitCtl_action is called when production ctl_action is exited.
func (s *BaseSecLangParserListener) ExitCtl_action(ctx *Ctl_actionContext) {}

// EnterTransformation_action_value is called when production transformation_action_value is entered.
func (s *BaseSecLangParserListener) EnterTransformation_action_value(ctx *Transformation_action_valueContext) {
}

// ExitTransformation_action_value is called when production transformation_action_value is exited.
func (s *BaseSecLangParserListener) ExitTransformation_action_value(ctx *Transformation_action_valueContext) {
}

// EnterCollection_value is called when production collection_value is entered.
func (s *BaseSecLangParserListener) EnterCollection_value(ctx *Collection_valueContext) {}

// ExitCollection_value is called when production collection_value is exited.
func (s *BaseSecLangParserListener) ExitCollection_value(ctx *Collection_valueContext) {}

// EnterSetvar_action is called when production setvar_action is entered.
func (s *BaseSecLangParserListener) EnterSetvar_action(ctx *Setvar_actionContext) {}

// ExitSetvar_action is called when production setvar_action is exited.
func (s *BaseSecLangParserListener) ExitSetvar_action(ctx *Setvar_actionContext) {}

// EnterCol_name is called when production col_name is entered.
func (s *BaseSecLangParserListener) EnterCol_name(ctx *Col_nameContext) {}

// ExitCol_name is called when production col_name is exited.
func (s *BaseSecLangParserListener) ExitCol_name(ctx *Col_nameContext) {}

// EnterSetvar_stmt is called when production setvar_stmt is entered.
func (s *BaseSecLangParserListener) EnterSetvar_stmt(ctx *Setvar_stmtContext) {}

// ExitSetvar_stmt is called when production setvar_stmt is exited.
func (s *BaseSecLangParserListener) ExitSetvar_stmt(ctx *Setvar_stmtContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseSecLangParserListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseSecLangParserListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterVar_assignment is called when production var_assignment is entered.
func (s *BaseSecLangParserListener) EnterVar_assignment(ctx *Var_assignmentContext) {}

// ExitVar_assignment is called when production var_assignment is exited.
func (s *BaseSecLangParserListener) ExitVar_assignment(ctx *Var_assignmentContext) {}

// EnterCtl_id is called when production ctl_id is entered.
func (s *BaseSecLangParserListener) EnterCtl_id(ctx *Ctl_idContext) {}

// ExitCtl_id is called when production ctl_id is exited.
func (s *BaseSecLangParserListener) ExitCtl_id(ctx *Ctl_idContext) {}
