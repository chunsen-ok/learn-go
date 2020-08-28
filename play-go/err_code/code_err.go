package err_code

const (

	// Class 00 - Successful Completion
	ERR_00000 = "successful completion"

	// Class 01 - Warning
	ERR_01000 = "warning"
	ERR_0100C = "dynamic result sets returned"
	ERR_01008 = "implicit zero bit padding"
	ERR_01003 = "null value eliminated in set function"
	ERR_01007 = "privilege not granted"
	ERR_01006 = "privilege not revoked"
	ERR_01004 = "string data right truncation"
	ERR_01P01 = "deprecated feature"

	// Class 02 - No Data (this is also a warning class per the SQL standard)
	ERR_02000 = "no data"
	ERR_02001 = "no additional dynamic result sets returned"

	// Class 03 - SQL Statement Not Yet Complete
	ERR_03000 = "sql statement not yet complete"

	// Class 08 - Connection Exception
	ERR_08000 = "connection exception"
	ERR_08003 = "connection does not exist"
	ERR_08006 = "connection failure"
	ERR_08001 = "sqlclient unable to establish sqlconnection"
	ERR_08004 = "sqlserver rejected establishment of sqlconnection"
	ERR_08007 = "transaction resolution unknown"
	ERR_08P01 = "protocol violation"

	// Class 09 - Triggered Action Exception
	ERR_09000 = "triggered action exception"

	// Class 0A - Feature Not Supported
	ERR_0A000 = "feature not supported"

	// Class 0B - Invalid Transaction Initiation
	ERR_0B000 = "invalid transaction initiation"

	// Class 0F - Locator Exception
	ERR_0F000 = "locator exception"
	ERR_0F001 = "invalid locator specification"

	// Class 0L - Invalid Grantor
	ERR_0L000 = "invalid grantor"
	ERR_0LP01 = "invalid grant operation"

	// Class 0P - Invalid Role Specification
	ERR_0P000 = "invalid role specification"

	// Class 0Z - Diagnostics Exception
	ERR_0Z000 = "diagnostics exception"
	ERR_0Z002 = "stacked diagnostics accessed without active handler"

	// Class 20 - Case Not Found
	ERR_20000 = "case not found"

	// Class 21 - Cardinality Violation
	ERR_21000 = "cardinality violation"

	// Class 22 - Data Exception
	ERR_22000 = "data exception"
	ERR_2202E = "array subscript error"
	ERR_22021 = "character not in repertoire"
	ERR_22008 = "datetime field overflow"
	ERR_22012 = "division by zero"
	ERR_22005 = "error in assignment"
	ERR_2200B = "escape character conflict"
	ERR_22022 = "indicator overflow"
	ERR_22015 = "interval field overflow"
	ERR_2201E = "invalid argument for logarithm"
	ERR_22014 = "invalid argument for ntile function"
	ERR_22016 = "invalid argument for nth value function"
	ERR_2201F = "invalid argument for power function"
	ERR_2201G = "invalid argument for width bucket function"
	ERR_22018 = "invalid character value for cast"
	ERR_22007 = "invalid datetime format"
	ERR_22019 = "invalid escape character"
	ERR_2200D = "invalid escape octet"
	ERR_22025 = "invalid escape sequence"
	ERR_22P06 = "nonstandard use of escape character"
	ERR_22010 = "invalid indicator parameter value"
	ERR_22023 = "invalid parameter value"
	ERR_2201B = "invalid regular expression"
	ERR_2201W = "invalid row count in limit clause"
	ERR_2201X = "invalid row count in result offset clause"
	ERR_2202H = "invalid tablesample argument"
	ERR_2202G = "invalid tablesample repeat"
	ERR_22009 = "invalid time zone displacement value"
	ERR_2200C = "invalid use of escape character"
	ERR_2200G = "most specific type mismatch"
	ERR_22004 = "null value not allowed"
	ERR_22002 = "null value no indicator parameter"
	ERR_22003 = "numeric value out of range"
	ERR_2200H = "sequence generator limit exceeded"
	ERR_22026 = "string data length mismatch"
	ERR_22001 = "string data right truncation"
	ERR_22011 = "substring error"
	ERR_22027 = "trim error"
	ERR_22024 = "unterminated c string"
	ERR_2200F = "zero length character string"
	ERR_22P01 = "floating point exception"
	ERR_22P02 = "invalid text representation"
	ERR_22P03 = "invalid binary representation"
	ERR_22P04 = "bad copy file format"
	ERR_22P05 = "untranslatable character"
	ERR_2200L = "not an xml document"
	ERR_2200M = "invalid xml document"
	ERR_2200N = "invalid xml content"
	ERR_2200S = "invalid xml comment"
	ERR_2200T = "invalid xml processing instruction"

	// Class 23 - Integrity Constraint Violation
	ERR_23000 = "integrity constraint violation"
	ERR_23001 = "restrict violation"
	ERR_23502 = "not null violation"
	ERR_23503 = "foreign key violation"
	ERR_23505 = "unique violation"
	ERR_23514 = "check violation"
	ERR_23P01 = "exclusion violation"

	// Class 24 - Invalid Cursor State
	ERR_24000 = "invalid cursor state"

	// Class 25 - Invalid Transaction State
	ERR_25000 = "invalid transaction state"
	ERR_25001 = "active sql transaction"
	ERR_25002 = "branch transaction already active"
	ERR_25008 = "held cursor requires same isolation level"
	ERR_25003 = "inappropriate access mode for branch transaction"
	ERR_25004 = "inappropriate isolation level for branch transaction"
	ERR_25005 = "no active sql transaction for branch transaction"
	ERR_25006 = "read only sql transaction"
	ERR_25007 = "schema and data statement mixing not supported"
	ERR_25P01 = "no active sql transaction"
	ERR_25P02 = "in failed sql transaction"
	ERR_25P03 = "idle in transaction session timeout"

	// Class 26 - Invalid SQL Statement Name
	ERR_26000 = "invalid sql statement name"

	// Class 27 - Triggered Data Change Violation
	ERR_27000 = "triggered data change violation"

	// Class 28 - Invalid Authorization Specification
	ERR_28000 = "invalid authorization specification"
	ERR_28P01 = "invalid password"

	// Class 2B - Dependent Privilege Descriptors Still Exist
	ERR_2B000 = "dependent privilege descriptors still exist"
	ERR_2BP01 = "dependent objects still exist"

	// Class 2D - Invalid Transaction Termination
	ERR_2D000 = "invalid transaction termination"

	// Class 2F - SQL Routine Exception
	ERR_2F000 = "sql routine exception"
	ERR_2F005 = "function executed no return statement"
	ERR_2F002 = "modifying sql data not permitted"
	ERR_2F003 = "prohibited sql statement attempted"
	ERR_2F004 = "reading sql data not permitted"

	// Class 34 - Invalid Cursor Name
	ERR_34000 = "invalid cursor name"

	// Class 38 - External Routine Exception
	ERR_38000 = "external routine exception"
	ERR_38001 = "containing sql not permitted"
	ERR_38002 = "modifying sql data not permitted"
	ERR_38003 = "prohibited sql statement attempted"
	ERR_38004 = "reading sql data not permitted"

	// Class 39 - External Routine Invocation Exception
	ERR_39000 = "external routine invocation exception"
	ERR_39001 = "invalid sqlstate returned"
	ERR_39004 = "null value not allowed"
	ERR_39P01 = "trigger protocol violated"
	ERR_39P02 = "srf protocol violated"
	ERR_39P03 = "event trigger protocol violated"

	// Class 3B - Savepoint Exception
	ERR_3B000 = "savepoint exception"
	ERR_3B001 = "invalid savepoint specification"

	// Class 3D - Invalid Catalog Name
	ERR_3D000 = "invalid catalog name"

	// Class 3F - Invalid Schema Name
	ERR_3F000 = "invalid schema name"

	// Class 40 - Transaction Rollback
	ERR_40000 = "transaction rollback"
	ERR_40002 = "transaction integrity constraint violation"
	ERR_40001 = "serialization failure"
	ERR_40003 = "statement completion unknown"
	ERR_40P01 = "deadlock detected"

	// Class 42 - Syntax Error or Access Rule Violation
	ERR_42000 = "syntax error or access rule violation"
	ERR_42601 = "syntax error"
	ERR_42501 = "insufficient privilege"
	ERR_42846 = "cannot coerce"
	ERR_42803 = "grouping error"
	ERR_42P20 = "windowing error"
	ERR_42P19 = "invalid recursion"
	ERR_42830 = "invalid foreign key"
	ERR_42602 = "invalid name"
	ERR_42622 = "name too long"
	ERR_42939 = "reserved name"
	ERR_42804 = "datatype mismatch"
	ERR_42P18 = "indeterminate datatype"
	ERR_42P21 = "collation mismatch"
	ERR_42P22 = "indeterminate collation"
	ERR_42809 = "wrong object type"
	ERR_428C9 = "generated always"
	ERR_42703 = "undefined column"
	ERR_42883 = "undefined function"
	ERR_42P01 = "undefined table"
	ERR_42P02 = "undefined parameter"
	ERR_42704 = "undefined object"
	ERR_42701 = "duplicate column"
	ERR_42P03 = "duplicate cursor"
	ERR_42P04 = "duplicate database"
	ERR_42723 = "duplicate function"
	ERR_42P05 = "duplicate prepared statement"
	ERR_42P06 = "duplicate schema"
	ERR_42P07 = "duplicate table"
	ERR_42712 = "duplicate alias"
	ERR_42710 = "duplicate object"
	ERR_42702 = "ambiguous column"
	ERR_42725 = "ambiguous function"
	ERR_42P08 = "ambiguous parameter"
	ERR_42P09 = "ambiguous alias"
	ERR_42P10 = "invalid column reference"
	ERR_42611 = "invalid column definition"
	ERR_42P11 = "invalid cursor definition"
	ERR_42P12 = "invalid database definition"
	ERR_42P13 = "invalid function definition"
	ERR_42P14 = "invalid prepared statement definition"
	ERR_42P15 = "invalid schema definition"
	ERR_42P16 = "invalid table definition"
	ERR_42P17 = "invalid object definition"

	// Class 44 - WITH CHECK OPTION Violation
	ERR_44000 = "with check option violation"

	// Class 53 - Insufficient Resources
	ERR_53000 = "insufficient resources"
	ERR_53100 = "disk full"
	ERR_53200 = "out of memory"
	ERR_53300 = "too many connections"
	ERR_53400 = "configuration limit exceeded"

	// Class 54 - Program Limit Exceeded
	ERR_54000 = "program limit exceeded"
	ERR_54001 = "statement too complex"
	ERR_54011 = "too many columns"
	ERR_54023 = "too many arguments"

	// Class 55 - Object Not In Prerequisite State
	ERR_55000 = "object not in prerequisite state"
	ERR_55006 = "object in use"
	ERR_55P02 = "cant change runtime param"
	ERR_55P03 = "lock not available"

	// Class 57 - Operator Intervention
	ERR_57000 = "operator intervention"
	ERR_57014 = "query canceled"
	ERR_57P01 = "admin shutdown"
	ERR_57P02 = "crash shutdown"
	ERR_57P03 = "cannot connect now"
	ERR_57P04 = "database dropped"

	// Class 58 - System Error (errors external to PostgreSQL itself)
	ERR_58000 = "system error"
	ERR_58030 = "io error"
	ERR_58P01 = "undefined file"
	ERR_58P02 = "duplicate file"

	// Class 72 - Snapshot Failure
	ERR_72000 = "snapshot too old"

	// Class F0 - Configuration File Error
	ERR_F0000 = "config file error"
	ERR_F0001 = "lock file exists"

	// Class HV - Foreign Data Wrapper Error (SQL/MED)
	ERR_HV000 = "fdw error"
	ERR_HV005 = "fdw column name not found"
	ERR_HV002 = "fdw dynamic parameter value needed"
	ERR_HV010 = "fdw function sequence error"
	ERR_HV021 = "fdw inconsistent descriptor information"
	ERR_HV024 = "fdw invalid attribute value"
	ERR_HV007 = "fdw invalid column name"
	ERR_HV008 = "fdw invalid column number"
	ERR_HV004 = "fdw invalid data type"
	ERR_HV006 = "fdw invalid data type descriptors"
	ERR_HV091 = "fdw invalid descriptor field identifier"
	ERR_HV00B = "fdw invalid handle"
	ERR_HV00C = "fdw invalid option index"
	ERR_HV00D = "fdw invalid option name"
	ERR_HV090 = "fdw invalid string length or buffer length"
	ERR_HV00A = "fdw invalid string format"
	ERR_HV009 = "fdw invalid use of null pointer"
	ERR_HV014 = "fdw too many handles"
	ERR_HV001 = "fdw out of memory"
	ERR_HV00P = "fdw no schemas"
	ERR_HV00J = "fdw option name not found"
	ERR_HV00K = "fdw reply handle"
	ERR_HV00Q = "fdw schema not found"
	ERR_HV00R = "fdw table not found"
	ERR_HV00L = "fdw unable to create execution"
	ERR_HV00M = "fdw unable to create reply"
	ERR_HV00N = "fdw unable to establish connection"

	// Class P0 - PL/pgSQL Error
	ERR_P0000 = "plpgsql error"
	ERR_P0001 = "raise exception"
	ERR_P0002 = "no data found"
	ERR_P0003 = "too many rows"
	ERR_P0004 = "assert failure"

	// Class XX - Internal Error
	ERR_XX000 = "internal error"
	ERR_XX001 = "data corrupted"
	ERR_XX002 = "index corrupted"

//
)

