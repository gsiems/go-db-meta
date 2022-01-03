package mariadb

import (
	"database/sql"

	m "github.com/gsiems/go-db-meta/model"
)

// ReferentialConstraints defines the query for obtaining the
// referential constraints for the (schemaName, tableName) parameters
// (as either the parent or child) and returns the results of executing
// the query
func ReferentialConstraints(db *sql.DB, schemaName, tableName string) ([]m.ReferentialConstraint, error) {

	q := `
SELECT col.table_catalog,
        col.table_schema,
        col.table_name,
        group_concat(col.column_name
            order by col.position_in_unique_constraint
            separator ', ') AS table_columns,
        con.constraint_name,
        con.unique_constraint_catalog AS ref_table_catalog,
        col.referenced_table_schema AS ref_table_schema,
        col.referenced_table_name AS ref_table_name,
        group_concat( col.referenced_column_name
            order by col.position_in_unique_constraint
            separator ', ' ) AS ref_table_columns,
        con.unique_constraint_name AS ref_constraint_name,
        con.match_option,
        con.update_rule,
        con.delete_rule,
        'YES' AS is_enforced,
        #is_deferrable,
        #initially_deferred,
        NULL AS comments
    FROM information_schema.referential_constraints con
    JOIN information_schema.key_column_usage col
         ON ( con.constraint_schema = col.table_schema
             AND con.table_name = col.table_name
             AND con.constraint_name = col.constraint_name )
    CROSS JOIN (
        SELECT coalesce ( ?, '' ) AS schema_name,
                coalesce ( ?, '' ) AS table_name
        ) AS args
    WHERE con.constraint_schema NOT IN ( 'information_schema', 'mysql', 'performance_schema', 'sys' )
        AND con.unique_constraint_schema NOT IN ( 'information_schema', 'mysql', 'performance_schema', 'sys' )
        AND ( ( ( col.table_schema = args.schema_name OR ( args.schema_name = '' AND args.table_name = '' ) )
                AND ( col.table_name = args.table_name OR args.table_name = '' ) )
            OR ( ( col.referenced_table_schema = args.schema_name OR ( args.schema_name = '' AND args.table_name = '' ) )
                AND ( col.referenced_table_name = args.table_name OR args.table_name = '' ) ) )
    GROUP BY col.table_catalog,
        col.table_schema,
        col.table_name,
        con.constraint_name,
        con.unique_constraint_catalog,
        col.referenced_table_schema,
        col.referenced_table_name,
        con.unique_constraint_name,
        con.match_option,
        con.update_rule,
        con.delete_rule
`
	return m.ReferentialConstraints(db, q, schemaName, tableName)
}


/*

CONSTRAINT_CATALOG CONSTRAINT_SCHEMA CONSTRAINT_NAME     TABLE_SCHEMA  TABLE_NAME   CONSTRAINT_TYPE
------------------ ----------------- ------------------- ------------- ------------ ---------------
def                classicmodels     PRIMARY             classicmodels employees    PRIMARY KEY
def                classicmodels     employees_ibfk_1    classicmodels employees    FOREIGN KEY
def                classicmodels     employees_ibfk_2    classicmodels employees    FOREIGN KEY
def                classicmodels     zzt                 classicmodels foo          UNIQUE
def                classicmodels     foo_ck              classicmodels foo          CHECK
def                classicmodels     PRIMARY             classicmodels payments     PRIMARY KEY
def                classicmodels     payments_ibfk_1     classicmodels payments     FOREIGN KEY
def                classicmodels     PRIMARY             classicmodels productlines PRIMARY KEY
def                classicmodels     PRIMARY             classicmodels products     PRIMARY KEY
def                classicmodels     products_ibfk_1     classicmodels products     FOREIGN KEY
def                classicmodels     PRIMARY             classicmodels offices      PRIMARY KEY
def                classicmodels     PRIMARY             classicmodels customers    PRIMARY KEY
def                classicmodels     customers_ibfk_1    classicmodels customers    FOREIGN KEY
def                classicmodels     PRIMARY             classicmodels orders       PRIMARY KEY
def                classicmodels     orders_ibfk_1       classicmodels orders       FOREIGN KEY
def                classicmodels     PRIMARY             classicmodels orderdetails PRIMARY KEY
def                classicmodels     orderdetails_ibfk_1 classicmodels orderdetails FOREIGN KEY
def                classicmodels     orderdetails_ibfk_2 classicmodels orderdetails FOREIGN KEY

*/
