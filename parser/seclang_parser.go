// Code generated from SecLangParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // SecLangParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type SecLangParser struct {
	*antlr.BaseParser
}

var SecLangParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func seclangparserParserInit() {
	staticData := &SecLangParserParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "'#'", "'+'", "'-'",
		"'*'", "'/'", "':='", "';'", "'<>'", "'<'", "'<='", "'>='", "'>'", "'('",
		"')'", "'accuracy'", "", "'append'", "'auditlog'", "'block'", "'capture'",
		"'chain'", "'ctl'", "'auditEngine'", "'auditLogParts'", "'requestBodyProcessor'",
		"'forceRequestBodyVariable'", "'requestBodyAccess'", "'ruleEngine'",
		"'ruleRemoveByTag'", "'ruleRemoveById'", "'ruleRemoveTargetById'", "'ruleRemoveTargetByTag'",
		"'deny'", "'deprecatevar'", "'drop'", "'exec'", "'expirevar'", "'id'",
		"'initcol'", "'logdata'", "'log'", "'maturity'", "'msg'", "'multiMatch'",
		"'noauditlog'", "'nolog'", "'pass'", "'pause'", "'phase'", "'prepend'",
		"'proxy'", "'redirect'", "'rev'", "'sanitiseArg'", "'sanitiseMatchedBytes'",
		"'sanitiseMatched'", "'sanitiseRequestHeader'", "'sanitiseResponseHeader'",
		"'setenv'", "'setrsc'", "'setsid'", "'setuid'", "'setvar'", "'severity'",
		"", "'skipAfter'", "'skip'", "'status'", "'tag'", "'ver'", "'xmlns'",
		"'t'", "", "", "", "'XML'", "'&'", "'beginsWith'", "'contains'", "'containsWord'",
		"'detectSQLi'", "'detectXSS'", "'endsWith'", "'eq'", "'fuzzyHash'",
		"'ge'", "'geoLookup'", "'gsbLookup'", "'gt'", "'inspectFile'", "", "'ipMatch'",
		"'le'", "'lt'", "", "'pm'", "'rbl'", "'rsub'", "'rx'", "'rxGlobal'",
		"'streq'", "'strmatch'", "'unconditionalMatch'", "'validateByteRange'",
		"'validateDTD'", "'validateHash'", "'validateSchema'", "'validateUrlEncoding'",
		"'validateUtf8Encoding'", "'verifyCC'", "'verifyCPF'", "'verifySSN'",
		"'verifySVNR'", "'within'", "", "'SecComponentSignature'", "'SecServerSignature'",
		"'SecWebAppId'", "'SecCacheTransformations'", "'SecChrootDir'", "'SecConnEngine'",
		"'SecHashEngine'", "'SecHashKey'", "'SecHashParam'", "'SecHashMethodRx'",
		"'SecHashMethodPm'", "'SecContentInjection'", "'SecArgumentSeparator'",
		"'SecAuditLogStorageDir'", "'SecAuditLogDirMode'", "'SecAuditEngine'",
		"'SecAuditLogFileMode'", "'SecAuditLog2'", "'SecAuditLog'", "'SecAuditLogFormat'",
		"'SecAuditLogParts'", "'SecAuditLogRelevantStatus'", "'SecAuditLogType'",
		"'SecDebugLog'", "'SecDebugLogLevel'", "'SecGeoLookupDb'", "'SecGsbLookupDb'",
		"'SecGuardianLog'", "'SecInterceptOnError'", "'SecConnReadStateLimit'",
		"'SecConnWriteStateLimit'", "'SecSensorId'", "'SecRuleInheritance'",
		"'SecRulePerfTime'", "'SecStreamInBodyInspection'", "'SecStreamOutBodyInspection'",
		"'SecPcreMatchLimit'", "'SecPcreMatchLimitRecursion'", "'SecArgumentsLimit'",
		"'SecRequestBodyJsonDepthLimit'", "'SecRequestBodyAccess'", "'SecRequestBodyInMemoryLimit'",
		"'SecRequestBodyLimit'", "'SecRequestBodyLimitAction'", "'SecRequestBodyNoFilesLimit'",
		"'SecResponseBodyAccess'", "'SecResponseBodyLimit'", "'SecResponseBodyLimitAction'",
		"'SecRuleEngine'", "'SecAction'", "'SecDefaultAction'", "'SecDisableBackendCompression'",
		"'SecMarker'", "'SecUnicodeMapFile'", "'Include'", "'SecCollectionTimeout'",
		"'SecHttpBlKey'", "'SecRemoteRules'", "'SecRemoteRulesFailAction'",
		"", "'SecRuleRemoveByMsg'", "'SecRuleRemoveByTag'", "'SecRuleUpdateTargetByTag'",
		"'SecRuleUpdateTargetByMsg'", "'SecRuleUpdateTargetById'", "'SecRuleUpdateActionById'",
		"'SecUploadKeepFiles'", "'SecTmpSaveUploadedFiles'", "'SecUploadDir'",
		"'SecUploadFileLimit'", "'SecUploadFileMode'", "'Abort'", "'DetectionOnly'",
		"'https'", "'Off'", "'On'", "", "'ProcessPartial'", "'Reject'", "'RelevantOnly'",
		"'Serial'", "'Warn'", "'SecXmlExternalEntity'", "'SecResponseBodyMimeType'",
		"'SecResponseBodyMimeTypesClear'", "'SecCookieFormat'", "'SecCookieV0Separator'",
		"'SecDataDir'", "'SecStatusEngine'", "'SecTmpDir'", "'SecRule'", "'SecRuleScript'",
		"", "", "", "", "'NATIVE'", "", "", "", "", "", "", "", "", "", "",
		"", "", "'.'", "", "'%{'", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "'@'",
	}
	staticData.SymbolicNames = []string{
		"", "QUOTE", "SINGLE_QUOTE", "EQUAL", "COLON", "EQUALS_PLUS", "EQUALS_MINUS",
		"COMMA", "PIPE", "CONFIG_VALUE_PATH", "NOT", "WS", "HASH", "PLUS", "MINUS",
		"STAR", "SLASH", "ASSIGN", "SEMI", "NOT_EQUAL", "LT", "LE", "GE", "GT",
		"LPAREN", "RPAREN", "ACTION_ACCURACY", "ACTION_ALLOW", "ACTION_APPEND",
		"ACTION_AUDIT_LOG", "ACTION_BLOCK", "ACTION_CAPTURE", "ACTION_CHAIN",
		"ACTION_CTL", "ACTION_CTL_AUDIT_ENGINE", "ACTION_CTL_AUDIT_LOG_PARTS",
		"ACTION_CTL_REQUEST_BODY_PROCESSOR", "ACTION_CTL_FORCE_REQ_BODY_VAR",
		"ACTION_CTL_REQUEST_BODY_ACCESS", "ACTION_CTL_RULE_ENGINE", "ACTION_CTL_RULE_REMOVE_BY_TAG",
		"ACTION_CTL_RULE_REMOVE_BY_ID", "ACTION_CTL_RULE_REMOVE_TARGET_BY_ID",
		"ACTION_CTL_RULE_REMOVE_TARGET_BY_TAG", "ACTION_DENY", "ACTION_DEPRECATE_VAR",
		"ACTION_DROP", "ACTION_EXEC", "ACTION_EXPIRE_VAR", "ACTION_ID", "ACTION_INITCOL",
		"ACTION_LOG_DATA", "ACTION_LOG", "ACTION_MATURITY", "ACTION_MSG", "ACTION_MULTI_MATCH",
		"ACTION_NO_AUDIT_LOG", "ACTION_NO_LOG", "ACTION_PASS", "ACTION_PAUSE",
		"ACTION_PHASE", "ACTION_PREPEND", "ACTION_PROXY", "ACTION_REDIRECT",
		"ACTION_REV", "ACTION_SANITISE_ARG", "ACTION_SANITISE_MATCHED_BYTES",
		"ACTION_SANITISE_MATCHED", "ACTION_SANITISE_REQUEST_HEADER", "ACTION_SANITISE_RESPONSE_HEADER",
		"ACTION_SETENV", "ACTION_SETRSC", "ACTION_SETSID", "ACTION_SETUID",
		"ACTION_SETVAR", "ACTION_SEVERITY", "ACTION_SEVERITY_VALUE", "ACTION_SKIP_AFTER",
		"ACTION_SKIP", "ACTION_STATUS", "ACTION_TAG", "ACTION_VER", "ACTION_XMLNS",
		"ACTION_TRANSFORMATION", "TRANSFORMATION_VALUE", "COLLECTION_NAME_ENUM",
		"VARIABLE_NAME_ENUM", "RUN_TIME_VAR_XML", "VAR_COUNT", "OPERATOR_BEGINS_WITH",
		"OPERATOR_CONTAINS", "OPERATOR_CONTAINS_WORD", "OPERATOR_DETECT_SQLI",
		"OPERATOR_DETECT_XSS", "OPERATOR_ENDS_WITH", "OPERATOR_EQ", "OPERATOR_FUZZY_HASH",
		"OPERATOR_GE", "OPERATOR_GEOLOOKUP", "OPERATOR_GSB_LOOKUP", "OPERATOR_GT",
		"OPERATOR_INSPECT_FILE", "OPERATOR_IP_MATCH_FROM_FILE", "OPERATOR_IP_MATCH",
		"OPERATOR_LE", "OPERATOR_LT", "OPERATOR_PM_FROM_FILE", "OPERATOR_PM",
		"OPERATOR_RBL", "OPERATOR_RSUB", "OPERATOR_RX", "OPERATOR_RX_GLOBAL",
		"OPERATOR_STR_EQ", "OPERATOR_STR_MATCH", "OPERATOR_UNCONDITIONAL_MATCH",
		"OPERATOR_VALIDATE_BYTE_RANGE", "OPERATOR_VALIDATE_DTD", "OPERATOR_VALIDATE_HASH",
		"OPERATOR_VALIDATE_SCHEMA", "OPERATOR_VALIDATE_URL_ENCODING", "OPERATOR_VALIDATE_UTF8_ENCODING",
		"OPERATOR_VERIFY_CC", "OPERATOR_VERIFY_CPF", "OPERATOR_VERIFY_SSN",
		"OPERATOR_VERIFY_SVNR", "OPERATOR_WITHIN", "AUDIT_PARTS", "CONFIG_COMPONENT_SIG",
		"CONFIG_SEC_SERVER_SIG", "CONFIG_SEC_WEB_APP_ID", "CONFIG_SEC_CACHE_TRANSFORMATIONS",
		"CONFIG_SEC_CHROOT_DIR", "CONFIG_CONN_ENGINE", "CONFIG_SEC_HASH_ENGINE",
		"CONFIG_SEC_HASH_KEY", "CONFIG_SEC_HASH_PARAM", "CONFIG_SEC_HASH_METHOD_RX",
		"CONFIG_SEC_HASH_METHOD_PM", "CONFIG_CONTENT_INJECTION", "CONFIG_SEC_ARGUMENT_SEPARATOR",
		"CONFIG_DIR_AUDIT_DIR", "CONFIG_DIR_AUDIT_DIR_MOD", "CONFIG_DIR_AUDIT_ENG",
		"CONFIG_DIR_AUDIT_FILE_MODE", "CONFIG_DIR_AUDIT_LOG2", "CONFIG_DIR_AUDIT_LOG",
		"CONFIG_DIR_AUDIT_LOG_FMT", "CONFIG_DIR_AUDIT_LOG_P", "CONFIG_DIR_AUDIT_STS",
		"CONFIG_DIR_AUDIT_TYPE", "CONFIG_DIR_DEBUG_LOG", "CONFIG_DIR_DEBUG_LVL",
		"CONFIG_DIR_GEO_DB", "CONFIG_DIR_GSB_DB", "CONFIG_SEC_GUARDIAN_LOG",
		"CONFIG_SEC_INTERCEPT_ON_ERROR", "CONFIG_SEC_CONN_R_STATE_LIMIT", "CONFIG_SEC_CONN_W_STATE_LIMIT",
		"CONFIG_SEC_SENSOR_ID", "CONFIG_SEC_RULE_INHERITANCE", "CONFIG_SEC_RULE_PERF_TIME",
		"CONFIG_SEC_STREAM_IN_BODY_INSPECTION", "CONFIG_SEC_STREAM_OUT_BODY_INSPECTION",
		"CONFIG_DIR_PCRE_MATCH_LIMIT", "CONFIG_DIR_PCRE_MATCH_LIMIT_RECURSION",
		"CONFIG_DIR_ARGS_LIMIT", "CONFIG_DIR_REQ_BODY_JSON_DEPTH_LIMIT", "CONFIG_DIR_REQ_BODY",
		"CONFIG_DIR_REQ_BODY_IN_MEMORY_LIMIT", "CONFIG_DIR_REQ_BODY_LIMIT",
		"CONFIG_DIR_REQ_BODY_LIMIT_ACTION", "CONFIG_DIR_REQ_BODY_NO_FILES_LIMIT",
		"CONFIG_DIR_RES_BODY", "CONFIG_DIR_RES_BODY_LIMIT", "CONFIG_DIR_RES_BODY_LIMIT_ACTION",
		"CONFIG_DIR_RULE_ENG", "CONFIG_DIR_SEC_ACTION", "CONFIG_DIR_SEC_DEFAULT_ACTION",
		"CONFIG_SEC_DISABLE_BACKEND_COMPRESS", "CONFIG_DIR_SEC_MARKER", "CONFIG_DIR_UNICODE_MAP_FILE",
		"CONFIG_INCLUDE", "CONFIG_SEC_COLLECTION_TIMEOUT", "CONFIG_SEC_HTTP_BLKEY",
		"CONFIG_SEC_REMOTE_RULES", "CONFIG_SEC_REMOTE_RULES_FAIL_ACTION", "CONFIG_SEC_RULE_REMOVE_BY_ID",
		"CONFIG_SEC_RULE_REMOVE_BY_MSG", "CONFIG_SEC_RULE_REMOVE_BY_TAG", "CONFIG_SEC_RULE_UPDATE_TARGET_BY_TAG",
		"CONFIG_SEC_RULE_UPDATE_TARGET_BY_MSG", "CONFIG_SEC_RULE_UPDATE_TARGET_BY_ID",
		"CONFIG_SEC_RULE_UPDATE_ACTION_BY_ID", "CONFIG_UPLOAD_KEEP_FILES", "CONFIG_UPLOAD_SAVE_TMP_FILES",
		"CONFIG_UPLOAD_DIR", "CONFIG_UPLOAD_FILE_LIMIT", "CONFIG_UPLOAD_FILE_MODE",
		"CONFIG_VALUE_ABORT", "CONFIG_VALUE_DETC", "CONFIG_VALUE_HTTPS", "CONFIG_VALUE_OFF",
		"CONFIG_VALUE_ON", "CONFIG_VALUE_PARALLEL", "CONFIG_VALUE_PROCESS_PARTIAL",
		"CONFIG_VALUE_REJECT", "CONFIG_VALUE_RELEVANT_ONLY", "CONFIG_VALUE_SERIAL",
		"CONFIG_VALUE_WARN", "CONFIG_XML_EXTERNAL_ENTITY", "CONFIG_DIR_RESPONSE_BODY_MP",
		"CONFIG_DIR_RESPONSE_BODY_MP_CLEAR", "CONFIG_DIR_SEC_COOKIE_FORMAT",
		"CONFIG_SEC_COOKIEV0_SEPARATOR", "CONFIG_DIR_SEC_DATA_DIR", "CONFIG_DIR_SEC_STATUS_ENGINE",
		"CONFIG_DIR_SEC_TMP_DIR", "CONFIG_DIR_SEC_RULE", "DIRECTIVE_SECRULESCRIPT",
		"OPTION_NAME", "SINGLE_QUOTE_BUT_SCAPED", "DOUBLE_SINGLE_QUOTE_BUT_SCAPED",
		"COMMA_BUT_SCAPED", "NATIVE", "NEWLINE", "VARIABLE_NAME", "IDENT", "INT",
		"DIGIT", "LETTER", "DICT_ELEMENT_REGEXP", "FREE_TEXT_QUOTE_MACRO_EXPANSION",
		"WS_STRING_MODE", "STRING", "MACRO_EXPANSION", "COLLECTION_NAME_SETVAR",
		"DOT", "COLLECTION_ELEMENT", "COLLECTION_WITH_MACRO", "VAR_ASSIGNMENT",
		"SPACE_SETVAR_ASSIGNMENT", "COMMA_SEPARATED_STRING", "WS_FILE_PATH_MODE",
		"XPATH_EXPRESSION", "XPATH_MODE_POP_CHARS", "ACTION_CTL_BODY_PROCESSOR_TYPE",
		"STRING_LITERAL", "SPACE_COL", "SPACE_VAR", "NEWLINE_VAR", "COLLECTION_ELEMENT_VALUE",
		"SPACE_COL_ELEM", "NEWLINE_COL_ELEM", "SKIP_CHARS", "OPERATOR_UNQUOTED_STRING",
		"AT", "OPERATOR_QUOTED_STRING", "COMMENT", "HASH_COMMENT_BLOCK", "BLOCK_COMMENT_END",
		"PIPE_DEFAULT", "COMMA_DEFAULT", "COLON_DEFAULT", "EQUAL_DEFAULT", "NOT_DEFAULT",
		"QUOTE_DEFAULT", "SINGLE_QUOTE_SETVAR",
	}
	staticData.RuleNames = []string{
		"configuration", "stmt", "comment_block", "comment", "engine_config_rule_directive",
		"engine_config_directive", "string_engine_config_directive", "sec_marker_directive",
		"engine_config_directive_with_param", "rule_script_directive", "file_path",
		"remove_rule_by_id", "remove_rule_by_id_values", "int_range", "range_start",
		"range_end", "string_remove_rules", "string_remove_rules_values", "update_target_rules",
		"update_action_rule", "id", "engine_config_sec_cache_transformations",
		"option_list", "option", "option_name", "engine_config_action_directive",
		"stmt_audit_log", "values", "action_ctl_target_value", "update_target_rules_values",
		"operator_not", "operator", "operator_name", "operator_value", "var_not",
		"var_count", "variables", "update_variables", "new_target", "var_stmt",
		"variable_enum", "ctl_variable_enum", "collection_enum", "ctl_collection_enum",
		"actions", "action", "action_only", "disruptive_action_only", "non_disruptive_action_only",
		"flow_action_only", "action_with_params", "metadata_action_with_params",
		"disruptive_action_with_params", "non_disruptive_action_with_params",
		"data_action_with_params", "flow_action_with_params", "action_value",
		"action_value_types", "string_literal", "ctl_action", "transformation_action_value",
		"collection_value", "ctl_collection_value", "setvar_action", "col_name",
		"setvar_stmt", "assignment", "var_assignment", "ctl_id",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 266, 669, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7,
		52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57, 7, 57,
		2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7, 62, 2,
		63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67, 2, 68,
		7, 68, 1, 0, 5, 0, 140, 8, 0, 10, 0, 12, 0, 143, 9, 0, 1, 0, 1, 0, 1, 1,
		3, 1, 148, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 154, 8, 1, 1, 1, 3, 1, 157,
		8, 1, 1, 1, 1, 1, 1, 1, 3, 1, 162, 8, 1, 1, 1, 3, 1, 165, 8, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 3, 1, 172, 8, 1, 1, 1, 3, 1, 175, 8, 1, 1, 1, 1, 1,
		4, 1, 179, 8, 1, 11, 1, 12, 1, 180, 1, 1, 3, 1, 184, 8, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 3, 1, 190, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1,
		198, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 205, 8, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 214, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 3, 1, 223, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 3, 1, 234, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 241,
		8, 1, 1, 1, 1, 1, 4, 1, 245, 8, 1, 11, 1, 12, 1, 246, 3, 1, 249, 8, 1,
		1, 2, 4, 2, 252, 8, 2, 11, 2, 12, 2, 253, 1, 2, 3, 2, 257, 8, 2, 1, 3,
		1, 3, 3, 3, 261, 8, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 293,
		8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 299, 8, 6, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 3, 12, 313, 8, 12,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 3,
		16, 325, 8, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 3, 18, 332, 8, 18, 1,
		19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 5, 22,
		344, 8, 22, 10, 22, 12, 22, 347, 9, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 3, 25, 359, 8, 25, 1, 26, 1, 26,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27,
		1, 27, 1, 27, 3, 27, 386, 8, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 3, 28, 394, 8, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28,
		1, 28, 3, 28, 404, 8, 28, 1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 410, 8, 28,
		3, 28, 412, 8, 28, 1, 29, 1, 29, 1, 29, 3, 29, 417, 8, 29, 1, 30, 1, 30,
		1, 31, 1, 31, 3, 31, 423, 8, 31, 1, 31, 1, 31, 1, 31, 3, 31, 428, 8, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 437, 8, 31, 1,
		32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 3, 33, 445, 8, 33, 1, 33, 1, 33,
		1, 33, 3, 33, 450, 8, 33, 5, 33, 452, 8, 33, 10, 33, 12, 33, 455, 9, 33,
		1, 33, 1, 33, 3, 33, 459, 8, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 3,
		36, 466, 8, 36, 1, 36, 3, 36, 469, 8, 36, 1, 36, 3, 36, 472, 8, 36, 1,
		36, 1, 36, 3, 36, 476, 8, 36, 1, 36, 1, 36, 3, 36, 480, 8, 36, 1, 36, 3,
		36, 483, 8, 36, 1, 36, 1, 36, 3, 36, 487, 8, 36, 5, 36, 489, 8, 36, 10,
		36, 12, 36, 492, 9, 36, 1, 37, 3, 37, 495, 8, 37, 1, 37, 3, 37, 498, 8,
		37, 1, 37, 3, 37, 501, 8, 37, 1, 37, 1, 37, 3, 37, 505, 8, 37, 1, 37, 1,
		37, 3, 37, 509, 8, 37, 1, 37, 3, 37, 512, 8, 37, 1, 37, 1, 37, 3, 37, 516,
		8, 37, 5, 37, 518, 8, 37, 10, 37, 12, 37, 521, 9, 37, 1, 38, 1, 38, 1,
		39, 1, 39, 1, 39, 1, 39, 3, 39, 529, 8, 39, 3, 39, 531, 8, 39, 1, 40, 1,
		40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 44,
		5, 44, 545, 8, 44, 10, 44, 12, 44, 548, 9, 44, 1, 44, 1, 44, 1, 45, 1,
		45, 1, 45, 3, 45, 555, 8, 45, 1, 45, 3, 45, 558, 8, 45, 1, 45, 1, 45, 1,
		45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 3, 45, 570, 8, 45,
		1, 46, 1, 46, 1, 46, 3, 46, 575, 8, 46, 1, 47, 1, 47, 1, 48, 1, 48, 1,
		49, 1, 49, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 3, 50, 588, 8, 50, 1, 51,
		1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 3, 51, 598, 8, 51, 1,
		52, 1, 52, 1, 53, 1, 53, 1, 54, 1, 54, 1, 55, 1, 55, 1, 56, 1, 56, 1, 56,
		1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 3, 56, 617, 8, 56, 1, 57, 1,
		57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 3, 57,
		630, 8, 57, 1, 58, 1, 58, 1, 59, 1, 59, 1, 60, 1, 60, 1, 61, 1, 61, 1,
		61, 3, 61, 641, 8, 61, 1, 62, 1, 62, 1, 62, 3, 62, 646, 8, 62, 1, 63, 1,
		63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 64, 1, 64, 1, 65, 1, 65, 1, 65, 4, 65,
		659, 8, 65, 11, 65, 12, 65, 660, 1, 66, 1, 66, 1, 67, 1, 67, 1, 68, 1,
		68, 1, 68, 0, 0, 69, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26,
		28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62,
		64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98,
		100, 102, 104, 106, 108, 110, 112, 114, 116, 118, 120, 122, 124, 126, 128,
		130, 132, 134, 136, 0, 15, 1, 1, 259, 259, 8, 0, 131, 139, 150, 167, 169,
		175, 178, 178, 180, 180, 182, 183, 185, 185, 209, 216, 3, 0, 225, 225,
		233, 233, 241, 241, 3, 0, 140, 149, 193, 197, 227, 227, 1, 0, 89, 125,
		2, 0, 85, 85, 87, 87, 5, 0, 27, 27, 30, 30, 44, 44, 46, 46, 58, 59, 5,
		0, 29, 29, 31, 31, 52, 52, 55, 57, 67, 67, 1, 0, 62, 63, 8, 0, 28, 28,
		33, 33, 45, 45, 47, 48, 50, 51, 61, 61, 65, 66, 68, 74, 2, 0, 79, 79, 82,
		82, 1, 0, 77, 78, 1, 0, 34, 43, 2, 0, 3, 3, 5, 6, 1, 0, 226, 227, 739,
		0, 141, 1, 0, 0, 0, 2, 248, 1, 0, 0, 0, 4, 251, 1, 0, 0, 0, 6, 258, 1,
		0, 0, 0, 8, 262, 1, 0, 0, 0, 10, 292, 1, 0, 0, 0, 12, 298, 1, 0, 0, 0,
		14, 300, 1, 0, 0, 0, 16, 302, 1, 0, 0, 0, 18, 304, 1, 0, 0, 0, 20, 306,
		1, 0, 0, 0, 22, 308, 1, 0, 0, 0, 24, 312, 1, 0, 0, 0, 26, 314, 1, 0, 0,
		0, 28, 318, 1, 0, 0, 0, 30, 320, 1, 0, 0, 0, 32, 324, 1, 0, 0, 0, 34, 326,
		1, 0, 0, 0, 36, 331, 1, 0, 0, 0, 38, 333, 1, 0, 0, 0, 40, 335, 1, 0, 0,
		0, 42, 337, 1, 0, 0, 0, 44, 339, 1, 0, 0, 0, 46, 350, 1, 0, 0, 0, 48, 354,
		1, 0, 0, 0, 50, 358, 1, 0, 0, 0, 52, 360, 1, 0, 0, 0, 54, 385, 1, 0, 0,
		0, 56, 411, 1, 0, 0, 0, 58, 416, 1, 0, 0, 0, 60, 418, 1, 0, 0, 0, 62, 436,
		1, 0, 0, 0, 64, 438, 1, 0, 0, 0, 66, 458, 1, 0, 0, 0, 68, 460, 1, 0, 0,
		0, 70, 462, 1, 0, 0, 0, 72, 465, 1, 0, 0, 0, 74, 494, 1, 0, 0, 0, 76, 522,
		1, 0, 0, 0, 78, 530, 1, 0, 0, 0, 80, 532, 1, 0, 0, 0, 82, 534, 1, 0, 0,
		0, 84, 536, 1, 0, 0, 0, 86, 538, 1, 0, 0, 0, 88, 540, 1, 0, 0, 0, 90, 569,
		1, 0, 0, 0, 92, 574, 1, 0, 0, 0, 94, 576, 1, 0, 0, 0, 96, 578, 1, 0, 0,
		0, 98, 580, 1, 0, 0, 0, 100, 587, 1, 0, 0, 0, 102, 597, 1, 0, 0, 0, 104,
		599, 1, 0, 0, 0, 106, 601, 1, 0, 0, 0, 108, 603, 1, 0, 0, 0, 110, 605,
		1, 0, 0, 0, 112, 616, 1, 0, 0, 0, 114, 629, 1, 0, 0, 0, 116, 631, 1, 0,
		0, 0, 118, 633, 1, 0, 0, 0, 120, 635, 1, 0, 0, 0, 122, 640, 1, 0, 0, 0,
		124, 645, 1, 0, 0, 0, 126, 647, 1, 0, 0, 0, 128, 653, 1, 0, 0, 0, 130,
		658, 1, 0, 0, 0, 132, 662, 1, 0, 0, 0, 134, 664, 1, 0, 0, 0, 136, 666,
		1, 0, 0, 0, 138, 140, 3, 2, 1, 0, 139, 138, 1, 0, 0, 0, 140, 143, 1, 0,
		0, 0, 141, 139, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 144, 1, 0, 0, 0,
		143, 141, 1, 0, 0, 0, 144, 145, 5, 0, 0, 1, 145, 1, 1, 0, 0, 0, 146, 148,
		3, 4, 2, 0, 147, 146, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 149, 1, 0,
		0, 0, 149, 150, 3, 8, 4, 0, 150, 151, 3, 72, 36, 0, 151, 153, 3, 62, 31,
		0, 152, 154, 3, 88, 44, 0, 153, 152, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0,
		154, 249, 1, 0, 0, 0, 155, 157, 3, 4, 2, 0, 156, 155, 1, 0, 0, 0, 156,
		157, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 159, 3, 18, 9, 0, 159, 161,
		3, 20, 10, 0, 160, 162, 3, 88, 44, 0, 161, 160, 1, 0, 0, 0, 161, 162, 1,
		0, 0, 0, 162, 249, 1, 0, 0, 0, 163, 165, 3, 4, 2, 0, 164, 163, 1, 0, 0,
		0, 164, 165, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 167, 3, 18, 9, 0, 167,
		168, 5, 1, 0, 0, 168, 169, 3, 20, 10, 0, 169, 171, 5, 1, 0, 0, 170, 172,
		3, 88, 44, 0, 171, 170, 1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172, 249, 1,
		0, 0, 0, 173, 175, 3, 4, 2, 0, 174, 173, 1, 0, 0, 0, 174, 175, 1, 0, 0,
		0, 175, 176, 1, 0, 0, 0, 176, 178, 3, 22, 11, 0, 177, 179, 3, 24, 12, 0,
		178, 177, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 178, 1, 0, 0, 0, 180,
		181, 1, 0, 0, 0, 181, 249, 1, 0, 0, 0, 182, 184, 3, 4, 2, 0, 183, 182,
		1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184, 185, 1, 0, 0, 0, 185, 186, 3, 32,
		16, 0, 186, 187, 3, 34, 17, 0, 187, 249, 1, 0, 0, 0, 188, 190, 3, 4, 2,
		0, 189, 188, 1, 0, 0, 0, 189, 190, 1, 0, 0, 0, 190, 191, 1, 0, 0, 0, 191,
		192, 3, 32, 16, 0, 192, 193, 5, 1, 0, 0, 193, 194, 3, 34, 17, 0, 194, 195,
		5, 1, 0, 0, 195, 249, 1, 0, 0, 0, 196, 198, 3, 4, 2, 0, 197, 196, 1, 0,
		0, 0, 197, 198, 1, 0, 0, 0, 198, 199, 1, 0, 0, 0, 199, 200, 3, 36, 18,
		0, 200, 201, 3, 58, 29, 0, 201, 202, 3, 74, 37, 0, 202, 249, 1, 0, 0, 0,
		203, 205, 3, 4, 2, 0, 204, 203, 1, 0, 0, 0, 204, 205, 1, 0, 0, 0, 205,
		206, 1, 0, 0, 0, 206, 207, 3, 36, 18, 0, 207, 208, 5, 1, 0, 0, 208, 209,
		3, 58, 29, 0, 209, 210, 5, 1, 0, 0, 210, 211, 3, 74, 37, 0, 211, 249, 1,
		0, 0, 0, 212, 214, 3, 4, 2, 0, 213, 212, 1, 0, 0, 0, 213, 214, 1, 0, 0,
		0, 214, 215, 1, 0, 0, 0, 215, 216, 3, 36, 18, 0, 216, 217, 3, 58, 29, 0,
		217, 218, 3, 74, 37, 0, 218, 219, 5, 8, 0, 0, 219, 220, 3, 76, 38, 0, 220,
		249, 1, 0, 0, 0, 221, 223, 3, 4, 2, 0, 222, 221, 1, 0, 0, 0, 222, 223,
		1, 0, 0, 0, 223, 224, 1, 0, 0, 0, 224, 225, 3, 36, 18, 0, 225, 226, 5,
		1, 0, 0, 226, 227, 3, 58, 29, 0, 227, 228, 5, 1, 0, 0, 228, 229, 3, 74,
		37, 0, 229, 230, 5, 8, 0, 0, 230, 231, 3, 76, 38, 0, 231, 249, 1, 0, 0,
		0, 232, 234, 3, 4, 2, 0, 233, 232, 1, 0, 0, 0, 233, 234, 1, 0, 0, 0, 234,
		235, 1, 0, 0, 0, 235, 236, 3, 38, 19, 0, 236, 237, 3, 40, 20, 0, 237, 238,
		3, 88, 44, 0, 238, 249, 1, 0, 0, 0, 239, 241, 3, 4, 2, 0, 240, 239, 1,
		0, 0, 0, 240, 241, 1, 0, 0, 0, 241, 242, 1, 0, 0, 0, 242, 249, 3, 10, 5,
		0, 243, 245, 3, 4, 2, 0, 244, 243, 1, 0, 0, 0, 245, 246, 1, 0, 0, 0, 246,
		244, 1, 0, 0, 0, 246, 247, 1, 0, 0, 0, 247, 249, 1, 0, 0, 0, 248, 147,
		1, 0, 0, 0, 248, 156, 1, 0, 0, 0, 248, 164, 1, 0, 0, 0, 248, 174, 1, 0,
		0, 0, 248, 183, 1, 0, 0, 0, 248, 189, 1, 0, 0, 0, 248, 197, 1, 0, 0, 0,
		248, 204, 1, 0, 0, 0, 248, 213, 1, 0, 0, 0, 248, 222, 1, 0, 0, 0, 248,
		233, 1, 0, 0, 0, 248, 240, 1, 0, 0, 0, 248, 244, 1, 0, 0, 0, 249, 3, 1,
		0, 0, 0, 250, 252, 3, 6, 3, 0, 251, 250, 1, 0, 0, 0, 252, 253, 1, 0, 0,
		0, 253, 251, 1, 0, 0, 0, 253, 254, 1, 0, 0, 0, 254, 256, 1, 0, 0, 0, 255,
		257, 7, 0, 0, 0, 256, 255, 1, 0, 0, 0, 256, 257, 1, 0, 0, 0, 257, 5, 1,
		0, 0, 0, 258, 260, 5, 12, 0, 0, 259, 261, 5, 257, 0, 0, 260, 259, 1, 0,
		0, 0, 260, 261, 1, 0, 0, 0, 261, 7, 1, 0, 0, 0, 262, 263, 5, 217, 0, 0,
		263, 9, 1, 0, 0, 0, 264, 265, 3, 52, 26, 0, 265, 266, 3, 54, 27, 0, 266,
		293, 1, 0, 0, 0, 267, 268, 3, 52, 26, 0, 268, 269, 5, 1, 0, 0, 269, 270,
		3, 54, 27, 0, 270, 271, 5, 1, 0, 0, 271, 293, 1, 0, 0, 0, 272, 273, 3,
		50, 25, 0, 273, 274, 3, 88, 44, 0, 274, 293, 1, 0, 0, 0, 275, 276, 3, 12,
		6, 0, 276, 277, 5, 1, 0, 0, 277, 278, 3, 54, 27, 0, 278, 279, 5, 1, 0,
		0, 279, 293, 1, 0, 0, 0, 280, 281, 3, 14, 7, 0, 281, 282, 5, 1, 0, 0, 282,
		283, 3, 54, 27, 0, 283, 284, 5, 1, 0, 0, 284, 293, 1, 0, 0, 0, 285, 286,
		3, 16, 8, 0, 286, 287, 3, 54, 27, 0, 287, 293, 1, 0, 0, 0, 288, 289, 3,
		42, 21, 0, 289, 290, 3, 54, 27, 0, 290, 291, 3, 44, 22, 0, 291, 293, 1,
		0, 0, 0, 292, 264, 1, 0, 0, 0, 292, 267, 1, 0, 0, 0, 292, 272, 1, 0, 0,
		0, 292, 275, 1, 0, 0, 0, 292, 280, 1, 0, 0, 0, 292, 285, 1, 0, 0, 0, 292,
		288, 1, 0, 0, 0, 293, 11, 1, 0, 0, 0, 294, 299, 1, 0, 0, 0, 295, 299, 5,
		127, 0, 0, 296, 299, 5, 128, 0, 0, 297, 299, 5, 129, 0, 0, 298, 294, 1,
		0, 0, 0, 298, 295, 1, 0, 0, 0, 298, 296, 1, 0, 0, 0, 298, 297, 1, 0, 0,
		0, 299, 13, 1, 0, 0, 0, 300, 301, 5, 179, 0, 0, 301, 15, 1, 0, 0, 0, 302,
		303, 7, 1, 0, 0, 303, 17, 1, 0, 0, 0, 304, 305, 5, 218, 0, 0, 305, 19,
		1, 0, 0, 0, 306, 307, 5, 9, 0, 0, 307, 21, 1, 0, 0, 0, 308, 309, 5, 186,
		0, 0, 309, 23, 1, 0, 0, 0, 310, 313, 5, 227, 0, 0, 311, 313, 3, 26, 13,
		0, 312, 310, 1, 0, 0, 0, 312, 311, 1, 0, 0, 0, 313, 25, 1, 0, 0, 0, 314,
		315, 3, 28, 14, 0, 315, 316, 5, 14, 0, 0, 316, 317, 3, 30, 15, 0, 317,
		27, 1, 0, 0, 0, 318, 319, 5, 227, 0, 0, 319, 29, 1, 0, 0, 0, 320, 321,
		5, 227, 0, 0, 321, 31, 1, 0, 0, 0, 322, 325, 5, 187, 0, 0, 323, 325, 5,
		188, 0, 0, 324, 322, 1, 0, 0, 0, 324, 323, 1, 0, 0, 0, 325, 33, 1, 0, 0,
		0, 326, 327, 7, 2, 0, 0, 327, 35, 1, 0, 0, 0, 328, 332, 5, 191, 0, 0, 329,
		332, 5, 190, 0, 0, 330, 332, 5, 189, 0, 0, 331, 328, 1, 0, 0, 0, 331, 329,
		1, 0, 0, 0, 331, 330, 1, 0, 0, 0, 332, 37, 1, 0, 0, 0, 333, 334, 5, 192,
		0, 0, 334, 39, 1, 0, 0, 0, 335, 336, 5, 227, 0, 0, 336, 41, 1, 0, 0, 0,
		337, 338, 5, 130, 0, 0, 338, 43, 1, 0, 0, 0, 339, 340, 5, 1, 0, 0, 340,
		345, 3, 46, 23, 0, 341, 342, 5, 7, 0, 0, 342, 344, 3, 46, 23, 0, 343, 341,
		1, 0, 0, 0, 344, 347, 1, 0, 0, 0, 345, 343, 1, 0, 0, 0, 345, 346, 1, 0,
		0, 0, 346, 348, 1, 0, 0, 0, 347, 345, 1, 0, 0, 0, 348, 349, 5, 1, 0, 0,
		349, 45, 1, 0, 0, 0, 350, 351, 3, 48, 24, 0, 351, 352, 5, 4, 0, 0, 352,
		353, 3, 54, 27, 0, 353, 47, 1, 0, 0, 0, 354, 355, 5, 219, 0, 0, 355, 49,
		1, 0, 0, 0, 356, 359, 5, 176, 0, 0, 357, 359, 5, 177, 0, 0, 358, 356, 1,
		0, 0, 0, 358, 357, 1, 0, 0, 0, 359, 51, 1, 0, 0, 0, 360, 361, 7, 3, 0,
		0, 361, 53, 1, 0, 0, 0, 362, 386, 5, 227, 0, 0, 363, 386, 3, 26, 13, 0,
		364, 386, 5, 202, 0, 0, 365, 386, 5, 201, 0, 0, 366, 386, 5, 207, 0, 0,
		367, 386, 5, 203, 0, 0, 368, 386, 5, 200, 0, 0, 369, 386, 5, 206, 0, 0,
		370, 386, 5, 223, 0, 0, 371, 386, 5, 198, 0, 0, 372, 386, 5, 208, 0, 0,
		373, 386, 5, 199, 0, 0, 374, 386, 5, 204, 0, 0, 375, 386, 5, 205, 0, 0,
		376, 377, 5, 9, 0, 0, 377, 386, 5, 227, 0, 0, 378, 386, 5, 9, 0, 0, 379,
		386, 5, 233, 0, 0, 380, 386, 5, 225, 0, 0, 381, 386, 5, 241, 0, 0, 382,
		386, 5, 245, 0, 0, 383, 386, 5, 126, 0, 0, 384, 386, 3, 56, 28, 0, 385,
		362, 1, 0, 0, 0, 385, 363, 1, 0, 0, 0, 385, 364, 1, 0, 0, 0, 385, 365,
		1, 0, 0, 0, 385, 366, 1, 0, 0, 0, 385, 367, 1, 0, 0, 0, 385, 368, 1, 0,
		0, 0, 385, 369, 1, 0, 0, 0, 385, 370, 1, 0, 0, 0, 385, 371, 1, 0, 0, 0,
		385, 372, 1, 0, 0, 0, 385, 373, 1, 0, 0, 0, 385, 374, 1, 0, 0, 0, 385,
		375, 1, 0, 0, 0, 385, 376, 1, 0, 0, 0, 385, 378, 1, 0, 0, 0, 385, 379,
		1, 0, 0, 0, 385, 380, 1, 0, 0, 0, 385, 381, 1, 0, 0, 0, 385, 382, 1, 0,
		0, 0, 385, 383, 1, 0, 0, 0, 385, 384, 1, 0, 0, 0, 386, 55, 1, 0, 0, 0,
		387, 394, 3, 136, 68, 0, 388, 389, 5, 2, 0, 0, 389, 390, 3, 116, 58, 0,
		390, 391, 5, 2, 0, 0, 391, 394, 1, 0, 0, 0, 392, 394, 5, 225, 0, 0, 393,
		387, 1, 0, 0, 0, 393, 388, 1, 0, 0, 0, 393, 392, 1, 0, 0, 0, 394, 395,
		1, 0, 0, 0, 395, 396, 5, 18, 0, 0, 396, 412, 3, 82, 41, 0, 397, 404, 3,
		136, 68, 0, 398, 399, 5, 2, 0, 0, 399, 400, 3, 116, 58, 0, 400, 401, 5,
		2, 0, 0, 401, 404, 1, 0, 0, 0, 402, 404, 5, 225, 0, 0, 403, 397, 1, 0,
		0, 0, 403, 398, 1, 0, 0, 0, 403, 402, 1, 0, 0, 0, 404, 405, 1, 0, 0, 0,
		405, 406, 5, 18, 0, 0, 406, 409, 3, 86, 43, 0, 407, 408, 5, 4, 0, 0, 408,
		410, 3, 124, 62, 0, 409, 407, 1, 0, 0, 0, 409, 410, 1, 0, 0, 0, 410, 412,
		1, 0, 0, 0, 411, 393, 1, 0, 0, 0, 411, 403, 1, 0, 0, 0, 412, 57, 1, 0,
		0, 0, 413, 417, 5, 227, 0, 0, 414, 417, 3, 26, 13, 0, 415, 417, 5, 233,
		0, 0, 416, 413, 1, 0, 0, 0, 416, 414, 1, 0, 0, 0, 416, 415, 1, 0, 0, 0,
		417, 59, 1, 0, 0, 0, 418, 419, 5, 10, 0, 0, 419, 61, 1, 0, 0, 0, 420, 422,
		5, 1, 0, 0, 421, 423, 3, 60, 30, 0, 422, 421, 1, 0, 0, 0, 422, 423, 1,
		0, 0, 0, 423, 424, 1, 0, 0, 0, 424, 425, 5, 255, 0, 0, 425, 427, 3, 64,
		32, 0, 426, 428, 3, 66, 33, 0, 427, 426, 1, 0, 0, 0, 427, 428, 1, 0, 0,
		0, 428, 429, 1, 0, 0, 0, 429, 430, 5, 1, 0, 0, 430, 437, 1, 0, 0, 0, 431,
		432, 5, 1, 0, 0, 432, 433, 3, 66, 33, 0, 433, 434, 5, 1, 0, 0, 434, 437,
		1, 0, 0, 0, 435, 437, 3, 66, 33, 0, 436, 420, 1, 0, 0, 0, 436, 431, 1,
		0, 0, 0, 436, 435, 1, 0, 0, 0, 437, 63, 1, 0, 0, 0, 438, 439, 7, 4, 0,
		0, 439, 65, 1, 0, 0, 0, 440, 459, 3, 80, 40, 0, 441, 459, 5, 233, 0, 0,
		442, 445, 5, 227, 0, 0, 443, 445, 3, 26, 13, 0, 444, 442, 1, 0, 0, 0, 444,
		443, 1, 0, 0, 0, 445, 453, 1, 0, 0, 0, 446, 449, 5, 7, 0, 0, 447, 450,
		5, 227, 0, 0, 448, 450, 3, 26, 13, 0, 449, 447, 1, 0, 0, 0, 449, 448, 1,
		0, 0, 0, 450, 452, 1, 0, 0, 0, 451, 446, 1, 0, 0, 0, 452, 455, 1, 0, 0,
		0, 453, 451, 1, 0, 0, 0, 453, 454, 1, 0, 0, 0, 454, 459, 1, 0, 0, 0, 455,
		453, 1, 0, 0, 0, 456, 459, 5, 254, 0, 0, 457, 459, 5, 256, 0, 0, 458, 440,
		1, 0, 0, 0, 458, 441, 1, 0, 0, 0, 458, 444, 1, 0, 0, 0, 458, 456, 1, 0,
		0, 0, 458, 457, 1, 0, 0, 0, 459, 67, 1, 0, 0, 0, 460, 461, 5, 10, 0, 0,
		461, 69, 1, 0, 0, 0, 462, 463, 5, 88, 0, 0, 463, 71, 1, 0, 0, 0, 464, 466,
		5, 1, 0, 0, 465, 464, 1, 0, 0, 0, 465, 466, 1, 0, 0, 0, 466, 468, 1, 0,
		0, 0, 467, 469, 3, 68, 34, 0, 468, 467, 1, 0, 0, 0, 468, 469, 1, 0, 0,
		0, 469, 471, 1, 0, 0, 0, 470, 472, 3, 70, 35, 0, 471, 470, 1, 0, 0, 0,
		471, 472, 1, 0, 0, 0, 472, 473, 1, 0, 0, 0, 473, 475, 3, 78, 39, 0, 474,
		476, 5, 1, 0, 0, 475, 474, 1, 0, 0, 0, 475, 476, 1, 0, 0, 0, 476, 490,
		1, 0, 0, 0, 477, 479, 5, 8, 0, 0, 478, 480, 5, 1, 0, 0, 479, 478, 1, 0,
		0, 0, 479, 480, 1, 0, 0, 0, 480, 482, 1, 0, 0, 0, 481, 483, 3, 68, 34,
		0, 482, 481, 1, 0, 0, 0, 482, 483, 1, 0, 0, 0, 483, 484, 1, 0, 0, 0, 484,
		486, 3, 78, 39, 0, 485, 487, 5, 1, 0, 0, 486, 485, 1, 0, 0, 0, 486, 487,
		1, 0, 0, 0, 487, 489, 1, 0, 0, 0, 488, 477, 1, 0, 0, 0, 489, 492, 1, 0,
		0, 0, 490, 488, 1, 0, 0, 0, 490, 491, 1, 0, 0, 0, 491, 73, 1, 0, 0, 0,
		492, 490, 1, 0, 0, 0, 493, 495, 5, 1, 0, 0, 494, 493, 1, 0, 0, 0, 494,
		495, 1, 0, 0, 0, 495, 497, 1, 0, 0, 0, 496, 498, 3, 68, 34, 0, 497, 496,
		1, 0, 0, 0, 497, 498, 1, 0, 0, 0, 498, 500, 1, 0, 0, 0, 499, 501, 3, 70,
		35, 0, 500, 499, 1, 0, 0, 0, 500, 501, 1, 0, 0, 0, 501, 502, 1, 0, 0, 0,
		502, 504, 3, 78, 39, 0, 503, 505, 5, 1, 0, 0, 504, 503, 1, 0, 0, 0, 504,
		505, 1, 0, 0, 0, 505, 519, 1, 0, 0, 0, 506, 508, 5, 7, 0, 0, 507, 509,
		5, 1, 0, 0, 508, 507, 1, 0, 0, 0, 508, 509, 1, 0, 0, 0, 509, 511, 1, 0,
		0, 0, 510, 512, 3, 68, 34, 0, 511, 510, 1, 0, 0, 0, 511, 512, 1, 0, 0,
		0, 512, 513, 1, 0, 0, 0, 513, 515, 3, 78, 39, 0, 514, 516, 5, 1, 0, 0,
		515, 514, 1, 0, 0, 0, 515, 516, 1, 0, 0, 0, 516, 518, 1, 0, 0, 0, 517,
		506, 1, 0, 0, 0, 518, 521, 1, 0, 0, 0, 519, 517, 1, 0, 0, 0, 519, 520,
		1, 0, 0, 0, 520, 75, 1, 0, 0, 0, 521, 519, 1, 0, 0, 0, 522, 523, 3, 78,
		39, 0, 523, 77, 1, 0, 0, 0, 524, 531, 3, 80, 40, 0, 525, 528, 3, 84, 42,
		0, 526, 527, 5, 4, 0, 0, 527, 529, 3, 122, 61, 0, 528, 526, 1, 0, 0, 0,
		528, 529, 1, 0, 0, 0, 529, 531, 1, 0, 0, 0, 530, 524, 1, 0, 0, 0, 530,
		525, 1, 0, 0, 0, 531, 79, 1, 0, 0, 0, 532, 533, 5, 86, 0, 0, 533, 81, 1,
		0, 0, 0, 534, 535, 5, 86, 0, 0, 535, 83, 1, 0, 0, 0, 536, 537, 7, 5, 0,
		0, 537, 85, 1, 0, 0, 0, 538, 539, 7, 5, 0, 0, 539, 87, 1, 0, 0, 0, 540,
		541, 5, 1, 0, 0, 541, 546, 3, 90, 45, 0, 542, 543, 5, 7, 0, 0, 543, 545,
		3, 90, 45, 0, 544, 542, 1, 0, 0, 0, 545, 548, 1, 0, 0, 0, 546, 544, 1,
		0, 0, 0, 546, 547, 1, 0, 0, 0, 547, 549, 1, 0, 0, 0, 548, 546, 1, 0, 0,
		0, 549, 550, 5, 1, 0, 0, 550, 89, 1, 0, 0, 0, 551, 552, 3, 100, 50, 0,
		552, 554, 5, 4, 0, 0, 553, 555, 5, 10, 0, 0, 554, 553, 1, 0, 0, 0, 554,
		555, 1, 0, 0, 0, 555, 557, 1, 0, 0, 0, 556, 558, 5, 3, 0, 0, 557, 556,
		1, 0, 0, 0, 557, 558, 1, 0, 0, 0, 558, 559, 1, 0, 0, 0, 559, 560, 3, 112,
		56, 0, 560, 570, 1, 0, 0, 0, 561, 562, 3, 100, 50, 0, 562, 563, 5, 4, 0,
		0, 563, 564, 3, 112, 56, 0, 564, 570, 1, 0, 0, 0, 565, 566, 5, 83, 0, 0,
		566, 567, 5, 4, 0, 0, 567, 570, 3, 120, 60, 0, 568, 570, 3, 92, 46, 0,
		569, 551, 1, 0, 0, 0, 569, 561, 1, 0, 0, 0, 569, 565, 1, 0, 0, 0, 569,
		568, 1, 0, 0, 0, 570, 91, 1, 0, 0, 0, 571, 575, 3, 94, 47, 0, 572, 575,
		3, 96, 48, 0, 573, 575, 3, 98, 49, 0, 574, 571, 1, 0, 0, 0, 574, 572, 1,
		0, 0, 0, 574, 573, 1, 0, 0, 0, 575, 93, 1, 0, 0, 0, 576, 577, 7, 6, 0,
		0, 577, 95, 1, 0, 0, 0, 578, 579, 7, 7, 0, 0, 579, 97, 1, 0, 0, 0, 580,
		581, 5, 32, 0, 0, 581, 99, 1, 0, 0, 0, 582, 588, 3, 102, 51, 0, 583, 588,
		3, 104, 52, 0, 584, 588, 3, 106, 53, 0, 585, 588, 3, 110, 55, 0, 586, 588,
		3, 108, 54, 0, 587, 582, 1, 0, 0, 0, 587, 583, 1, 0, 0, 0, 587, 584, 1,
		0, 0, 0, 587, 585, 1, 0, 0, 0, 587, 586, 1, 0, 0, 0, 588, 101, 1, 0, 0,
		0, 589, 598, 5, 60, 0, 0, 590, 598, 5, 49, 0, 0, 591, 598, 5, 53, 0, 0,
		592, 598, 5, 54, 0, 0, 593, 598, 5, 64, 0, 0, 594, 598, 5, 75, 0, 0, 595,
		598, 5, 80, 0, 0, 596, 598, 5, 81, 0, 0, 597, 589, 1, 0, 0, 0, 597, 590,
		1, 0, 0, 0, 597, 591, 1, 0, 0, 0, 597, 592, 1, 0, 0, 0, 597, 593, 1, 0,
		0, 0, 597, 594, 1, 0, 0, 0, 597, 595, 1, 0, 0, 0, 597, 596, 1, 0, 0, 0,
		598, 103, 1, 0, 0, 0, 599, 600, 7, 8, 0, 0, 600, 105, 1, 0, 0, 0, 601,
		602, 7, 9, 0, 0, 602, 107, 1, 0, 0, 0, 603, 604, 7, 10, 0, 0, 604, 109,
		1, 0, 0, 0, 605, 606, 7, 11, 0, 0, 606, 111, 1, 0, 0, 0, 607, 617, 3, 114,
		57, 0, 608, 609, 5, 2, 0, 0, 609, 610, 3, 114, 57, 0, 610, 611, 5, 2, 0,
		0, 611, 617, 1, 0, 0, 0, 612, 613, 5, 2, 0, 0, 613, 614, 3, 116, 58, 0,
		614, 615, 5, 2, 0, 0, 615, 617, 1, 0, 0, 0, 616, 607, 1, 0, 0, 0, 616,
		608, 1, 0, 0, 0, 616, 612, 1, 0, 0, 0, 617, 113, 1, 0, 0, 0, 618, 630,
		5, 227, 0, 0, 619, 630, 3, 122, 61, 0, 620, 630, 3, 126, 63, 0, 621, 622,
		3, 118, 59, 0, 622, 623, 3, 132, 66, 0, 623, 624, 3, 54, 27, 0, 624, 630,
		1, 0, 0, 0, 625, 630, 5, 225, 0, 0, 626, 630, 5, 76, 0, 0, 627, 630, 5,
		231, 0, 0, 628, 630, 5, 241, 0, 0, 629, 618, 1, 0, 0, 0, 629, 619, 1, 0,
		0, 0, 629, 620, 1, 0, 0, 0, 629, 621, 1, 0, 0, 0, 629, 625, 1, 0, 0, 0,
		629, 626, 1, 0, 0, 0, 629, 627, 1, 0, 0, 0, 629, 628, 1, 0, 0, 0, 630,
		115, 1, 0, 0, 0, 631, 632, 5, 246, 0, 0, 632, 117, 1, 0, 0, 0, 633, 634,
		7, 12, 0, 0, 634, 119, 1, 0, 0, 0, 635, 636, 5, 84, 0, 0, 636, 121, 1,
		0, 0, 0, 637, 641, 1, 0, 0, 0, 638, 641, 5, 243, 0, 0, 639, 641, 5, 250,
		0, 0, 640, 637, 1, 0, 0, 0, 640, 638, 1, 0, 0, 0, 640, 639, 1, 0, 0, 0,
		641, 123, 1, 0, 0, 0, 642, 646, 1, 0, 0, 0, 643, 646, 5, 243, 0, 0, 644,
		646, 5, 250, 0, 0, 645, 642, 1, 0, 0, 0, 645, 643, 1, 0, 0, 0, 645, 644,
		1, 0, 0, 0, 646, 125, 1, 0, 0, 0, 647, 648, 3, 128, 64, 0, 648, 649, 5,
		236, 0, 0, 649, 650, 3, 130, 65, 0, 650, 651, 3, 132, 66, 0, 651, 652,
		3, 134, 67, 0, 652, 127, 1, 0, 0, 0, 653, 654, 5, 235, 0, 0, 654, 129,
		1, 0, 0, 0, 655, 659, 5, 237, 0, 0, 656, 657, 5, 238, 0, 0, 657, 659, 5,
		234, 0, 0, 658, 655, 1, 0, 0, 0, 658, 656, 1, 0, 0, 0, 659, 660, 1, 0,
		0, 0, 660, 658, 1, 0, 0, 0, 660, 661, 1, 0, 0, 0, 661, 131, 1, 0, 0, 0,
		662, 663, 7, 13, 0, 0, 663, 133, 1, 0, 0, 0, 664, 665, 5, 239, 0, 0, 665,
		135, 1, 0, 0, 0, 666, 667, 7, 14, 0, 0, 667, 137, 1, 0, 0, 0, 73, 141,
		147, 153, 156, 161, 164, 171, 174, 180, 183, 189, 197, 204, 213, 222, 233,
		240, 246, 248, 253, 256, 260, 292, 298, 312, 324, 331, 345, 358, 385, 393,
		403, 409, 411, 416, 422, 427, 436, 444, 449, 453, 458, 465, 468, 471, 475,
		479, 482, 486, 490, 494, 497, 500, 504, 508, 511, 515, 519, 528, 530, 546,
		554, 557, 569, 574, 587, 597, 616, 629, 640, 645, 658, 660,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// SecLangParserInit initializes any static state used to implement SecLangParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSecLangParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SecLangParserInit() {
	staticData := &SecLangParserParserStaticData
	staticData.once.Do(seclangparserParserInit)
}

// NewSecLangParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSecLangParser(input antlr.TokenStream) *SecLangParser {
	SecLangParserInit()
	this := new(SecLangParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &SecLangParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "SecLangParser.g4"

	return this
}

// SecLangParser tokens.
const (
	SecLangParserEOF                                   = antlr.TokenEOF
	SecLangParserQUOTE                                 = 1
	SecLangParserSINGLE_QUOTE                          = 2
	SecLangParserEQUAL                                 = 3
	SecLangParserCOLON                                 = 4
	SecLangParserEQUALS_PLUS                           = 5
	SecLangParserEQUALS_MINUS                          = 6
	SecLangParserCOMMA                                 = 7
	SecLangParserPIPE                                  = 8
	SecLangParserCONFIG_VALUE_PATH                     = 9
	SecLangParserNOT                                   = 10
	SecLangParserWS                                    = 11
	SecLangParserHASH                                  = 12
	SecLangParserPLUS                                  = 13
	SecLangParserMINUS                                 = 14
	SecLangParserSTAR                                  = 15
	SecLangParserSLASH                                 = 16
	SecLangParserASSIGN                                = 17
	SecLangParserSEMI                                  = 18
	SecLangParserNOT_EQUAL                             = 19
	SecLangParserLT                                    = 20
	SecLangParserLE                                    = 21
	SecLangParserGE                                    = 22
	SecLangParserGT                                    = 23
	SecLangParserLPAREN                                = 24
	SecLangParserRPAREN                                = 25
	SecLangParserACTION_ACCURACY                       = 26
	SecLangParserACTION_ALLOW                          = 27
	SecLangParserACTION_APPEND                         = 28
	SecLangParserACTION_AUDIT_LOG                      = 29
	SecLangParserACTION_BLOCK                          = 30
	SecLangParserACTION_CAPTURE                        = 31
	SecLangParserACTION_CHAIN                          = 32
	SecLangParserACTION_CTL                            = 33
	SecLangParserACTION_CTL_AUDIT_ENGINE               = 34
	SecLangParserACTION_CTL_AUDIT_LOG_PARTS            = 35
	SecLangParserACTION_CTL_REQUEST_BODY_PROCESSOR     = 36
	SecLangParserACTION_CTL_FORCE_REQ_BODY_VAR         = 37
	SecLangParserACTION_CTL_REQUEST_BODY_ACCESS        = 38
	SecLangParserACTION_CTL_RULE_ENGINE                = 39
	SecLangParserACTION_CTL_RULE_REMOVE_BY_TAG         = 40
	SecLangParserACTION_CTL_RULE_REMOVE_BY_ID          = 41
	SecLangParserACTION_CTL_RULE_REMOVE_TARGET_BY_ID   = 42
	SecLangParserACTION_CTL_RULE_REMOVE_TARGET_BY_TAG  = 43
	SecLangParserACTION_DENY                           = 44
	SecLangParserACTION_DEPRECATE_VAR                  = 45
	SecLangParserACTION_DROP                           = 46
	SecLangParserACTION_EXEC                           = 47
	SecLangParserACTION_EXPIRE_VAR                     = 48
	SecLangParserACTION_ID                             = 49
	SecLangParserACTION_INITCOL                        = 50
	SecLangParserACTION_LOG_DATA                       = 51
	SecLangParserACTION_LOG                            = 52
	SecLangParserACTION_MATURITY                       = 53
	SecLangParserACTION_MSG                            = 54
	SecLangParserACTION_MULTI_MATCH                    = 55
	SecLangParserACTION_NO_AUDIT_LOG                   = 56
	SecLangParserACTION_NO_LOG                         = 57
	SecLangParserACTION_PASS                           = 58
	SecLangParserACTION_PAUSE                          = 59
	SecLangParserACTION_PHASE                          = 60
	SecLangParserACTION_PREPEND                        = 61
	SecLangParserACTION_PROXY                          = 62
	SecLangParserACTION_REDIRECT                       = 63
	SecLangParserACTION_REV                            = 64
	SecLangParserACTION_SANITISE_ARG                   = 65
	SecLangParserACTION_SANITISE_MATCHED_BYTES         = 66
	SecLangParserACTION_SANITISE_MATCHED               = 67
	SecLangParserACTION_SANITISE_REQUEST_HEADER        = 68
	SecLangParserACTION_SANITISE_RESPONSE_HEADER       = 69
	SecLangParserACTION_SETENV                         = 70
	SecLangParserACTION_SETRSC                         = 71
	SecLangParserACTION_SETSID                         = 72
	SecLangParserACTION_SETUID                         = 73
	SecLangParserACTION_SETVAR                         = 74
	SecLangParserACTION_SEVERITY                       = 75
	SecLangParserACTION_SEVERITY_VALUE                 = 76
	SecLangParserACTION_SKIP_AFTER                     = 77
	SecLangParserACTION_SKIP                           = 78
	SecLangParserACTION_STATUS                         = 79
	SecLangParserACTION_TAG                            = 80
	SecLangParserACTION_VER                            = 81
	SecLangParserACTION_XMLNS                          = 82
	SecLangParserACTION_TRANSFORMATION                 = 83
	SecLangParserTRANSFORMATION_VALUE                  = 84
	SecLangParserCOLLECTION_NAME_ENUM                  = 85
	SecLangParserVARIABLE_NAME_ENUM                    = 86
	SecLangParserRUN_TIME_VAR_XML                      = 87
	SecLangParserVAR_COUNT                             = 88
	SecLangParserOPERATOR_BEGINS_WITH                  = 89
	SecLangParserOPERATOR_CONTAINS                     = 90
	SecLangParserOPERATOR_CONTAINS_WORD                = 91
	SecLangParserOPERATOR_DETECT_SQLI                  = 92
	SecLangParserOPERATOR_DETECT_XSS                   = 93
	SecLangParserOPERATOR_ENDS_WITH                    = 94
	SecLangParserOPERATOR_EQ                           = 95
	SecLangParserOPERATOR_FUZZY_HASH                   = 96
	SecLangParserOPERATOR_GE                           = 97
	SecLangParserOPERATOR_GEOLOOKUP                    = 98
	SecLangParserOPERATOR_GSB_LOOKUP                   = 99
	SecLangParserOPERATOR_GT                           = 100
	SecLangParserOPERATOR_INSPECT_FILE                 = 101
	SecLangParserOPERATOR_IP_MATCH_FROM_FILE           = 102
	SecLangParserOPERATOR_IP_MATCH                     = 103
	SecLangParserOPERATOR_LE                           = 104
	SecLangParserOPERATOR_LT                           = 105
	SecLangParserOPERATOR_PM_FROM_FILE                 = 106
	SecLangParserOPERATOR_PM                           = 107
	SecLangParserOPERATOR_RBL                          = 108
	SecLangParserOPERATOR_RSUB                         = 109
	SecLangParserOPERATOR_RX                           = 110
	SecLangParserOPERATOR_RX_GLOBAL                    = 111
	SecLangParserOPERATOR_STR_EQ                       = 112
	SecLangParserOPERATOR_STR_MATCH                    = 113
	SecLangParserOPERATOR_UNCONDITIONAL_MATCH          = 114
	SecLangParserOPERATOR_VALIDATE_BYTE_RANGE          = 115
	SecLangParserOPERATOR_VALIDATE_DTD                 = 116
	SecLangParserOPERATOR_VALIDATE_HASH                = 117
	SecLangParserOPERATOR_VALIDATE_SCHEMA              = 118
	SecLangParserOPERATOR_VALIDATE_URL_ENCODING        = 119
	SecLangParserOPERATOR_VALIDATE_UTF8_ENCODING       = 120
	SecLangParserOPERATOR_VERIFY_CC                    = 121
	SecLangParserOPERATOR_VERIFY_CPF                   = 122
	SecLangParserOPERATOR_VERIFY_SSN                   = 123
	SecLangParserOPERATOR_VERIFY_SVNR                  = 124
	SecLangParserOPERATOR_WITHIN                       = 125
	SecLangParserAUDIT_PARTS                           = 126
	SecLangParserCONFIG_COMPONENT_SIG                  = 127
	SecLangParserCONFIG_SEC_SERVER_SIG                 = 128
	SecLangParserCONFIG_SEC_WEB_APP_ID                 = 129
	SecLangParserCONFIG_SEC_CACHE_TRANSFORMATIONS      = 130
	SecLangParserCONFIG_SEC_CHROOT_DIR                 = 131
	SecLangParserCONFIG_CONN_ENGINE                    = 132
	SecLangParserCONFIG_SEC_HASH_ENGINE                = 133
	SecLangParserCONFIG_SEC_HASH_KEY                   = 134
	SecLangParserCONFIG_SEC_HASH_PARAM                 = 135
	SecLangParserCONFIG_SEC_HASH_METHOD_RX             = 136
	SecLangParserCONFIG_SEC_HASH_METHOD_PM             = 137
	SecLangParserCONFIG_CONTENT_INJECTION              = 138
	SecLangParserCONFIG_SEC_ARGUMENT_SEPARATOR         = 139
	SecLangParserCONFIG_DIR_AUDIT_DIR                  = 140
	SecLangParserCONFIG_DIR_AUDIT_DIR_MOD              = 141
	SecLangParserCONFIG_DIR_AUDIT_ENG                  = 142
	SecLangParserCONFIG_DIR_AUDIT_FILE_MODE            = 143
	SecLangParserCONFIG_DIR_AUDIT_LOG2                 = 144
	SecLangParserCONFIG_DIR_AUDIT_LOG                  = 145
	SecLangParserCONFIG_DIR_AUDIT_LOG_FMT              = 146
	SecLangParserCONFIG_DIR_AUDIT_LOG_P                = 147
	SecLangParserCONFIG_DIR_AUDIT_STS                  = 148
	SecLangParserCONFIG_DIR_AUDIT_TYPE                 = 149
	SecLangParserCONFIG_DIR_DEBUG_LOG                  = 150
	SecLangParserCONFIG_DIR_DEBUG_LVL                  = 151
	SecLangParserCONFIG_DIR_GEO_DB                     = 152
	SecLangParserCONFIG_DIR_GSB_DB                     = 153
	SecLangParserCONFIG_SEC_GUARDIAN_LOG               = 154
	SecLangParserCONFIG_SEC_INTERCEPT_ON_ERROR         = 155
	SecLangParserCONFIG_SEC_CONN_R_STATE_LIMIT         = 156
	SecLangParserCONFIG_SEC_CONN_W_STATE_LIMIT         = 157
	SecLangParserCONFIG_SEC_SENSOR_ID                  = 158
	SecLangParserCONFIG_SEC_RULE_INHERITANCE           = 159
	SecLangParserCONFIG_SEC_RULE_PERF_TIME             = 160
	SecLangParserCONFIG_SEC_STREAM_IN_BODY_INSPECTION  = 161
	SecLangParserCONFIG_SEC_STREAM_OUT_BODY_INSPECTION = 162
	SecLangParserCONFIG_DIR_PCRE_MATCH_LIMIT           = 163
	SecLangParserCONFIG_DIR_PCRE_MATCH_LIMIT_RECURSION = 164
	SecLangParserCONFIG_DIR_ARGS_LIMIT                 = 165
	SecLangParserCONFIG_DIR_REQ_BODY_JSON_DEPTH_LIMIT  = 166
	SecLangParserCONFIG_DIR_REQ_BODY                   = 167
	SecLangParserCONFIG_DIR_REQ_BODY_IN_MEMORY_LIMIT   = 168
	SecLangParserCONFIG_DIR_REQ_BODY_LIMIT             = 169
	SecLangParserCONFIG_DIR_REQ_BODY_LIMIT_ACTION      = 170
	SecLangParserCONFIG_DIR_REQ_BODY_NO_FILES_LIMIT    = 171
	SecLangParserCONFIG_DIR_RES_BODY                   = 172
	SecLangParserCONFIG_DIR_RES_BODY_LIMIT             = 173
	SecLangParserCONFIG_DIR_RES_BODY_LIMIT_ACTION      = 174
	SecLangParserCONFIG_DIR_RULE_ENG                   = 175
	SecLangParserCONFIG_DIR_SEC_ACTION                 = 176
	SecLangParserCONFIG_DIR_SEC_DEFAULT_ACTION         = 177
	SecLangParserCONFIG_SEC_DISABLE_BACKEND_COMPRESS   = 178
	SecLangParserCONFIG_DIR_SEC_MARKER                 = 179
	SecLangParserCONFIG_DIR_UNICODE_MAP_FILE           = 180
	SecLangParserCONFIG_INCLUDE                        = 181
	SecLangParserCONFIG_SEC_COLLECTION_TIMEOUT         = 182
	SecLangParserCONFIG_SEC_HTTP_BLKEY                 = 183
	SecLangParserCONFIG_SEC_REMOTE_RULES               = 184
	SecLangParserCONFIG_SEC_REMOTE_RULES_FAIL_ACTION   = 185
	SecLangParserCONFIG_SEC_RULE_REMOVE_BY_ID          = 186
	SecLangParserCONFIG_SEC_RULE_REMOVE_BY_MSG         = 187
	SecLangParserCONFIG_SEC_RULE_REMOVE_BY_TAG         = 188
	SecLangParserCONFIG_SEC_RULE_UPDATE_TARGET_BY_TAG  = 189
	SecLangParserCONFIG_SEC_RULE_UPDATE_TARGET_BY_MSG  = 190
	SecLangParserCONFIG_SEC_RULE_UPDATE_TARGET_BY_ID   = 191
	SecLangParserCONFIG_SEC_RULE_UPDATE_ACTION_BY_ID   = 192
	SecLangParserCONFIG_UPLOAD_KEEP_FILES              = 193
	SecLangParserCONFIG_UPLOAD_SAVE_TMP_FILES          = 194
	SecLangParserCONFIG_UPLOAD_DIR                     = 195
	SecLangParserCONFIG_UPLOAD_FILE_LIMIT              = 196
	SecLangParserCONFIG_UPLOAD_FILE_MODE               = 197
	SecLangParserCONFIG_VALUE_ABORT                    = 198
	SecLangParserCONFIG_VALUE_DETC                     = 199
	SecLangParserCONFIG_VALUE_HTTPS                    = 200
	SecLangParserCONFIG_VALUE_OFF                      = 201
	SecLangParserCONFIG_VALUE_ON                       = 202
	SecLangParserCONFIG_VALUE_PARALLEL                 = 203
	SecLangParserCONFIG_VALUE_PROCESS_PARTIAL          = 204
	SecLangParserCONFIG_VALUE_REJECT                   = 205
	SecLangParserCONFIG_VALUE_RELEVANT_ONLY            = 206
	SecLangParserCONFIG_VALUE_SERIAL                   = 207
	SecLangParserCONFIG_VALUE_WARN                     = 208
	SecLangParserCONFIG_XML_EXTERNAL_ENTITY            = 209
	SecLangParserCONFIG_DIR_RESPONSE_BODY_MP           = 210
	SecLangParserCONFIG_DIR_RESPONSE_BODY_MP_CLEAR     = 211
	SecLangParserCONFIG_DIR_SEC_COOKIE_FORMAT          = 212
	SecLangParserCONFIG_SEC_COOKIEV0_SEPARATOR         = 213
	SecLangParserCONFIG_DIR_SEC_DATA_DIR               = 214
	SecLangParserCONFIG_DIR_SEC_STATUS_ENGINE          = 215
	SecLangParserCONFIG_DIR_SEC_TMP_DIR                = 216
	SecLangParserCONFIG_DIR_SEC_RULE                   = 217
	SecLangParserDIRECTIVE_SECRULESCRIPT               = 218
	SecLangParserOPTION_NAME                           = 219
	SecLangParserSINGLE_QUOTE_BUT_SCAPED               = 220
	SecLangParserDOUBLE_SINGLE_QUOTE_BUT_SCAPED        = 221
	SecLangParserCOMMA_BUT_SCAPED                      = 222
	SecLangParserNATIVE                                = 223
	SecLangParserNEWLINE                               = 224
	SecLangParserVARIABLE_NAME                         = 225
	SecLangParserIDENT                                 = 226
	SecLangParserINT                                   = 227
	SecLangParserDIGIT                                 = 228
	SecLangParserLETTER                                = 229
	SecLangParserDICT_ELEMENT_REGEXP                   = 230
	SecLangParserFREE_TEXT_QUOTE_MACRO_EXPANSION       = 231
	SecLangParserWS_STRING_MODE                        = 232
	SecLangParserSTRING                                = 233
	SecLangParserMACRO_EXPANSION                       = 234
	SecLangParserCOLLECTION_NAME_SETVAR                = 235
	SecLangParserDOT                                   = 236
	SecLangParserCOLLECTION_ELEMENT                    = 237
	SecLangParserCOLLECTION_WITH_MACRO                 = 238
	SecLangParserVAR_ASSIGNMENT                        = 239
	SecLangParserSPACE_SETVAR_ASSIGNMENT               = 240
	SecLangParserCOMMA_SEPARATED_STRING                = 241
	SecLangParserWS_FILE_PATH_MODE                     = 242
	SecLangParserXPATH_EXPRESSION                      = 243
	SecLangParserXPATH_MODE_POP_CHARS                  = 244
	SecLangParserACTION_CTL_BODY_PROCESSOR_TYPE        = 245
	SecLangParserSTRING_LITERAL                        = 246
	SecLangParserSPACE_COL                             = 247
	SecLangParserSPACE_VAR                             = 248
	SecLangParserNEWLINE_VAR                           = 249
	SecLangParserCOLLECTION_ELEMENT_VALUE              = 250
	SecLangParserSPACE_COL_ELEM                        = 251
	SecLangParserNEWLINE_COL_ELEM                      = 252
	SecLangParserSKIP_CHARS                            = 253
	SecLangParserOPERATOR_UNQUOTED_STRING              = 254
	SecLangParserAT                                    = 255
	SecLangParserOPERATOR_QUOTED_STRING                = 256
	SecLangParserCOMMENT                               = 257
	SecLangParserHASH_COMMENT_BLOCK                    = 258
	SecLangParserBLOCK_COMMENT_END                     = 259
	SecLangParserPIPE_DEFAULT                          = 260
	SecLangParserCOMMA_DEFAULT                         = 261
	SecLangParserCOLON_DEFAULT                         = 262
	SecLangParserEQUAL_DEFAULT                         = 263
	SecLangParserNOT_DEFAULT                           = 264
	SecLangParserQUOTE_DEFAULT                         = 265
	SecLangParserSINGLE_QUOTE_SETVAR                   = 266
)

