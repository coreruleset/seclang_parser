// Code generated from SecLangParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // SecLangParser
import "github.com/antlr4-go/antlr/v4"

// SecLangParserListener is a complete listener for a parse tree produced by SecLangParser.
type SecLangParserListener interface {
	antlr.ParseTreeListener

	// EnterConfiguration is called when entering the configuration production.
	EnterConfiguration(c *ConfigurationContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// EnterEngine_config_rule_directive is called when entering the engine_config_rule_directive production.
	EnterEngine_config_rule_directive(c *Engine_config_rule_directiveContext)

	// EnterEngine_config_directive is called when entering the engine_config_directive production.
	EnterEngine_config_directive(c *Engine_config_directiveContext)

	// EnterString_engine_config_directive is called when entering the string_engine_config_directive production.
	EnterString_engine_config_directive(c *String_engine_config_directiveContext)

	// EnterSec_marker_directive is called when entering the sec_marker_directive production.
	EnterSec_marker_directive(c *Sec_marker_directiveContext)

	// EnterEngine_config_directive_with_param is called when entering the engine_config_directive_with_param production.
	EnterEngine_config_directive_with_param(c *Engine_config_directive_with_paramContext)

	// EnterRule_script_directive is called when entering the rule_script_directive production.
	EnterRule_script_directive(c *Rule_script_directiveContext)

	// EnterFile_path is called when entering the file_path production.
	EnterFile_path(c *File_pathContext)

	// EnterRemove_rule_by_id is called when entering the remove_rule_by_id production.
	EnterRemove_rule_by_id(c *Remove_rule_by_idContext)

	// EnterRemove_rule_by_id_int is called when entering the remove_rule_by_id_int production.
	EnterRemove_rule_by_id_int(c *Remove_rule_by_id_intContext)

	// EnterRemove_rule_by_id_int_range is called when entering the remove_rule_by_id_int_range production.
	EnterRemove_rule_by_id_int_range(c *Remove_rule_by_id_int_rangeContext)

	// EnterInt_range is called when entering the int_range production.
	EnterInt_range(c *Int_rangeContext)

	// EnterRange_start is called when entering the range_start production.
	EnterRange_start(c *Range_startContext)

	// EnterRange_end is called when entering the range_end production.
	EnterRange_end(c *Range_endContext)

	// EnterRemove_rule_by_msg is called when entering the remove_rule_by_msg production.
	EnterRemove_rule_by_msg(c *Remove_rule_by_msgContext)

	// EnterRemove_rule_by_tag is called when entering the remove_rule_by_tag production.
	EnterRemove_rule_by_tag(c *Remove_rule_by_tagContext)

	// EnterString_remove_rules_values is called when entering the string_remove_rules_values production.
	EnterString_remove_rules_values(c *String_remove_rules_valuesContext)

	// EnterUpdate_target_by_id is called when entering the update_target_by_id production.
	EnterUpdate_target_by_id(c *Update_target_by_idContext)

	// EnterUpdate_target_by_msg is called when entering the update_target_by_msg production.
	EnterUpdate_target_by_msg(c *Update_target_by_msgContext)

	// EnterUpdate_target_by_tag is called when entering the update_target_by_tag production.
	EnterUpdate_target_by_tag(c *Update_target_by_tagContext)

	// EnterUpdate_action_rule is called when entering the update_action_rule production.
	EnterUpdate_action_rule(c *Update_action_ruleContext)

	// EnterId is called when entering the id production.
	EnterId(c *IdContext)

	// EnterEngine_config_sec_cache_transformations is called when entering the engine_config_sec_cache_transformations production.
	EnterEngine_config_sec_cache_transformations(c *Engine_config_sec_cache_transformationsContext)

	// EnterOption_list is called when entering the option_list production.
	EnterOption_list(c *Option_listContext)

	// EnterOption is called when entering the option production.
	EnterOption(c *OptionContext)

	// EnterOption_name is called when entering the option_name production.
	EnterOption_name(c *Option_nameContext)

	// EnterConfig_dir_sec_action is called when entering the config_dir_sec_action production.
	EnterConfig_dir_sec_action(c *Config_dir_sec_actionContext)

	// EnterConfig_dir_sec_default_action is called when entering the config_dir_sec_default_action production.
	EnterConfig_dir_sec_default_action(c *Config_dir_sec_default_actionContext)

	// EnterStmt_audit_log is called when entering the stmt_audit_log production.
	EnterStmt_audit_log(c *Stmt_audit_logContext)

	// EnterValues is called when entering the values production.
	EnterValues(c *ValuesContext)

	// EnterAction_ctl_target_value is called when entering the action_ctl_target_value production.
	EnterAction_ctl_target_value(c *Action_ctl_target_valueContext)

	// EnterUpdate_target_rules_values is called when entering the update_target_rules_values production.
	EnterUpdate_target_rules_values(c *Update_target_rules_valuesContext)

	// EnterOperator_not is called when entering the operator_not production.
	EnterOperator_not(c *Operator_notContext)

	// EnterOperator is called when entering the operator production.
	EnterOperator(c *OperatorContext)

	// EnterOperator_name is called when entering the operator_name production.
	EnterOperator_name(c *Operator_nameContext)

	// EnterOperator_value is called when entering the operator_value production.
	EnterOperator_value(c *Operator_valueContext)

	// EnterVar_not is called when entering the var_not production.
	EnterVar_not(c *Var_notContext)

	// EnterVar_count is called when entering the var_count production.
	EnterVar_count(c *Var_countContext)

	// EnterVariables is called when entering the variables production.
	EnterVariables(c *VariablesContext)

	// EnterUpdate_variables is called when entering the update_variables production.
	EnterUpdate_variables(c *Update_variablesContext)

	// EnterNew_target is called when entering the new_target production.
	EnterNew_target(c *New_targetContext)

	// EnterVar_stmt is called when entering the var_stmt production.
	EnterVar_stmt(c *Var_stmtContext)

	// EnterVariable_enum is called when entering the variable_enum production.
	EnterVariable_enum(c *Variable_enumContext)

	// EnterCollection_enum is called when entering the collection_enum production.
	EnterCollection_enum(c *Collection_enumContext)

	// EnterActions is called when entering the actions production.
	EnterActions(c *ActionsContext)

	// EnterAction is called when entering the action production.
	EnterAction(c *ActionContext)

	// EnterAction_only is called when entering the action_only production.
	EnterAction_only(c *Action_onlyContext)

	// EnterDisruptive_action_only is called when entering the disruptive_action_only production.
	EnterDisruptive_action_only(c *Disruptive_action_onlyContext)

	// EnterNon_disruptive_action_only is called when entering the non_disruptive_action_only production.
	EnterNon_disruptive_action_only(c *Non_disruptive_action_onlyContext)

	// EnterFlow_action_only is called when entering the flow_action_only production.
	EnterFlow_action_only(c *Flow_action_onlyContext)

	// EnterAction_with_params is called when entering the action_with_params production.
	EnterAction_with_params(c *Action_with_paramsContext)

	// EnterACTION_PHASE is called when entering the ACTION_PHASE production.
	EnterACTION_PHASE(c *ACTION_PHASEContext)

	// EnterACTION_ID is called when entering the ACTION_ID production.
	EnterACTION_ID(c *ACTION_IDContext)

	// EnterACTION_MATURITY is called when entering the ACTION_MATURITY production.
	EnterACTION_MATURITY(c *ACTION_MATURITYContext)

	// EnterACTION_MSG is called when entering the ACTION_MSG production.
	EnterACTION_MSG(c *ACTION_MSGContext)

	// EnterACTION_REV is called when entering the ACTION_REV production.
	EnterACTION_REV(c *ACTION_REVContext)

	// EnterACTION_SEVERITY is called when entering the ACTION_SEVERITY production.
	EnterACTION_SEVERITY(c *ACTION_SEVERITYContext)

	// EnterACTION_TAG is called when entering the ACTION_TAG production.
	EnterACTION_TAG(c *ACTION_TAGContext)

	// EnterACTION_VER is called when entering the ACTION_VER production.
	EnterACTION_VER(c *ACTION_VERContext)

	// EnterDisruptive_action_with_params is called when entering the disruptive_action_with_params production.
	EnterDisruptive_action_with_params(c *Disruptive_action_with_paramsContext)

	// EnterNon_disruptive_action_with_params is called when entering the non_disruptive_action_with_params production.
	EnterNon_disruptive_action_with_params(c *Non_disruptive_action_with_paramsContext)

	// EnterData_action_with_params is called when entering the data_action_with_params production.
	EnterData_action_with_params(c *Data_action_with_paramsContext)

	// EnterFlow_action_with_params is called when entering the flow_action_with_params production.
	EnterFlow_action_with_params(c *Flow_action_with_paramsContext)

	// EnterAction_value is called when entering the action_value production.
	EnterAction_value(c *Action_valueContext)

	// EnterAction_value_types is called when entering the action_value_types production.
	EnterAction_value_types(c *Action_value_typesContext)

	// EnterString_literal is called when entering the string_literal production.
	EnterString_literal(c *String_literalContext)

	// EnterCtl_action is called when entering the ctl_action production.
	EnterCtl_action(c *Ctl_actionContext)

	// EnterTransformation_action_value is called when entering the transformation_action_value production.
	EnterTransformation_action_value(c *Transformation_action_valueContext)

	// EnterCollection_value is called when entering the collection_value production.
	EnterCollection_value(c *Collection_valueContext)

	// EnterSetvar_action is called when entering the setvar_action production.
	EnterSetvar_action(c *Setvar_actionContext)

	// EnterCol_name is called when entering the col_name production.
	EnterCol_name(c *Col_nameContext)

	// EnterSetvar_stmt is called when entering the setvar_stmt production.
	EnterSetvar_stmt(c *Setvar_stmtContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterVar_assignment is called when entering the var_assignment production.
	EnterVar_assignment(c *Var_assignmentContext)

	// EnterCtl_id is called when entering the ctl_id production.
	EnterCtl_id(c *Ctl_idContext)

	// ExitConfiguration is called when exiting the configuration production.
	ExitConfiguration(c *ConfigurationContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)

	// ExitEngine_config_rule_directive is called when exiting the engine_config_rule_directive production.
	ExitEngine_config_rule_directive(c *Engine_config_rule_directiveContext)

	// ExitEngine_config_directive is called when exiting the engine_config_directive production.
	ExitEngine_config_directive(c *Engine_config_directiveContext)

	// ExitString_engine_config_directive is called when exiting the string_engine_config_directive production.
	ExitString_engine_config_directive(c *String_engine_config_directiveContext)

	// ExitSec_marker_directive is called when exiting the sec_marker_directive production.
	ExitSec_marker_directive(c *Sec_marker_directiveContext)

	// ExitEngine_config_directive_with_param is called when exiting the engine_config_directive_with_param production.
	ExitEngine_config_directive_with_param(c *Engine_config_directive_with_paramContext)

	// ExitRule_script_directive is called when exiting the rule_script_directive production.
	ExitRule_script_directive(c *Rule_script_directiveContext)

	// ExitFile_path is called when exiting the file_path production.
	ExitFile_path(c *File_pathContext)

	// ExitRemove_rule_by_id is called when exiting the remove_rule_by_id production.
	ExitRemove_rule_by_id(c *Remove_rule_by_idContext)

	// ExitRemove_rule_by_id_int is called when exiting the remove_rule_by_id_int production.
	ExitRemove_rule_by_id_int(c *Remove_rule_by_id_intContext)

	// ExitRemove_rule_by_id_int_range is called when exiting the remove_rule_by_id_int_range production.
	ExitRemove_rule_by_id_int_range(c *Remove_rule_by_id_int_rangeContext)

	// ExitInt_range is called when exiting the int_range production.
	ExitInt_range(c *Int_rangeContext)

	// ExitRange_start is called when exiting the range_start production.
	ExitRange_start(c *Range_startContext)

	// ExitRange_end is called when exiting the range_end production.
	ExitRange_end(c *Range_endContext)

	// ExitRemove_rule_by_msg is called when exiting the remove_rule_by_msg production.
	ExitRemove_rule_by_msg(c *Remove_rule_by_msgContext)

	// ExitRemove_rule_by_tag is called when exiting the remove_rule_by_tag production.
	ExitRemove_rule_by_tag(c *Remove_rule_by_tagContext)

	// ExitString_remove_rules_values is called when exiting the string_remove_rules_values production.
	ExitString_remove_rules_values(c *String_remove_rules_valuesContext)

	// ExitUpdate_target_by_id is called when exiting the update_target_by_id production.
	ExitUpdate_target_by_id(c *Update_target_by_idContext)

	// ExitUpdate_target_by_msg is called when exiting the update_target_by_msg production.
	ExitUpdate_target_by_msg(c *Update_target_by_msgContext)

	// ExitUpdate_target_by_tag is called when exiting the update_target_by_tag production.
	ExitUpdate_target_by_tag(c *Update_target_by_tagContext)

	// ExitUpdate_action_rule is called when exiting the update_action_rule production.
	ExitUpdate_action_rule(c *Update_action_ruleContext)

	// ExitId is called when exiting the id production.
	ExitId(c *IdContext)

	// ExitEngine_config_sec_cache_transformations is called when exiting the engine_config_sec_cache_transformations production.
	ExitEngine_config_sec_cache_transformations(c *Engine_config_sec_cache_transformationsContext)

	// ExitOption_list is called when exiting the option_list production.
	ExitOption_list(c *Option_listContext)

	// ExitOption is called when exiting the option production.
	ExitOption(c *OptionContext)

	// ExitOption_name is called when exiting the option_name production.
	ExitOption_name(c *Option_nameContext)

	// ExitConfig_dir_sec_action is called when exiting the config_dir_sec_action production.
	ExitConfig_dir_sec_action(c *Config_dir_sec_actionContext)

	// ExitConfig_dir_sec_default_action is called when exiting the config_dir_sec_default_action production.
	ExitConfig_dir_sec_default_action(c *Config_dir_sec_default_actionContext)

	// ExitStmt_audit_log is called when exiting the stmt_audit_log production.
	ExitStmt_audit_log(c *Stmt_audit_logContext)

	// ExitValues is called when exiting the values production.
	ExitValues(c *ValuesContext)

	// ExitAction_ctl_target_value is called when exiting the action_ctl_target_value production.
	ExitAction_ctl_target_value(c *Action_ctl_target_valueContext)

	// ExitUpdate_target_rules_values is called when exiting the update_target_rules_values production.
	ExitUpdate_target_rules_values(c *Update_target_rules_valuesContext)

	// ExitOperator_not is called when exiting the operator_not production.
	ExitOperator_not(c *Operator_notContext)

	// ExitOperator is called when exiting the operator production.
	ExitOperator(c *OperatorContext)

	// ExitOperator_name is called when exiting the operator_name production.
	ExitOperator_name(c *Operator_nameContext)

	// ExitOperator_value is called when exiting the operator_value production.
	ExitOperator_value(c *Operator_valueContext)

	// ExitVar_not is called when exiting the var_not production.
	ExitVar_not(c *Var_notContext)

	// ExitVar_count is called when exiting the var_count production.
	ExitVar_count(c *Var_countContext)

	// ExitVariables is called when exiting the variables production.
	ExitVariables(c *VariablesContext)

	// ExitUpdate_variables is called when exiting the update_variables production.
	ExitUpdate_variables(c *Update_variablesContext)

	// ExitNew_target is called when exiting the new_target production.
	ExitNew_target(c *New_targetContext)

	// ExitVar_stmt is called when exiting the var_stmt production.
	ExitVar_stmt(c *Var_stmtContext)

	// ExitVariable_enum is called when exiting the variable_enum production.
	ExitVariable_enum(c *Variable_enumContext)

	// ExitCollection_enum is called when exiting the collection_enum production.
	ExitCollection_enum(c *Collection_enumContext)

	// ExitActions is called when exiting the actions production.
	ExitActions(c *ActionsContext)

	// ExitAction is called when exiting the action production.
	ExitAction(c *ActionContext)

	// ExitAction_only is called when exiting the action_only production.
	ExitAction_only(c *Action_onlyContext)

	// ExitDisruptive_action_only is called when exiting the disruptive_action_only production.
	ExitDisruptive_action_only(c *Disruptive_action_onlyContext)

	// ExitNon_disruptive_action_only is called when exiting the non_disruptive_action_only production.
	ExitNon_disruptive_action_only(c *Non_disruptive_action_onlyContext)

	// ExitFlow_action_only is called when exiting the flow_action_only production.
	ExitFlow_action_only(c *Flow_action_onlyContext)

	// ExitAction_with_params is called when exiting the action_with_params production.
	ExitAction_with_params(c *Action_with_paramsContext)

	// ExitACTION_PHASE is called when exiting the ACTION_PHASE production.
	ExitACTION_PHASE(c *ACTION_PHASEContext)

	// ExitACTION_ID is called when exiting the ACTION_ID production.
	ExitACTION_ID(c *ACTION_IDContext)

	// ExitACTION_MATURITY is called when exiting the ACTION_MATURITY production.
	ExitACTION_MATURITY(c *ACTION_MATURITYContext)

	// ExitACTION_MSG is called when exiting the ACTION_MSG production.
	ExitACTION_MSG(c *ACTION_MSGContext)

	// ExitACTION_REV is called when exiting the ACTION_REV production.
	ExitACTION_REV(c *ACTION_REVContext)

	// ExitACTION_SEVERITY is called when exiting the ACTION_SEVERITY production.
	ExitACTION_SEVERITY(c *ACTION_SEVERITYContext)

	// ExitACTION_TAG is called when exiting the ACTION_TAG production.
	ExitACTION_TAG(c *ACTION_TAGContext)

	// ExitACTION_VER is called when exiting the ACTION_VER production.
	ExitACTION_VER(c *ACTION_VERContext)

	// ExitDisruptive_action_with_params is called when exiting the disruptive_action_with_params production.
	ExitDisruptive_action_with_params(c *Disruptive_action_with_paramsContext)

	// ExitNon_disruptive_action_with_params is called when exiting the non_disruptive_action_with_params production.
	ExitNon_disruptive_action_with_params(c *Non_disruptive_action_with_paramsContext)

	// ExitData_action_with_params is called when exiting the data_action_with_params production.
	ExitData_action_with_params(c *Data_action_with_paramsContext)

	// ExitFlow_action_with_params is called when exiting the flow_action_with_params production.
	ExitFlow_action_with_params(c *Flow_action_with_paramsContext)

	// ExitAction_value is called when exiting the action_value production.
	ExitAction_value(c *Action_valueContext)

	// ExitAction_value_types is called when exiting the action_value_types production.
	ExitAction_value_types(c *Action_value_typesContext)

	// ExitString_literal is called when exiting the string_literal production.
	ExitString_literal(c *String_literalContext)

	// ExitCtl_action is called when exiting the ctl_action production.
	ExitCtl_action(c *Ctl_actionContext)

	// ExitTransformation_action_value is called when exiting the transformation_action_value production.
	ExitTransformation_action_value(c *Transformation_action_valueContext)

	// ExitCollection_value is called when exiting the collection_value production.
	ExitCollection_value(c *Collection_valueContext)

	// ExitSetvar_action is called when exiting the setvar_action production.
	ExitSetvar_action(c *Setvar_actionContext)

	// ExitCol_name is called when exiting the col_name production.
	ExitCol_name(c *Col_nameContext)

	// ExitSetvar_stmt is called when exiting the setvar_stmt production.
	ExitSetvar_stmt(c *Setvar_stmtContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitVar_assignment is called when exiting the var_assignment production.
	ExitVar_assignment(c *Var_assignmentContext)

	// ExitCtl_id is called when exiting the ctl_id production.
	ExitCtl_id(c *Ctl_idContext)
}
