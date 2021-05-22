package ora

import (
	m "github.com/gsiems/go-db-meta/model"
)

// Types defines the query for obtaining a list of user defined types
// for the (schema) parameters and returns the results of executing the
// query
func Types(db *m.DB, schema string) ([]m.Type, error) {

	q := `
WITH args AS (
    SELECT $1 AS schema_name
        FROM dual
)
SELECT sys_context ( 'userenv', 'DB_NAME' ) AS type_catalog,
        obj.owner AS type_schema,
        obj.type_name AS type_name,
        obj.owner AS type_owner,
        NULL AS comment
    FROM dba_types obj
    CROSS JOIN args
    WHERE ( obj.owner = args.schema_name
            OR args.schema_name IS NULL )
        AND obj.owner NOT IN (
                'APPQOSSYS', 'AWR_STAGE', 'CSMIG', 'CTXSYS', 'DBSNMP',
                'DIP', 'DMSYS', 'DSSYS', 'EXFSYS', 'LBACSYS', 'MDSYS',
                'OLAPSYS', 'ORACLE_OCM', 'ORDPLUGINS', 'ORDSYS', 'OUTLN',
                'PERFSTAT', 'PUBLIC', 'SQLTXPLAIN', 'SYS', 'SYSMAN',
                'SYSTEM', 'TRACESVR', 'TSMSYS', 'WMSYS', 'XDB' )
        AND obj.owner NOT LIKE '%$%'
`
	return db.Types(q, schema)
}

/*

| OWNER | TABLE_NAME     | COLUMN_NAME        | DATA_TYPE | DATA_LENGTH | COLUMN_ID |
| ----- | -------------- | ------------------ | --------- | ----------- | --------- |
| sys   | dba_types      | owner              | varchar2  |          30 |         1 |
| sys   | dba_types      | type_name          | varchar2  |          30 |         2 |
| sys   | dba_types      | type_oid           | raw       |          16 |         3 |
| sys   | dba_types      | typecode           | varchar2  |          30 |         4 |
| sys   | dba_types      | attributes         | number    |          22 |         5 |
| sys   | dba_types      | methods            | number    |          22 |         6 |
| sys   | dba_types      | predefined         | varchar2  |           3 |         7 |
| sys   | dba_types      | incomplete         | varchar2  |           3 |         8 |
| sys   | dba_types      | final              | varchar2  |           3 |         9 |
| sys   | dba_types      | instantiable       | varchar2  |           3 |        10 |
| sys   | dba_types      | supertype_owner    | varchar2  |          30 |        11 |
| sys   | dba_types      | supertype_name     | varchar2  |          30 |        12 |
| sys   | dba_types      | local_attributes   | number    |          22 |        13 |
| sys   | dba_types      | local_methods      | number    |          22 |        14 |
| sys   | dba_types      | typeid             | raw       |          16 |        15 |
| sys   | dba_type_attrs | owner              | varchar2  |          30 |         1 |
| sys   | dba_type_attrs | type_name          | varchar2  |          30 |         2 |
| sys   | dba_type_attrs | attr_name          | varchar2  |          30 |         3 |
| sys   | dba_type_attrs | attr_type_mod      | varchar2  |           7 |         4 |
| sys   | dba_type_attrs | attr_type_owner    | varchar2  |          30 |         5 |
| sys   | dba_type_attrs | attr_type_name     | varchar2  |          30 |         6 |
| sys   | dba_type_attrs | length             | number    |          22 |         7 |
| sys   | dba_type_attrs | precision          | number    |          22 |         8 |
| sys   | dba_type_attrs | scale              | number    |          22 |         9 |
| sys   | dba_type_attrs | character_set_name | varchar2  |          44 |        10 |
| sys   | dba_type_attrs | attr_no            | number    |          22 |        11 |
| sys   | dba_type_attrs | inherited          | varchar2  |          3  |        12 |

*/