// SecLangParser rules.
const (
	SecLangParserRULE_configuration                           = 0
	SecLangParserRULE_stmt                                    = 1
	SecLangParserRULE_comment_block                           = 2
	SecLangParserRULE_comment                                 = 3
	SecLangParserRULE_engine_config_rule_directive            = 4
	SecLangParserRULE_engine_config_directive                 = 5
	SecLangParserRULE_string_engine_config_directive          = 6
	SecLangParserRULE_sec_marker_directive                    = 7
	SecLangParserRULE_engine_config_directive_with_param      = 8
	SecLangParserRULE_rule_script_directive                   = 9
	SecLangParserRULE_file_path                               = 10
	SecLangParserRULE_remove_rule_by_id                       = 11
	SecLangParserRULE_remove_rule_by_id_values                = 12
	SecLangParserRULE_int_range                               = 13
	SecLangParserRULE_range_start                             = 14
	SecLangParserRULE_range_end                               = 15
	SecLangParserRULE_string_remove_rules                     = 16
	SecLangParserRULE_string_remove_rules_values              = 17
	SecLangParserRULE_update_target_rules                     = 18
	SecLangParserRULE_update_action_rule                      = 19
	SecLangParserRULE_id                                      = 20
	SecLangParserRULE_engine_config_sec_cache_transformations = 21
	SecLangParserRULE_option_list                             = 22
	SecLangParserRULE_option                                  = 23
	SecLangParserRULE_option_name                             = 24
	SecLangParserRULE_engine_config_action_directive          = 25
	SecLangParserRULE_stmt_audit_log                          = 26
	SecLangParserRULE_values                                  = 27
	SecLangParserRULE_action_ctl_target_value                 = 28
	SecLangParserRULE_update_target_rules_values              = 29
	SecLangParserRULE_operator_not                            = 30
	SecLangParserRULE_operator                                = 31
	SecLangParserRULE_operator_name                           = 32
	SecLangParserRULE_operator_value                          = 33
	SecLangParserRULE_var_not                                 = 34
	SecLangParserRULE_var_count                               = 35
	SecLangParserRULE_variables                               = 36
	SecLangParserRULE_update_variables                        = 37
	SecLangParserRULE_new_target                              = 38
	SecLangParserRULE_var_stmt                                = 39
	SecLangParserRULE_variable_enum                           = 40
	SecLangParserRULE_ctl_variable_enum                       = 41
	SecLangParserRULE_collection_enum                         = 42
	SecLangParserRULE_ctl_collection_enum                     = 43
	SecLangParserRULE_actions                                 = 44
	SecLangParserRULE_action                                  = 45
	SecLangParserRULE_action_only                             = 46
	SecLangParserRULE_disruptive_action_only                  = 47
	SecLangParserRULE_non_disruptive_action_only              = 48
	SecLangParserRULE_flow_action_only                        = 49
	SecLangParserRULE_action_with_params                      = 50
	SecLangParserRULE_metadata_action_with_params             = 51
	SecLangParserRULE_disruptive_action_with_params           = 52
	SecLangParserRULE_non_disruptive_action_with_params       = 53
	SecLangParserRULE_data_action_with_params                 = 54
	SecLangParserRULE_flow_action_with_params                 = 55
	SecLangParserRULE_action_value                            = 56
	SecLangParserRULE_action_value_types                      = 57
	SecLangParserRULE_string_literal                          = 58
	SecLangParserRULE_ctl_action                              = 59
	SecLangParserRULE_transformation_action_value             = 60
	SecLangParserRULE_collection_value                        = 61
	SecLangParserRULE_ctl_collection_value                    = 62
	SecLangParserRULE_setvar_action                           = 63
	SecLangParserRULE_col_name                                = 64
	SecLangParserRULE_setvar_stmt                             = 65
	SecLangParserRULE_assignment                              = 66
	SecLangParserRULE_var_assignment                          = 67
	SecLangParserRULE_ctl_id                                  = 68
)

// IConfigurationContext is an interface to support dynamic dispatch.
type IConfigurationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

	// IsConfigurationContext differentiates from other interfaces.
	IsConfigurationContext()
}

type ConfigurationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConfigurationContext() *ConfigurationContext {
	var p = new(ConfigurationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_configuration
	return p
}

func InitEmptyConfigurationContext(p *ConfigurationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_configuration
}

func (*ConfigurationContext) IsConfigurationContext() {}

func NewConfigurationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConfigurationContext {
	var p = new(ConfigurationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_configuration

	return p
}

func (s *ConfigurationContext) GetParser() antlr.Parser { return s.parser }

func (s *ConfigurationContext) EOF() antlr.TerminalNode {
	return s.GetToken(SecLangParserEOF, 0)
}

func (s *ConfigurationContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ConfigurationContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *ConfigurationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConfigurationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConfigurationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterConfiguration(s)
	}
}

func (s *ConfigurationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitConfiguration(s)
	}
}

func (p *SecLangParser) Configuration() (localctx IConfigurationContext) {
	localctx = NewConfigurationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SecLangParserRULE_configuration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SecLangParserQUOTE || _la == SecLangParserHASH || ((int64((_la-127)) & ^0x3f) == 0 && ((int64(1)<<(_la-127))&-162131785608593409) != 0) || ((int64((_la-191)) & ^0x3f) == 0 && ((int64(1)<<(_la-191))&68987650175) != 0) {
		{
			p.SetState(138)
			p.Stmt()
		}

		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(144)
		p.Match(SecLangParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Engine_config_rule_directive() IEngine_config_rule_directiveContext
	Variables() IVariablesContext
	Operator() IOperatorContext
	AllComment_block() []IComment_blockContext
	Comment_block(i int) IComment_blockContext
	Actions() IActionsContext
	Rule_script_directive() IRule_script_directiveContext
	File_path() IFile_pathContext
	AllQUOTE() []antlr.TerminalNode
	QUOTE(i int) antlr.TerminalNode
	Remove_rule_by_id() IRemove_rule_by_idContext
	AllRemove_rule_by_id_values() []IRemove_rule_by_id_valuesContext
	Remove_rule_by_id_values(i int) IRemove_rule_by_id_valuesContext
	String_remove_rules() IString_remove_rulesContext
	String_remove_rules_values() IString_remove_rules_valuesContext
	Update_target_rules() IUpdate_target_rulesContext
	Update_target_rules_values() IUpdate_target_rules_valuesContext
	Update_variables() IUpdate_variablesContext
	PIPE() antlr.TerminalNode
	New_target() INew_targetContext
	Update_action_rule() IUpdate_action_ruleContext
	Id() IIdContext
	Engine_config_directive() IEngine_config_directiveContext

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_stmt
	return p
}

func InitEmptyStmtContext(p *StmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_stmt
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) Engine_config_rule_directive() IEngine_config_rule_directiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEngine_config_rule_directiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEngine_config_rule_directiveContext)
}

func (s *StmtContext) Variables() IVariablesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariablesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariablesContext)
}

func (s *StmtContext) Operator() IOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperatorContext)
}

func (s *StmtContext) AllComment_block() []IComment_blockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IComment_blockContext); ok {
			len++
		}
	}

	tst := make([]IComment_blockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IComment_blockContext); ok {
			tst[i] = t.(IComment_blockContext)
			i++
		}
	}

	return tst
}

func (s *StmtContext) Comment_block(i int) IComment_blockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComment_blockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComment_blockContext)
}

func (s *StmtContext) Actions() IActionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionsContext)
}

func (s *StmtContext) Rule_script_directive() IRule_script_directiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRule_script_directiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRule_script_directiveContext)
}

func (s *StmtContext) File_path() IFile_pathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFile_pathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFile_pathContext)
}

func (s *StmtContext) AllQUOTE() []antlr.TerminalNode {
	return s.GetTokens(SecLangParserQUOTE)
}

func (s *StmtContext) QUOTE(i int) antlr.TerminalNode {
	return s.GetToken(SecLangParserQUOTE, i)
}

func (s *StmtContext) Remove_rule_by_id() IRemove_rule_by_idContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRemove_rule_by_idContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRemove_rule_by_idContext)
}

func (s *StmtContext) AllRemove_rule_by_id_values() []IRemove_rule_by_id_valuesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRemove_rule_by_id_valuesContext); ok {
			len++
		}
	}

	tst := make([]IRemove_rule_by_id_valuesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRemove_rule_by_id_valuesContext); ok {
			tst[i] = t.(IRemove_rule_by_id_valuesContext)
			i++
		}
	}

	return tst
}

func (s *StmtContext) Remove_rule_by_id_values(i int) IRemove_rule_by_id_valuesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRemove_rule_by_id_valuesContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRemove_rule_by_id_valuesContext)
}

func (s *StmtContext) String_remove_rules() IString_remove_rulesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IString_remove_rulesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IString_remove_rulesContext)
}

func (s *StmtContext) String_remove_rules_values() IString_remove_rules_valuesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IString_remove_rules_valuesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IString_remove_rules_valuesContext)
}

func (s *StmtContext) Update_target_rules() IUpdate_target_rulesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUpdate_target_rulesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUpdate_target_rulesContext)
}

func (s *StmtContext) Update_target_rules_values() IUpdate_target_rules_valuesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUpdate_target_rules_valuesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUpdate_target_rules_valuesContext)
}

func (s *StmtContext) Update_variables() IUpdate_variablesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUpdate_variablesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUpdate_variablesContext)
}

func (s *StmtContext) PIPE() antlr.TerminalNode {
	return s.GetToken(SecLangParserPIPE, 0)
}

func (s *StmtContext) New_target() INew_targetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INew_targetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INew_targetContext)
}

func (s *StmtContext) Update_action_rule() IUpdate_action_ruleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUpdate_action_ruleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUpdate_action_ruleContext)
}

func (s *StmtContext) Id() IIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *StmtContext) Engine_config_directive() IEngine_config_directiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEngine_config_directiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEngine_config_directiveContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitStmt(s)
	}
}

func (p *SecLangParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SecLangParserRULE_stmt)
	var _la int

	var _alt int

	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SecLangParserHASH {
			{
				p.SetState(146)
				p.Comment_block()
			}

		}
		{
			p.SetState(149)
			p.Engine_config_rule_directive()
		}
		{
			p.SetState(150)
			p.Variables()
		}
		{
			p.SetState(151)
			p.Operator()
		}
		p.SetState(153)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(152)
				p.Actions()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SecLangParserHASH {
			{
				p.SetState(155)
				p.Comment_block()
			}

		}
		{
			p.SetState(158)
			p.Rule_script_directive()
		}
		{
			p.SetState(159)
			p.File_path()
		}
		p.SetState(161)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(160)
				p.Actions()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SecLangParserHASH {
			{
				p.SetState(163)
				p.Comment_block()
			}

		}
		{
			p.SetState(166)
			p.Rule_script_directive()
		}
		{
			p.SetState(167)
			p.Match(SecLangParserQUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(168)
			p.File_path()
		}
		{
			p.SetState(169)
			p.Match(SecLangParserQUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(171)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(170)
				p.Actions()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SecLangParserHASH {
			{
				p.SetState(173)
				p.Comment_block()
			}

		}
		{
			p.SetState(176)
			p.Remove_rule_by_id()
		}
		p.SetState(178)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(177)
					p.Remove_rule_by_id_values()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(180)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		p.SetState(183)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SecLangParserHASH {
			{
				p.SetState(182)
				p.Comment_block()
			}

		}
		{
			p.SetState(185)
			p.String_remove_rules()
		}
		{
			p.SetState(186)
			p.String_remove_rules_values()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		p.SetState(189)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SecLangParserHASH {
			{
				p.SetState(188)
				p.Comment_block()
			}

		}
		{
			p.SetState(191)
			p.String_remove_rules()
		}
		{
			p.SetState(192)
			p.Match(SecLangParserQUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(193)
			p.String_remove_rules_values()
		}
		{
			p.SetState(194)
			p.Match(SecLangParserQUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		p.SetState(197)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SecLangParserHASH {
			{
				p.SetState(196)
				p.Comment_block()
			}

		}
		{
			p.SetState(199)
			p.Update_target_rules()
		}
		{
			p.SetState(200)
			p.Update_target_rules_values()
		}
		{
			p.SetState(201)
			p.Update_variables()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		p.SetState(204)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SecLangParserHASH {
			{
				p.SetState(203)
				p.Comment_block()
			}

		}
		{
			p.SetState(206)
			p.Update_target_rules()
		}
		{
			p.SetState(207)
			p.Match(SecLangParserQUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(208)
			p.Update_target_rules_values()
		}
		{
			p.SetState(209)
			p.Match(SecLangParserQUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(210)
			p.Update_variables()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		p.SetState(213)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SecLangParserHASH {
			{
				p.SetState(212)
				p.Comment_block()
			}

		}
		{
			p.SetState(215)
			p.Update_target_rules()
		}
		{
			p.SetState(216)
			p.Update_target_rules_values()
		}
		{
			p.SetState(217)
			p.Update_variables()
		}
		{
			p.SetState(218)
			p.Match(SecLangParserPIPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(219)
			p.New_target()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		p.SetState(222)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SecLangParserHASH {
			{
				p.SetState(221)
				p.Comment_block()
			}

		}
		{
			p.SetState(224)
			p.Update_target_rules()
		}
		{
			p.SetState(225)
			p.Match(SecLangParserQUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(226)
			p.Update_target_rules_values()
		}
		{
			p.SetState(227)
			p.Match(SecLangParserQUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(228)
			p.Update_variables()
		}
		{
			p.SetState(229)
			p.Match(SecLangParserPIPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(230)
			p.New_target()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		p.SetState(233)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SecLangParserHASH {
			{
				p.SetState(232)
				p.Comment_block()
			}

		}
		{
			p.SetState(235)
			p.Update_action_rule()
		}
		{
			p.SetState(236)
			p.Id()
		}
		{
			p.SetState(237)
			p.Actions()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		p.SetState(240)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SecLangParserHASH {
			{
				p.SetState(239)
				p.Comment_block()
			}

		}
		{
			p.SetState(242)
			p.Engine_config_directive()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		p.SetState(244)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(243)
					p.Comment_block()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(246)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComment_blockContext is an interface to support dynamic dispatch.
type IComment_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllComment() []ICommentContext
	Comment(i int) ICommentContext
	BLOCK_COMMENT_END() antlr.TerminalNode
	EOF() antlr.TerminalNode

	// IsComment_blockContext differentiates from other interfaces.
	IsComment_blockContext()
}

type Comment_blockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComment_blockContext() *Comment_blockContext {
	var p = new(Comment_blockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_comment_block
	return p
}

func InitEmptyComment_blockContext(p *Comment_blockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_comment_block
}

func (*Comment_blockContext) IsComment_blockContext() {}

func NewComment_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Comment_blockContext {
	var p = new(Comment_blockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_comment_block

	return p
}

func (s *Comment_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *Comment_blockContext) AllComment() []ICommentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICommentContext); ok {
			len++
		}
	}

	tst := make([]ICommentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICommentContext); ok {
			tst[i] = t.(ICommentContext)
			i++
		}
	}

	return tst
}

func (s *Comment_blockContext) Comment(i int) ICommentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *Comment_blockContext) BLOCK_COMMENT_END() antlr.TerminalNode {
	return s.GetToken(SecLangParserBLOCK_COMMENT_END, 0)
}

func (s *Comment_blockContext) EOF() antlr.TerminalNode {
	return s.GetToken(SecLangParserEOF, 0)
}

func (s *Comment_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Comment_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Comment_blockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterComment_block(s)
	}
}

func (s *Comment_blockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitComment_block(s)
	}
}

func (p *SecLangParser) Comment_block() (localctx IComment_blockContext) {
	localctx = NewComment_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SecLangParserRULE_comment_block)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(250)
				p.Comment()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(253)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(256)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(255)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SecLangParserEOF || _la == SecLangParserBLOCK_COMMENT_END) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICommentContext is an interface to support dynamic dispatch.
type ICommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	HASH() antlr.TerminalNode
	COMMENT() antlr.TerminalNode

	// IsCommentContext differentiates from other interfaces.
	IsCommentContext()
}

type CommentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentContext() *CommentContext {
	var p = new(CommentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_comment
	return p
}

func InitEmptyCommentContext(p *CommentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_comment
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) HASH() antlr.TerminalNode {
	return s.GetToken(SecLangParserHASH, 0)
}

func (s *CommentContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(SecLangParserCOMMENT, 0)
}

func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterComment(s)
	}
}

func (s *CommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitComment(s)
	}
}

func (p *SecLangParser) Comment() (localctx ICommentContext) {
	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SecLangParserRULE_comment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(258)
		p.Match(SecLangParserHASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SecLangParserCOMMENT {
		{
			p.SetState(259)
			p.Match(SecLangParserCOMMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEngine_config_rule_directiveContext is an interface to support dynamic dispatch.
type IEngine_config_rule_directiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONFIG_DIR_SEC_RULE() antlr.TerminalNode

	// IsEngine_config_rule_directiveContext differentiates from other interfaces.
	IsEngine_config_rule_directiveContext()
}

type Engine_config_rule_directiveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEngine_config_rule_directiveContext() *Engine_config_rule_directiveContext {
	var p = new(Engine_config_rule_directiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_engine_config_rule_directive
	return p
}

func InitEmptyEngine_config_rule_directiveContext(p *Engine_config_rule_directiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_engine_config_rule_directive
}

func (*Engine_config_rule_directiveContext) IsEngine_config_rule_directiveContext() {}

func NewEngine_config_rule_directiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Engine_config_rule_directiveContext {
	var p = new(Engine_config_rule_directiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_engine_config_rule_directive

	return p
}

func (s *Engine_config_rule_directiveContext) GetParser() antlr.Parser { return s.parser }

func (s *Engine_config_rule_directiveContext) CONFIG_DIR_SEC_RULE() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_DIR_SEC_RULE, 0)
}

func (s *Engine_config_rule_directiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Engine_config_rule_directiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Engine_config_rule_directiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterEngine_config_rule_directive(s)
	}
}

func (s *Engine_config_rule_directiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitEngine_config_rule_directive(s)
	}
}

func (p *SecLangParser) Engine_config_rule_directive() (localctx IEngine_config_rule_directiveContext) {
	localctx = NewEngine_config_rule_directiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SecLangParserRULE_engine_config_rule_directive)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(262)
		p.Match(SecLangParserCONFIG_DIR_SEC_RULE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEngine_config_directiveContext is an interface to support dynamic dispatch.
type IEngine_config_directiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Stmt_audit_log() IStmt_audit_logContext
	Values() IValuesContext
	AllQUOTE() []antlr.TerminalNode
	QUOTE(i int) antlr.TerminalNode
	Engine_config_action_directive() IEngine_config_action_directiveContext
	Actions() IActionsContext
	String_engine_config_directive() IString_engine_config_directiveContext
	Sec_marker_directive() ISec_marker_directiveContext
	Engine_config_directive_with_param() IEngine_config_directive_with_paramContext
	Engine_config_sec_cache_transformations() IEngine_config_sec_cache_transformationsContext
	Option_list() IOption_listContext

	// IsEngine_config_directiveContext differentiates from other interfaces.
	IsEngine_config_directiveContext()
}

type Engine_config_directiveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEngine_config_directiveContext() *Engine_config_directiveContext {
	var p = new(Engine_config_directiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_engine_config_directive
	return p
}

func InitEmptyEngine_config_directiveContext(p *Engine_config_directiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_engine_config_directive
}

func (*Engine_config_directiveContext) IsEngine_config_directiveContext() {}

func NewEngine_config_directiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Engine_config_directiveContext {
	var p = new(Engine_config_directiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_engine_config_directive

	return p
}

func (s *Engine_config_directiveContext) GetParser() antlr.Parser { return s.parser }

func (s *Engine_config_directiveContext) Stmt_audit_log() IStmt_audit_logContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmt_audit_logContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmt_audit_logContext)
}

func (s *Engine_config_directiveContext) Values() IValuesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValuesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValuesContext)
}

func (s *Engine_config_directiveContext) AllQUOTE() []antlr.TerminalNode {
	return s.GetTokens(SecLangParserQUOTE)
}

func (s *Engine_config_directiveContext) QUOTE(i int) antlr.TerminalNode {
	return s.GetToken(SecLangParserQUOTE, i)
}

func (s *Engine_config_directiveContext) Engine_config_action_directive() IEngine_config_action_directiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEngine_config_action_directiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEngine_config_action_directiveContext)
}

func (s *Engine_config_directiveContext) Actions() IActionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionsContext)
}

func (s *Engine_config_directiveContext) String_engine_config_directive() IString_engine_config_directiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IString_engine_config_directiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IString_engine_config_directiveContext)
}

func (s *Engine_config_directiveContext) Sec_marker_directive() ISec_marker_directiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISec_marker_directiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISec_marker_directiveContext)
}

func (s *Engine_config_directiveContext) Engine_config_directive_with_param() IEngine_config_directive_with_paramContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEngine_config_directive_with_paramContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEngine_config_directive_with_paramContext)
}

func (s *Engine_config_directiveContext) Engine_config_sec_cache_transformations() IEngine_config_sec_cache_transformationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEngine_config_sec_cache_transformationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEngine_config_sec_cache_transformationsContext)
}

func (s *Engine_config_directiveContext) Option_list() IOption_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOption_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOption_listContext)
}

func (s *Engine_config_directiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Engine_config_directiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Engine_config_directiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterEngine_config_directive(s)
	}
}

func (s *Engine_config_directiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitEngine_config_directive(s)
	}
}

func (p *SecLangParser) Engine_config_directive() (localctx IEngine_config_directiveContext) {
	localctx = NewEngine_config_directiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SecLangParserRULE_engine_config_directive)
	p.SetState(292)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(264)
			p.Stmt_audit_log()
		}
		{
			p.SetState(265)
			p.Values()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(267)
			p.Stmt_audit_log()
		}
		{
			p.SetState(268)
			p.Match(SecLangParserQUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(269)
			p.Values()
		}
		{
			p.SetState(270)
			p.Match(SecLangParserQUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(272)
			p.Engine_config_action_directive()
		}
		{
			p.SetState(273)
			p.Actions()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(275)
			p.String_engine_config_directive()
		}
		{
			p.SetState(276)
			p.Match(SecLangParserQUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(277)
			p.Values()
		}
		{
			p.SetState(278)
			p.Match(SecLangParserQUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(280)
			p.Sec_marker_directive()
		}
		{
			p.SetState(281)
			p.Match(SecLangParserQUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(282)
			p.Values()
		}
		{
			p.SetState(283)
			p.Match(SecLangParserQUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(285)
			p.Engine_config_directive_with_param()
		}
		{
			p.SetState(286)
			p.Values()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(288)
			p.Engine_config_sec_cache_transformations()
		}
		{
			p.SetState(289)
			p.Values()
		}
		{
			p.SetState(290)
			p.Option_list()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IString_engine_config_directiveContext is an interface to support dynamic dispatch.
type IString_engine_config_directiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONFIG_COMPONENT_SIG() antlr.TerminalNode
	CONFIG_SEC_SERVER_SIG() antlr.TerminalNode
	CONFIG_SEC_WEB_APP_ID() antlr.TerminalNode

	// IsString_engine_config_directiveContext differentiates from other interfaces.
	IsString_engine_config_directiveContext()
}

type String_engine_config_directiveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyString_engine_config_directiveContext() *String_engine_config_directiveContext {
	var p = new(String_engine_config_directiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_string_engine_config_directive
	return p
}

func InitEmptyString_engine_config_directiveContext(p *String_engine_config_directiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_string_engine_config_directive
}

func (*String_engine_config_directiveContext) IsString_engine_config_directiveContext() {}

func NewString_engine_config_directiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_engine_config_directiveContext {
	var p = new(String_engine_config_directiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_string_engine_config_directive

	return p
}

func (s *String_engine_config_directiveContext) GetParser() antlr.Parser { return s.parser }

func (s *String_engine_config_directiveContext) CONFIG_COMPONENT_SIG() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_COMPONENT_SIG, 0)
}

func (s *String_engine_config_directiveContext) CONFIG_SEC_SERVER_SIG() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_SEC_SERVER_SIG, 0)
}

func (s *String_engine_config_directiveContext) CONFIG_SEC_WEB_APP_ID() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_SEC_WEB_APP_ID, 0)
}

func (s *String_engine_config_directiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_engine_config_directiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_engine_config_directiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterString_engine_config_directive(s)
	}
}

func (s *String_engine_config_directiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitString_engine_config_directive(s)
	}
}

func (p *SecLangParser) String_engine_config_directive() (localctx IString_engine_config_directiveContext) {
	localctx = NewString_engine_config_directiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SecLangParserRULE_string_engine_config_directive)
	p.SetState(298)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SecLangParserQUOTE:
		p.EnterOuterAlt(localctx, 1)

	case SecLangParserCONFIG_COMPONENT_SIG:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(295)
			p.Match(SecLangParserCONFIG_COMPONENT_SIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SecLangParserCONFIG_SEC_SERVER_SIG:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(296)
			p.Match(SecLangParserCONFIG_SEC_SERVER_SIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SecLangParserCONFIG_SEC_WEB_APP_ID:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(297)
			p.Match(SecLangParserCONFIG_SEC_WEB_APP_ID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISec_marker_directiveContext is an interface to support dynamic dispatch.
type ISec_marker_directiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONFIG_DIR_SEC_MARKER() antlr.TerminalNode

	// IsSec_marker_directiveContext differentiates from other interfaces.
	IsSec_marker_directiveContext()
}

type Sec_marker_directiveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySec_marker_directiveContext() *Sec_marker_directiveContext {
	var p = new(Sec_marker_directiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_sec_marker_directive
	return p
}

func InitEmptySec_marker_directiveContext(p *Sec_marker_directiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_sec_marker_directive
}

func (*Sec_marker_directiveContext) IsSec_marker_directiveContext() {}

func NewSec_marker_directiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sec_marker_directiveContext {
	var p = new(Sec_marker_directiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_sec_marker_directive

	return p
}

func (s *Sec_marker_directiveContext) GetParser() antlr.Parser { return s.parser }

func (s *Sec_marker_directiveContext) CONFIG_DIR_SEC_MARKER() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_DIR_SEC_MARKER, 0)
}

func (s *Sec_marker_directiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sec_marker_directiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sec_marker_directiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterSec_marker_directive(s)
	}
}

func (s *Sec_marker_directiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitSec_marker_directive(s)
	}
}

