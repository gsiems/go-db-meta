package sqlite

import (
	"database/sql"

	m "github.com/gsiems/go-db-meta/model"
)

// ReferentialConstraints defines the query for obtaining the
// referential constraints for the (schemaName, tableName) parameters
// (as either the parent or child) and returns the results of executing
// the query
func ReferentialConstraints(db *sql.DB, schemaName, tableName string) ([]m.ReferentialConstraint, error) {

	var r []m.ReferentialConstraint

	catName, err := catalogName(db)
	if err != nil {
		return r, err
	}

	q := `
SELECT args.table_catalog,
        args.table_schema,
        con.table_name,
        con.column_names,
        idx_fk.index_name AS constraint_name,
        args.table_catalog AS ref_table_catalog,
        args.table_schema AS ref_table_schema,
        con.ref_table_name,
        con.ref_column_names,
        idx_uniq.index_name AS ref_constraint_name,
        con.match_option,
        con.update_rule,
        con.delete_rule,
        'YES' AS is_enforced,
        --is_deferrable,
        --initially_deferred,
        '' AS comments
    FROM (
        SELECT fk_col.table_name,
                group_concat ( fk_col.column_name, ', ' ) AS column_names,
                --fk_col.table_name || '_fk' || fk_col.id AS constraint_name,
                fk_col.ref_table_name,
                group_concat ( fk_col.ref_column_name, ', ' ) AS ref_column_names,
                fk_col.match_option,
                fk_col.update_rule,
                fk_col.delete_rule
            FROM (
                SELECT tab.name AS table_name,
                        con.id,
                        con."table" AS ref_table_name,
                        con."from" AS column_name,
                        con."to" AS ref_column_name,
                        con.seq AS ordinal_position,
                        con."match" AS match_option,
                        con.on_update AS update_rule,
                        con.on_delete AS delete_rule
                    FROM sqlite_master AS tab
                    JOIN pragma_foreign_key_list ( tab.name ) AS con
                    WHERE tab.type IN ( 'table' )
                        AND substr ( tab.name, 1, 7 ) <>  'sqlite_'
                    ORDER BY tab.name,
                        con.id,
                        con.seq
                ) AS fk_col
            GROUP BY fk_col.table_name,
                fk_col.ref_table_name,
                fk_col.id,
                fk_col.match_option,
                fk_col.update_rule,
                fk_col.delete_rule
        ) AS con
    CROSS JOIN (
        SELECT '` + catName.String + `' AS table_catalog,
                coalesce ( $1, '' ) AS table_schema,
                coalesce ( $2, '' ) AS table_name
        ) AS args
    LEFT JOIN (
        SELECT idx_col.table_name,
                idx_col.index_name,
                group_concat ( idx_col.column_name, ', ' ) AS column_names
            FROM (
                SELECT tab.tbl_name AS table_name,
                        tab.name AS index_name,
                        col.name AS column_name,
                        col.seqno AS ordinal_position
                    FROM sqlite_master AS tab
                    JOIN pragma_index_info ( tab.name ) AS col
                    WHERE tab.type IN ( 'index' )
                        AND substr ( tab.name, 1, 7 ) <>  'sqlite_'
                    ORDER BY tab.tbl_name,
                        tab.name,
                        col.seqno
                ) AS idx_col
            GROUP BY idx_col.table_name,
                idx_col.index_name
        ) AS idx_fk
        ON ( con.table_name = idx_fk.table_name
            AND con.column_names = idx_fk.column_names )
    LEFT JOIN (
        SELECT table_name,
                index_name,
                group_concat ( column_name, ', ' ) AS column_names
            FROM (
                SELECT m.name AS table_name,
                        con.name AS index_name,
                        col.name AS column_name,
                        col.seqno AS ordinal_position
                    FROM sqlite_master AS m
                    JOIN pragma_index_list ( m.name ) AS con
                    JOIN pragma_index_info ( con.name ) AS col
                    WHERE con."unique" = 1
                ) AS idx_col
            GROUP BY idx_col.table_name,
                idx_col.index_name
        ) AS idx_uniq
        ON ( con.table_name = idx_uniq.table_name
            AND con.column_names = idx_uniq.column_names )
    WHERE ( con.table_name = args.table_name OR args.table_name = '' )
        OR ( con.ref_table_name = args.table_name OR args.table_name = '' )
`
	return m.ReferentialConstraints(db, q, schemaName, tableName)
}