const ERR_CODES = `Class 00 — Successful Completion
00000	successful_completion
Class 01 — Warning
01000	warning
0100C	dynamic_result_sets_returned
01008	implicit_zero_bit_padding
01003	null_value_eliminated_in_set_function
01007	privilege_not_granted
01006	privilege_not_revoked
01004	string_data_right_truncation
01P01	deprecated_feature
Class 02 — No Data (this is also a warning class per the SQL standard)
02000	no_data
02001	no_additional_dynamic_result_sets_returned
Class 03 — SQL Statement Not Yet Complete
03000	sql_statement_not_yet_complete
Class 08 — Connection Exception
08000	connection_exception
08003	connection_does_not_exist
08006	connection_failure
08001	sqlclient_unable_to_establish_sqlconnection
08004	sqlserver_rejected_establishment_of_sqlconnection
08007	transaction_resolution_unknown
08P01	protocol_violation
Class 09 — Triggered Action Exception
09000	triggered_action_exception
Class 0A — Feature Not Supported
0A000	feature_not_supported
Class 0B — Invalid Transaction Initiation
0B000	invalid_transaction_initiation
Class 0F — Locator Exception
0F000	locator_exception
0F001	invalid_locator_specification
Class 0L — Invalid Grantor
0L000	invalid_grantor
0LP01	invalid_grant_operation
Class 0P — Invalid Role Specification
0P000	invalid_role_specification
Class 0Z — Diagnostics Exception
0Z000	diagnostics_exception
0Z002	stacked_diagnostics_accessed_without_active_handler
Class 20 — Case Not Found
20000	case_not_found
Class 21 — Cardinality Violation
21000	cardinality_violation
Class 22 — Data Exception
22000	data_exception
2202E	array_subscript_error
22021	character_not_in_repertoire
22008	datetime_field_overflow
22012	division_by_zero
22005	error_in_assignment
2200B	escape_character_conflict
22022	indicator_overflow
22015	interval_field_overflow
2201E	invalid_argument_for_logarithm
22014	invalid_argument_for_ntile_function
22016	invalid_argument_for_nth_value_function
2201F	invalid_argument_for_power_function
2201G	invalid_argument_for_width_bucket_function
22018	invalid_character_value_for_cast
22007	invalid_datetime_format
22019	invalid_escape_character
2200D	invalid_escape_octet
22025	invalid_escape_sequence
22P06	nonstandard_use_of_escape_character
22010	invalid_indicator_parameter_value
22023	invalid_parameter_value
2201B	invalid_regular_expression
2201W	invalid_row_count_in_limit_clause
2201X	invalid_row_count_in_result_offset_clause
2202H	invalid_tablesample_argument
2202G	invalid_tablesample_repeat
22009	invalid_time_zone_displacement_value
2200C	invalid_use_of_escape_character
2200G	most_specific_type_mismatch
22004	null_value_not_allowed
22002	null_value_no_indicator_parameter
22003	numeric_value_out_of_range
2200H	sequence_generator_limit_exceeded
22026	string_data_length_mismatch
22001	string_data_right_truncation
22011	substring_error
22027	trim_error
22024	unterminated_c_string
2200F	zero_length_character_string
22P01	floating_point_exception
22P02	invalid_text_representation
22P03	invalid_binary_representation
22P04	bad_copy_file_format
22P05	untranslatable_character
2200L	not_an_xml_document
2200M	invalid_xml_document
2200N	invalid_xml_content
2200S	invalid_xml_comment
2200T	invalid_xml_processing_instruction
Class 23 — Integrity Constraint Violation
23000	integrity_constraint_violation
23001	restrict_violation
23502	not_null_violation
23503	foreign_key_violation
23505	unique_violation
23514	check_violation
23P01	exclusion_violation
Class 24 — Invalid Cursor State
24000	invalid_cursor_state
Class 25 — Invalid Transaction State
25000	invalid_transaction_state
25001	active_sql_transaction
25002	branch_transaction_already_active
25008	held_cursor_requires_same_isolation_level
25003	inappropriate_access_mode_for_branch_transaction
25004	inappropriate_isolation_level_for_branch_transaction
25005	no_active_sql_transaction_for_branch_transaction
25006	read_only_sql_transaction
25007	schema_and_data_statement_mixing_not_supported
25P01	no_active_sql_transaction
25P02	in_failed_sql_transaction
25P03	idle_in_transaction_session_timeout
Class 26 — Invalid SQL Statement Name
26000	invalid_sql_statement_name
Class 27 — Triggered Data Change Violation
27000	triggered_data_change_violation
Class 28 — Invalid Authorization Specification
28000	invalid_authorization_specification
28P01	invalid_password
Class 2B — Dependent Privilege Descriptors Still Exist
2B000	dependent_privilege_descriptors_still_exist
2BP01	dependent_objects_still_exist
Class 2D — Invalid Transaction Termination
2D000	invalid_transaction_termination
Class 2F — SQL Routine Exception
2F000	sql_routine_exception
2F005	function_executed_no_return_statement
2F002	modifying_sql_data_not_permitted
2F003	prohibited_sql_statement_attempted
2F004	reading_sql_data_not_permitted
Class 34 — Invalid Cursor Name
34000	invalid_cursor_name
Class 38 — External Routine Exception
38000	external_routine_exception
38001	containing_sql_not_permitted
38002	modifying_sql_data_not_permitted
38003	prohibited_sql_statement_attempted
38004	reading_sql_data_not_permitted
Class 39 — External Routine Invocation Exception
39000	external_routine_invocation_exception
39001	invalid_sqlstate_returned
39004	null_value_not_allowed
39P01	trigger_protocol_violated
39P02	srf_protocol_violated
39P03	event_trigger_protocol_violated
Class 3B — Savepoint Exception
3B000	savepoint_exception
3B001	invalid_savepoint_specification
Class 3D — Invalid Catalog Name
3D000	invalid_catalog_name
Class 3F — Invalid Schema Name
3F000	invalid_schema_name
Class 40 — Transaction Rollback
40000	transaction_rollback
40002	transaction_integrity_constraint_violation
40001	serialization_failure
40003	statement_completion_unknown
40P01	deadlock_detected
Class 42 — Syntax Error or Access Rule Violation
42000	syntax_error_or_access_rule_violation
42601	syntax_error
42501	insufficient_privilege
42846	cannot_coerce
42803	grouping_error
42P20	windowing_error
42P19	invalid_recursion
42830	invalid_foreign_key
42602	invalid_name
42622	name_too_long
42939	reserved_name
42804	datatype_mismatch
42P18	indeterminate_datatype
42P21	collation_mismatch
42P22	indeterminate_collation
42809	wrong_object_type
428C9	generated_always
42703	undefined_column
42883	undefined_function
42P01	undefined_table
42P02	undefined_parameter
42704	undefined_object
42701	duplicate_column
42P03	duplicate_cursor
42P04	duplicate_database
42723	duplicate_function
42P05	duplicate_prepared_statement
42P06	duplicate_schema
42P07	duplicate_table
42712	duplicate_alias
42710	duplicate_object
42702	ambiguous_column
42725	ambiguous_function
42P08	ambiguous_parameter
42P09	ambiguous_alias
42P10	invalid_column_reference
42611	invalid_column_definition
42P11	invalid_cursor_definition
42P12	invalid_database_definition
42P13	invalid_function_definition
42P14	invalid_prepared_statement_definition
42P15	invalid_schema_definition
42P16	invalid_table_definition
42P17	invalid_object_definition
Class 44 — WITH CHECK OPTION Violation
44000	with_check_option_violation
Class 53 — Insufficient Resources
53000	insufficient_resources
53100	disk_full
53200	out_of_memory
53300	too_many_connections
53400	configuration_limit_exceeded
Class 54 — Program Limit Exceeded
54000	program_limit_exceeded
54001	statement_too_complex
54011	too_many_columns
54023	too_many_arguments
Class 55 — Object Not In Prerequisite State
55000	object_not_in_prerequisite_state
55006	object_in_use
55P02	cant_change_runtime_param
55P03	lock_not_available
Class 57 — Operator Intervention
57000	operator_intervention
57014	query_canceled
57P01	admin_shutdown
57P02	crash_shutdown
57P03	cannot_connect_now
57P04	database_dropped
Class 58 — System Error (errors external to PostgreSQL itself)
58000	system_error
58030	io_error
58P01	undefined_file
58P02	duplicate_file
Class 72 — Snapshot Failure
72000	snapshot_too_old
Class F0 — Configuration File Error
F0000	config_file_error
F0001	lock_file_exists
Class HV — Foreign Data Wrapper Error (SQL/MED)
HV000	fdw_error
HV005	fdw_column_name_not_found
HV002	fdw_dynamic_parameter_value_needed
HV010	fdw_function_sequence_error
HV021	fdw_inconsistent_descriptor_information
HV024	fdw_invalid_attribute_value
HV007	fdw_invalid_column_name
HV008	fdw_invalid_column_number
HV004	fdw_invalid_data_type
HV006	fdw_invalid_data_type_descriptors
HV091	fdw_invalid_descriptor_field_identifier
HV00B	fdw_invalid_handle
HV00C	fdw_invalid_option_index
HV00D	fdw_invalid_option_name
HV090	fdw_invalid_string_length_or_buffer_length
HV00A	fdw_invalid_string_format
HV009	fdw_invalid_use_of_null_pointer
HV014	fdw_too_many_handles
HV001	fdw_out_of_memory
HV00P	fdw_no_schemas
HV00J	fdw_option_name_not_found
HV00K	fdw_reply_handle
HV00Q	fdw_schema_not_found
HV00R	fdw_table_not_found
HV00L	fdw_unable_to_create_execution
HV00M	fdw_unable_to_create_reply
HV00N	fdw_unable_to_establish_connection
Class P0 — PL/pgSQL Error
P0000	plpgsql_error
P0001	raise_exception
P0002	no_data_found
P0003	too_many_rows
P0004	assert_failure
Class XX — Internal Error
XX000	internal_error
XX001	data_corrupted
XX002	index_corrupted
`