func (p *SecLangParser) Sec_marker_directive() (localctx ISec_marker_directiveContext) {
	localctx = NewSec_marker_directiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SecLangParserRULE_sec_marker_directive)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(300)
		p.Match(SecLangParserCONFIG_DIR_SEC_MARKER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEngine_config_directive_with_paramContext is an interface to support dynamic dispatch.
type IEngine_config_directive_with_paramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONFIG_CONN_ENGINE() antlr.TerminalNode
	CONFIG_CONTENT_INJECTION() antlr.TerminalNode
	CONFIG_DIR_ARGS_LIMIT() antlr.TerminalNode
	CONFIG_DIR_DEBUG_LOG() antlr.TerminalNode
	CONFIG_DIR_DEBUG_LVL() antlr.TerminalNode
	CONFIG_DIR_GEO_DB() antlr.TerminalNode
	CONFIG_DIR_GSB_DB() antlr.TerminalNode
	CONFIG_DIR_PCRE_MATCH_LIMIT() antlr.TerminalNode
	CONFIG_DIR_PCRE_MATCH_LIMIT_RECURSION() antlr.TerminalNode
	CONFIG_DIR_REQ_BODY() antlr.TerminalNode
	CONFIG_DIR_REQ_BODY_JSON_DEPTH_LIMIT() antlr.TerminalNode
	CONFIG_DIR_REQ_BODY_LIMIT() antlr.TerminalNode
	CONFIG_DIR_REQ_BODY_LIMIT_ACTION() antlr.TerminalNode
	CONFIG_DIR_REQ_BODY_NO_FILES_LIMIT() antlr.TerminalNode
	CONFIG_DIR_RESPONSE_BODY_MP() antlr.TerminalNode
	CONFIG_DIR_RESPONSE_BODY_MP_CLEAR() antlr.TerminalNode
	CONFIG_DIR_RES_BODY() antlr.TerminalNode
	CONFIG_DIR_RES_BODY_LIMIT() antlr.TerminalNode
	CONFIG_DIR_RES_BODY_LIMIT_ACTION() antlr.TerminalNode
	CONFIG_DIR_RULE_ENG() antlr.TerminalNode
	CONFIG_DIR_SEC_COOKIE_FORMAT() antlr.TerminalNode
	CONFIG_DIR_SEC_DATA_DIR() antlr.TerminalNode
	CONFIG_DIR_SEC_STATUS_ENGINE() antlr.TerminalNode
	CONFIG_DIR_SEC_TMP_DIR() antlr.TerminalNode
	CONFIG_DIR_UNICODE_MAP_FILE() antlr.TerminalNode
	CONFIG_SEC_ARGUMENT_SEPARATOR() antlr.TerminalNode
	CONFIG_SEC_CHROOT_DIR() antlr.TerminalNode
	CONFIG_SEC_COLLECTION_TIMEOUT() antlr.TerminalNode
	CONFIG_SEC_CONN_R_STATE_LIMIT() antlr.TerminalNode
	CONFIG_SEC_CONN_W_STATE_LIMIT() antlr.TerminalNode
	CONFIG_SEC_COOKIEV0_SEPARATOR() antlr.TerminalNode
	CONFIG_SEC_DISABLE_BACKEND_COMPRESS() antlr.TerminalNode
	CONFIG_SEC_GUARDIAN_LOG() antlr.TerminalNode
	CONFIG_SEC_HASH_ENGINE() antlr.TerminalNode
	CONFIG_SEC_HASH_KEY() antlr.TerminalNode
	CONFIG_SEC_HASH_METHOD_PM() antlr.TerminalNode
	CONFIG_SEC_HASH_METHOD_RX() antlr.TerminalNode
	CONFIG_SEC_HASH_PARAM() antlr.TerminalNode
	CONFIG_SEC_HTTP_BLKEY() antlr.TerminalNode
	CONFIG_SEC_INTERCEPT_ON_ERROR() antlr.TerminalNode
	CONFIG_SEC_REMOTE_RULES_FAIL_ACTION() antlr.TerminalNode
	CONFIG_SEC_RULE_INHERITANCE() antlr.TerminalNode
	CONFIG_SEC_RULE_PERF_TIME() antlr.TerminalNode
	CONFIG_SEC_SENSOR_ID() antlr.TerminalNode
	CONFIG_SEC_STREAM_IN_BODY_INSPECTION() antlr.TerminalNode
	CONFIG_SEC_STREAM_OUT_BODY_INSPECTION() antlr.TerminalNode
	CONFIG_XML_EXTERNAL_ENTITY() antlr.TerminalNode

	// IsEngine_config_directive_with_paramContext differentiates from other interfaces.
	IsEngine_config_directive_with_paramContext()
}

type Engine_config_directive_with_paramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEngine_config_directive_with_paramContext() *Engine_config_directive_with_paramContext {
	var p = new(Engine_config_directive_with_paramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_engine_config_directive_with_param
	return p
}

func InitEmptyEngine_config_directive_with_paramContext(p *Engine_config_directive_with_paramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_engine_config_directive_with_param
}

func (*Engine_config_directive_with_paramContext) IsEngine_config_directive_with_paramContext() {}

func NewEngine_config_directive_with_paramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Engine_config_directive_with_paramContext {
	var p = new(Engine_config_directive_with_paramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_engine_config_directive_with_param

	return p
}

func (s *Engine_config_directive_with_paramContext) GetParser() antlr.Parser { return s.parser }

func (s *Engine_config_directive_with_paramContext) CONFIG_CONN_ENGINE() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_CONN_ENGINE, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_CONTENT_INJECTION() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_CONTENT_INJECTION, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_DIR_ARGS_LIMIT() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_DIR_ARGS_LIMIT, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_DIR_DEBUG_LOG() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_DIR_DEBUG_LOG, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_DIR_DEBUG_LVL() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_DIR_DEBUG_LVL, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_DIR_GEO_DB() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_DIR_GEO_DB, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_DIR_GSB_DB() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_DIR_GSB_DB, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_DIR_PCRE_MATCH_LIMIT() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_DIR_PCRE_MATCH_LIMIT, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_DIR_PCRE_MATCH_LIMIT_RECURSION() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_DIR_PCRE_MATCH_LIMIT_RECURSION, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_DIR_REQ_BODY() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_DIR_REQ_BODY, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_DIR_REQ_BODY_JSON_DEPTH_LIMIT() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_DIR_REQ_BODY_JSON_DEPTH_LIMIT, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_DIR_REQ_BODY_LIMIT() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_DIR_REQ_BODY_LIMIT, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_DIR_REQ_BODY_LIMIT_ACTION() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_DIR_REQ_BODY_LIMIT_ACTION, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_DIR_REQ_BODY_NO_FILES_LIMIT() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_DIR_REQ_BODY_NO_FILES_LIMIT, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_DIR_RESPONSE_BODY_MP() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_DIR_RESPONSE_BODY_MP, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_DIR_RESPONSE_BODY_MP_CLEAR() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_DIR_RESPONSE_BODY_MP_CLEAR, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_DIR_RES_BODY() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_DIR_RES_BODY, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_DIR_RES_BODY_LIMIT() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_DIR_RES_BODY_LIMIT, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_DIR_RES_BODY_LIMIT_ACTION() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_DIR_RES_BODY_LIMIT_ACTION, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_DIR_RULE_ENG() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_DIR_RULE_ENG, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_DIR_SEC_COOKIE_FORMAT() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_DIR_SEC_COOKIE_FORMAT, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_DIR_SEC_DATA_DIR() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_DIR_SEC_DATA_DIR, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_DIR_SEC_STATUS_ENGINE() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_DIR_SEC_STATUS_ENGINE, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_DIR_SEC_TMP_DIR() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_DIR_SEC_TMP_DIR, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_DIR_UNICODE_MAP_FILE() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_DIR_UNICODE_MAP_FILE, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_SEC_ARGUMENT_SEPARATOR() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_SEC_ARGUMENT_SEPARATOR, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_SEC_CHROOT_DIR() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_SEC_CHROOT_DIR, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_SEC_COLLECTION_TIMEOUT() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_SEC_COLLECTION_TIMEOUT, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_SEC_CONN_R_STATE_LIMIT() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_SEC_CONN_R_STATE_LIMIT, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_SEC_CONN_W_STATE_LIMIT() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_SEC_CONN_W_STATE_LIMIT, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_SEC_COOKIEV0_SEPARATOR() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_SEC_COOKIEV0_SEPARATOR, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_SEC_DISABLE_BACKEND_COMPRESS() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_SEC_DISABLE_BACKEND_COMPRESS, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_SEC_GUARDIAN_LOG() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_SEC_GUARDIAN_LOG, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_SEC_HASH_ENGINE() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_SEC_HASH_ENGINE, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_SEC_HASH_KEY() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_SEC_HASH_KEY, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_SEC_HASH_METHOD_PM() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_SEC_HASH_METHOD_PM, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_SEC_HASH_METHOD_RX() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_SEC_HASH_METHOD_RX, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_SEC_HASH_PARAM() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_SEC_HASH_PARAM, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_SEC_HTTP_BLKEY() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_SEC_HTTP_BLKEY, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_SEC_INTERCEPT_ON_ERROR() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_SEC_INTERCEPT_ON_ERROR, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_SEC_REMOTE_RULES_FAIL_ACTION() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_SEC_REMOTE_RULES_FAIL_ACTION, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_SEC_RULE_INHERITANCE() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_SEC_RULE_INHERITANCE, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_SEC_RULE_PERF_TIME() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_SEC_RULE_PERF_TIME, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_SEC_SENSOR_ID() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_SEC_SENSOR_ID, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_SEC_STREAM_IN_BODY_INSPECTION() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_SEC_STREAM_IN_BODY_INSPECTION, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_SEC_STREAM_OUT_BODY_INSPECTION() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_SEC_STREAM_OUT_BODY_INSPECTION, 0)
}

func (s *Engine_config_directive_with_paramContext) CONFIG_XML_EXTERNAL_ENTITY() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_XML_EXTERNAL_ENTITY, 0)
}

func (s *Engine_config_directive_with_paramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Engine_config_directive_with_paramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Engine_config_directive_with_paramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterEngine_config_directive_with_param(s)
	}
}

func (s *Engine_config_directive_with_paramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitEngine_config_directive_with_param(s)
	}
}

func (p *SecLangParser) Engine_config_directive_with_param() (localctx IEngine_config_directive_with_paramContext) {
	localctx = NewEngine_config_directive_with_paramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SecLangParserRULE_engine_config_directive_with_param)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(302)
		_la = p.GetTokenStream().LA(1)

		if !(((int64((_la-131)) & ^0x3f) == 0 && ((int64(1)<<(_la-131))&25508532324925951) != 0) || ((int64((_la-209)) & ^0x3f) == 0 && ((int64(1)<<(_la-209))&255) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRule_script_directiveContext is an interface to support dynamic dispatch.
type IRule_script_directiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DIRECTIVE_SECRULESCRIPT() antlr.TerminalNode

	// IsRule_script_directiveContext differentiates from other interfaces.
	IsRule_script_directiveContext()
}

type Rule_script_directiveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRule_script_directiveContext() *Rule_script_directiveContext {
	var p = new(Rule_script_directiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_rule_script_directive
	return p
}

func InitEmptyRule_script_directiveContext(p *Rule_script_directiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_rule_script_directive
}

func (*Rule_script_directiveContext) IsRule_script_directiveContext() {}

func NewRule_script_directiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Rule_script_directiveContext {
	var p = new(Rule_script_directiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_rule_script_directive

	return p
}

func (s *Rule_script_directiveContext) GetParser() antlr.Parser { return s.parser }

func (s *Rule_script_directiveContext) DIRECTIVE_SECRULESCRIPT() antlr.TerminalNode {
	return s.GetToken(SecLangParserDIRECTIVE_SECRULESCRIPT, 0)
}

func (s *Rule_script_directiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Rule_script_directiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Rule_script_directiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterRule_script_directive(s)
	}
}

func (s *Rule_script_directiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitRule_script_directive(s)
	}
}

func (p *SecLangParser) Rule_script_directive() (localctx IRule_script_directiveContext) {
	localctx = NewRule_script_directiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SecLangParserRULE_rule_script_directive)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(304)
		p.Match(SecLangParserDIRECTIVE_SECRULESCRIPT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFile_pathContext is an interface to support dynamic dispatch.
type IFile_pathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONFIG_VALUE_PATH() antlr.TerminalNode

	// IsFile_pathContext differentiates from other interfaces.
	IsFile_pathContext()
}

type File_pathContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFile_pathContext() *File_pathContext {
	var p = new(File_pathContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_file_path
	return p
}

func InitEmptyFile_pathContext(p *File_pathContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_file_path
}

func (*File_pathContext) IsFile_pathContext() {}

func NewFile_pathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *File_pathContext {
	var p = new(File_pathContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_file_path

	return p
}

func (s *File_pathContext) GetParser() antlr.Parser { return s.parser }

func (s *File_pathContext) CONFIG_VALUE_PATH() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_VALUE_PATH, 0)
}

func (s *File_pathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *File_pathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *File_pathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterFile_path(s)
	}
}

func (s *File_pathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitFile_path(s)
	}
}

func (p *SecLangParser) File_path() (localctx IFile_pathContext) {
	localctx = NewFile_pathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SecLangParserRULE_file_path)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(306)
		p.Match(SecLangParserCONFIG_VALUE_PATH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRemove_rule_by_idContext is an interface to support dynamic dispatch.
type IRemove_rule_by_idContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONFIG_SEC_RULE_REMOVE_BY_ID() antlr.TerminalNode

	// IsRemove_rule_by_idContext differentiates from other interfaces.
	IsRemove_rule_by_idContext()
}

type Remove_rule_by_idContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRemove_rule_by_idContext() *Remove_rule_by_idContext {
	var p = new(Remove_rule_by_idContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_remove_rule_by_id
	return p
}

func InitEmptyRemove_rule_by_idContext(p *Remove_rule_by_idContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_remove_rule_by_id
}

func (*Remove_rule_by_idContext) IsRemove_rule_by_idContext() {}

func NewRemove_rule_by_idContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Remove_rule_by_idContext {
	var p = new(Remove_rule_by_idContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_remove_rule_by_id

	return p
}

func (s *Remove_rule_by_idContext) GetParser() antlr.Parser { return s.parser }

func (s *Remove_rule_by_idContext) CONFIG_SEC_RULE_REMOVE_BY_ID() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_SEC_RULE_REMOVE_BY_ID, 0)
}

func (s *Remove_rule_by_idContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Remove_rule_by_idContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Remove_rule_by_idContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterRemove_rule_by_id(s)
	}
}

func (s *Remove_rule_by_idContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitRemove_rule_by_id(s)
	}
}

func (p *SecLangParser) Remove_rule_by_id() (localctx IRemove_rule_by_idContext) {
	localctx = NewRemove_rule_by_idContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SecLangParserRULE_remove_rule_by_id)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(308)
		p.Match(SecLangParserCONFIG_SEC_RULE_REMOVE_BY_ID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRemove_rule_by_id_valuesContext is an interface to support dynamic dispatch.
type IRemove_rule_by_id_valuesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsRemove_rule_by_id_valuesContext differentiates from other interfaces.
	IsRemove_rule_by_id_valuesContext()
}

type Remove_rule_by_id_valuesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRemove_rule_by_id_valuesContext() *Remove_rule_by_id_valuesContext {
	var p = new(Remove_rule_by_id_valuesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_remove_rule_by_id_values
	return p
}

func InitEmptyRemove_rule_by_id_valuesContext(p *Remove_rule_by_id_valuesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_remove_rule_by_id_values
}

func (*Remove_rule_by_id_valuesContext) IsRemove_rule_by_id_valuesContext() {}

func NewRemove_rule_by_id_valuesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Remove_rule_by_id_valuesContext {
	var p = new(Remove_rule_by_id_valuesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_remove_rule_by_id_values

	return p
}

func (s *Remove_rule_by_id_valuesContext) GetParser() antlr.Parser { return s.parser }

func (s *Remove_rule_by_id_valuesContext) CopyAll(ctx *Remove_rule_by_id_valuesContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Remove_rule_by_id_valuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Remove_rule_by_id_valuesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Remove_rule_by_id_intContext struct {
	Remove_rule_by_id_valuesContext
}

func NewRemove_rule_by_id_intContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Remove_rule_by_id_intContext {
	var p = new(Remove_rule_by_id_intContext)

	InitEmptyRemove_rule_by_id_valuesContext(&p.Remove_rule_by_id_valuesContext)
	p.parser = parser
	p.CopyAll(ctx.(*Remove_rule_by_id_valuesContext))

	return p
}

func (s *Remove_rule_by_id_intContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Remove_rule_by_id_intContext) INT() antlr.TerminalNode {
	return s.GetToken(SecLangParserINT, 0)
}

func (s *Remove_rule_by_id_intContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterRemove_rule_by_id_int(s)
	}
}

func (s *Remove_rule_by_id_intContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitRemove_rule_by_id_int(s)
	}
}

type Remove_rule_by_id_int_rangeContext struct {
	Remove_rule_by_id_valuesContext
}

func NewRemove_rule_by_id_int_rangeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Remove_rule_by_id_int_rangeContext {
	var p = new(Remove_rule_by_id_int_rangeContext)

	InitEmptyRemove_rule_by_id_valuesContext(&p.Remove_rule_by_id_valuesContext)
	p.parser = parser
	p.CopyAll(ctx.(*Remove_rule_by_id_valuesContext))

	return p
}

func (s *Remove_rule_by_id_int_rangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Remove_rule_by_id_int_rangeContext) Int_range() IInt_rangeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInt_rangeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInt_rangeContext)
}

func (s *Remove_rule_by_id_int_rangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterRemove_rule_by_id_int_range(s)
	}
}

func (s *Remove_rule_by_id_int_rangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitRemove_rule_by_id_int_range(s)
	}
}

func (p *SecLangParser) Remove_rule_by_id_values() (localctx IRemove_rule_by_id_valuesContext) {
	localctx = NewRemove_rule_by_id_valuesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SecLangParserRULE_remove_rule_by_id_values)
	p.SetState(312)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		localctx = NewRemove_rule_by_id_intContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(310)
			p.Match(SecLangParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewRemove_rule_by_id_int_rangeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(311)
			p.Int_range()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInt_rangeContext is an interface to support dynamic dispatch.
type IInt_rangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Range_start() IRange_startContext
	MINUS() antlr.TerminalNode
	Range_end() IRange_endContext

	// IsInt_rangeContext differentiates from other interfaces.
	IsInt_rangeContext()
}

type Int_rangeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInt_rangeContext() *Int_rangeContext {
	var p = new(Int_rangeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_int_range
	return p
}

func InitEmptyInt_rangeContext(p *Int_rangeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_int_range
}

func (*Int_rangeContext) IsInt_rangeContext() {}

func NewInt_rangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Int_rangeContext {
	var p = new(Int_rangeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_int_range

	return p
}

func (s *Int_rangeContext) GetParser() antlr.Parser { return s.parser }

func (s *Int_rangeContext) Range_start() IRange_startContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRange_startContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRange_startContext)
}

func (s *Int_rangeContext) MINUS() antlr.TerminalNode {
	return s.GetToken(SecLangParserMINUS, 0)
}

func (s *Int_rangeContext) Range_end() IRange_endContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRange_endContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRange_endContext)
}

func (s *Int_rangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Int_rangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Int_rangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterInt_range(s)
	}
}

func (s *Int_rangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitInt_range(s)
	}
}

func (p *SecLangParser) Int_range() (localctx IInt_rangeContext) {
	localctx = NewInt_rangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SecLangParserRULE_int_range)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(314)
		p.Range_start()
	}
	{
		p.SetState(315)
		p.Match(SecLangParserMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(316)
		p.Range_end()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRange_startContext is an interface to support dynamic dispatch.
type IRange_startContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode

	// IsRange_startContext differentiates from other interfaces.
	IsRange_startContext()
}

type Range_startContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRange_startContext() *Range_startContext {
	var p = new(Range_startContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_range_start
	return p
}

func InitEmptyRange_startContext(p *Range_startContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_range_start
}

func (*Range_startContext) IsRange_startContext() {}

func NewRange_startContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Range_startContext {
	var p = new(Range_startContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_range_start

	return p
}

func (s *Range_startContext) GetParser() antlr.Parser { return s.parser }

func (s *Range_startContext) INT() antlr.TerminalNode {
	return s.GetToken(SecLangParserINT, 0)
}

func (s *Range_startContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Range_startContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Range_startContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterRange_start(s)
	}
}

func (s *Range_startContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitRange_start(s)
	}
}

func (p *SecLangParser) Range_start() (localctx IRange_startContext) {
	localctx = NewRange_startContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SecLangParserRULE_range_start)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(318)
		p.Match(SecLangParserINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRange_endContext is an interface to support dynamic dispatch.
type IRange_endContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode

	// IsRange_endContext differentiates from other interfaces.
	IsRange_endContext()
}

type Range_endContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRange_endContext() *Range_endContext {
	var p = new(Range_endContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_range_end
	return p
}

func InitEmptyRange_endContext(p *Range_endContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_range_end
}

func (*Range_endContext) IsRange_endContext() {}

func NewRange_endContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Range_endContext {
	var p = new(Range_endContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_range_end

	return p
}

func (s *Range_endContext) GetParser() antlr.Parser { return s.parser }

func (s *Range_endContext) INT() antlr.TerminalNode {
	return s.GetToken(SecLangParserINT, 0)
}

func (s *Range_endContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Range_endContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Range_endContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterRange_end(s)
	}
}

func (s *Range_endContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitRange_end(s)
	}
}

func (p *SecLangParser) Range_end() (localctx IRange_endContext) {
	localctx = NewRange_endContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SecLangParserRULE_range_end)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(320)
		p.Match(SecLangParserINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IString_remove_rulesContext is an interface to support dynamic dispatch.
type IString_remove_rulesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsString_remove_rulesContext differentiates from other interfaces.
	IsString_remove_rulesContext()
}

type String_remove_rulesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyString_remove_rulesContext() *String_remove_rulesContext {
	var p = new(String_remove_rulesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_string_remove_rules
	return p
}

func InitEmptyString_remove_rulesContext(p *String_remove_rulesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_string_remove_rules
}

func (*String_remove_rulesContext) IsString_remove_rulesContext() {}

func NewString_remove_rulesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_remove_rulesContext {
	var p = new(String_remove_rulesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_string_remove_rules

	return p
}

func (s *String_remove_rulesContext) GetParser() antlr.Parser { return s.parser }

func (s *String_remove_rulesContext) CopyAll(ctx *String_remove_rulesContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *String_remove_rulesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_remove_rulesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Remove_rule_by_msgContext struct {
	String_remove_rulesContext
}

func NewRemove_rule_by_msgContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Remove_rule_by_msgContext {
	var p = new(Remove_rule_by_msgContext)

	InitEmptyString_remove_rulesContext(&p.String_remove_rulesContext)
	p.parser = parser
	p.CopyAll(ctx.(*String_remove_rulesContext))

	return p
}

func (s *Remove_rule_by_msgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Remove_rule_by_msgContext) CONFIG_SEC_RULE_REMOVE_BY_MSG() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_SEC_RULE_REMOVE_BY_MSG, 0)
}

func (s *Remove_rule_by_msgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterRemove_rule_by_msg(s)
	}
}

func (s *Remove_rule_by_msgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitRemove_rule_by_msg(s)
	}
}

type Remove_rule_by_tagContext struct {
	String_remove_rulesContext
}

func NewRemove_rule_by_tagContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Remove_rule_by_tagContext {
	var p = new(Remove_rule_by_tagContext)

	InitEmptyString_remove_rulesContext(&p.String_remove_rulesContext)
	p.parser = parser
	p.CopyAll(ctx.(*String_remove_rulesContext))

	return p
}

func (s *Remove_rule_by_tagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Remove_rule_by_tagContext) CONFIG_SEC_RULE_REMOVE_BY_TAG() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_SEC_RULE_REMOVE_BY_TAG, 0)
}

func (s *Remove_rule_by_tagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterRemove_rule_by_tag(s)
	}
}

func (s *Remove_rule_by_tagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitRemove_rule_by_tag(s)
	}
}

func (p *SecLangParser) String_remove_rules() (localctx IString_remove_rulesContext) {
	localctx = NewString_remove_rulesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SecLangParserRULE_string_remove_rules)
	p.SetState(324)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SecLangParserCONFIG_SEC_RULE_REMOVE_BY_MSG:
		localctx = NewRemove_rule_by_msgContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(322)
			p.Match(SecLangParserCONFIG_SEC_RULE_REMOVE_BY_MSG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SecLangParserCONFIG_SEC_RULE_REMOVE_BY_TAG:
		localctx = NewRemove_rule_by_tagContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(323)
			p.Match(SecLangParserCONFIG_SEC_RULE_REMOVE_BY_TAG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IString_remove_rules_valuesContext is an interface to support dynamic dispatch.
type IString_remove_rules_valuesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	VARIABLE_NAME() antlr.TerminalNode
	COMMA_SEPARATED_STRING() antlr.TerminalNode

	// IsString_remove_rules_valuesContext differentiates from other interfaces.
	IsString_remove_rules_valuesContext()
}

type String_remove_rules_valuesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyString_remove_rules_valuesContext() *String_remove_rules_valuesContext {
	var p = new(String_remove_rules_valuesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_string_remove_rules_values
	return p
}

func InitEmptyString_remove_rules_valuesContext(p *String_remove_rules_valuesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_string_remove_rules_values
}

func (*String_remove_rules_valuesContext) IsString_remove_rules_valuesContext() {}

func NewString_remove_rules_valuesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_remove_rules_valuesContext {
	var p = new(String_remove_rules_valuesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_string_remove_rules_values

	return p
}

func (s *String_remove_rules_valuesContext) GetParser() antlr.Parser { return s.parser }

func (s *String_remove_rules_valuesContext) STRING() antlr.TerminalNode {
	return s.GetToken(SecLangParserSTRING, 0)
}

func (s *String_remove_rules_valuesContext) VARIABLE_NAME() antlr.TerminalNode {
	return s.GetToken(SecLangParserVARIABLE_NAME, 0)
}

func (s *String_remove_rules_valuesContext) COMMA_SEPARATED_STRING() antlr.TerminalNode {
	return s.GetToken(SecLangParserCOMMA_SEPARATED_STRING, 0)
}

func (s *String_remove_rules_valuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_remove_rules_valuesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_remove_rules_valuesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterString_remove_rules_values(s)
	}
}

func (s *String_remove_rules_valuesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitString_remove_rules_values(s)
	}
}

func (p *SecLangParser) String_remove_rules_values() (localctx IString_remove_rules_valuesContext) {
	localctx = NewString_remove_rules_valuesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SecLangParserRULE_string_remove_rules_values)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(326)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-225)) & ^0x3f) == 0 && ((int64(1)<<(_la-225))&65793) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUpdate_target_rulesContext is an interface to support dynamic dispatch.
type IUpdate_target_rulesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsUpdate_target_rulesContext differentiates from other interfaces.
	IsUpdate_target_rulesContext()
}

type Update_target_rulesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdate_target_rulesContext() *Update_target_rulesContext {
	var p = new(Update_target_rulesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_update_target_rules
	return p
}

func InitEmptyUpdate_target_rulesContext(p *Update_target_rulesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_update_target_rules
}

func (*Update_target_rulesContext) IsUpdate_target_rulesContext() {}

func NewUpdate_target_rulesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Update_target_rulesContext {
	var p = new(Update_target_rulesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_update_target_rules

	return p
}

func (s *Update_target_rulesContext) GetParser() antlr.Parser { return s.parser }

func (s *Update_target_rulesContext) CopyAll(ctx *Update_target_rulesContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Update_target_rulesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Update_target_rulesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Update_target_by_tagContext struct {
	Update_target_rulesContext
}

func NewUpdate_target_by_tagContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Update_target_by_tagContext {
	var p = new(Update_target_by_tagContext)

	InitEmptyUpdate_target_rulesContext(&p.Update_target_rulesContext)
	p.parser = parser
	p.CopyAll(ctx.(*Update_target_rulesContext))

	return p
}

func (s *Update_target_by_tagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Update_target_by_tagContext) CONFIG_SEC_RULE_UPDATE_TARGET_BY_TAG() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_SEC_RULE_UPDATE_TARGET_BY_TAG, 0)
}

func (s *Update_target_by_tagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterUpdate_target_by_tag(s)
	}
}

func (s *Update_target_by_tagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitUpdate_target_by_tag(s)
	}
}

type Update_target_by_idContext struct {
	Update_target_rulesContext
}

func NewUpdate_target_by_idContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Update_target_by_idContext {
	var p = new(Update_target_by_idContext)

	InitEmptyUpdate_target_rulesContext(&p.Update_target_rulesContext)
	p.parser = parser
	p.CopyAll(ctx.(*Update_target_rulesContext))

	return p
}

func (s *Update_target_by_idContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Update_target_by_idContext) CONFIG_SEC_RULE_UPDATE_TARGET_BY_ID() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_SEC_RULE_UPDATE_TARGET_BY_ID, 0)
}

func (s *Update_target_by_idContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterUpdate_target_by_id(s)
	}
}

func (s *Update_target_by_idContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitUpdate_target_by_id(s)
	}
}

type Update_target_by_msgContext struct {
	Update_target_rulesContext
}

func NewUpdate_target_by_msgContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Update_target_by_msgContext {
	var p = new(Update_target_by_msgContext)

	InitEmptyUpdate_target_rulesContext(&p.Update_target_rulesContext)
	p.parser = parser
	p.CopyAll(ctx.(*Update_target_rulesContext))

	return p
}

func (s *Update_target_by_msgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Update_target_by_msgContext) CONFIG_SEC_RULE_UPDATE_TARGET_BY_MSG() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_SEC_RULE_UPDATE_TARGET_BY_MSG, 0)
}

func (s *Update_target_by_msgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterUpdate_target_by_msg(s)
	}
}

func (s *Update_target_by_msgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitUpdate_target_by_msg(s)
	}
}

func (p *SecLangParser) Update_target_rules() (localctx IUpdate_target_rulesContext) {
	localctx = NewUpdate_target_rulesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SecLangParserRULE_update_target_rules)
	p.SetState(331)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SecLangParserCONFIG_SEC_RULE_UPDATE_TARGET_BY_ID:
		localctx = NewUpdate_target_by_idContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(328)
			p.Match(SecLangParserCONFIG_SEC_RULE_UPDATE_TARGET_BY_ID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SecLangParserCONFIG_SEC_RULE_UPDATE_TARGET_BY_MSG:
		localctx = NewUpdate_target_by_msgContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(329)
			p.Match(SecLangParserCONFIG_SEC_RULE_UPDATE_TARGET_BY_MSG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SecLangParserCONFIG_SEC_RULE_UPDATE_TARGET_BY_TAG:
		localctx = NewUpdate_target_by_tagContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(330)
			p.Match(SecLangParserCONFIG_SEC_RULE_UPDATE_TARGET_BY_TAG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUpdate_action_ruleContext is an interface to support dynamic dispatch.
type IUpdate_action_ruleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONFIG_SEC_RULE_UPDATE_ACTION_BY_ID() antlr.TerminalNode

	// IsUpdate_action_ruleContext differentiates from other interfaces.
	IsUpdate_action_ruleContext()
}

type Update_action_ruleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdate_action_ruleContext() *Update_action_ruleContext {
	var p = new(Update_action_ruleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_update_action_rule
	return p
}

func InitEmptyUpdate_action_ruleContext(p *Update_action_ruleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_update_action_rule
}

func (*Update_action_ruleContext) IsUpdate_action_ruleContext() {}

func NewUpdate_action_ruleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Update_action_ruleContext {
	var p = new(Update_action_ruleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_update_action_rule

	return p
}

func (s *Update_action_ruleContext) GetParser() antlr.Parser { return s.parser }

func (s *Update_action_ruleContext) CONFIG_SEC_RULE_UPDATE_ACTION_BY_ID() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_SEC_RULE_UPDATE_ACTION_BY_ID, 0)
}

func (s *Update_action_ruleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Update_action_ruleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Update_action_ruleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterUpdate_action_rule(s)
	}
}

func (s *Update_action_ruleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitUpdate_action_rule(s)
	}
}

func (p *SecLangParser) Update_action_rule() (localctx IUpdate_action_ruleContext) {
	localctx = NewUpdate_action_ruleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SecLangParserRULE_update_action_rule)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(333)
		p.Match(SecLangParserCONFIG_SEC_RULE_UPDATE_ACTION_BY_ID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdContext is an interface to support dynamic dispatch.
type IIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode

	// IsIdContext differentiates from other interfaces.
	IsIdContext()
}

type IdContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdContext() *IdContext {
	var p = new(IdContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_id
	return p
}

func InitEmptyIdContext(p *IdContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_id
}

func (*IdContext) IsIdContext() {}

func NewIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdContext {
	var p = new(IdContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_id

	return p
}

func (s *IdContext) GetParser() antlr.Parser { return s.parser }

func (s *IdContext) INT() antlr.TerminalNode {
	return s.GetToken(SecLangParserINT, 0)
}

func (s *IdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterId(s)
	}
}

func (s *IdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitId(s)
	}
}

func (p *SecLangParser) Id() (localctx IIdContext) {
	localctx = NewIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SecLangParserRULE_id)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(335)
		p.Match(SecLangParserINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEngine_config_sec_cache_transformationsContext is an interface to support dynamic dispatch.
type IEngine_config_sec_cache_transformationsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONFIG_SEC_CACHE_TRANSFORMATIONS() antlr.TerminalNode

	// IsEngine_config_sec_cache_transformationsContext differentiates from other interfaces.
	IsEngine_config_sec_cache_transformationsContext()
}

type Engine_config_sec_cache_transformationsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEngine_config_sec_cache_transformationsContext() *Engine_config_sec_cache_transformationsContext {
	var p = new(Engine_config_sec_cache_transformationsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_engine_config_sec_cache_transformations
	return p
}

func InitEmptyEngine_config_sec_cache_transformationsContext(p *Engine_config_sec_cache_transformationsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_engine_config_sec_cache_transformations
}

func (*Engine_config_sec_cache_transformationsContext) IsEngine_config_sec_cache_transformationsContext() {
}

func NewEngine_config_sec_cache_transformationsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Engine_config_sec_cache_transformationsContext {
	var p = new(Engine_config_sec_cache_transformationsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_engine_config_sec_cache_transformations

	return p
}

func (s *Engine_config_sec_cache_transformationsContext) GetParser() antlr.Parser { return s.parser }

func (s *Engine_config_sec_cache_transformationsContext) CONFIG_SEC_CACHE_TRANSFORMATIONS() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_SEC_CACHE_TRANSFORMATIONS, 0)
}

func (s *Engine_config_sec_cache_transformationsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Engine_config_sec_cache_transformationsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Engine_config_sec_cache_transformationsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterEngine_config_sec_cache_transformations(s)
	}
}

func (s *Engine_config_sec_cache_transformationsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitEngine_config_sec_cache_transformations(s)
	}
}

func (p *SecLangParser) Engine_config_sec_cache_transformations() (localctx IEngine_config_sec_cache_transformationsContext) {
	localctx = NewEngine_config_sec_cache_transformationsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SecLangParserRULE_engine_config_sec_cache_transformations)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(337)
		p.Match(SecLangParserCONFIG_SEC_CACHE_TRANSFORMATIONS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOption_listContext is an interface to support dynamic dispatch.
type IOption_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQUOTE() []antlr.TerminalNode
	QUOTE(i int) antlr.TerminalNode
	AllOption() []IOptionContext
	Option(i int) IOptionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsOption_listContext differentiates from other interfaces.
	IsOption_listContext()
}

type Option_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOption_listContext() *Option_listContext {
	var p = new(Option_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_option_list
	return p
}

func InitEmptyOption_listContext(p *Option_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_option_list
}

func (*Option_listContext) IsOption_listContext() {}

func NewOption_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Option_listContext {
	var p = new(Option_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_option_list

	return p
}

func (s *Option_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Option_listContext) AllQUOTE() []antlr.TerminalNode {
	return s.GetTokens(SecLangParserQUOTE)
}

func (s *Option_listContext) QUOTE(i int) antlr.TerminalNode {
	return s.GetToken(SecLangParserQUOTE, i)
}

func (s *Option_listContext) AllOption() []IOptionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOptionContext); ok {
			len++
		}
	}

	tst := make([]IOptionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOptionContext); ok {
			tst[i] = t.(IOptionContext)
			i++
		}
	}

	return tst
}

func (s *Option_listContext) Option(i int) IOptionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionContext)
}

func (s *Option_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SecLangParserCOMMA)
}

func (s *Option_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SecLangParserCOMMA, i)
}

func (s *Option_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Option_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Option_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterOption_list(s)
	}
}

func (s *Option_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitOption_list(s)
	}
}

func (p *SecLangParser) Option_list() (localctx IOption_listContext) {
	localctx = NewOption_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SecLangParserRULE_option_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(339)
		p.Match(SecLangParserQUOTE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(340)
		p.Option()
	}
	p.SetState(345)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SecLangParserCOMMA {
		{
			p.SetState(341)
			p.Match(SecLangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(342)
			p.Option()
		}

		p.SetState(347)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(348)
		p.Match(SecLangParserQUOTE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOptionContext is an interface to support dynamic dispatch.
type IOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Option_name() IOption_nameContext
	COLON() antlr.TerminalNode
	Values() IValuesContext

	// IsOptionContext differentiates from other interfaces.
	IsOptionContext()
}

type OptionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionContext() *OptionContext {
	var p = new(OptionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_option
	return p
}

func InitEmptyOptionContext(p *OptionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_option
}

func (*OptionContext) IsOptionContext() {}

func NewOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionContext {
	var p = new(OptionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_option

	return p
}

func (s *OptionContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionContext) Option_name() IOption_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOption_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOption_nameContext)
}

func (s *OptionContext) COLON() antlr.TerminalNode {
	return s.GetToken(SecLangParserCOLON, 0)
}

func (s *OptionContext) Values() IValuesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValuesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValuesContext)
}

func (s *OptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterOption(s)
	}
}

func (s *OptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitOption(s)
	}
}

func (p *SecLangParser) Option() (localctx IOptionContext) {
	localctx = NewOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SecLangParserRULE_option)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(350)
		p.Option_name()
	}
	{
		p.SetState(351)
		p.Match(SecLangParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(352)
		p.Values()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOption_nameContext is an interface to support dynamic dispatch.
type IOption_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPTION_NAME() antlr.TerminalNode

	// IsOption_nameContext differentiates from other interfaces.
	IsOption_nameContext()
}

type Option_nameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOption_nameContext() *Option_nameContext {
	var p = new(Option_nameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_option_name
	return p
}

func InitEmptyOption_nameContext(p *Option_nameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_option_name
}

func (*Option_nameContext) IsOption_nameContext() {}

func NewOption_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Option_nameContext {
	var p = new(Option_nameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_option_name

	return p
}

func (s *Option_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Option_nameContext) OPTION_NAME() antlr.TerminalNode {
	return s.GetToken(SecLangParserOPTION_NAME, 0)
}

func (s *Option_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Option_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Option_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterOption_name(s)
	}
}

func (s *Option_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitOption_name(s)
	}
}

func (p *SecLangParser) Option_name() (localctx IOption_nameContext) {
	localctx = NewOption_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SecLangParserRULE_option_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(354)
		p.Match(SecLangParserOPTION_NAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEngine_config_action_directiveContext is an interface to support dynamic dispatch.
type IEngine_config_action_directiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsEngine_config_action_directiveContext differentiates from other interfaces.
	IsEngine_config_action_directiveContext()
}

type Engine_config_action_directiveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEngine_config_action_directiveContext() *Engine_config_action_directiveContext {
	var p = new(Engine_config_action_directiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_engine_config_action_directive
	return p
}

func InitEmptyEngine_config_action_directiveContext(p *Engine_config_action_directiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_engine_config_action_directive
}

func (*Engine_config_action_directiveContext) IsEngine_config_action_directiveContext() {}

func NewEngine_config_action_directiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Engine_config_action_directiveContext {
	var p = new(Engine_config_action_directiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_engine_config_action_directive

	return p
}

func (s *Engine_config_action_directiveContext) GetParser() antlr.Parser { return s.parser }

func (s *Engine_config_action_directiveContext) CopyAll(ctx *Engine_config_action_directiveContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Engine_config_action_directiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Engine_config_action_directiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Config_dir_sec_actionContext struct {
	Engine_config_action_directiveContext
}

func NewConfig_dir_sec_actionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Config_dir_sec_actionContext {
	var p = new(Config_dir_sec_actionContext)

	InitEmptyEngine_config_action_directiveContext(&p.Engine_config_action_directiveContext)
	p.parser = parser
	p.CopyAll(ctx.(*Engine_config_action_directiveContext))

	return p
}

func (s *Config_dir_sec_actionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Config_dir_sec_actionContext) CONFIG_DIR_SEC_ACTION() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_DIR_SEC_ACTION, 0)
}

func (s *Config_dir_sec_actionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterConfig_dir_sec_action(s)
	}
}

func (s *Config_dir_sec_actionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitConfig_dir_sec_action(s)
	}
}

type Config_dir_sec_default_actionContext struct {
	Engine_config_action_directiveContext
}

func NewConfig_dir_sec_default_actionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Config_dir_sec_default_actionContext {
	var p = new(Config_dir_sec_default_actionContext)

	InitEmptyEngine_config_action_directiveContext(&p.Engine_config_action_directiveContext)
	p.parser = parser
	p.CopyAll(ctx.(*Engine_config_action_directiveContext))

	return p
}

func (s *Config_dir_sec_default_actionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Config_dir_sec_default_actionContext) CONFIG_DIR_SEC_DEFAULT_ACTION() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_DIR_SEC_DEFAULT_ACTION, 0)
}

func (s *Config_dir_sec_default_actionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterConfig_dir_sec_default_action(s)
	}
}

func (s *Config_dir_sec_default_actionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitConfig_dir_sec_default_action(s)
	}
}

func (p *SecLangParser) Engine_config_action_directive() (localctx IEngine_config_action_directiveContext) {
	localctx = NewEngine_config_action_directiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SecLangParserRULE_engine_config_action_directive)
	p.SetState(358)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SecLangParserCONFIG_DIR_SEC_ACTION:
		localctx = NewConfig_dir_sec_actionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(356)
			p.Match(SecLangParserCONFIG_DIR_SEC_ACTION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SecLangParserCONFIG_DIR_SEC_DEFAULT_ACTION:
		localctx = NewConfig_dir_sec_default_actionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(357)
			p.Match(SecLangParserCONFIG_DIR_SEC_DEFAULT_ACTION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStmt_audit_logContext is an interface to support dynamic dispatch.
type IStmt_audit_logContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONFIG_DIR_AUDIT_DIR_MOD() antlr.TerminalNode
	CONFIG_DIR_AUDIT_DIR() antlr.TerminalNode
	CONFIG_DIR_AUDIT_ENG() antlr.TerminalNode
	CONFIG_DIR_AUDIT_FILE_MODE() antlr.TerminalNode
	CONFIG_DIR_AUDIT_LOG2() antlr.TerminalNode
	CONFIG_DIR_AUDIT_LOG_P() antlr.TerminalNode
	CONFIG_DIR_AUDIT_LOG() antlr.TerminalNode
	CONFIG_DIR_AUDIT_LOG_FMT() antlr.TerminalNode
	CONFIG_DIR_AUDIT_STS() antlr.TerminalNode
	CONFIG_DIR_AUDIT_TYPE() antlr.TerminalNode
	CONFIG_UPLOAD_KEEP_FILES() antlr.TerminalNode
	CONFIG_UPLOAD_FILE_LIMIT() antlr.TerminalNode
	CONFIG_UPLOAD_FILE_MODE() antlr.TerminalNode
	CONFIG_UPLOAD_DIR() antlr.TerminalNode
	CONFIG_UPLOAD_SAVE_TMP_FILES() antlr.TerminalNode
	INT() antlr.TerminalNode

	// IsStmt_audit_logContext differentiates from other interfaces.
	IsStmt_audit_logContext()
}

type Stmt_audit_logContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmt_audit_logContext() *Stmt_audit_logContext {
	var p = new(Stmt_audit_logContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_stmt_audit_log
	return p
}

func InitEmptyStmt_audit_logContext(p *Stmt_audit_logContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_stmt_audit_log
}

func (*Stmt_audit_logContext) IsStmt_audit_logContext() {}

func NewStmt_audit_logContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Stmt_audit_logContext {
	var p = new(Stmt_audit_logContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_stmt_audit_log

	return p
}

func (s *Stmt_audit_logContext) GetParser() antlr.Parser { return s.parser }

func (s *Stmt_audit_logContext) CONFIG_DIR_AUDIT_DIR_MOD() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_DIR_AUDIT_DIR_MOD, 0)
}

func (s *Stmt_audit_logContext) CONFIG_DIR_AUDIT_DIR() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_DIR_AUDIT_DIR, 0)
}

func (s *Stmt_audit_logContext) CONFIG_DIR_AUDIT_ENG() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_DIR_AUDIT_ENG, 0)
}

func (s *Stmt_audit_logContext) CONFIG_DIR_AUDIT_FILE_MODE() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_DIR_AUDIT_FILE_MODE, 0)
}

func (s *Stmt_audit_logContext) CONFIG_DIR_AUDIT_LOG2() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_DIR_AUDIT_LOG2, 0)
}

func (s *Stmt_audit_logContext) CONFIG_DIR_AUDIT_LOG_P() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_DIR_AUDIT_LOG_P, 0)
}

func (s *Stmt_audit_logContext) CONFIG_DIR_AUDIT_LOG() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_DIR_AUDIT_LOG, 0)
}

func (s *Stmt_audit_logContext) CONFIG_DIR_AUDIT_LOG_FMT() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_DIR_AUDIT_LOG_FMT, 0)
}

func (s *Stmt_audit_logContext) CONFIG_DIR_AUDIT_STS() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_DIR_AUDIT_STS, 0)
}

func (s *Stmt_audit_logContext) CONFIG_DIR_AUDIT_TYPE() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_DIR_AUDIT_TYPE, 0)
}

func (s *Stmt_audit_logContext) CONFIG_UPLOAD_KEEP_FILES() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_UPLOAD_KEEP_FILES, 0)
}

func (s *Stmt_audit_logContext) CONFIG_UPLOAD_FILE_LIMIT() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_UPLOAD_FILE_LIMIT, 0)
}

func (s *Stmt_audit_logContext) CONFIG_UPLOAD_FILE_MODE() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_UPLOAD_FILE_MODE, 0)
}

func (s *Stmt_audit_logContext) CONFIG_UPLOAD_DIR() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_UPLOAD_DIR, 0)
}

func (s *Stmt_audit_logContext) CONFIG_UPLOAD_SAVE_TMP_FILES() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_UPLOAD_SAVE_TMP_FILES, 0)
}

func (s *Stmt_audit_logContext) INT() antlr.TerminalNode {
	return s.GetToken(SecLangParserINT, 0)
}

func (s *Stmt_audit_logContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Stmt_audit_logContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Stmt_audit_logContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterStmt_audit_log(s)
	}
}

func (s *Stmt_audit_logContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitStmt_audit_log(s)
	}
}

func (p *SecLangParser) Stmt_audit_log() (localctx IStmt_audit_logContext) {
	localctx = NewStmt_audit_logContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, SecLangParserRULE_stmt_audit_log)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(360)
		_la = p.GetTokenStream().LA(1)

		if !(((int64((_la-140)) & ^0x3f) == 0 && ((int64(1)<<(_la-140))&279223176896971775) != 0) || _la == SecLangParserINT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValuesContext is an interface to support dynamic dispatch.
type IValuesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	Int_range() IInt_rangeContext
	CONFIG_VALUE_ON() antlr.TerminalNode
	CONFIG_VALUE_OFF() antlr.TerminalNode
	CONFIG_VALUE_SERIAL() antlr.TerminalNode
	CONFIG_VALUE_PARALLEL() antlr.TerminalNode
	CONFIG_VALUE_HTTPS() antlr.TerminalNode
	CONFIG_VALUE_RELEVANT_ONLY() antlr.TerminalNode
	NATIVE() antlr.TerminalNode
	CONFIG_VALUE_ABORT() antlr.TerminalNode
	CONFIG_VALUE_WARN() antlr.TerminalNode
	CONFIG_VALUE_DETC() antlr.TerminalNode
	CONFIG_VALUE_PROCESS_PARTIAL() antlr.TerminalNode
	CONFIG_VALUE_REJECT() antlr.TerminalNode
	CONFIG_VALUE_PATH() antlr.TerminalNode
	STRING() antlr.TerminalNode
	VARIABLE_NAME() antlr.TerminalNode
	COMMA_SEPARATED_STRING() antlr.TerminalNode
	ACTION_CTL_BODY_PROCESSOR_TYPE() antlr.TerminalNode
	AUDIT_PARTS() antlr.TerminalNode
	Action_ctl_target_value() IAction_ctl_target_valueContext

	// IsValuesContext differentiates from other interfaces.
	IsValuesContext()
}

type ValuesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValuesContext() *ValuesContext {
	var p = new(ValuesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_values
	return p
}

func InitEmptyValuesContext(p *ValuesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_values
}

func (*ValuesContext) IsValuesContext() {}

func NewValuesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValuesContext {
	var p = new(ValuesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_values

	return p
}

func (s *ValuesContext) GetParser() antlr.Parser { return s.parser }

func (s *ValuesContext) INT() antlr.TerminalNode {
	return s.GetToken(SecLangParserINT, 0)
}

func (s *ValuesContext) Int_range() IInt_rangeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInt_rangeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInt_rangeContext)
}

func (s *ValuesContext) CONFIG_VALUE_ON() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_VALUE_ON, 0)
}

func (s *ValuesContext) CONFIG_VALUE_OFF() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_VALUE_OFF, 0)
}

func (s *ValuesContext) CONFIG_VALUE_SERIAL() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_VALUE_SERIAL, 0)
}

func (s *ValuesContext) CONFIG_VALUE_PARALLEL() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_VALUE_PARALLEL, 0)
}

func (s *ValuesContext) CONFIG_VALUE_HTTPS() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_VALUE_HTTPS, 0)
}

func (s *ValuesContext) CONFIG_VALUE_RELEVANT_ONLY() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_VALUE_RELEVANT_ONLY, 0)
}

func (s *ValuesContext) NATIVE() antlr.TerminalNode {
	return s.GetToken(SecLangParserNATIVE, 0)
}

func (s *ValuesContext) CONFIG_VALUE_ABORT() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_VALUE_ABORT, 0)
}

func (s *ValuesContext) CONFIG_VALUE_WARN() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_VALUE_WARN, 0)
}

func (s *ValuesContext) CONFIG_VALUE_DETC() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_VALUE_DETC, 0)
}

func (s *ValuesContext) CONFIG_VALUE_PROCESS_PARTIAL() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_VALUE_PROCESS_PARTIAL, 0)
}

func (s *ValuesContext) CONFIG_VALUE_REJECT() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_VALUE_REJECT, 0)
}

func (s *ValuesContext) CONFIG_VALUE_PATH() antlr.TerminalNode {
	return s.GetToken(SecLangParserCONFIG_VALUE_PATH, 0)
}

func (s *ValuesContext) STRING() antlr.TerminalNode {
	return s.GetToken(SecLangParserSTRING, 0)
}

func (s *ValuesContext) VARIABLE_NAME() antlr.TerminalNode {
	return s.GetToken(SecLangParserVARIABLE_NAME, 0)
}

func (s *ValuesContext) COMMA_SEPARATED_STRING() antlr.TerminalNode {
	return s.GetToken(SecLangParserCOMMA_SEPARATED_STRING, 0)
}

func (s *ValuesContext) ACTION_CTL_BODY_PROCESSOR_TYPE() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_CTL_BODY_PROCESSOR_TYPE, 0)
}

func (s *ValuesContext) AUDIT_PARTS() antlr.TerminalNode {
	return s.GetToken(SecLangParserAUDIT_PARTS, 0)
}

func (s *ValuesContext) Action_ctl_target_value() IAction_ctl_target_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAction_ctl_target_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAction_ctl_target_valueContext)
}

func (s *ValuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValuesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValuesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterValues(s)
	}
}

func (s *ValuesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitValues(s)
	}
}