var dbErrs map[string]string

func init() {
	dbErrs = make(map[string]string)

	// Class 00 - Successful Completion
	dbErrs["00000"] = "successful completion"

	// Class 01 - Warning
	dbErrs["01000"] = "warning"
	dbErrs["0100C"] = "dynamic result sets returned"
	dbErrs["01008"] = "implicit zero bit padding"
	dbErrs["01003"] = "null value eliminated in set function"
	dbErrs["01007"] = "privilege not granted"
	dbErrs["01006"] = "privilege not revoked"
	dbErrs["01004"] = "string data right truncation"
	dbErrs["01P01"] = "deprecated feature"

	// Class 02 - No Data (this is also a warning class per the SQL standard)
	dbErrs["02000"] = "no data"
	dbErrs["02001"] = "no additional dynamic result sets returned"

	// Class 03 - SQL Statement Not Yet Complete
	dbErrs["03000"] = "sql statement not yet complete"

	// Class 08 - Connection Exception
	dbErrs["08000"] = "connection exception"
	dbErrs["08003"] = "connection does not exist"
	dbErrs["08006"] = "connection failure"
	dbErrs["08001"] = "sqlclient unable to establish sqlconnection"
	dbErrs["08004"] = "sqlserver rejected establishment of sqlconnection"
	dbErrs["08007"] = "transaction resolution unknown"
	dbErrs["08P01"] = "protocol violation"

	// Class 09 - Triggered Action Exception
	dbErrs["09000"] = "triggered action exception"

	// Class 0A - Feature Not Supported
	dbErrs["0A000"] = "feature not supported"

	// Class 0B - Invalid Transaction Initiation
	dbErrs["0B000"] = "invalid transaction initiation"

	// Class 0F - Locator Exception
	dbErrs["0F000"] = "locator exception"
	dbErrs["0F001"] = "invalid locator specification"

	// Class 0L - Invalid Grantor
	dbErrs["0L000"] = "invalid grantor"
	dbErrs["0LP01"] = "invalid grant operation"

	// Class 0P - Invalid Role Specification
	dbErrs["0P000"] = "invalid role specification"

	// Class 0Z - Diagnostics Exception
	dbErrs["0Z000"] = "diagnostics exception"
	dbErrs["0Z002"] = "stacked diagnostics accessed without active handler"

	// Class 20 - Case Not Found
	dbErrs["20000"] = "case not found"

	// Class 21 - Cardinality Violation
	dbErrs["21000"] = "cardinality violation"

	// Class 22 - Data Exception
	dbErrs["22000"] = "data exception"
	dbErrs["2202E"] = "array subscript error"
	dbErrs["22021"] = "character not in repertoire"
	dbErrs["22008"] = "datetime field overflow"
	dbErrs["22012"] = "division by zero"
	dbErrs["22005"] = "error in assignment"
	dbErrs["2200B"] = "escape character conflict"
	dbErrs["22022"] = "indicator overflow"
	dbErrs["22015"] = "interval field overflow"
	dbErrs["2201E"] = "invalid argument for logarithm"
	dbErrs["22014"] = "invalid argument for ntile function"
	dbErrs["22016"] = "invalid argument for nth value function"
	dbErrs["2201F"] = "invalid argument for power function"
	dbErrs["2201G"] = "invalid argument for width bucket function"
	dbErrs["22018"] = "invalid character value for cast"
	dbErrs["22007"] = "invalid datetime format"
	dbErrs["22019"] = "invalid escape character"
	dbErrs["2200D"] = "invalid escape octet"
	dbErrs["22025"] = "invalid escape sequence"
	dbErrs["22P06"] = "nonstandard use of escape character"
	dbErrs["22010"] = "invalid indicator parameter value"
	dbErrs["22023"] = "invalid parameter value"
	dbErrs["2201B"] = "invalid regular expression"
	dbErrs["2201W"] = "invalid row count in limit clause"
	dbErrs["2201X"] = "invalid row count in result offset clause"
	dbErrs["2202H"] = "invalid tablesample argument"
	dbErrs["2202G"] = "invalid tablesample repeat"
	dbErrs["22009"] = "invalid time zone displacement value"
	dbErrs["2200C"] = "invalid use of escape character"
	dbErrs["2200G"] = "most specific type mismatch"
	dbErrs["22004"] = "null value not allowed"
	dbErrs["22002"] = "null value no indicator parameter"
	dbErrs["22003"] = "numeric value out of range"
	dbErrs["2200H"] = "sequence generator limit exceeded"
	dbErrs["22026"] = "string data length mismatch"
	dbErrs["22001"] = "string data right truncation"
	dbErrs["22011"] = "substring error"
	dbErrs["22027"] = "trim error"
	dbErrs["22024"] = "unterminated c string"
	dbErrs["2200F"] = "zero length character string"
	dbErrs["22P01"] = "floating point exception"
	dbErrs["22P02"] = "invalid text representation"
	dbErrs["22P03"] = "invalid binary representation"
	dbErrs["22P04"] = "bad copy file format"
	dbErrs["22P05"] = "untranslatable character"
	dbErrs["2200L"] = "not an xml document"
	dbErrs["2200M"] = "invalid xml document"
	dbErrs["2200N"] = "invalid xml content"
	dbErrs["2200S"] = "invalid xml comment"
	dbErrs["2200T"] = "invalid xml processing instruction"

	// Class 23 - Integrity Constraint Violation
	dbErrs["23000"] = "integrity constraint violation"
	dbErrs["23001"] = "restrict violation"
	dbErrs["23502"] = "not null violation"
	dbErrs["23503"] = "foreign key violation"
	dbErrs["23505"] = "unique violation"
	dbErrs["23514"] = "check violation"
	dbErrs["23P01"] = "exclusion violation"

	// Class 24 - Invalid Cursor State
	dbErrs["24000"] = "invalid cursor state"

	// Class 25 - Invalid Transaction State
	dbErrs["25000"] = "invalid transaction state"
	dbErrs["25001"] = "active sql transaction"
	dbErrs["25002"] = "branch transaction already active"
	dbErrs["25008"] = "held cursor requires same isolation level"
	dbErrs["25003"] = "inappropriate access mode for branch transaction"
	dbErrs["25004"] = "inappropriate isolation level for branch transaction"
	dbErrs["25005"] = "no active sql transaction for branch transaction"
	dbErrs["25006"] = "read only sql transaction"
	dbErrs["25007"] = "schema and data statement mixing not supported"
	dbErrs["25P01"] = "no active sql transaction"
	dbErrs["25P02"] = "in failed sql transaction"
	dbErrs["25P03"] = "idle in transaction session timeout"

	// Class 26 - Invalid SQL Statement Name
	dbErrs["26000"] = "invalid sql statement name"

	// Class 27 - Triggered Data Change Violation
	dbErrs["27000"] = "triggered data change violation"

	// Class 28 - Invalid Authorization Specification
	dbErrs["28000"] = "invalid authorization specification"
	dbErrs["28P01"] = "invalid password"

	// Class 2B - Dependent Privilege Descriptors Still Exist
	dbErrs["2B000"] = "dependent privilege descriptors still exist"
	dbErrs["2BP01"] = "dependent objects still exist"

	// Class 2D - Invalid Transaction Termination
	dbErrs["2D000"] = "invalid transaction termination"

	// Class 2F - SQL Routine Exception
	dbErrs["2F000"] = "sql routine exception"
	dbErrs["2F005"] = "function executed no return statement"
	dbErrs["2F002"] = "modifying sql data not permitted"
	dbErrs["2F003"] = "prohibited sql statement attempted"
	dbErrs["2F004"] = "reading sql data not permitted"

	// Class 34 - Invalid Cursor Name
	dbErrs["34000"] = "invalid cursor name"

	// Class 38 - External Routine Exception
	dbErrs["38000"] = "external routine exception"
	dbErrs["38001"] = "containing sql not permitted"
	dbErrs["38002"] = "modifying sql data not permitted"
	dbErrs["38003"] = "prohibited sql statement attempted"
	dbErrs["38004"] = "reading sql data not permitted"

	// Class 39 - External Routine Invocation Exception
	dbErrs["39000"] = "external routine invocation exception"
	dbErrs["39001"] = "invalid sqlstate returned"
	dbErrs["39004"] = "null value not allowed"
	dbErrs["39P01"] = "trigger protocol violated"
	dbErrs["39P02"] = "srf protocol violated"
	dbErrs["39P03"] = "event trigger protocol violated"

	// Class 3B - Savepoint Exception
	dbErrs["3B000"] = "savepoint exception"
	dbErrs["3B001"] = "invalid savepoint specification"

	// Class 3D - Invalid Catalog Name
	dbErrs["3D000"] = "invalid catalog name"

	// Class 3F - Invalid Schema Name
	dbErrs["3F000"] = "invalid schema name"

	// Class 40 - Transaction Rollback
	dbErrs["40000"] = "transaction rollback"
	dbErrs["40002"] = "transaction integrity constraint violation"
	dbErrs["40001"] = "serialization failure"
	dbErrs["40003"] = "statement completion unknown"
	dbErrs["40P01"] = "deadlock detected"

	// Class 42 - Syntax Error or Access Rule Violation
	dbErrs["42000"] = "syntax error or access rule violation"
	dbErrs["42601"] = "syntax error"
	dbErrs["42501"] = "insufficient privilege"
	dbErrs["42846"] = "cannot coerce"
	dbErrs["42803"] = "grouping error"
	dbErrs["42P20"] = "windowing error"
	dbErrs["42P19"] = "invalid recursion"
	dbErrs["42830"] = "invalid foreign key"
	dbErrs["42602"] = "invalid name"
	dbErrs["42622"] = "name too long"
	dbErrs["42939"] = "reserved name"
	dbErrs["42804"] = "datatype mismatch"
	dbErrs["42P18"] = "indeterminate datatype"
	dbErrs["42P21"] = "collation mismatch"
	dbErrs["42P22"] = "indeterminate collation"
	dbErrs["42809"] = "wrong object type"
	dbErrs["428C9"] = "generated always"
	dbErrs["42703"] = "undefined column"
	dbErrs["42883"] = "undefined function"
	dbErrs["42P01"] = "undefined table"
	dbErrs["42P02"] = "undefined parameter"
	dbErrs["42704"] = "undefined object"
	dbErrs["42701"] = "duplicate column"
	dbErrs["42P03"] = "duplicate cursor"
	dbErrs["42P04"] = "duplicate database"
	dbErrs["42723"] = "duplicate function"
	dbErrs["42P05"] = "duplicate prepared statement"
	dbErrs["42P06"] = "duplicate schema"
	dbErrs["42P07"] = "duplicate table"
	dbErrs["42712"] = "duplicate alias"
	dbErrs["42710"] = "duplicate object"
	dbErrs["42702"] = "ambiguous column"
	dbErrs["42725"] = "ambiguous function"
	dbErrs["42P08"] = "ambiguous parameter"
	dbErrs["42P09"] = "ambiguous alias"
	dbErrs["42P10"] = "invalid column reference"
	dbErrs["42611"] = "invalid column definition"
	dbErrs["42P11"] = "invalid cursor definition"
	dbErrs["42P12"] = "invalid database definition"
	dbErrs["42P13"] = "invalid function definition"
	dbErrs["42P14"] = "invalid prepared statement definition"
	dbErrs["42P15"] = "invalid schema definition"
	dbErrs["42P16"] = "invalid table definition"
	dbErrs["42P17"] = "invalid object definition"

	// Class 44 - WITH CHECK OPTION Violation
	dbErrs["44000"] = "with check option violation"

	// Class 53 - Insufficient Resources
	dbErrs["53000"] = "insufficient resources"
	dbErrs["53100"] = "disk full"
	dbErrs["53200"] = "out of memory"
	dbErrs["53300"] = "too many connections"
	dbErrs["53400"] = "configuration limit exceeded"

	// Class 54 - Program Limit Exceeded
	dbErrs["54000"] = "program limit exceeded"
	dbErrs["54001"] = "statement too complex"
	dbErrs["54011"] = "too many columns"
	dbErrs["54023"] = "too many arguments"

	// Class 55 - Object Not In Prerequisite State
	dbErrs["55000"] = "object not in prerequisite state"
	dbErrs["55006"] = "object in use"
	dbErrs["55P02"] = "cant change runtime param"
	dbErrs["55P03"] = "lock not available"

	// Class 57 - Operator Intervention
	dbErrs["57000"] = "operator intervention"
	dbErrs["57014"] = "query canceled"
	dbErrs["57P01"] = "admin shutdown"
	dbErrs["57P02"] = "crash shutdown"
	dbErrs["57P03"] = "cannot connect now"
	dbErrs["57P04"] = "database dropped"

	// Class 58 - System Error (errors external to PostgreSQL itself)
	dbErrs["58000"] = "system error"
	dbErrs["58030"] = "io error"
	dbErrs["58P01"] = "undefined file"
	dbErrs["58P02"] = "duplicate file"

	// Class 72 - Snapshot Failure
	dbErrs["72000"] = "snapshot too old"

	// Class F0 - Configuration File Error
	dbErrs["F0000"] = "config file error"
	dbErrs["F0001"] = "lock file exists"

	// Class HV - Foreign Data Wrapper Error (SQL/MED)
	dbErrs["HV000"] = "fdw error"
	dbErrs["HV005"] = "fdw column name not found"
	dbErrs["HV002"] = "fdw dynamic parameter value needed"
	dbErrs["HV010"] = "fdw function sequence error"
	dbErrs["HV021"] = "fdw inconsistent descriptor information"
	dbErrs["HV024"] = "fdw invalid attribute value"
	dbErrs["HV007"] = "fdw invalid column name"
	dbErrs["HV008"] = "fdw invalid column number"
	dbErrs["HV004"] = "fdw invalid data type"
	dbErrs["HV006"] = "fdw invalid data type descriptors"
	dbErrs["HV091"] = "fdw invalid descriptor field identifier"
	dbErrs["HV00B"] = "fdw invalid handle"
	dbErrs["HV00C"] = "fdw invalid option index"
	dbErrs["HV00D"] = "fdw invalid option name"
	dbErrs["HV090"] = "fdw invalid string length or buffer length"
	dbErrs["HV00A"] = "fdw invalid string format"
	dbErrs["HV009"] = "fdw invalid use of null pointer"
	dbErrs["HV014"] = "fdw too many handles"
	dbErrs["HV001"] = "fdw out of memory"
	dbErrs["HV00P"] = "fdw no schemas"
	dbErrs["HV00J"] = "fdw option name not found"
	dbErrs["HV00K"] = "fdw reply handle"
	dbErrs["HV00Q"] = "fdw schema not found"
	dbErrs["HV00R"] = "fdw table not found"
	dbErrs["HV00L"] = "fdw unable to create execution"
	dbErrs["HV00M"] = "fdw unable to create reply"
	dbErrs["HV00N"] = "fdw unable to establish connection"

	// Class P0 - PL/pgSQL Error
	dbErrs["P0000"] = "plpgsql error"
	dbErrs["P0001"] = "raise exception"
	dbErrs["P0002"] = "no data found"
	dbErrs["P0003"] = "too many rows"
	dbErrs["P0004"] = "assert failure"

	// Class XX - Internal Error
	dbErrs["XX000"] = "internal error"
	dbErrs["XX001"] = "data corrupted"
	dbErrs["XX002"] = "index corrupted"
}

func DBErrs(code string) string {
	return dbErrs[code]
}