func (p *SecLangParser) Values() (localctx IValuesContext) {
	localctx = NewValuesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, SecLangParserRULE_values)
	p.SetState(385)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(362)
			p.Match(SecLangParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(363)
			p.Int_range()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(364)
			p.Match(SecLangParserCONFIG_VALUE_ON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(365)
			p.Match(SecLangParserCONFIG_VALUE_OFF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(366)
			p.Match(SecLangParserCONFIG_VALUE_SERIAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(367)
			p.Match(SecLangParserCONFIG_VALUE_PARALLEL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(368)
			p.Match(SecLangParserCONFIG_VALUE_HTTPS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(369)
			p.Match(SecLangParserCONFIG_VALUE_RELEVANT_ONLY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(370)
			p.Match(SecLangParserNATIVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(371)
			p.Match(SecLangParserCONFIG_VALUE_ABORT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(372)
			p.Match(SecLangParserCONFIG_VALUE_WARN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(373)
			p.Match(SecLangParserCONFIG_VALUE_DETC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(374)
			p.Match(SecLangParserCONFIG_VALUE_PROCESS_PARTIAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(375)
			p.Match(SecLangParserCONFIG_VALUE_REJECT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(376)
			p.Match(SecLangParserCONFIG_VALUE_PATH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(377)
			p.Match(SecLangParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(378)
			p.Match(SecLangParserCONFIG_VALUE_PATH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(379)
			p.Match(SecLangParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 18:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(380)
			p.Match(SecLangParserVARIABLE_NAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 19:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(381)
			p.Match(SecLangParserCOMMA_SEPARATED_STRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 20:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(382)
			p.Match(SecLangParserACTION_CTL_BODY_PROCESSOR_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 21:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(383)
			p.Match(SecLangParserAUDIT_PARTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 22:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(384)
			p.Action_ctl_target_value()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAction_ctl_target_valueContext is an interface to support dynamic dispatch.
type IAction_ctl_target_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SEMI() antlr.TerminalNode
	Ctl_variable_enum() ICtl_variable_enumContext
	Ctl_id() ICtl_idContext
	AllSINGLE_QUOTE() []antlr.TerminalNode
	SINGLE_QUOTE(i int) antlr.TerminalNode
	String_literal() IString_literalContext
	VARIABLE_NAME() antlr.TerminalNode
	Ctl_collection_enum() ICtl_collection_enumContext
	COLON() antlr.TerminalNode
	Ctl_collection_value() ICtl_collection_valueContext

	// IsAction_ctl_target_valueContext differentiates from other interfaces.
	IsAction_ctl_target_valueContext()
}

type Action_ctl_target_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAction_ctl_target_valueContext() *Action_ctl_target_valueContext {
	var p = new(Action_ctl_target_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_action_ctl_target_value
	return p
}

func InitEmptyAction_ctl_target_valueContext(p *Action_ctl_target_valueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_action_ctl_target_value
}

func (*Action_ctl_target_valueContext) IsAction_ctl_target_valueContext() {}

func NewAction_ctl_target_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Action_ctl_target_valueContext {
	var p = new(Action_ctl_target_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_action_ctl_target_value

	return p
}

func (s *Action_ctl_target_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Action_ctl_target_valueContext) SEMI() antlr.TerminalNode {
	return s.GetToken(SecLangParserSEMI, 0)
}

func (s *Action_ctl_target_valueContext) Ctl_variable_enum() ICtl_variable_enumContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICtl_variable_enumContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICtl_variable_enumContext)
}

func (s *Action_ctl_target_valueContext) Ctl_id() ICtl_idContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICtl_idContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICtl_idContext)
}

func (s *Action_ctl_target_valueContext) AllSINGLE_QUOTE() []antlr.TerminalNode {
	return s.GetTokens(SecLangParserSINGLE_QUOTE)
}

func (s *Action_ctl_target_valueContext) SINGLE_QUOTE(i int) antlr.TerminalNode {
	return s.GetToken(SecLangParserSINGLE_QUOTE, i)
}

func (s *Action_ctl_target_valueContext) String_literal() IString_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IString_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IString_literalContext)
}

func (s *Action_ctl_target_valueContext) VARIABLE_NAME() antlr.TerminalNode {
	return s.GetToken(SecLangParserVARIABLE_NAME, 0)
}

func (s *Action_ctl_target_valueContext) Ctl_collection_enum() ICtl_collection_enumContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICtl_collection_enumContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICtl_collection_enumContext)
}

func (s *Action_ctl_target_valueContext) COLON() antlr.TerminalNode {
	return s.GetToken(SecLangParserCOLON, 0)
}

func (s *Action_ctl_target_valueContext) Ctl_collection_value() ICtl_collection_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICtl_collection_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICtl_collection_valueContext)
}

func (s *Action_ctl_target_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Action_ctl_target_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Action_ctl_target_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterAction_ctl_target_value(s)
	}
}

func (s *Action_ctl_target_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitAction_ctl_target_value(s)
	}
}

func (p *SecLangParser) Action_ctl_target_value() (localctx IAction_ctl_target_valueContext) {
	localctx = NewAction_ctl_target_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, SecLangParserRULE_action_ctl_target_value)
	var _la int

	p.SetState(411)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(393)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case SecLangParserIDENT, SecLangParserINT:
			{
				p.SetState(387)
				p.Ctl_id()
			}

		case SecLangParserSINGLE_QUOTE:
			{
				p.SetState(388)
				p.Match(SecLangParserSINGLE_QUOTE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(389)
				p.String_literal()
			}
			{
				p.SetState(390)
				p.Match(SecLangParserSINGLE_QUOTE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case SecLangParserVARIABLE_NAME:
			{
				p.SetState(392)
				p.Match(SecLangParserVARIABLE_NAME)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}
		{
			p.SetState(395)
			p.Match(SecLangParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(396)
			p.Ctl_variable_enum()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(403)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case SecLangParserIDENT, SecLangParserINT:
			{
				p.SetState(397)
				p.Ctl_id()
			}

		case SecLangParserSINGLE_QUOTE:
			{
				p.SetState(398)
				p.Match(SecLangParserSINGLE_QUOTE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(399)
				p.String_literal()
			}
			{
				p.SetState(400)
				p.Match(SecLangParserSINGLE_QUOTE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case SecLangParserVARIABLE_NAME:
			{
				p.SetState(402)
				p.Match(SecLangParserVARIABLE_NAME)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}
		{
			p.SetState(405)
			p.Match(SecLangParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(406)
			p.Ctl_collection_enum()
		}
		p.SetState(409)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SecLangParserCOLON {
			{
				p.SetState(407)
				p.Match(SecLangParserCOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(408)
				p.Ctl_collection_value()
			}

		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUpdate_target_rules_valuesContext is an interface to support dynamic dispatch.
type IUpdate_target_rules_valuesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	Int_range() IInt_rangeContext
	STRING() antlr.TerminalNode

	// IsUpdate_target_rules_valuesContext differentiates from other interfaces.
	IsUpdate_target_rules_valuesContext()
}

type Update_target_rules_valuesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdate_target_rules_valuesContext() *Update_target_rules_valuesContext {
	var p = new(Update_target_rules_valuesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_update_target_rules_values
	return p
}

func InitEmptyUpdate_target_rules_valuesContext(p *Update_target_rules_valuesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_update_target_rules_values
}

func (*Update_target_rules_valuesContext) IsUpdate_target_rules_valuesContext() {}

func NewUpdate_target_rules_valuesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Update_target_rules_valuesContext {
	var p = new(Update_target_rules_valuesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_update_target_rules_values

	return p
}

func (s *Update_target_rules_valuesContext) GetParser() antlr.Parser { return s.parser }

func (s *Update_target_rules_valuesContext) INT() antlr.TerminalNode {
	return s.GetToken(SecLangParserINT, 0)
}

func (s *Update_target_rules_valuesContext) Int_range() IInt_rangeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInt_rangeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInt_rangeContext)
}

func (s *Update_target_rules_valuesContext) STRING() antlr.TerminalNode {
	return s.GetToken(SecLangParserSTRING, 0)
}

func (s *Update_target_rules_valuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Update_target_rules_valuesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Update_target_rules_valuesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterUpdate_target_rules_values(s)
	}
}

func (s *Update_target_rules_valuesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitUpdate_target_rules_values(s)
	}
}

func (p *SecLangParser) Update_target_rules_values() (localctx IUpdate_target_rules_valuesContext) {
	localctx = NewUpdate_target_rules_valuesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, SecLangParserRULE_update_target_rules_values)
	p.SetState(416)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(413)
			p.Match(SecLangParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(414)
			p.Int_range()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(415)
			p.Match(SecLangParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOperator_notContext is an interface to support dynamic dispatch.
type IOperator_notContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NOT() antlr.TerminalNode

	// IsOperator_notContext differentiates from other interfaces.
	IsOperator_notContext()
}

type Operator_notContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperator_notContext() *Operator_notContext {
	var p = new(Operator_notContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_operator_not
	return p
}

func InitEmptyOperator_notContext(p *Operator_notContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_operator_not
}

func (*Operator_notContext) IsOperator_notContext() {}

func NewOperator_notContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Operator_notContext {
	var p = new(Operator_notContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_operator_not

	return p
}

func (s *Operator_notContext) GetParser() antlr.Parser { return s.parser }

func (s *Operator_notContext) NOT() antlr.TerminalNode {
	return s.GetToken(SecLangParserNOT, 0)
}

func (s *Operator_notContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Operator_notContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Operator_notContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterOperator_not(s)
	}
}

func (s *Operator_notContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitOperator_not(s)
	}
}

func (p *SecLangParser) Operator_not() (localctx IOperator_notContext) {
	localctx = NewOperator_notContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, SecLangParserRULE_operator_not)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(418)
		p.Match(SecLangParserNOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOperatorContext is an interface to support dynamic dispatch.
type IOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQUOTE() []antlr.TerminalNode
	QUOTE(i int) antlr.TerminalNode
	AT() antlr.TerminalNode
	Operator_name() IOperator_nameContext
	Operator_not() IOperator_notContext
	Operator_value() IOperator_valueContext

	// IsOperatorContext differentiates from other interfaces.
	IsOperatorContext()
}

type OperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorContext() *OperatorContext {
	var p = new(OperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_operator
	return p
}

func InitEmptyOperatorContext(p *OperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_operator
}

func (*OperatorContext) IsOperatorContext() {}

func NewOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorContext {
	var p = new(OperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_operator

	return p
}

func (s *OperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *OperatorContext) AllQUOTE() []antlr.TerminalNode {
	return s.GetTokens(SecLangParserQUOTE)
}

func (s *OperatorContext) QUOTE(i int) antlr.TerminalNode {
	return s.GetToken(SecLangParserQUOTE, i)
}

func (s *OperatorContext) AT() antlr.TerminalNode {
	return s.GetToken(SecLangParserAT, 0)
}

func (s *OperatorContext) Operator_name() IOperator_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperator_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperator_nameContext)
}

func (s *OperatorContext) Operator_not() IOperator_notContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperator_notContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperator_notContext)
}

func (s *OperatorContext) Operator_value() IOperator_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperator_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperator_valueContext)
}

func (s *OperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterOperator(s)
	}
}

func (s *OperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitOperator(s)
	}
}

func (p *SecLangParser) Operator() (localctx IOperatorContext) {
	localctx = NewOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, SecLangParserRULE_operator)
	var _la int

	p.SetState(436)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(420)
			p.Match(SecLangParserQUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(422)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SecLangParserNOT {
			{
				p.SetState(421)
				p.Operator_not()
			}

		}
		{
			p.SetState(424)
			p.Match(SecLangParserAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(425)
			p.Operator_name()
		}
		p.SetState(427)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SecLangParserVARIABLE_NAME_ENUM || ((int64((_la-227)) & ^0x3f) == 0 && ((int64(1)<<(_la-227))&671088705) != 0) {
			{
				p.SetState(426)
				p.Operator_value()
			}

		}
		{
			p.SetState(429)
			p.Match(SecLangParserQUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(431)
			p.Match(SecLangParserQUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(432)
			p.Operator_value()
		}
		{
			p.SetState(433)
			p.Match(SecLangParserQUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(435)
			p.Operator_value()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOperator_nameContext is an interface to support dynamic dispatch.
type IOperator_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPERATOR_UNCONDITIONAL_MATCH() antlr.TerminalNode
	OPERATOR_DETECT_SQLI() antlr.TerminalNode
	OPERATOR_DETECT_XSS() antlr.TerminalNode
	OPERATOR_VALIDATE_URL_ENCODING() antlr.TerminalNode
	OPERATOR_VALIDATE_UTF8_ENCODING() antlr.TerminalNode
	OPERATOR_INSPECT_FILE() antlr.TerminalNode
	OPERATOR_FUZZY_HASH() antlr.TerminalNode
	OPERATOR_VALIDATE_BYTE_RANGE() antlr.TerminalNode
	OPERATOR_VALIDATE_DTD() antlr.TerminalNode
	OPERATOR_VALIDATE_HASH() antlr.TerminalNode
	OPERATOR_VALIDATE_SCHEMA() antlr.TerminalNode
	OPERATOR_VERIFY_CC() antlr.TerminalNode
	OPERATOR_VERIFY_CPF() antlr.TerminalNode
	OPERATOR_VERIFY_SSN() antlr.TerminalNode
	OPERATOR_VERIFY_SVNR() antlr.TerminalNode
	OPERATOR_GSB_LOOKUP() antlr.TerminalNode
	OPERATOR_RSUB() antlr.TerminalNode
	OPERATOR_WITHIN() antlr.TerminalNode
	OPERATOR_CONTAINS_WORD() antlr.TerminalNode
	OPERATOR_CONTAINS() antlr.TerminalNode
	OPERATOR_ENDS_WITH() antlr.TerminalNode
	OPERATOR_EQ() antlr.TerminalNode
	OPERATOR_GE() antlr.TerminalNode
	OPERATOR_GT() antlr.TerminalNode
	OPERATOR_IP_MATCH_FROM_FILE() antlr.TerminalNode
	OPERATOR_IP_MATCH() antlr.TerminalNode
	OPERATOR_LE() antlr.TerminalNode
	OPERATOR_LT() antlr.TerminalNode
	OPERATOR_PM_FROM_FILE() antlr.TerminalNode
	OPERATOR_PM() antlr.TerminalNode
	OPERATOR_RBL() antlr.TerminalNode
	OPERATOR_RX() antlr.TerminalNode
	OPERATOR_RX_GLOBAL() antlr.TerminalNode
	OPERATOR_STR_EQ() antlr.TerminalNode
	OPERATOR_STR_MATCH() antlr.TerminalNode
	OPERATOR_BEGINS_WITH() antlr.TerminalNode
	OPERATOR_GEOLOOKUP() antlr.TerminalNode

	// IsOperator_nameContext differentiates from other interfaces.
	IsOperator_nameContext()
}

type Operator_nameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperator_nameContext() *Operator_nameContext {
	var p = new(Operator_nameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_operator_name
	return p
}

func InitEmptyOperator_nameContext(p *Operator_nameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_operator_name
}

func (*Operator_nameContext) IsOperator_nameContext() {}

func NewOperator_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Operator_nameContext {
	var p = new(Operator_nameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_operator_name

	return p
}

func (s *Operator_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Operator_nameContext) OPERATOR_UNCONDITIONAL_MATCH() antlr.TerminalNode {
	return s.GetToken(SecLangParserOPERATOR_UNCONDITIONAL_MATCH, 0)
}

func (s *Operator_nameContext) OPERATOR_DETECT_SQLI() antlr.TerminalNode {
	return s.GetToken(SecLangParserOPERATOR_DETECT_SQLI, 0)
}

func (s *Operator_nameContext) OPERATOR_DETECT_XSS() antlr.TerminalNode {
	return s.GetToken(SecLangParserOPERATOR_DETECT_XSS, 0)
}

func (s *Operator_nameContext) OPERATOR_VALIDATE_URL_ENCODING() antlr.TerminalNode {
	return s.GetToken(SecLangParserOPERATOR_VALIDATE_URL_ENCODING, 0)
}

func (s *Operator_nameContext) OPERATOR_VALIDATE_UTF8_ENCODING() antlr.TerminalNode {
	return s.GetToken(SecLangParserOPERATOR_VALIDATE_UTF8_ENCODING, 0)
}

func (s *Operator_nameContext) OPERATOR_INSPECT_FILE() antlr.TerminalNode {
	return s.GetToken(SecLangParserOPERATOR_INSPECT_FILE, 0)
}

func (s *Operator_nameContext) OPERATOR_FUZZY_HASH() antlr.TerminalNode {
	return s.GetToken(SecLangParserOPERATOR_FUZZY_HASH, 0)
}

func (s *Operator_nameContext) OPERATOR_VALIDATE_BYTE_RANGE() antlr.TerminalNode {
	return s.GetToken(SecLangParserOPERATOR_VALIDATE_BYTE_RANGE, 0)
}

func (s *Operator_nameContext) OPERATOR_VALIDATE_DTD() antlr.TerminalNode {
	return s.GetToken(SecLangParserOPERATOR_VALIDATE_DTD, 0)
}

func (s *Operator_nameContext) OPERATOR_VALIDATE_HASH() antlr.TerminalNode {
	return s.GetToken(SecLangParserOPERATOR_VALIDATE_HASH, 0)
}

func (s *Operator_nameContext) OPERATOR_VALIDATE_SCHEMA() antlr.TerminalNode {
	return s.GetToken(SecLangParserOPERATOR_VALIDATE_SCHEMA, 0)
}

func (s *Operator_nameContext) OPERATOR_VERIFY_CC() antlr.TerminalNode {
	return s.GetToken(SecLangParserOPERATOR_VERIFY_CC, 0)
}

func (s *Operator_nameContext) OPERATOR_VERIFY_CPF() antlr.TerminalNode {
	return s.GetToken(SecLangParserOPERATOR_VERIFY_CPF, 0)
}

func (s *Operator_nameContext) OPERATOR_VERIFY_SSN() antlr.TerminalNode {
	return s.GetToken(SecLangParserOPERATOR_VERIFY_SSN, 0)
}

func (s *Operator_nameContext) OPERATOR_VERIFY_SVNR() antlr.TerminalNode {
	return s.GetToken(SecLangParserOPERATOR_VERIFY_SVNR, 0)
}

func (s *Operator_nameContext) OPERATOR_GSB_LOOKUP() antlr.TerminalNode {
	return s.GetToken(SecLangParserOPERATOR_GSB_LOOKUP, 0)
}

func (s *Operator_nameContext) OPERATOR_RSUB() antlr.TerminalNode {
	return s.GetToken(SecLangParserOPERATOR_RSUB, 0)
}

func (s *Operator_nameContext) OPERATOR_WITHIN() antlr.TerminalNode {
	return s.GetToken(SecLangParserOPERATOR_WITHIN, 0)
}

func (s *Operator_nameContext) OPERATOR_CONTAINS_WORD() antlr.TerminalNode {
	return s.GetToken(SecLangParserOPERATOR_CONTAINS_WORD, 0)
}

func (s *Operator_nameContext) OPERATOR_CONTAINS() antlr.TerminalNode {
	return s.GetToken(SecLangParserOPERATOR_CONTAINS, 0)
}

func (s *Operator_nameContext) OPERATOR_ENDS_WITH() antlr.TerminalNode {
	return s.GetToken(SecLangParserOPERATOR_ENDS_WITH, 0)
}

func (s *Operator_nameContext) OPERATOR_EQ() antlr.TerminalNode {
	return s.GetToken(SecLangParserOPERATOR_EQ, 0)
}

func (s *Operator_nameContext) OPERATOR_GE() antlr.TerminalNode {
	return s.GetToken(SecLangParserOPERATOR_GE, 0)
}

func (s *Operator_nameContext) OPERATOR_GT() antlr.TerminalNode {
	return s.GetToken(SecLangParserOPERATOR_GT, 0)
}

func (s *Operator_nameContext) OPERATOR_IP_MATCH_FROM_FILE() antlr.TerminalNode {
	return s.GetToken(SecLangParserOPERATOR_IP_MATCH_FROM_FILE, 0)
}

func (s *Operator_nameContext) OPERATOR_IP_MATCH() antlr.TerminalNode {
	return s.GetToken(SecLangParserOPERATOR_IP_MATCH, 0)
}

func (s *Operator_nameContext) OPERATOR_LE() antlr.TerminalNode {
	return s.GetToken(SecLangParserOPERATOR_LE, 0)
}

func (s *Operator_nameContext) OPERATOR_LT() antlr.TerminalNode {
	return s.GetToken(SecLangParserOPERATOR_LT, 0)
}

func (s *Operator_nameContext) OPERATOR_PM_FROM_FILE() antlr.TerminalNode {
	return s.GetToken(SecLangParserOPERATOR_PM_FROM_FILE, 0)
}

func (s *Operator_nameContext) OPERATOR_PM() antlr.TerminalNode {
	return s.GetToken(SecLangParserOPERATOR_PM, 0)
}

func (s *Operator_nameContext) OPERATOR_RBL() antlr.TerminalNode {
	return s.GetToken(SecLangParserOPERATOR_RBL, 0)
}

func (s *Operator_nameContext) OPERATOR_RX() antlr.TerminalNode {
	return s.GetToken(SecLangParserOPERATOR_RX, 0)
}

func (s *Operator_nameContext) OPERATOR_RX_GLOBAL() antlr.TerminalNode {
	return s.GetToken(SecLangParserOPERATOR_RX_GLOBAL, 0)
}

func (s *Operator_nameContext) OPERATOR_STR_EQ() antlr.TerminalNode {
	return s.GetToken(SecLangParserOPERATOR_STR_EQ, 0)
}

func (s *Operator_nameContext) OPERATOR_STR_MATCH() antlr.TerminalNode {
	return s.GetToken(SecLangParserOPERATOR_STR_MATCH, 0)
}

func (s *Operator_nameContext) OPERATOR_BEGINS_WITH() antlr.TerminalNode {
	return s.GetToken(SecLangParserOPERATOR_BEGINS_WITH, 0)
}

func (s *Operator_nameContext) OPERATOR_GEOLOOKUP() antlr.TerminalNode {
	return s.GetToken(SecLangParserOPERATOR_GEOLOOKUP, 0)
}

func (s *Operator_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Operator_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Operator_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterOperator_name(s)
	}
}

func (s *Operator_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitOperator_name(s)
	}
}

func (p *SecLangParser) Operator_name() (localctx IOperator_nameContext) {
	localctx = NewOperator_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, SecLangParserRULE_operator_name)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(438)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-89)) & ^0x3f) == 0 && ((int64(1)<<(_la-89))&137438953471) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOperator_valueContext is an interface to support dynamic dispatch.
type IOperator_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Variable_enum() IVariable_enumContext
	STRING() antlr.TerminalNode
	AllINT() []antlr.TerminalNode
	INT(i int) antlr.TerminalNode
	AllInt_range() []IInt_rangeContext
	Int_range(i int) IInt_rangeContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	OPERATOR_UNQUOTED_STRING() antlr.TerminalNode
	OPERATOR_QUOTED_STRING() antlr.TerminalNode

	// IsOperator_valueContext differentiates from other interfaces.
	IsOperator_valueContext()
}

type Operator_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperator_valueContext() *Operator_valueContext {
	var p = new(Operator_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_operator_value
	return p
}

func InitEmptyOperator_valueContext(p *Operator_valueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_operator_value
}

func (*Operator_valueContext) IsOperator_valueContext() {}

func NewOperator_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Operator_valueContext {
	var p = new(Operator_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_operator_value

	return p
}

func (s *Operator_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Operator_valueContext) Variable_enum() IVariable_enumContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariable_enumContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariable_enumContext)
}

func (s *Operator_valueContext) STRING() antlr.TerminalNode {
	return s.GetToken(SecLangParserSTRING, 0)
}

func (s *Operator_valueContext) AllINT() []antlr.TerminalNode {
	return s.GetTokens(SecLangParserINT)
}

func (s *Operator_valueContext) INT(i int) antlr.TerminalNode {
	return s.GetToken(SecLangParserINT, i)
}

func (s *Operator_valueContext) AllInt_range() []IInt_rangeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInt_rangeContext); ok {
			len++
		}
	}

	tst := make([]IInt_rangeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInt_rangeContext); ok {
			tst[i] = t.(IInt_rangeContext)
			i++
		}
	}

	return tst
}

func (s *Operator_valueContext) Int_range(i int) IInt_rangeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInt_rangeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInt_rangeContext)
}

func (s *Operator_valueContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SecLangParserCOMMA)
}

func (s *Operator_valueContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SecLangParserCOMMA, i)
}

func (s *Operator_valueContext) OPERATOR_UNQUOTED_STRING() antlr.TerminalNode {
	return s.GetToken(SecLangParserOPERATOR_UNQUOTED_STRING, 0)
}

func (s *Operator_valueContext) OPERATOR_QUOTED_STRING() antlr.TerminalNode {
	return s.GetToken(SecLangParserOPERATOR_QUOTED_STRING, 0)
}

func (s *Operator_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Operator_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Operator_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterOperator_value(s)
	}
}

func (s *Operator_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitOperator_value(s)
	}
}

func (p *SecLangParser) Operator_value() (localctx IOperator_valueContext) {
	localctx = NewOperator_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, SecLangParserRULE_operator_value)
	var _la int

	p.SetState(458)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SecLangParserVARIABLE_NAME_ENUM:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(440)
			p.Variable_enum()
		}

	case SecLangParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(441)
			p.Match(SecLangParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SecLangParserINT:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(444)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(442)
				p.Match(SecLangParserINT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case 2:
			{
				p.SetState(443)
				p.Int_range()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}
		p.SetState(453)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SecLangParserCOMMA {
			{
				p.SetState(446)
				p.Match(SecLangParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(449)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(447)
					p.Match(SecLangParserINT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 2:
				{
					p.SetState(448)
					p.Int_range()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

			p.SetState(455)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case SecLangParserOPERATOR_UNQUOTED_STRING:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(456)
			p.Match(SecLangParserOPERATOR_UNQUOTED_STRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SecLangParserOPERATOR_QUOTED_STRING:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(457)
			p.Match(SecLangParserOPERATOR_QUOTED_STRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVar_notContext is an interface to support dynamic dispatch.
type IVar_notContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NOT() antlr.TerminalNode

	// IsVar_notContext differentiates from other interfaces.
	IsVar_notContext()
}

type Var_notContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_notContext() *Var_notContext {
	var p = new(Var_notContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_var_not
	return p
}

func InitEmptyVar_notContext(p *Var_notContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_var_not
}

func (*Var_notContext) IsVar_notContext() {}

func NewVar_notContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_notContext {
	var p = new(Var_notContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_var_not

	return p
}

func (s *Var_notContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_notContext) NOT() antlr.TerminalNode {
	return s.GetToken(SecLangParserNOT, 0)
}

func (s *Var_notContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_notContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_notContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterVar_not(s)
	}
}

func (s *Var_notContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitVar_not(s)
	}
}

func (p *SecLangParser) Var_not() (localctx IVar_notContext) {
	localctx = NewVar_notContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, SecLangParserRULE_var_not)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(460)
		p.Match(SecLangParserNOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVar_countContext is an interface to support dynamic dispatch.
type IVar_countContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR_COUNT() antlr.TerminalNode

	// IsVar_countContext differentiates from other interfaces.
	IsVar_countContext()
}

type Var_countContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_countContext() *Var_countContext {
	var p = new(Var_countContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_var_count
	return p
}

func InitEmptyVar_countContext(p *Var_countContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_var_count
}

func (*Var_countContext) IsVar_countContext() {}

func NewVar_countContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_countContext {
	var p = new(Var_countContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_var_count

	return p
}

func (s *Var_countContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_countContext) VAR_COUNT() antlr.TerminalNode {
	return s.GetToken(SecLangParserVAR_COUNT, 0)
}

func (s *Var_countContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_countContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_countContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterVar_count(s)
	}
}

func (s *Var_countContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitVar_count(s)
	}
}

func (p *SecLangParser) Var_count() (localctx IVar_countContext) {
	localctx = NewVar_countContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, SecLangParserRULE_var_count)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(462)
		p.Match(SecLangParserVAR_COUNT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariablesContext is an interface to support dynamic dispatch.
type IVariablesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVar_stmt() []IVar_stmtContext
	Var_stmt(i int) IVar_stmtContext
	AllQUOTE() []antlr.TerminalNode
	QUOTE(i int) antlr.TerminalNode
	AllVar_not() []IVar_notContext
	Var_not(i int) IVar_notContext
	Var_count() IVar_countContext
	AllPIPE() []antlr.TerminalNode
	PIPE(i int) antlr.TerminalNode

	// IsVariablesContext differentiates from other interfaces.
	IsVariablesContext()
}

type VariablesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariablesContext() *VariablesContext {
	var p = new(VariablesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_variables
	return p
}

func InitEmptyVariablesContext(p *VariablesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_variables
}

func (*VariablesContext) IsVariablesContext() {}

func NewVariablesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariablesContext {
	var p = new(VariablesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_variables

	return p
}

func (s *VariablesContext) GetParser() antlr.Parser { return s.parser }

func (s *VariablesContext) AllVar_stmt() []IVar_stmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVar_stmtContext); ok {
			len++
		}
	}

	tst := make([]IVar_stmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVar_stmtContext); ok {
			tst[i] = t.(IVar_stmtContext)
			i++
		}
	}

	return tst
}

func (s *VariablesContext) Var_stmt(i int) IVar_stmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_stmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_stmtContext)
}

func (s *VariablesContext) AllQUOTE() []antlr.TerminalNode {
	return s.GetTokens(SecLangParserQUOTE)
}

func (s *VariablesContext) QUOTE(i int) antlr.TerminalNode {
	return s.GetToken(SecLangParserQUOTE, i)
}

func (s *VariablesContext) AllVar_not() []IVar_notContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVar_notContext); ok {
			len++
		}
	}

	tst := make([]IVar_notContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVar_notContext); ok {
			tst[i] = t.(IVar_notContext)
			i++
		}
	}

	return tst
}

func (s *VariablesContext) Var_not(i int) IVar_notContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_notContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_notContext)
}

func (s *VariablesContext) Var_count() IVar_countContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_countContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_countContext)
}

func (s *VariablesContext) AllPIPE() []antlr.TerminalNode {
	return s.GetTokens(SecLangParserPIPE)
}

func (s *VariablesContext) PIPE(i int) antlr.TerminalNode {
	return s.GetToken(SecLangParserPIPE, i)
}

func (s *VariablesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariablesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariablesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterVariables(s)
	}
}

func (s *VariablesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitVariables(s)
	}
}

func (p *SecLangParser) Variables() (localctx IVariablesContext) {
	localctx = NewVariablesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, SecLangParserRULE_variables)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(465)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SecLangParserQUOTE {
		{
			p.SetState(464)
			p.Match(SecLangParserQUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(468)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SecLangParserNOT {
		{
			p.SetState(467)
			p.Var_not()
		}

	}
	p.SetState(471)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SecLangParserVAR_COUNT {
		{
			p.SetState(470)
			p.Var_count()
		}

	}
	{
		p.SetState(473)
		p.Var_stmt()
	}
	p.SetState(475)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(474)
			p.Match(SecLangParserQUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(490)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SecLangParserPIPE {
		{
			p.SetState(477)
			p.Match(SecLangParserPIPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(479)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SecLangParserQUOTE {
			{
				p.SetState(478)
				p.Match(SecLangParserQUOTE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(482)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SecLangParserNOT {
			{
				p.SetState(481)
				p.Var_not()
			}

		}
		{
			p.SetState(484)
			p.Var_stmt()
		}
		p.SetState(486)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 48, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(485)
				p.Match(SecLangParserQUOTE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

		p.SetState(492)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUpdate_variablesContext is an interface to support dynamic dispatch.
type IUpdate_variablesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVar_stmt() []IVar_stmtContext
	Var_stmt(i int) IVar_stmtContext
	AllQUOTE() []antlr.TerminalNode
	QUOTE(i int) antlr.TerminalNode
	AllVar_not() []IVar_notContext
	Var_not(i int) IVar_notContext
	Var_count() IVar_countContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsUpdate_variablesContext differentiates from other interfaces.
	IsUpdate_variablesContext()
}

type Update_variablesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdate_variablesContext() *Update_variablesContext {
	var p = new(Update_variablesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_update_variables
	return p
}

func InitEmptyUpdate_variablesContext(p *Update_variablesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_update_variables
}

func (*Update_variablesContext) IsUpdate_variablesContext() {}

func NewUpdate_variablesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Update_variablesContext {
	var p = new(Update_variablesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_update_variables

	return p
}

func (s *Update_variablesContext) GetParser() antlr.Parser { return s.parser }

func (s *Update_variablesContext) AllVar_stmt() []IVar_stmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVar_stmtContext); ok {
			len++
		}
	}

	tst := make([]IVar_stmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVar_stmtContext); ok {
			tst[i] = t.(IVar_stmtContext)
			i++
		}
	}

	return tst
}

func (s *Update_variablesContext) Var_stmt(i int) IVar_stmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_stmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_stmtContext)
}

func (s *Update_variablesContext) AllQUOTE() []antlr.TerminalNode {
	return s.GetTokens(SecLangParserQUOTE)
}

func (s *Update_variablesContext) QUOTE(i int) antlr.TerminalNode {
	return s.GetToken(SecLangParserQUOTE, i)
}

func (s *Update_variablesContext) AllVar_not() []IVar_notContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVar_notContext); ok {
			len++
		}
	}

	tst := make([]IVar_notContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVar_notContext); ok {
			tst[i] = t.(IVar_notContext)
			i++
		}
	}

	return tst
}

func (s *Update_variablesContext) Var_not(i int) IVar_notContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_notContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_notContext)
}

func (s *Update_variablesContext) Var_count() IVar_countContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_countContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_countContext)
}

func (s *Update_variablesContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SecLangParserCOMMA)
}

func (s *Update_variablesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SecLangParserCOMMA, i)
}

func (s *Update_variablesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Update_variablesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Update_variablesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterUpdate_variables(s)
	}
}

func (s *Update_variablesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitUpdate_variables(s)
	}
}

func (p *SecLangParser) Update_variables() (localctx IUpdate_variablesContext) {
	localctx = NewUpdate_variablesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, SecLangParserRULE_update_variables)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(494)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SecLangParserQUOTE {
		{
			p.SetState(493)
			p.Match(SecLangParserQUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(497)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SecLangParserNOT {
		{
			p.SetState(496)
			p.Var_not()
		}

	}
	p.SetState(500)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SecLangParserVAR_COUNT {
		{
			p.SetState(499)
			p.Var_count()
		}

	}
	{
		p.SetState(502)
		p.Var_stmt()
	}
	p.SetState(504)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 53, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(503)
			p.Match(SecLangParserQUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(519)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SecLangParserCOMMA {
		{
			p.SetState(506)
			p.Match(SecLangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(508)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SecLangParserQUOTE {
			{
				p.SetState(507)
				p.Match(SecLangParserQUOTE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(511)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SecLangParserNOT {
			{
				p.SetState(510)
				p.Var_not()
			}

		}
		{
			p.SetState(513)
			p.Var_stmt()
		}
		p.SetState(515)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 56, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(514)
				p.Match(SecLangParserQUOTE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

		p.SetState(521)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INew_targetContext is an interface to support dynamic dispatch.
type INew_targetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Var_stmt() IVar_stmtContext

	// IsNew_targetContext differentiates from other interfaces.
	IsNew_targetContext()
}

type New_targetContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNew_targetContext() *New_targetContext {
	var p = new(New_targetContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_new_target
	return p
}

func InitEmptyNew_targetContext(p *New_targetContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_new_target
}

func (*New_targetContext) IsNew_targetContext() {}

func NewNew_targetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *New_targetContext {
	var p = new(New_targetContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_new_target

	return p
}

func (s *New_targetContext) GetParser() antlr.Parser { return s.parser }

func (s *New_targetContext) Var_stmt() IVar_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_stmtContext)
}

func (s *New_targetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *New_targetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *New_targetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterNew_target(s)
	}
}

func (s *New_targetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitNew_target(s)
	}
}

func (p *SecLangParser) New_target() (localctx INew_targetContext) {
	localctx = NewNew_targetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, SecLangParserRULE_new_target)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(522)
		p.Var_stmt()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVar_stmtContext is an interface to support dynamic dispatch.
type IVar_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Variable_enum() IVariable_enumContext
	Collection_enum() ICollection_enumContext
	COLON() antlr.TerminalNode
	Collection_value() ICollection_valueContext

	// IsVar_stmtContext differentiates from other interfaces.
	IsVar_stmtContext()
}

type Var_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_stmtContext() *Var_stmtContext {
	var p = new(Var_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_var_stmt
	return p
}

func InitEmptyVar_stmtContext(p *Var_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_var_stmt
}

func (*Var_stmtContext) IsVar_stmtContext() {}

func NewVar_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_stmtContext {
	var p = new(Var_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_var_stmt

	return p
}

func (s *Var_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_stmtContext) Variable_enum() IVariable_enumContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariable_enumContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariable_enumContext)
}

func (s *Var_stmtContext) Collection_enum() ICollection_enumContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICollection_enumContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICollection_enumContext)
}

func (s *Var_stmtContext) COLON() antlr.TerminalNode {
	return s.GetToken(SecLangParserCOLON, 0)
}

func (s *Var_stmtContext) Collection_value() ICollection_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICollection_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICollection_valueContext)
}

func (s *Var_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterVar_stmt(s)
	}
}

func (s *Var_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitVar_stmt(s)
	}
}

func (p *SecLangParser) Var_stmt() (localctx IVar_stmtContext) {
	localctx = NewVar_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, SecLangParserRULE_var_stmt)
	var _la int

	p.SetState(530)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SecLangParserVARIABLE_NAME_ENUM:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(524)
			p.Variable_enum()
		}

	case SecLangParserCOLLECTION_NAME_ENUM, SecLangParserRUN_TIME_VAR_XML:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(525)
			p.Collection_enum()
		}
		p.SetState(528)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SecLangParserCOLON {
			{
				p.SetState(526)
				p.Match(SecLangParserCOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(527)
				p.Collection_value()
			}

		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariable_enumContext is an interface to support dynamic dispatch.
type IVariable_enumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VARIABLE_NAME_ENUM() antlr.TerminalNode

	// IsVariable_enumContext differentiates from other interfaces.
	IsVariable_enumContext()
}

type Variable_enumContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariable_enumContext() *Variable_enumContext {
	var p = new(Variable_enumContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_variable_enum
	return p
}

func InitEmptyVariable_enumContext(p *Variable_enumContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_variable_enum
}

func (*Variable_enumContext) IsVariable_enumContext() {}

func NewVariable_enumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Variable_enumContext {
	var p = new(Variable_enumContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_variable_enum

	return p
}

func (s *Variable_enumContext) GetParser() antlr.Parser { return s.parser }

func (s *Variable_enumContext) VARIABLE_NAME_ENUM() antlr.TerminalNode {
	return s.GetToken(SecLangParserVARIABLE_NAME_ENUM, 0)
}

func (s *Variable_enumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Variable_enumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Variable_enumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterVariable_enum(s)
	}
}

func (s *Variable_enumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitVariable_enum(s)
	}
}

func (p *SecLangParser) Variable_enum() (localctx IVariable_enumContext) {
	localctx = NewVariable_enumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, SecLangParserRULE_variable_enum)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(532)
		p.Match(SecLangParserVARIABLE_NAME_ENUM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICtl_variable_enumContext is an interface to support dynamic dispatch.
type ICtl_variable_enumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VARIABLE_NAME_ENUM() antlr.TerminalNode

	// IsCtl_variable_enumContext differentiates from other interfaces.
	IsCtl_variable_enumContext()
}

type Ctl_variable_enumContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCtl_variable_enumContext() *Ctl_variable_enumContext {
	var p = new(Ctl_variable_enumContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_ctl_variable_enum
	return p
}

func InitEmptyCtl_variable_enumContext(p *Ctl_variable_enumContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_ctl_variable_enum
}

func (*Ctl_variable_enumContext) IsCtl_variable_enumContext() {}

func NewCtl_variable_enumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ctl_variable_enumContext {
	var p = new(Ctl_variable_enumContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_ctl_variable_enum

	return p
}

func (s *Ctl_variable_enumContext) GetParser() antlr.Parser { return s.parser }

func (s *Ctl_variable_enumContext) VARIABLE_NAME_ENUM() antlr.TerminalNode {
	return s.GetToken(SecLangParserVARIABLE_NAME_ENUM, 0)
}

func (s *Ctl_variable_enumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ctl_variable_enumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ctl_variable_enumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterCtl_variable_enum(s)
	}
}

func (s *Ctl_variable_enumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitCtl_variable_enum(s)
	}
}

func (p *SecLangParser) Ctl_variable_enum() (localctx ICtl_variable_enumContext) {
	localctx = NewCtl_variable_enumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, SecLangParserRULE_ctl_variable_enum)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(534)
		p.Match(SecLangParserVARIABLE_NAME_ENUM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICollection_enumContext is an interface to support dynamic dispatch.
type ICollection_enumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COLLECTION_NAME_ENUM() antlr.TerminalNode
	RUN_TIME_VAR_XML() antlr.TerminalNode

	// IsCollection_enumContext differentiates from other interfaces.
	IsCollection_enumContext()
}

type Collection_enumContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCollection_enumContext() *Collection_enumContext {
	var p = new(Collection_enumContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_collection_enum
	return p
}

func InitEmptyCollection_enumContext(p *Collection_enumContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_collection_enum
}

func (*Collection_enumContext) IsCollection_enumContext() {}

func NewCollection_enumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Collection_enumContext {
	var p = new(Collection_enumContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_collection_enum

	return p
}

func (s *Collection_enumContext) GetParser() antlr.Parser { return s.parser }

func (s *Collection_enumContext) COLLECTION_NAME_ENUM() antlr.TerminalNode {
	return s.GetToken(SecLangParserCOLLECTION_NAME_ENUM, 0)
}

func (s *Collection_enumContext) RUN_TIME_VAR_XML() antlr.TerminalNode {
	return s.GetToken(SecLangParserRUN_TIME_VAR_XML, 0)
}

func (s *Collection_enumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Collection_enumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Collection_enumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterCollection_enum(s)
	}
}

func (s *Collection_enumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitCollection_enum(s)
	}
}

func (p *SecLangParser) Collection_enum() (localctx ICollection_enumContext) {
	localctx = NewCollection_enumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, SecLangParserRULE_collection_enum)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(536)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SecLangParserCOLLECTION_NAME_ENUM || _la == SecLangParserRUN_TIME_VAR_XML) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICtl_collection_enumContext is an interface to support dynamic dispatch.
type ICtl_collection_enumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COLLECTION_NAME_ENUM() antlr.TerminalNode
	RUN_TIME_VAR_XML() antlr.TerminalNode

	// IsCtl_collection_enumContext differentiates from other interfaces.
	IsCtl_collection_enumContext()
}

type Ctl_collection_enumContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCtl_collection_enumContext() *Ctl_collection_enumContext {
	var p = new(Ctl_collection_enumContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_ctl_collection_enum
	return p
}

func InitEmptyCtl_collection_enumContext(p *Ctl_collection_enumContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_ctl_collection_enum
}

func (*Ctl_collection_enumContext) IsCtl_collection_enumContext() {}

func NewCtl_collection_enumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ctl_collection_enumContext {
	var p = new(Ctl_collection_enumContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_ctl_collection_enum

	return p
}

func (s *Ctl_collection_enumContext) GetParser() antlr.Parser { return s.parser }

func (s *Ctl_collection_enumContext) COLLECTION_NAME_ENUM() antlr.TerminalNode {
	return s.GetToken(SecLangParserCOLLECTION_NAME_ENUM, 0)
}

func (s *Ctl_collection_enumContext) RUN_TIME_VAR_XML() antlr.TerminalNode {
	return s.GetToken(SecLangParserRUN_TIME_VAR_XML, 0)
}

func (s *Ctl_collection_enumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ctl_collection_enumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ctl_collection_enumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterCtl_collection_enum(s)
	}
}

func (s *Ctl_collection_enumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitCtl_collection_enum(s)
	}
}

func (p *SecLangParser) Ctl_collection_enum() (localctx ICtl_collection_enumContext) {
	localctx = NewCtl_collection_enumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, SecLangParserRULE_ctl_collection_enum)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(538)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SecLangParserCOLLECTION_NAME_ENUM || _la == SecLangParserRUN_TIME_VAR_XML) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IActionsContext is an interface to support dynamic dispatch.
type IActionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQUOTE() []antlr.TerminalNode
	QUOTE(i int) antlr.TerminalNode
	AllAction_() []IActionContext
	Action_(i int) IActionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsActionsContext differentiates from other interfaces.
	IsActionsContext()
}

type ActionsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActionsContext() *ActionsContext {
	var p = new(ActionsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_actions
	return p
}

func InitEmptyActionsContext(p *ActionsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_actions
}

func (*ActionsContext) IsActionsContext() {}

func NewActionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActionsContext {
	var p = new(ActionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_actions

	return p
}

func (s *ActionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ActionsContext) AllQUOTE() []antlr.TerminalNode {
	return s.GetTokens(SecLangParserQUOTE)
}

func (s *ActionsContext) QUOTE(i int) antlr.TerminalNode {
	return s.GetToken(SecLangParserQUOTE, i)
}

func (s *ActionsContext) AllAction_() []IActionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IActionContext); ok {
			len++
		}
	}

	tst := make([]IActionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IActionContext); ok {
			tst[i] = t.(IActionContext)
			i++
		}
	}

	return tst
}

func (s *ActionsContext) Action_(i int) IActionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionContext)
}

func (s *ActionsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SecLangParserCOMMA)
}

func (s *ActionsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SecLangParserCOMMA, i)
}

func (s *ActionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterActions(s)
	}
}

func (s *ActionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitActions(s)
	}
}

func (p *SecLangParser) Actions() (localctx IActionsContext) {
	localctx = NewActionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, SecLangParserRULE_actions)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(540)
		p.Match(SecLangParserQUOTE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(541)
		p.Action_()
	}
	p.SetState(546)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SecLangParserCOMMA {
		{
			p.SetState(542)
			p.Match(SecLangParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(543)
			p.Action_()
		}

		p.SetState(548)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(549)
		p.Match(SecLangParserQUOTE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IActionContext is an interface to support dynamic dispatch.
type IActionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Action_with_params() IAction_with_paramsContext
	COLON() antlr.TerminalNode
	Action_value() IAction_valueContext
	NOT() antlr.TerminalNode
	EQUAL() antlr.TerminalNode
	ACTION_TRANSFORMATION() antlr.TerminalNode
	Transformation_action_value() ITransformation_action_valueContext
	Action_only() IAction_onlyContext

	// IsActionContext differentiates from other interfaces.
	IsActionContext()
}

type ActionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActionContext() *ActionContext {
	var p = new(ActionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_action
	return p
}

func InitEmptyActionContext(p *ActionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_action
}

func (*ActionContext) IsActionContext() {}

func NewActionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActionContext {
	var p = new(ActionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_action

	return p
}

func (s *ActionContext) GetParser() antlr.Parser { return s.parser }

func (s *ActionContext) Action_with_params() IAction_with_paramsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAction_with_paramsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAction_with_paramsContext)
}

func (s *ActionContext) COLON() antlr.TerminalNode {
	return s.GetToken(SecLangParserCOLON, 0)
}

func (s *ActionContext) Action_value() IAction_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAction_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAction_valueContext)
}

func (s *ActionContext) NOT() antlr.TerminalNode {
	return s.GetToken(SecLangParserNOT, 0)
}

func (s *ActionContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(SecLangParserEQUAL, 0)
}

func (s *ActionContext) ACTION_TRANSFORMATION() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_TRANSFORMATION, 0)
}

func (s *ActionContext) Transformation_action_value() ITransformation_action_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITransformation_action_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITransformation_action_valueContext)
}

func (s *ActionContext) Action_only() IAction_onlyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAction_onlyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAction_onlyContext)
}

func (s *ActionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterAction(s)
	}
}

func (s *ActionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitAction(s)
	}
}

func (p *SecLangParser) Action_() (localctx IActionContext) {
	localctx = NewActionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, SecLangParserRULE_action)
	var _la int

	p.SetState(569)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 63, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(551)
			p.Action_with_params()
		}
		{
			p.SetState(552)
			p.Match(SecLangParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(554)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SecLangParserNOT {
			{
				p.SetState(553)
				p.Match(SecLangParserNOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(557)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SecLangParserEQUAL {
			{
				p.SetState(556)
				p.Match(SecLangParserEQUAL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(559)
			p.Action_value()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(561)
			p.Action_with_params()
		}
		{
			p.SetState(562)
			p.Match(SecLangParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(563)
			p.Action_value()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(565)
			p.Match(SecLangParserACTION_TRANSFORMATION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(566)
			p.Match(SecLangParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(567)
			p.Transformation_action_value()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(568)
			p.Action_only()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAction_onlyContext is an interface to support dynamic dispatch.
type IAction_onlyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Disruptive_action_only() IDisruptive_action_onlyContext
	Non_disruptive_action_only() INon_disruptive_action_onlyContext
	Flow_action_only() IFlow_action_onlyContext

	// IsAction_onlyContext differentiates from other interfaces.
	IsAction_onlyContext()
}

type Action_onlyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAction_onlyContext() *Action_onlyContext {
	var p = new(Action_onlyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_action_only
	return p
}

func InitEmptyAction_onlyContext(p *Action_onlyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_action_only
}

func (*Action_onlyContext) IsAction_onlyContext() {}

func NewAction_onlyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Action_onlyContext {
	var p = new(Action_onlyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_action_only

	return p
}

func (s *Action_onlyContext) GetParser() antlr.Parser { return s.parser }

func (s *Action_onlyContext) Disruptive_action_only() IDisruptive_action_onlyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDisruptive_action_onlyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDisruptive_action_onlyContext)
}

func (s *Action_onlyContext) Non_disruptive_action_only() INon_disruptive_action_onlyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INon_disruptive_action_onlyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INon_disruptive_action_onlyContext)
}

func (s *Action_onlyContext) Flow_action_only() IFlow_action_onlyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFlow_action_onlyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFlow_action_onlyContext)
}

func (s *Action_onlyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Action_onlyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Action_onlyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterAction_only(s)
	}
}

func (s *Action_onlyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitAction_only(s)
	}
}

func (p *SecLangParser) Action_only() (localctx IAction_onlyContext) {
	localctx = NewAction_onlyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, SecLangParserRULE_action_only)
	p.SetState(574)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SecLangParserACTION_ALLOW, SecLangParserACTION_BLOCK, SecLangParserACTION_DENY, SecLangParserACTION_DROP, SecLangParserACTION_PASS, SecLangParserACTION_PAUSE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(571)
			p.Disruptive_action_only()
		}

	case SecLangParserACTION_AUDIT_LOG, SecLangParserACTION_CAPTURE, SecLangParserACTION_LOG, SecLangParserACTION_MULTI_MATCH, SecLangParserACTION_NO_AUDIT_LOG, SecLangParserACTION_NO_LOG, SecLangParserACTION_SANITISE_MATCHED:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(572)
			p.Non_disruptive_action_only()
		}

	case SecLangParserACTION_CHAIN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(573)
			p.Flow_action_only()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDisruptive_action_onlyContext is an interface to support dynamic dispatch.
type IDisruptive_action_onlyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ACTION_ALLOW() antlr.TerminalNode
	ACTION_BLOCK() antlr.TerminalNode
	ACTION_DENY() antlr.TerminalNode
	ACTION_DROP() antlr.TerminalNode
	ACTION_PASS() antlr.TerminalNode
	ACTION_PAUSE() antlr.TerminalNode

	// IsDisruptive_action_onlyContext differentiates from other interfaces.
	IsDisruptive_action_onlyContext()
}

type Disruptive_action_onlyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDisruptive_action_onlyContext() *Disruptive_action_onlyContext {
	var p = new(Disruptive_action_onlyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_disruptive_action_only
	return p
}

func InitEmptyDisruptive_action_onlyContext(p *Disruptive_action_onlyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_disruptive_action_only
}

func (*Disruptive_action_onlyContext) IsDisruptive_action_onlyContext() {}

func NewDisruptive_action_onlyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Disruptive_action_onlyContext {
	var p = new(Disruptive_action_onlyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_disruptive_action_only

	return p
}

func (s *Disruptive_action_onlyContext) GetParser() antlr.Parser { return s.parser }

func (s *Disruptive_action_onlyContext) ACTION_ALLOW() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_ALLOW, 0)
}

func (s *Disruptive_action_onlyContext) ACTION_BLOCK() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_BLOCK, 0)
}

func (s *Disruptive_action_onlyContext) ACTION_DENY() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_DENY, 0)
}

func (s *Disruptive_action_onlyContext) ACTION_DROP() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_DROP, 0)
}

func (s *Disruptive_action_onlyContext) ACTION_PASS() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_PASS, 0)
}

func (s *Disruptive_action_onlyContext) ACTION_PAUSE() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_PAUSE, 0)
}

func (s *Disruptive_action_onlyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Disruptive_action_onlyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Disruptive_action_onlyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterDisruptive_action_only(s)
	}
}

func (s *Disruptive_action_onlyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitDisruptive_action_only(s)
	}
}

func (p *SecLangParser) Disruptive_action_only() (localctx IDisruptive_action_onlyContext) {
	localctx = NewDisruptive_action_onlyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, SecLangParserRULE_disruptive_action_only)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(576)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&864779090593316864) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INon_disruptive_action_onlyContext is an interface to support dynamic dispatch.
type INon_disruptive_action_onlyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ACTION_AUDIT_LOG() antlr.TerminalNode
	ACTION_CAPTURE() antlr.TerminalNode
	ACTION_SANITISE_MATCHED() antlr.TerminalNode
	ACTION_LOG() antlr.TerminalNode
	ACTION_MULTI_MATCH() antlr.TerminalNode
	ACTION_NO_AUDIT_LOG() antlr.TerminalNode
	ACTION_NO_LOG() antlr.TerminalNode

	// IsNon_disruptive_action_onlyContext differentiates from other interfaces.
	IsNon_disruptive_action_onlyContext()
}

type Non_disruptive_action_onlyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNon_disruptive_action_onlyContext() *Non_disruptive_action_onlyContext {
	var p = new(Non_disruptive_action_onlyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_non_disruptive_action_only
	return p
}

func InitEmptyNon_disruptive_action_onlyContext(p *Non_disruptive_action_onlyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_non_disruptive_action_only
}

func (*Non_disruptive_action_onlyContext) IsNon_disruptive_action_onlyContext() {}

func NewNon_disruptive_action_onlyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Non_disruptive_action_onlyContext {
	var p = new(Non_disruptive_action_onlyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_non_disruptive_action_only

	return p
}

func (s *Non_disruptive_action_onlyContext) GetParser() antlr.Parser { return s.parser }

func (s *Non_disruptive_action_onlyContext) ACTION_AUDIT_LOG() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_AUDIT_LOG, 0)
}

func (s *Non_disruptive_action_onlyContext) ACTION_CAPTURE() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_CAPTURE, 0)
}

func (s *Non_disruptive_action_onlyContext) ACTION_SANITISE_MATCHED() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_SANITISE_MATCHED, 0)
}

func (s *Non_disruptive_action_onlyContext) ACTION_LOG() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_LOG, 0)
}

func (s *Non_disruptive_action_onlyContext) ACTION_MULTI_MATCH() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_MULTI_MATCH, 0)
}

func (s *Non_disruptive_action_onlyContext) ACTION_NO_AUDIT_LOG() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_NO_AUDIT_LOG, 0)
}

func (s *Non_disruptive_action_onlyContext) ACTION_NO_LOG() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_NO_LOG, 0)
}

func (s *Non_disruptive_action_onlyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Non_disruptive_action_onlyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Non_disruptive_action_onlyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterNon_disruptive_action_only(s)
	}
}

func (s *Non_disruptive_action_onlyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitNon_disruptive_action_only(s)
	}
}

func (p *SecLangParser) Non_disruptive_action_only() (localctx INon_disruptive_action_onlyContext) {
	localctx = NewNon_disruptive_action_onlyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, SecLangParserRULE_non_disruptive_action_only)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(578)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-29)) & ^0x3f) == 0 && ((int64(1)<<(_la-29))&275356057605) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFlow_action_onlyContext is an interface to support dynamic dispatch.
type IFlow_action_onlyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ACTION_CHAIN() antlr.TerminalNode

	// IsFlow_action_onlyContext differentiates from other interfaces.
	IsFlow_action_onlyContext()
}

type Flow_action_onlyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFlow_action_onlyContext() *Flow_action_onlyContext {
	var p = new(Flow_action_onlyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_flow_action_only
	return p
}

func InitEmptyFlow_action_onlyContext(p *Flow_action_onlyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_flow_action_only
}

func (*Flow_action_onlyContext) IsFlow_action_onlyContext() {}

func NewFlow_action_onlyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Flow_action_onlyContext {
	var p = new(Flow_action_onlyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_flow_action_only

	return p
}

func (s *Flow_action_onlyContext) GetParser() antlr.Parser { return s.parser }

func (s *Flow_action_onlyContext) ACTION_CHAIN() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_CHAIN, 0)
}

func (s *Flow_action_onlyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Flow_action_onlyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Flow_action_onlyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterFlow_action_only(s)
	}
}

func (s *Flow_action_onlyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitFlow_action_only(s)
	}
}

func (p *SecLangParser) Flow_action_only() (localctx IFlow_action_onlyContext) {
	localctx = NewFlow_action_onlyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, SecLangParserRULE_flow_action_only)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(580)
		p.Match(SecLangParserACTION_CHAIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAction_with_paramsContext is an interface to support dynamic dispatch.
type IAction_with_paramsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Metadata_action_with_params() IMetadata_action_with_paramsContext
	Disruptive_action_with_params() IDisruptive_action_with_paramsContext
	Non_disruptive_action_with_params() INon_disruptive_action_with_paramsContext
	Flow_action_with_params() IFlow_action_with_paramsContext
	Data_action_with_params() IData_action_with_paramsContext

	// IsAction_with_paramsContext differentiates from other interfaces.
	IsAction_with_paramsContext()
}

type Action_with_paramsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAction_with_paramsContext() *Action_with_paramsContext {
	var p = new(Action_with_paramsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_action_with_params
	return p
}

func InitEmptyAction_with_paramsContext(p *Action_with_paramsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_action_with_params
}

func (*Action_with_paramsContext) IsAction_with_paramsContext() {}

func NewAction_with_paramsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Action_with_paramsContext {
	var p = new(Action_with_paramsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_action_with_params

	return p
}

func (s *Action_with_paramsContext) GetParser() antlr.Parser { return s.parser }

func (s *Action_with_paramsContext) Metadata_action_with_params() IMetadata_action_with_paramsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMetadata_action_with_paramsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMetadata_action_with_paramsContext)
}

func (s *Action_with_paramsContext) Disruptive_action_with_params() IDisruptive_action_with_paramsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDisruptive_action_with_paramsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDisruptive_action_with_paramsContext)
}

func (s *Action_with_paramsContext) Non_disruptive_action_with_params() INon_disruptive_action_with_paramsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INon_disruptive_action_with_paramsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INon_disruptive_action_with_paramsContext)
}

func (s *Action_with_paramsContext) Flow_action_with_params() IFlow_action_with_paramsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFlow_action_with_paramsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFlow_action_with_paramsContext)
}

func (s *Action_with_paramsContext) Data_action_with_params() IData_action_with_paramsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IData_action_with_paramsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IData_action_with_paramsContext)
}

func (s *Action_with_paramsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Action_with_paramsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Action_with_paramsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterAction_with_params(s)
	}
}

func (s *Action_with_paramsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitAction_with_params(s)
	}
}

func (p *SecLangParser) Action_with_params() (localctx IAction_with_paramsContext) {
	localctx = NewAction_with_paramsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, SecLangParserRULE_action_with_params)
	p.SetState(587)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SecLangParserACTION_ID, SecLangParserACTION_MATURITY, SecLangParserACTION_MSG, SecLangParserACTION_PHASE, SecLangParserACTION_REV, SecLangParserACTION_SEVERITY, SecLangParserACTION_TAG, SecLangParserACTION_VER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(582)
			p.Metadata_action_with_params()
		}

	case SecLangParserACTION_PROXY, SecLangParserACTION_REDIRECT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(583)
			p.Disruptive_action_with_params()
		}

	case SecLangParserACTION_APPEND, SecLangParserACTION_CTL, SecLangParserACTION_DEPRECATE_VAR, SecLangParserACTION_EXEC, SecLangParserACTION_EXPIRE_VAR, SecLangParserACTION_INITCOL, SecLangParserACTION_LOG_DATA, SecLangParserACTION_PREPEND, SecLangParserACTION_SANITISE_ARG, SecLangParserACTION_SANITISE_MATCHED_BYTES, SecLangParserACTION_SANITISE_REQUEST_HEADER, SecLangParserACTION_SANITISE_RESPONSE_HEADER, SecLangParserACTION_SETENV, SecLangParserACTION_SETRSC, SecLangParserACTION_SETSID, SecLangParserACTION_SETUID, SecLangParserACTION_SETVAR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(584)
			p.Non_disruptive_action_with_params()
		}

	case SecLangParserACTION_SKIP_AFTER, SecLangParserACTION_SKIP:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(585)
			p.Flow_action_with_params()
		}

	case SecLangParserACTION_STATUS, SecLangParserACTION_XMLNS:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(586)
			p.Data_action_with_params()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMetadata_action_with_paramsContext is an interface to support dynamic dispatch.
type IMetadata_action_with_paramsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsMetadata_action_with_paramsContext differentiates from other interfaces.
	IsMetadata_action_with_paramsContext()
}

type Metadata_action_with_paramsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMetadata_action_with_paramsContext() *Metadata_action_with_paramsContext {
	var p = new(Metadata_action_with_paramsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_metadata_action_with_params
	return p
}

func InitEmptyMetadata_action_with_paramsContext(p *Metadata_action_with_paramsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_metadata_action_with_params
}

func (*Metadata_action_with_paramsContext) IsMetadata_action_with_paramsContext() {}

func NewMetadata_action_with_paramsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Metadata_action_with_paramsContext {
	var p = new(Metadata_action_with_paramsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_metadata_action_with_params

	return p
}

func (s *Metadata_action_with_paramsContext) GetParser() antlr.Parser { return s.parser }

func (s *Metadata_action_with_paramsContext) CopyAll(ctx *Metadata_action_with_paramsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Metadata_action_with_paramsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Metadata_action_with_paramsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ACTION_MATURITYContext struct {
	Metadata_action_with_paramsContext
}

func NewACTION_MATURITYContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ACTION_MATURITYContext {
	var p = new(ACTION_MATURITYContext)

	InitEmptyMetadata_action_with_paramsContext(&p.Metadata_action_with_paramsContext)
	p.parser = parser
	p.CopyAll(ctx.(*Metadata_action_with_paramsContext))

	return p
}

func (s *ACTION_MATURITYContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ACTION_MATURITYContext) ACTION_MATURITY() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_MATURITY, 0)
}

func (s *ACTION_MATURITYContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterACTION_MATURITY(s)
	}
}

func (s *ACTION_MATURITYContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitACTION_MATURITY(s)
	}
}

type ACTION_REVContext struct {
	Metadata_action_with_paramsContext
}

func NewACTION_REVContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ACTION_REVContext {
	var p = new(ACTION_REVContext)

	InitEmptyMetadata_action_with_paramsContext(&p.Metadata_action_with_paramsContext)
	p.parser = parser
	p.CopyAll(ctx.(*Metadata_action_with_paramsContext))

	return p
}

func (s *ACTION_REVContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ACTION_REVContext) ACTION_REV() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_REV, 0)
}

func (s *ACTION_REVContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterACTION_REV(s)
	}
}

func (s *ACTION_REVContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitACTION_REV(s)
	}
}

type ACTION_VERContext struct {
	Metadata_action_with_paramsContext
}

func NewACTION_VERContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ACTION_VERContext {
	var p = new(ACTION_VERContext)

	InitEmptyMetadata_action_with_paramsContext(&p.Metadata_action_with_paramsContext)
	p.parser = parser
	p.CopyAll(ctx.(*Metadata_action_with_paramsContext))

	return p
}

func (s *ACTION_VERContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ACTION_VERContext) ACTION_VER() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_VER, 0)
}

func (s *ACTION_VERContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterACTION_VER(s)
	}
}

func (s *ACTION_VERContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitACTION_VER(s)
	}
}

type ACTION_SEVERITYContext struct {
	Metadata_action_with_paramsContext
}

func NewACTION_SEVERITYContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ACTION_SEVERITYContext {
	var p = new(ACTION_SEVERITYContext)

	InitEmptyMetadata_action_with_paramsContext(&p.Metadata_action_with_paramsContext)
	p.parser = parser
	p.CopyAll(ctx.(*Metadata_action_with_paramsContext))

	return p
}

func (s *ACTION_SEVERITYContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ACTION_SEVERITYContext) ACTION_SEVERITY() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_SEVERITY, 0)
}

func (s *ACTION_SEVERITYContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterACTION_SEVERITY(s)
	}
}

func (s *ACTION_SEVERITYContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitACTION_SEVERITY(s)
	}
}

type ACTION_MSGContext struct {
	Metadata_action_with_paramsContext
}

func NewACTION_MSGContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ACTION_MSGContext {
	var p = new(ACTION_MSGContext)

	InitEmptyMetadata_action_with_paramsContext(&p.Metadata_action_with_paramsContext)
	p.parser = parser
	p.CopyAll(ctx.(*Metadata_action_with_paramsContext))

	return p
}

func (s *ACTION_MSGContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ACTION_MSGContext) ACTION_MSG() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_MSG, 0)
}

func (s *ACTION_MSGContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterACTION_MSG(s)
	}
}

func (s *ACTION_MSGContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitACTION_MSG(s)
	}
}

type ACTION_PHASEContext struct {
	Metadata_action_with_paramsContext
}

func NewACTION_PHASEContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ACTION_PHASEContext {
	var p = new(ACTION_PHASEContext)

	InitEmptyMetadata_action_with_paramsContext(&p.Metadata_action_with_paramsContext)
	p.parser = parser
	p.CopyAll(ctx.(*Metadata_action_with_paramsContext))

	return p
}

func (s *ACTION_PHASEContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ACTION_PHASEContext) ACTION_PHASE() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_PHASE, 0)
}

func (s *ACTION_PHASEContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterACTION_PHASE(s)
	}
}

func (s *ACTION_PHASEContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitACTION_PHASE(s)
	}
}

type ACTION_IDContext struct {
	Metadata_action_with_paramsContext
}

func NewACTION_IDContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ACTION_IDContext {
	var p = new(ACTION_IDContext)

	InitEmptyMetadata_action_with_paramsContext(&p.Metadata_action_with_paramsContext)
	p.parser = parser
	p.CopyAll(ctx.(*Metadata_action_with_paramsContext))

	return p
}

func (s *ACTION_IDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ACTION_IDContext) ACTION_ID() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_ID, 0)
}

func (s *ACTION_IDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterACTION_ID(s)
	}
}

func (s *ACTION_IDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitACTION_ID(s)
	}
}

type ACTION_TAGContext struct {
	Metadata_action_with_paramsContext
}

func NewACTION_TAGContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ACTION_TAGContext {
	var p = new(ACTION_TAGContext)

	InitEmptyMetadata_action_with_paramsContext(&p.Metadata_action_with_paramsContext)
	p.parser = parser
	p.CopyAll(ctx.(*Metadata_action_with_paramsContext))

	return p
}

func (s *ACTION_TAGContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ACTION_TAGContext) ACTION_TAG() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_TAG, 0)
}

func (s *ACTION_TAGContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterACTION_TAG(s)
	}
}

func (s *ACTION_TAGContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitACTION_TAG(s)
	}
}

func (p *SecLangParser) Metadata_action_with_params() (localctx IMetadata_action_with_paramsContext) {
	localctx = NewMetadata_action_with_paramsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, SecLangParserRULE_metadata_action_with_params)
	p.SetState(597)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SecLangParserACTION_PHASE:
		localctx = NewACTION_PHASEContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(589)
			p.Match(SecLangParserACTION_PHASE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SecLangParserACTION_ID:
		localctx = NewACTION_IDContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(590)
			p.Match(SecLangParserACTION_ID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SecLangParserACTION_MATURITY:
		localctx = NewACTION_MATURITYContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(591)
			p.Match(SecLangParserACTION_MATURITY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SecLangParserACTION_MSG:
		localctx = NewACTION_MSGContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(592)
			p.Match(SecLangParserACTION_MSG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SecLangParserACTION_REV:
		localctx = NewACTION_REVContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(593)
			p.Match(SecLangParserACTION_REV)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SecLangParserACTION_SEVERITY:
		localctx = NewACTION_SEVERITYContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(594)
			p.Match(SecLangParserACTION_SEVERITY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SecLangParserACTION_TAG:
		localctx = NewACTION_TAGContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(595)
			p.Match(SecLangParserACTION_TAG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SecLangParserACTION_VER:
		localctx = NewACTION_VERContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(596)
			p.Match(SecLangParserACTION_VER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDisruptive_action_with_paramsContext is an interface to support dynamic dispatch.
type IDisruptive_action_with_paramsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ACTION_PROXY() antlr.TerminalNode
	ACTION_REDIRECT() antlr.TerminalNode

	// IsDisruptive_action_with_paramsContext differentiates from other interfaces.
	IsDisruptive_action_with_paramsContext()
}

type Disruptive_action_with_paramsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDisruptive_action_with_paramsContext() *Disruptive_action_with_paramsContext {
	var p = new(Disruptive_action_with_paramsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_disruptive_action_with_params
	return p
}

func InitEmptyDisruptive_action_with_paramsContext(p *Disruptive_action_with_paramsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_disruptive_action_with_params
}

func (*Disruptive_action_with_paramsContext) IsDisruptive_action_with_paramsContext() {}

func NewDisruptive_action_with_paramsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Disruptive_action_with_paramsContext {
	var p = new(Disruptive_action_with_paramsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_disruptive_action_with_params

	return p
}

func (s *Disruptive_action_with_paramsContext) GetParser() antlr.Parser { return s.parser }

func (s *Disruptive_action_with_paramsContext) ACTION_PROXY() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_PROXY, 0)
}

func (s *Disruptive_action_with_paramsContext) ACTION_REDIRECT() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_REDIRECT, 0)
}

func (s *Disruptive_action_with_paramsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Disruptive_action_with_paramsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Disruptive_action_with_paramsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterDisruptive_action_with_params(s)
	}
}

func (s *Disruptive_action_with_paramsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitDisruptive_action_with_params(s)
	}
}

func (p *SecLangParser) Disruptive_action_with_params() (localctx IDisruptive_action_with_paramsContext) {
	localctx = NewDisruptive_action_with_paramsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, SecLangParserRULE_disruptive_action_with_params)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(599)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SecLangParserACTION_PROXY || _la == SecLangParserACTION_REDIRECT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INon_disruptive_action_with_paramsContext is an interface to support dynamic dispatch.
type INon_disruptive_action_with_paramsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ACTION_APPEND() antlr.TerminalNode
	ACTION_CTL() antlr.TerminalNode
	ACTION_EXEC() antlr.TerminalNode
	ACTION_EXPIRE_VAR() antlr.TerminalNode
	ACTION_DEPRECATE_VAR() antlr.TerminalNode
	ACTION_INITCOL() antlr.TerminalNode
	ACTION_LOG_DATA() antlr.TerminalNode
	ACTION_PREPEND() antlr.TerminalNode
	ACTION_SANITISE_ARG() antlr.TerminalNode
	ACTION_SANITISE_MATCHED_BYTES() antlr.TerminalNode
	ACTION_SANITISE_REQUEST_HEADER() antlr.TerminalNode
	ACTION_SANITISE_RESPONSE_HEADER() antlr.TerminalNode
	ACTION_SETUID() antlr.TerminalNode
	ACTION_SETRSC() antlr.TerminalNode
	ACTION_SETSID() antlr.TerminalNode
	ACTION_SETENV() antlr.TerminalNode
	ACTION_SETVAR() antlr.TerminalNode

	// IsNon_disruptive_action_with_paramsContext differentiates from other interfaces.
	IsNon_disruptive_action_with_paramsContext()
}

type Non_disruptive_action_with_paramsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNon_disruptive_action_with_paramsContext() *Non_disruptive_action_with_paramsContext {
	var p = new(Non_disruptive_action_with_paramsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_non_disruptive_action_with_params
	return p
}

func InitEmptyNon_disruptive_action_with_paramsContext(p *Non_disruptive_action_with_paramsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_non_disruptive_action_with_params
}

func (*Non_disruptive_action_with_paramsContext) IsNon_disruptive_action_with_paramsContext() {}

func NewNon_disruptive_action_with_paramsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Non_disruptive_action_with_paramsContext {
	var p = new(Non_disruptive_action_with_paramsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_non_disruptive_action_with_params

	return p
}

func (s *Non_disruptive_action_with_paramsContext) GetParser() antlr.Parser { return s.parser }

func (s *Non_disruptive_action_with_paramsContext) ACTION_APPEND() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_APPEND, 0)
}

func (s *Non_disruptive_action_with_paramsContext) ACTION_CTL() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_CTL, 0)
}

func (s *Non_disruptive_action_with_paramsContext) ACTION_EXEC() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_EXEC, 0)
}

func (s *Non_disruptive_action_with_paramsContext) ACTION_EXPIRE_VAR() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_EXPIRE_VAR, 0)
}

func (s *Non_disruptive_action_with_paramsContext) ACTION_DEPRECATE_VAR() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_DEPRECATE_VAR, 0)
}

func (s *Non_disruptive_action_with_paramsContext) ACTION_INITCOL() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_INITCOL, 0)
}

func (s *Non_disruptive_action_with_paramsContext) ACTION_LOG_DATA() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_LOG_DATA, 0)
}

func (s *Non_disruptive_action_with_paramsContext) ACTION_PREPEND() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_PREPEND, 0)
}

func (s *Non_disruptive_action_with_paramsContext) ACTION_SANITISE_ARG() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_SANITISE_ARG, 0)
}

func (s *Non_disruptive_action_with_paramsContext) ACTION_SANITISE_MATCHED_BYTES() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_SANITISE_MATCHED_BYTES, 0)
}

func (s *Non_disruptive_action_with_paramsContext) ACTION_SANITISE_REQUEST_HEADER() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_SANITISE_REQUEST_HEADER, 0)
}

func (s *Non_disruptive_action_with_paramsContext) ACTION_SANITISE_RESPONSE_HEADER() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_SANITISE_RESPONSE_HEADER, 0)
}

func (s *Non_disruptive_action_with_paramsContext) ACTION_SETUID() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_SETUID, 0)
}

func (s *Non_disruptive_action_with_paramsContext) ACTION_SETRSC() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_SETRSC, 0)
}

func (s *Non_disruptive_action_with_paramsContext) ACTION_SETSID() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_SETSID, 0)
}

func (s *Non_disruptive_action_with_paramsContext) ACTION_SETENV() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_SETENV, 0)
}

func (s *Non_disruptive_action_with_paramsContext) ACTION_SETVAR() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_SETVAR, 0)
}

func (s *Non_disruptive_action_with_paramsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Non_disruptive_action_with_paramsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Non_disruptive_action_with_paramsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterNon_disruptive_action_with_params(s)
	}
}

func (s *Non_disruptive_action_with_paramsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitNon_disruptive_action_with_params(s)
	}
}

func (p *SecLangParser) Non_disruptive_action_with_params() (localctx INon_disruptive_action_with_paramsContext) {
	localctx = NewNon_disruptive_action_with_paramsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, SecLangParserRULE_non_disruptive_action_with_params)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(601)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-28)) & ^0x3f) == 0 && ((int64(1)<<(_la-28))&140058897809441) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IData_action_with_paramsContext is an interface to support dynamic dispatch.
type IData_action_with_paramsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ACTION_XMLNS() antlr.TerminalNode
	ACTION_STATUS() antlr.TerminalNode

	// IsData_action_with_paramsContext differentiates from other interfaces.
	IsData_action_with_paramsContext()
}

type Data_action_with_paramsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyData_action_with_paramsContext() *Data_action_with_paramsContext {
	var p = new(Data_action_with_paramsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_data_action_with_params
	return p
}

func InitEmptyData_action_with_paramsContext(p *Data_action_with_paramsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_data_action_with_params
}

func (*Data_action_with_paramsContext) IsData_action_with_paramsContext() {}

func NewData_action_with_paramsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Data_action_with_paramsContext {
	var p = new(Data_action_with_paramsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_data_action_with_params

	return p
}

func (s *Data_action_with_paramsContext) GetParser() antlr.Parser { return s.parser }

func (s *Data_action_with_paramsContext) ACTION_XMLNS() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_XMLNS, 0)
}

func (s *Data_action_with_paramsContext) ACTION_STATUS() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_STATUS, 0)
}

func (s *Data_action_with_paramsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Data_action_with_paramsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Data_action_with_paramsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterData_action_with_params(s)
	}
}

func (s *Data_action_with_paramsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitData_action_with_params(s)
	}
}

func (p *SecLangParser) Data_action_with_params() (localctx IData_action_with_paramsContext) {
	localctx = NewData_action_with_paramsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, SecLangParserRULE_data_action_with_params)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(603)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SecLangParserACTION_STATUS || _la == SecLangParserACTION_XMLNS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFlow_action_with_paramsContext is an interface to support dynamic dispatch.
type IFlow_action_with_paramsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ACTION_SKIP() antlr.TerminalNode
	ACTION_SKIP_AFTER() antlr.TerminalNode

	// IsFlow_action_with_paramsContext differentiates from other interfaces.
	IsFlow_action_with_paramsContext()
}

type Flow_action_with_paramsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFlow_action_with_paramsContext() *Flow_action_with_paramsContext {
	var p = new(Flow_action_with_paramsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_flow_action_with_params
	return p
}

func InitEmptyFlow_action_with_paramsContext(p *Flow_action_with_paramsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_flow_action_with_params
}

func (*Flow_action_with_paramsContext) IsFlow_action_with_paramsContext() {}

func NewFlow_action_with_paramsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Flow_action_with_paramsContext {
	var p = new(Flow_action_with_paramsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_flow_action_with_params

	return p
}

func (s *Flow_action_with_paramsContext) GetParser() antlr.Parser { return s.parser }

func (s *Flow_action_with_paramsContext) ACTION_SKIP() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_SKIP, 0)
}

func (s *Flow_action_with_paramsContext) ACTION_SKIP_AFTER() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_SKIP_AFTER, 0)
}

func (s *Flow_action_with_paramsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Flow_action_with_paramsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Flow_action_with_paramsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterFlow_action_with_params(s)
	}
}

func (s *Flow_action_with_paramsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitFlow_action_with_params(s)
	}
}

func (p *SecLangParser) Flow_action_with_params() (localctx IFlow_action_with_paramsContext) {
	localctx = NewFlow_action_with_paramsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, SecLangParserRULE_flow_action_with_params)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(605)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SecLangParserACTION_SKIP_AFTER || _la == SecLangParserACTION_SKIP) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAction_valueContext is an interface to support dynamic dispatch.
type IAction_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Action_value_types() IAction_value_typesContext
	AllSINGLE_QUOTE() []antlr.TerminalNode
	SINGLE_QUOTE(i int) antlr.TerminalNode
	String_literal() IString_literalContext

	// IsAction_valueContext differentiates from other interfaces.
	IsAction_valueContext()
}

type Action_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAction_valueContext() *Action_valueContext {
	var p = new(Action_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_action_value
	return p
}

func InitEmptyAction_valueContext(p *Action_valueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_action_value
}

func (*Action_valueContext) IsAction_valueContext() {}

func NewAction_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Action_valueContext {
	var p = new(Action_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_action_value

	return p
}

func (s *Action_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Action_valueContext) Action_value_types() IAction_value_typesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAction_value_typesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAction_value_typesContext)
}

func (s *Action_valueContext) AllSINGLE_QUOTE() []antlr.TerminalNode {
	return s.GetTokens(SecLangParserSINGLE_QUOTE)
}

func (s *Action_valueContext) SINGLE_QUOTE(i int) antlr.TerminalNode {
	return s.GetToken(SecLangParserSINGLE_QUOTE, i)
}

func (s *Action_valueContext) String_literal() IString_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IString_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IString_literalContext)
}

func (s *Action_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Action_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Action_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterAction_value(s)
	}
}

func (s *Action_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitAction_value(s)
	}
}

func (p *SecLangParser) Action_value() (localctx IAction_valueContext) {
	localctx = NewAction_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, SecLangParserRULE_action_value)
	p.SetState(616)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 67, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(607)
			p.Action_value_types()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(608)
			p.Match(SecLangParserSINGLE_QUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(609)
			p.Action_value_types()
		}
		{
			p.SetState(610)
			p.Match(SecLangParserSINGLE_QUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(612)
			p.Match(SecLangParserSINGLE_QUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(613)
			p.String_literal()
		}
		{
			p.SetState(614)
			p.Match(SecLangParserSINGLE_QUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAction_value_typesContext is an interface to support dynamic dispatch.
type IAction_value_typesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	Collection_value() ICollection_valueContext
	Setvar_action() ISetvar_actionContext
	Ctl_action() ICtl_actionContext
	Assignment() IAssignmentContext
	Values() IValuesContext
	VARIABLE_NAME() antlr.TerminalNode
	ACTION_SEVERITY_VALUE() antlr.TerminalNode
	FREE_TEXT_QUOTE_MACRO_EXPANSION() antlr.TerminalNode
	COMMA_SEPARATED_STRING() antlr.TerminalNode

	// IsAction_value_typesContext differentiates from other interfaces.
	IsAction_value_typesContext()
}

type Action_value_typesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAction_value_typesContext() *Action_value_typesContext {
	var p = new(Action_value_typesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_action_value_types
	return p
}

func InitEmptyAction_value_typesContext(p *Action_value_typesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_action_value_types
}

func (*Action_value_typesContext) IsAction_value_typesContext() {}

func NewAction_value_typesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Action_value_typesContext {
	var p = new(Action_value_typesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_action_value_types

	return p
}

func (s *Action_value_typesContext) GetParser() antlr.Parser { return s.parser }

func (s *Action_value_typesContext) INT() antlr.TerminalNode {
	return s.GetToken(SecLangParserINT, 0)
}

func (s *Action_value_typesContext) Collection_value() ICollection_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICollection_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICollection_valueContext)
}

func (s *Action_value_typesContext) Setvar_action() ISetvar_actionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISetvar_actionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISetvar_actionContext)
}

func (s *Action_value_typesContext) Ctl_action() ICtl_actionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICtl_actionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICtl_actionContext)
}

func (s *Action_value_typesContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *Action_value_typesContext) Values() IValuesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValuesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValuesContext)
}

func (s *Action_value_typesContext) VARIABLE_NAME() antlr.TerminalNode {
	return s.GetToken(SecLangParserVARIABLE_NAME, 0)
}

func (s *Action_value_typesContext) ACTION_SEVERITY_VALUE() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_SEVERITY_VALUE, 0)
}

func (s *Action_value_typesContext) FREE_TEXT_QUOTE_MACRO_EXPANSION() antlr.TerminalNode {
	return s.GetToken(SecLangParserFREE_TEXT_QUOTE_MACRO_EXPANSION, 0)
}

func (s *Action_value_typesContext) COMMA_SEPARATED_STRING() antlr.TerminalNode {
	return s.GetToken(SecLangParserCOMMA_SEPARATED_STRING, 0)
}

func (s *Action_value_typesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Action_value_typesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Action_value_typesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterAction_value_types(s)
	}
}

func (s *Action_value_typesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitAction_value_types(s)
	}
}

func (p *SecLangParser) Action_value_types() (localctx IAction_value_typesContext) {
	localctx = NewAction_value_typesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, SecLangParserRULE_action_value_types)
	p.SetState(629)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SecLangParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(618)
			p.Match(SecLangParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SecLangParserQUOTE, SecLangParserSINGLE_QUOTE, SecLangParserCOMMA, SecLangParserXPATH_EXPRESSION, SecLangParserCOLLECTION_ELEMENT_VALUE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(619)
			p.Collection_value()
		}

	case SecLangParserCOLLECTION_NAME_SETVAR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(620)
			p.Setvar_action()
		}

	case SecLangParserACTION_CTL_AUDIT_ENGINE, SecLangParserACTION_CTL_AUDIT_LOG_PARTS, SecLangParserACTION_CTL_REQUEST_BODY_PROCESSOR, SecLangParserACTION_CTL_FORCE_REQ_BODY_VAR, SecLangParserACTION_CTL_REQUEST_BODY_ACCESS, SecLangParserACTION_CTL_RULE_ENGINE, SecLangParserACTION_CTL_RULE_REMOVE_BY_TAG, SecLangParserACTION_CTL_RULE_REMOVE_BY_ID, SecLangParserACTION_CTL_RULE_REMOVE_TARGET_BY_ID, SecLangParserACTION_CTL_RULE_REMOVE_TARGET_BY_TAG:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(621)
			p.Ctl_action()
		}
		{
			p.SetState(622)
			p.Assignment()
		}
		{
			p.SetState(623)
			p.Values()
		}

	case SecLangParserVARIABLE_NAME:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(625)
			p.Match(SecLangParserVARIABLE_NAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SecLangParserACTION_SEVERITY_VALUE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(626)
			p.Match(SecLangParserACTION_SEVERITY_VALUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SecLangParserFREE_TEXT_QUOTE_MACRO_EXPANSION:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(627)
			p.Match(SecLangParserFREE_TEXT_QUOTE_MACRO_EXPANSION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SecLangParserCOMMA_SEPARATED_STRING:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(628)
			p.Match(SecLangParserCOMMA_SEPARATED_STRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IString_literalContext is an interface to support dynamic dispatch.
type IString_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING_LITERAL() antlr.TerminalNode

	// IsString_literalContext differentiates from other interfaces.
	IsString_literalContext()
}

type String_literalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyString_literalContext() *String_literalContext {
	var p = new(String_literalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_string_literal
	return p
}

func InitEmptyString_literalContext(p *String_literalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_string_literal
}

func (*String_literalContext) IsString_literalContext() {}

func NewString_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_literalContext {
	var p = new(String_literalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_string_literal

	return p
}

func (s *String_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *String_literalContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(SecLangParserSTRING_LITERAL, 0)
}

func (s *String_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterString_literal(s)
	}
}

func (s *String_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitString_literal(s)
	}
}

func (p *SecLangParser) String_literal() (localctx IString_literalContext) {
	localctx = NewString_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, SecLangParserRULE_string_literal)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(631)
		p.Match(SecLangParserSTRING_LITERAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICtl_actionContext is an interface to support dynamic dispatch.
type ICtl_actionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ACTION_CTL_FORCE_REQ_BODY_VAR() antlr.TerminalNode
	ACTION_CTL_REQUEST_BODY_ACCESS() antlr.TerminalNode
	ACTION_CTL_RULE_ENGINE() antlr.TerminalNode
	ACTION_CTL_RULE_REMOVE_BY_ID() antlr.TerminalNode
	ACTION_CTL_RULE_REMOVE_BY_TAG() antlr.TerminalNode
	ACTION_CTL_RULE_REMOVE_TARGET_BY_ID() antlr.TerminalNode
	ACTION_CTL_RULE_REMOVE_TARGET_BY_TAG() antlr.TerminalNode
	ACTION_CTL_AUDIT_ENGINE() antlr.TerminalNode
	ACTION_CTL_AUDIT_LOG_PARTS() antlr.TerminalNode
	ACTION_CTL_REQUEST_BODY_PROCESSOR() antlr.TerminalNode

	// IsCtl_actionContext differentiates from other interfaces.
	IsCtl_actionContext()
}

type Ctl_actionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCtl_actionContext() *Ctl_actionContext {
	var p = new(Ctl_actionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_ctl_action
	return p
}

func InitEmptyCtl_actionContext(p *Ctl_actionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_ctl_action
}

func (*Ctl_actionContext) IsCtl_actionContext() {}

func NewCtl_actionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ctl_actionContext {
	var p = new(Ctl_actionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_ctl_action

	return p
}

func (s *Ctl_actionContext) GetParser() antlr.Parser { return s.parser }

func (s *Ctl_actionContext) ACTION_CTL_FORCE_REQ_BODY_VAR() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_CTL_FORCE_REQ_BODY_VAR, 0)
}

func (s *Ctl_actionContext) ACTION_CTL_REQUEST_BODY_ACCESS() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_CTL_REQUEST_BODY_ACCESS, 0)
}

func (s *Ctl_actionContext) ACTION_CTL_RULE_ENGINE() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_CTL_RULE_ENGINE, 0)
}

func (s *Ctl_actionContext) ACTION_CTL_RULE_REMOVE_BY_ID() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_CTL_RULE_REMOVE_BY_ID, 0)
}

func (s *Ctl_actionContext) ACTION_CTL_RULE_REMOVE_BY_TAG() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_CTL_RULE_REMOVE_BY_TAG, 0)
}

func (s *Ctl_actionContext) ACTION_CTL_RULE_REMOVE_TARGET_BY_ID() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_CTL_RULE_REMOVE_TARGET_BY_ID, 0)
}

func (s *Ctl_actionContext) ACTION_CTL_RULE_REMOVE_TARGET_BY_TAG() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_CTL_RULE_REMOVE_TARGET_BY_TAG, 0)
}

func (s *Ctl_actionContext) ACTION_CTL_AUDIT_ENGINE() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_CTL_AUDIT_ENGINE, 0)
}

func (s *Ctl_actionContext) ACTION_CTL_AUDIT_LOG_PARTS() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_CTL_AUDIT_LOG_PARTS, 0)
}

func (s *Ctl_actionContext) ACTION_CTL_REQUEST_BODY_PROCESSOR() antlr.TerminalNode {
	return s.GetToken(SecLangParserACTION_CTL_REQUEST_BODY_PROCESSOR, 0)
}

func (s *Ctl_actionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ctl_actionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ctl_actionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterCtl_action(s)
	}
}

func (s *Ctl_actionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitCtl_action(s)
	}
}

func (p *SecLangParser) Ctl_action() (localctx ICtl_actionContext) {
	localctx = NewCtl_actionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, SecLangParserRULE_ctl_action)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(633)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17575006175232) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITransformation_action_valueContext is an interface to support dynamic dispatch.
type ITransformation_action_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TRANSFORMATION_VALUE() antlr.TerminalNode

	// IsTransformation_action_valueContext differentiates from other interfaces.
	IsTransformation_action_valueContext()
}

type Transformation_action_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTransformation_action_valueContext() *Transformation_action_valueContext {
	var p = new(Transformation_action_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_transformation_action_value
	return p
}

func InitEmptyTransformation_action_valueContext(p *Transformation_action_valueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_transformation_action_value
}

func (*Transformation_action_valueContext) IsTransformation_action_valueContext() {}

func NewTransformation_action_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Transformation_action_valueContext {
	var p = new(Transformation_action_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_transformation_action_value

	return p
}

func (s *Transformation_action_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Transformation_action_valueContext) TRANSFORMATION_VALUE() antlr.TerminalNode {
	return s.GetToken(SecLangParserTRANSFORMATION_VALUE, 0)
}

func (s *Transformation_action_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Transformation_action_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Transformation_action_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterTransformation_action_value(s)
	}
}

func (s *Transformation_action_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitTransformation_action_value(s)
	}
}

func (p *SecLangParser) Transformation_action_value() (localctx ITransformation_action_valueContext) {
	localctx = NewTransformation_action_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, SecLangParserRULE_transformation_action_value)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(635)
		p.Match(SecLangParserTRANSFORMATION_VALUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICollection_valueContext is an interface to support dynamic dispatch.
type ICollection_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	XPATH_EXPRESSION() antlr.TerminalNode
	COLLECTION_ELEMENT_VALUE() antlr.TerminalNode

	// IsCollection_valueContext differentiates from other interfaces.
	IsCollection_valueContext()
}

type Collection_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCollection_valueContext() *Collection_valueContext {
	var p = new(Collection_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_collection_value
	return p
}

func InitEmptyCollection_valueContext(p *Collection_valueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_collection_value
}

func (*Collection_valueContext) IsCollection_valueContext() {}

func NewCollection_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Collection_valueContext {
	var p = new(Collection_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_collection_value

	return p
}

func (s *Collection_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Collection_valueContext) XPATH_EXPRESSION() antlr.TerminalNode {
	return s.GetToken(SecLangParserXPATH_EXPRESSION, 0)
}

func (s *Collection_valueContext) COLLECTION_ELEMENT_VALUE() antlr.TerminalNode {
	return s.GetToken(SecLangParserCOLLECTION_ELEMENT_VALUE, 0)
}

func (s *Collection_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Collection_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Collection_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterCollection_value(s)
	}
}

func (s *Collection_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitCollection_value(s)
	}
}

func (p *SecLangParser) Collection_value() (localctx ICollection_valueContext) {
	localctx = NewCollection_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, SecLangParserRULE_collection_value)
	p.SetState(640)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SecLangParserEOF, SecLangParserQUOTE, SecLangParserSINGLE_QUOTE, SecLangParserCOMMA, SecLangParserPIPE, SecLangParserHASH, SecLangParserVARIABLE_NAME_ENUM, SecLangParserCONFIG_COMPONENT_SIG, SecLangParserCONFIG_SEC_SERVER_SIG, SecLangParserCONFIG_SEC_WEB_APP_ID, SecLangParserCONFIG_SEC_CACHE_TRANSFORMATIONS, SecLangParserCONFIG_SEC_CHROOT_DIR, SecLangParserCONFIG_CONN_ENGINE, SecLangParserCONFIG_SEC_HASH_ENGINE, SecLangParserCONFIG_SEC_HASH_KEY, SecLangParserCONFIG_SEC_HASH_PARAM, SecLangParserCONFIG_SEC_HASH_METHOD_RX, SecLangParserCONFIG_SEC_HASH_METHOD_PM, SecLangParserCONFIG_CONTENT_INJECTION, SecLangParserCONFIG_SEC_ARGUMENT_SEPARATOR, SecLangParserCONFIG_DIR_AUDIT_DIR, SecLangParserCONFIG_DIR_AUDIT_DIR_MOD, SecLangParserCONFIG_DIR_AUDIT_ENG, SecLangParserCONFIG_DIR_AUDIT_FILE_MODE, SecLangParserCONFIG_DIR_AUDIT_LOG2, SecLangParserCONFIG_DIR_AUDIT_LOG, SecLangParserCONFIG_DIR_AUDIT_LOG_FMT, SecLangParserCONFIG_DIR_AUDIT_LOG_P, SecLangParserCONFIG_DIR_AUDIT_STS, SecLangParserCONFIG_DIR_AUDIT_TYPE, SecLangParserCONFIG_DIR_DEBUG_LOG, SecLangParserCONFIG_DIR_DEBUG_LVL, SecLangParserCONFIG_DIR_GEO_DB, SecLangParserCONFIG_DIR_GSB_DB, SecLangParserCONFIG_SEC_GUARDIAN_LOG, SecLangParserCONFIG_SEC_INTERCEPT_ON_ERROR, SecLangParserCONFIG_SEC_CONN_R_STATE_LIMIT, SecLangParserCONFIG_SEC_CONN_W_STATE_LIMIT, SecLangParserCONFIG_SEC_SENSOR_ID, SecLangParserCONFIG_SEC_RULE_INHERITANCE, SecLangParserCONFIG_SEC_RULE_PERF_TIME, SecLangParserCONFIG_SEC_STREAM_IN_BODY_INSPECTION, SecLangParserCONFIG_SEC_STREAM_OUT_BODY_INSPECTION, SecLangParserCONFIG_DIR_PCRE_MATCH_LIMIT, SecLangParserCONFIG_DIR_PCRE_MATCH_LIMIT_RECURSION, SecLangParserCONFIG_DIR_ARGS_LIMIT, SecLangParserCONFIG_DIR_REQ_BODY_JSON_DEPTH_LIMIT, SecLangParserCONFIG_DIR_REQ_BODY, SecLangParserCONFIG_DIR_REQ_BODY_LIMIT, SecLangParserCONFIG_DIR_REQ_BODY_LIMIT_ACTION, SecLangParserCONFIG_DIR_REQ_BODY_NO_FILES_LIMIT, SecLangParserCONFIG_DIR_RES_BODY, SecLangParserCONFIG_DIR_RES_BODY_LIMIT, SecLangParserCONFIG_DIR_RES_BODY_LIMIT_ACTION, SecLangParserCONFIG_DIR_RULE_ENG, SecLangParserCONFIG_DIR_SEC_ACTION, SecLangParserCONFIG_DIR_SEC_DEFAULT_ACTION, SecLangParserCONFIG_SEC_DISABLE_BACKEND_COMPRESS, SecLangParserCONFIG_DIR_SEC_MARKER, SecLangParserCONFIG_DIR_UNICODE_MAP_FILE, SecLangParserCONFIG_SEC_COLLECTION_TIMEOUT, SecLangParserCONFIG_SEC_HTTP_BLKEY, SecLangParserCONFIG_SEC_REMOTE_RULES_FAIL_ACTION, SecLangParserCONFIG_SEC_RULE_REMOVE_BY_ID, SecLangParserCONFIG_SEC_RULE_REMOVE_BY_MSG, SecLangParserCONFIG_SEC_RULE_REMOVE_BY_TAG, SecLangParserCONFIG_SEC_RULE_UPDATE_TARGET_BY_TAG, SecLangParserCONFIG_SEC_RULE_UPDATE_TARGET_BY_MSG, SecLangParserCONFIG_SEC_RULE_UPDATE_TARGET_BY_ID, SecLangParserCONFIG_SEC_RULE_UPDATE_ACTION_BY_ID, SecLangParserCONFIG_UPLOAD_KEEP_FILES, SecLangParserCONFIG_UPLOAD_SAVE_TMP_FILES, SecLangParserCONFIG_UPLOAD_DIR, SecLangParserCONFIG_UPLOAD_FILE_LIMIT, SecLangParserCONFIG_UPLOAD_FILE_MODE, SecLangParserCONFIG_XML_EXTERNAL_ENTITY, SecLangParserCONFIG_DIR_RESPONSE_BODY_MP, SecLangParserCONFIG_DIR_RESPONSE_BODY_MP_CLEAR, SecLangParserCONFIG_DIR_SEC_COOKIE_FORMAT, SecLangParserCONFIG_SEC_COOKIEV0_SEPARATOR, SecLangParserCONFIG_DIR_SEC_DATA_DIR, SecLangParserCONFIG_DIR_SEC_STATUS_ENGINE, SecLangParserCONFIG_DIR_SEC_TMP_DIR, SecLangParserCONFIG_DIR_SEC_RULE, SecLangParserDIRECTIVE_SECRULESCRIPT, SecLangParserINT, SecLangParserSTRING, SecLangParserOPERATOR_UNQUOTED_STRING, SecLangParserOPERATOR_QUOTED_STRING:
		p.EnterOuterAlt(localctx, 1)

	case SecLangParserXPATH_EXPRESSION:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(638)
			p.Match(SecLangParserXPATH_EXPRESSION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SecLangParserCOLLECTION_ELEMENT_VALUE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(639)
			p.Match(SecLangParserCOLLECTION_ELEMENT_VALUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICtl_collection_valueContext is an interface to support dynamic dispatch.
type ICtl_collection_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	XPATH_EXPRESSION() antlr.TerminalNode
	COLLECTION_ELEMENT_VALUE() antlr.TerminalNode

	// IsCtl_collection_valueContext differentiates from other interfaces.
	IsCtl_collection_valueContext()
}

type Ctl_collection_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCtl_collection_valueContext() *Ctl_collection_valueContext {
	var p = new(Ctl_collection_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_ctl_collection_value
	return p
}

func InitEmptyCtl_collection_valueContext(p *Ctl_collection_valueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_ctl_collection_value
}

func (*Ctl_collection_valueContext) IsCtl_collection_valueContext() {}

func NewCtl_collection_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ctl_collection_valueContext {
	var p = new(Ctl_collection_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_ctl_collection_value

	return p
}

func (s *Ctl_collection_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Ctl_collection_valueContext) XPATH_EXPRESSION() antlr.TerminalNode {
	return s.GetToken(SecLangParserXPATH_EXPRESSION, 0)
}

func (s *Ctl_collection_valueContext) COLLECTION_ELEMENT_VALUE() antlr.TerminalNode {
	return s.GetToken(SecLangParserCOLLECTION_ELEMENT_VALUE, 0)
}

func (s *Ctl_collection_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ctl_collection_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ctl_collection_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterCtl_collection_value(s)
	}
}

func (s *Ctl_collection_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitCtl_collection_value(s)
	}
}

func (p *SecLangParser) Ctl_collection_value() (localctx ICtl_collection_valueContext) {
	localctx = NewCtl_collection_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, SecLangParserRULE_ctl_collection_value)
	p.SetState(645)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SecLangParserEOF, SecLangParserQUOTE, SecLangParserSINGLE_QUOTE, SecLangParserCOMMA, SecLangParserHASH, SecLangParserCONFIG_COMPONENT_SIG, SecLangParserCONFIG_SEC_SERVER_SIG, SecLangParserCONFIG_SEC_WEB_APP_ID, SecLangParserCONFIG_SEC_CACHE_TRANSFORMATIONS, SecLangParserCONFIG_SEC_CHROOT_DIR, SecLangParserCONFIG_CONN_ENGINE, SecLangParserCONFIG_SEC_HASH_ENGINE, SecLangParserCONFIG_SEC_HASH_KEY, SecLangParserCONFIG_SEC_HASH_PARAM, SecLangParserCONFIG_SEC_HASH_METHOD_RX, SecLangParserCONFIG_SEC_HASH_METHOD_PM, SecLangParserCONFIG_CONTENT_INJECTION, SecLangParserCONFIG_SEC_ARGUMENT_SEPARATOR, SecLangParserCONFIG_DIR_AUDIT_DIR, SecLangParserCONFIG_DIR_AUDIT_DIR_MOD, SecLangParserCONFIG_DIR_AUDIT_ENG, SecLangParserCONFIG_DIR_AUDIT_FILE_MODE, SecLangParserCONFIG_DIR_AUDIT_LOG2, SecLangParserCONFIG_DIR_AUDIT_LOG, SecLangParserCONFIG_DIR_AUDIT_LOG_FMT, SecLangParserCONFIG_DIR_AUDIT_LOG_P, SecLangParserCONFIG_DIR_AUDIT_STS, SecLangParserCONFIG_DIR_AUDIT_TYPE, SecLangParserCONFIG_DIR_DEBUG_LOG, SecLangParserCONFIG_DIR_DEBUG_LVL, SecLangParserCONFIG_DIR_GEO_DB, SecLangParserCONFIG_DIR_GSB_DB, SecLangParserCONFIG_SEC_GUARDIAN_LOG, SecLangParserCONFIG_SEC_INTERCEPT_ON_ERROR, SecLangParserCONFIG_SEC_CONN_R_STATE_LIMIT, SecLangParserCONFIG_SEC_CONN_W_STATE_LIMIT, SecLangParserCONFIG_SEC_SENSOR_ID, SecLangParserCONFIG_SEC_RULE_INHERITANCE, SecLangParserCONFIG_SEC_RULE_PERF_TIME, SecLangParserCONFIG_SEC_STREAM_IN_BODY_INSPECTION, SecLangParserCONFIG_SEC_STREAM_OUT_BODY_INSPECTION, SecLangParserCONFIG_DIR_PCRE_MATCH_LIMIT, SecLangParserCONFIG_DIR_PCRE_MATCH_LIMIT_RECURSION, SecLangParserCONFIG_DIR_ARGS_LIMIT, SecLangParserCONFIG_DIR_REQ_BODY_JSON_DEPTH_LIMIT, SecLangParserCONFIG_DIR_REQ_BODY, SecLangParserCONFIG_DIR_REQ_BODY_LIMIT, SecLangParserCONFIG_DIR_REQ_BODY_LIMIT_ACTION, SecLangParserCONFIG_DIR_REQ_BODY_NO_FILES_LIMIT, SecLangParserCONFIG_DIR_RES_BODY, SecLangParserCONFIG_DIR_RES_BODY_LIMIT, SecLangParserCONFIG_DIR_RES_BODY_LIMIT_ACTION, SecLangParserCONFIG_DIR_RULE_ENG, SecLangParserCONFIG_DIR_SEC_ACTION, SecLangParserCONFIG_DIR_SEC_DEFAULT_ACTION, SecLangParserCONFIG_SEC_DISABLE_BACKEND_COMPRESS, SecLangParserCONFIG_DIR_SEC_MARKER, SecLangParserCONFIG_DIR_UNICODE_MAP_FILE, SecLangParserCONFIG_SEC_COLLECTION_TIMEOUT, SecLangParserCONFIG_SEC_HTTP_BLKEY, SecLangParserCONFIG_SEC_REMOTE_RULES_FAIL_ACTION, SecLangParserCONFIG_SEC_RULE_REMOVE_BY_ID, SecLangParserCONFIG_SEC_RULE_REMOVE_BY_MSG, SecLangParserCONFIG_SEC_RULE_REMOVE_BY_TAG, SecLangParserCONFIG_SEC_RULE_UPDATE_TARGET_BY_TAG, SecLangParserCONFIG_SEC_RULE_UPDATE_TARGET_BY_MSG, SecLangParserCONFIG_SEC_RULE_UPDATE_TARGET_BY_ID, SecLangParserCONFIG_SEC_RULE_UPDATE_ACTION_BY_ID, SecLangParserCONFIG_UPLOAD_KEEP_FILES, SecLangParserCONFIG_UPLOAD_SAVE_TMP_FILES, SecLangParserCONFIG_UPLOAD_DIR, SecLangParserCONFIG_UPLOAD_FILE_LIMIT, SecLangParserCONFIG_UPLOAD_FILE_MODE, SecLangParserCONFIG_XML_EXTERNAL_ENTITY, SecLangParserCONFIG_DIR_RESPONSE_BODY_MP, SecLangParserCONFIG_DIR_RESPONSE_BODY_MP_CLEAR, SecLangParserCONFIG_DIR_SEC_COOKIE_FORMAT, SecLangParserCONFIG_SEC_COOKIEV0_SEPARATOR, SecLangParserCONFIG_DIR_SEC_DATA_DIR, SecLangParserCONFIG_DIR_SEC_STATUS_ENGINE, SecLangParserCONFIG_DIR_SEC_TMP_DIR, SecLangParserCONFIG_DIR_SEC_RULE, SecLangParserDIRECTIVE_SECRULESCRIPT, SecLangParserINT:
		p.EnterOuterAlt(localctx, 1)

	case SecLangParserXPATH_EXPRESSION:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(643)
			p.Match(SecLangParserXPATH_EXPRESSION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SecLangParserCOLLECTION_ELEMENT_VALUE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(644)
			p.Match(SecLangParserCOLLECTION_ELEMENT_VALUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISetvar_actionContext is an interface to support dynamic dispatch.
type ISetvar_actionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Col_name() ICol_nameContext
	DOT() antlr.TerminalNode
	Setvar_stmt() ISetvar_stmtContext
	Assignment() IAssignmentContext
	Var_assignment() IVar_assignmentContext

	// IsSetvar_actionContext differentiates from other interfaces.
	IsSetvar_actionContext()
}

type Setvar_actionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetvar_actionContext() *Setvar_actionContext {
	var p = new(Setvar_actionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_setvar_action
	return p
}

func InitEmptySetvar_actionContext(p *Setvar_actionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_setvar_action
}

func (*Setvar_actionContext) IsSetvar_actionContext() {}

func NewSetvar_actionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Setvar_actionContext {
	var p = new(Setvar_actionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_setvar_action

	return p
}

func (s *Setvar_actionContext) GetParser() antlr.Parser { return s.parser }

func (s *Setvar_actionContext) Col_name() ICol_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICol_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICol_nameContext)
}

func (s *Setvar_actionContext) DOT() antlr.TerminalNode {
	return s.GetToken(SecLangParserDOT, 0)
}

func (s *Setvar_actionContext) Setvar_stmt() ISetvar_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISetvar_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISetvar_stmtContext)
}

func (s *Setvar_actionContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *Setvar_actionContext) Var_assignment() IVar_assignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_assignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_assignmentContext)
}

func (s *Setvar_actionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Setvar_actionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Setvar_actionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterSetvar_action(s)
	}
}

func (s *Setvar_actionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitSetvar_action(s)
	}
}

func (p *SecLangParser) Setvar_action() (localctx ISetvar_actionContext) {
	localctx = NewSetvar_actionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 126, SecLangParserRULE_setvar_action)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(647)
		p.Col_name()
	}
	{
		p.SetState(648)
		p.Match(SecLangParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(649)
		p.Setvar_stmt()
	}
	{
		p.SetState(650)
		p.Assignment()
	}
	{
		p.SetState(651)
		p.Var_assignment()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICol_nameContext is an interface to support dynamic dispatch.
type ICol_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COLLECTION_NAME_SETVAR() antlr.TerminalNode

	// IsCol_nameContext differentiates from other interfaces.
	IsCol_nameContext()
}

type Col_nameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCol_nameContext() *Col_nameContext {
	var p = new(Col_nameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_col_name
	return p
}

func InitEmptyCol_nameContext(p *Col_nameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_col_name
}

func (*Col_nameContext) IsCol_nameContext() {}

func NewCol_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Col_nameContext {
	var p = new(Col_nameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_col_name

	return p
}

func (s *Col_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Col_nameContext) COLLECTION_NAME_SETVAR() antlr.TerminalNode {
	return s.GetToken(SecLangParserCOLLECTION_NAME_SETVAR, 0)
}

func (s *Col_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Col_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Col_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterCol_name(s)
	}
}

func (s *Col_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitCol_name(s)
	}
}

func (p *SecLangParser) Col_name() (localctx ICol_nameContext) {
	localctx = NewCol_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 128, SecLangParserRULE_col_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(653)
		p.Match(SecLangParserCOLLECTION_NAME_SETVAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISetvar_stmtContext is an interface to support dynamic dispatch.
type ISetvar_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllCOLLECTION_ELEMENT() []antlr.TerminalNode
	COLLECTION_ELEMENT(i int) antlr.TerminalNode
	AllCOLLECTION_WITH_MACRO() []antlr.TerminalNode
	COLLECTION_WITH_MACRO(i int) antlr.TerminalNode
	AllMACRO_EXPANSION() []antlr.TerminalNode
	MACRO_EXPANSION(i int) antlr.TerminalNode

	// IsSetvar_stmtContext differentiates from other interfaces.
	IsSetvar_stmtContext()
}

type Setvar_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetvar_stmtContext() *Setvar_stmtContext {
	var p = new(Setvar_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_setvar_stmt
	return p
}

func InitEmptySetvar_stmtContext(p *Setvar_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_setvar_stmt
}

func (*Setvar_stmtContext) IsSetvar_stmtContext() {}

func NewSetvar_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Setvar_stmtContext {
	var p = new(Setvar_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_setvar_stmt

	return p
}

func (s *Setvar_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Setvar_stmtContext) AllCOLLECTION_ELEMENT() []antlr.TerminalNode {
	return s.GetTokens(SecLangParserCOLLECTION_ELEMENT)
}

func (s *Setvar_stmtContext) COLLECTION_ELEMENT(i int) antlr.TerminalNode {
	return s.GetToken(SecLangParserCOLLECTION_ELEMENT, i)
}

func (s *Setvar_stmtContext) AllCOLLECTION_WITH_MACRO() []antlr.TerminalNode {
	return s.GetTokens(SecLangParserCOLLECTION_WITH_MACRO)
}

func (s *Setvar_stmtContext) COLLECTION_WITH_MACRO(i int) antlr.TerminalNode {
	return s.GetToken(SecLangParserCOLLECTION_WITH_MACRO, i)
}

func (s *Setvar_stmtContext) AllMACRO_EXPANSION() []antlr.TerminalNode {
	return s.GetTokens(SecLangParserMACRO_EXPANSION)
}

func (s *Setvar_stmtContext) MACRO_EXPANSION(i int) antlr.TerminalNode {
	return s.GetToken(SecLangParserMACRO_EXPANSION, i)
}

func (s *Setvar_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Setvar_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Setvar_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterSetvar_stmt(s)
	}
}

func (s *Setvar_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitSetvar_stmt(s)
	}
}

func (p *SecLangParser) Setvar_stmt() (localctx ISetvar_stmtContext) {
	localctx = NewSetvar_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 130, SecLangParserRULE_setvar_stmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(658)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SecLangParserCOLLECTION_ELEMENT || _la == SecLangParserCOLLECTION_WITH_MACRO {
		p.SetState(658)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case SecLangParserCOLLECTION_ELEMENT:
			{
				p.SetState(655)
				p.Match(SecLangParserCOLLECTION_ELEMENT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case SecLangParserCOLLECTION_WITH_MACRO:
			{
				p.SetState(656)
				p.Match(SecLangParserCOLLECTION_WITH_MACRO)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(657)
				p.Match(SecLangParserMACRO_EXPANSION)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(660)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EQUAL() antlr.TerminalNode
	EQUALS_PLUS() antlr.TerminalNode
	EQUALS_MINUS() antlr.TerminalNode

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(SecLangParserEQUAL, 0)
}

func (s *AssignmentContext) EQUALS_PLUS() antlr.TerminalNode {
	return s.GetToken(SecLangParserEQUALS_PLUS, 0)
}

func (s *AssignmentContext) EQUALS_MINUS() antlr.TerminalNode {
	return s.GetToken(SecLangParserEQUALS_MINUS, 0)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (p *SecLangParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 132, SecLangParserRULE_assignment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(662)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&104) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVar_assignmentContext is an interface to support dynamic dispatch.
type IVar_assignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR_ASSIGNMENT() antlr.TerminalNode

	// IsVar_assignmentContext differentiates from other interfaces.
	IsVar_assignmentContext()
}

type Var_assignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_assignmentContext() *Var_assignmentContext {
	var p = new(Var_assignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_var_assignment
	return p
}

func InitEmptyVar_assignmentContext(p *Var_assignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_var_assignment
}

func (*Var_assignmentContext) IsVar_assignmentContext() {}

func NewVar_assignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_assignmentContext {
	var p = new(Var_assignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_var_assignment

	return p
}

func (s *Var_assignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_assignmentContext) VAR_ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(SecLangParserVAR_ASSIGNMENT, 0)
}

func (s *Var_assignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_assignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_assignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterVar_assignment(s)
	}
}

func (s *Var_assignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitVar_assignment(s)
	}
}

func (p *SecLangParser) Var_assignment() (localctx IVar_assignmentContext) {
	localctx = NewVar_assignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 134, SecLangParserRULE_var_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(664)
		p.Match(SecLangParserVAR_ASSIGNMENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICtl_idContext is an interface to support dynamic dispatch.
type ICtl_idContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	IDENT() antlr.TerminalNode

	// IsCtl_idContext differentiates from other interfaces.
	IsCtl_idContext()
}

type Ctl_idContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCtl_idContext() *Ctl_idContext {
	var p = new(Ctl_idContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_ctl_id
	return p
}

func InitEmptyCtl_idContext(p *Ctl_idContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SecLangParserRULE_ctl_id
}

func (*Ctl_idContext) IsCtl_idContext() {}

func NewCtl_idContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ctl_idContext {
	var p = new(Ctl_idContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SecLangParserRULE_ctl_id

	return p
}

func (s *Ctl_idContext) GetParser() antlr.Parser { return s.parser }

func (s *Ctl_idContext) INT() antlr.TerminalNode {
	return s.GetToken(SecLangParserINT, 0)
}

func (s *Ctl_idContext) IDENT() antlr.TerminalNode {
	return s.GetToken(SecLangParserIDENT, 0)
}

func (s *Ctl_idContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ctl_idContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ctl_idContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.EnterCtl_id(s)
	}
}

func (s *Ctl_idContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SecLangParserListener); ok {
		listenerT.ExitCtl_id(s)
	}
}

func (p *SecLangParser) Ctl_id() (localctx ICtl_idContext) {
	localctx = NewCtl_idContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 136, SecLangParserRULE_ctl_id)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(666)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SecLangParserIDENT || _la == SecLangParserINT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
