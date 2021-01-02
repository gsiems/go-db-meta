# INFORMATION_SCHEMA

 * Ref: ISO standard 9075:2003 (draft) Part 11: Information and Definition Schemas
 * Not all DBMS engines have an information_schema
 * The information_schema for DBMS engines that do have an information_schema vary somewhat significantly
 * https://en.wikipedia.org/wiki/Information_schema
 * https://db-engines.com/en/ranking
 * Compared the information_schema for the following databases:
   * [H2](https://h2database.com/html/main.html)
   * [HyperSQL](http://hsqldb.org/)
   * [MariaDB](https://mariadb.com/)
   * [Microsoft SQL Server](https://docs.microsoft.com/en-us/sql/?view=sql-server-ver15)
   * [PostgreSQL](https://www.postgresql.org/)
 * Databases to compare were selected primarily by virtue of being readily available. Desire to support the database was also a factor.

# Data

 * Using data loaded into PostgreSQL for analysis [schemata.pgdump](schemata.pgdump)

## Databases to analyze and/or support

    SELECT DISTINCT coalesce ( ic.label, e.label ) AS "Label",
            e.engine_name AS "DBMS",
            CASE
                WHEN ic.label IS NULL THEN 'no'
                ELSE 'yes'
                END AS "Compare?",
            e.has_go_driver AS "Driver?",
            e.has_information_schema AS "IS?",
            CASE
                WHEN label = 'sql2003' THEN "SQL:2003 Standard"
                ELSE e.remarks
                END AS "Remarks",
            dbe.ranking AS "Ranking", -- as of Dec, 2020
            dbe.db_model AS "DB Model"
        FROM information_columns ic
        FULL JOIN engines e
        ON ( ic.label = e.label )
        LEFT JOIN db_engines_ranking dbe
            ON ( e.engine_name = dbe.engine_name )
        WHERE coalesce ( ic.label, e.label, '' ) <> ''
        ORDER BY dbe.ranking NULLS FIRST ;

| Label   | DBMS                 | Compare? | Driver? | IS? | Remarks                        | Ranking | DB Model                |
| ------- | -------------------- | -------- | ------- | --- | ------------------------------ | ------- | ----------------------- |
| sql2003 |                      | yes      |         |     |                                |         |                         |
| ora     | Oracle               | no       | yes     | no  |                                | 1       | Relational, Multi-model |
| mysql   | MySQL                | no       | yes     | yes | driver compatible with MariaDB | 2       | Relational, Multi-model |
| mssql   | Microsoft SQL Server | yes      | yes     | yes | Version 2019                   | 3       | Relational, Multi-model |
| pg      | PostgreSQL           | yes      | yes     | yes | Version 12.5                   | 4       | Relational, Multi-model |
| sqlite  | SQLite               | no       | yes     | no  | Version 3.23.1                 | 9       | Relational              |
| mariadb | MariaDB              | yes      | yes     | yes | Version 10.3                   | 12      | Relational, Multi-model |
| h2      | H2                   | yes      | yes     | yes | Version 1.4.200                | 49      | Relational              |
| hsqldb  | HyperSQL             | yes      |         | yes | Version 2.5.1                  | 67      | Relational              |


## Most supported INFORMATION_SCHEMA tables/views

 * for the databases analyzed

    WITH base AS (
        SELECT label,
                upper ( table_name ) AS table_name
            FROM information_columns
            WHERE upper (table_schema) = 'INFORMATION_SCHEMA'
            GROUP BY label,
                upper ( table_name )
    )
    SELECT table_name AS "Table Name",
            string_agg ( label, ', ' ORDER BY label desc ) AS "Matches",
            count (*) AS "Qty"
        FROM base
        GROUP BY table_name
        HAVING count (*) > 2
        ORDER BY 3 desc,
            1 ;

|
| Table Name                            | Matches                                 | Qty |
| ------------------------------------- | --------------------------------------- | --- |
| COLUMN_PRIVILEGES                     | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| COLUMNS                               | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| KEY_COLUMN_USAGE                      | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| REFERENTIAL_CONSTRAINTS               | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| SCHEMATA                              | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| TABLE_CONSTRAINTS                     | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| TABLE_PRIVILEGES                      | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| TABLES                                | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| VIEWS                                 | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| CHECK_CONSTRAINTS                     | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| COLLATIONS                            | sql2003, pg, mariadb, hsqldb, h2        | 5   |
| DOMAINS                               | sql2003, pg, mssql, hsqldb, h2          | 5   |
| PARAMETERS                            | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| ROUTINES                              | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| SEQUENCES                             | sql2003, pg, mssql, hsqldb, h2          | 5   |
| TRIGGERS                              | sql2003, pg, mariadb, hsqldb, h2        | 5   |
| APPLICABLE_ROLES                      | sql2003, pg, mariadb, hsqldb            | 4   |
| CHARACTER_SETS                        | sql2003, pg, mariadb, hsqldb            | 4   |
| COLUMN_DOMAIN_USAGE                   | sql2003, pg, mssql, hsqldb              | 4   |
| CONSTRAINT_COLUMN_USAGE               | sql2003, pg, mssql, hsqldb              | 4   |
| CONSTRAINT_TABLE_USAGE                | sql2003, pg, mssql, hsqldb              | 4   |
| DOMAIN_CONSTRAINTS                    | sql2003, pg, mssql, hsqldb              | 4   |
| ENABLED_ROLES                         | sql2003, pg, mariadb, hsqldb            | 4   |
| VIEW_COLUMN_USAGE                     | sql2003, pg, mssql, hsqldb              | 4   |
| VIEW_TABLE_USAGE                      | sql2003, pg, mssql, hsqldb              | 4   |
| ADMINISTRABLE_ROLE_AUTHORIZATIONS     | sql2003, pg, hsqldb                     | 3   |
| CHECK_CONSTRAINT_ROUTINE_USAGE        | sql2003, pg, hsqldb                     | 3   |
| COLLATION_CHARACTER_SET_APPLICABILITY | sql2003, pg, mariadb                    | 3   |
| COLUMN_COLUMN_USAGE                   | sql2003, pg, hsqldb                     | 3   |
| COLUMN_UDT_USAGE                      | sql2003, pg, hsqldb                     | 3   |
| DATA_TYPE_PRIVILEGES                  | sql2003, pg, hsqldb                     | 3   |
| ELEMENT_TYPES                         | sql2003, pg, hsqldb                     | 3   |
| INFORMATION_SCHEMA_CATALOG_NAME       | sql2003, pg, hsqldb                     | 3   |
| ROLE_COLUMN_GRANTS                    | sql2003, pg, hsqldb                     | 3   |
| ROLE_ROUTINE_GRANTS                   | sql2003, pg, hsqldb                     | 3   |
| ROLE_TABLE_GRANTS                     | sql2003, pg, hsqldb                     | 3   |
| ROLE_UDT_GRANTS                       | sql2003, pg, hsqldb                     | 3   |
| ROLE_USAGE_GRANTS                     | sql2003, pg, hsqldb                     | 3   |
| ROUTINE_PRIVILEGES                    | sql2003, pg, hsqldb                     | 3   |
| SQL_FEATURES                          | sql2003, pg, hsqldb                     | 3   |
| SQL_IMPLEMENTATION_INFO               | sql2003, pg, hsqldb                     | 3   |
| SQL_PACKAGES                          | sql2003, pg, hsqldb                     | 3   |
| SQL_PARTS                             | sql2003, pg, hsqldb                     | 3   |
| SQL_SIZING                            | sql2003, pg, hsqldb                     | 3   |
| SQL_SIZING_PROFILES                   | sql2003, pg, hsqldb                     | 3   |
| TRIGGERED_UPDATE_COLUMNS              | sql2003, pg, hsqldb                     | 3   |
| UDT_PRIVILEGES                        | sql2003, pg, hsqldb                     | 3   |
| USAGE_PRIVILEGES                      | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | sql2003, pg, hsqldb                     | 3   |
| VIEW_ROUTINE_USAGE                    | sql2003, pg, hsqldb                     | 3   |


## Supported INFORMATION_SCHEMA columns

    WITH base AS (
        SELECT label,
                upper ( table_name ) AS table_name,
                upper ( column_name ) AS column_name,
                CASE
                    WHEN label = 'sql2003' THEN ordinal_position
                    END AS ordinal_position
            FROM schemata.information_columns
            WHERE upper ( table_schema ) = 'INFORMATION_SCHEMA'
    )
    SELECT table_name AS "Table Name",
           column_name AS "Column Name",
           max ( ordinal_position ) AS "Position",
            string_agg ( label, ', ' ORDER BY label desc ) AS "Matches",
            count (*) AS "Qty"
        FROM base
        GROUP BY table_name,
                column_name
        HAVING count (*) > 1
        ORDER BY 1,
                3,
                2 ;

| Table Name                            | Column Name                       | Position | Matches                                 | Qty |
| ------------------------------------- | --------------------------------- | -------- | --------------------------------------- | --- |
| ADMINISTRABLE_ROLE_AUTHORIZATIONS     | GRANTEE                           | 1        | sql2003, pg, hsqldb                     | 3   |
| ADMINISTRABLE_ROLE_AUTHORIZATIONS     | ROLE_NAME                         | 2        | sql2003, pg, hsqldb                     | 3   |
| ADMINISTRABLE_ROLE_AUTHORIZATIONS     | IS_GRANTABLE                      | 3        | sql2003, pg, hsqldb                     | 3   |
| APPLICABLE_ROLES                      | GRANTEE                           | 1        | sql2003, pg, mariadb, hsqldb            | 4   |
| APPLICABLE_ROLES                      | ROLE_NAME                         | 2        | sql2003, pg, mariadb, hsqldb            | 4   |
| APPLICABLE_ROLES                      | IS_GRANTABLE                      | 3        | sql2003, pg, mariadb, hsqldb            | 4   |
| ASSERTIONS                            | CONSTRAINT_CATALOG                | 1        | sql2003, hsqldb                         | 2   |
| ASSERTIONS                            | CONSTRAINT_SCHEMA                 | 2        | sql2003, hsqldb                         | 2   |
| ASSERTIONS                            | CONSTRAINT_NAME                   | 3        | sql2003, hsqldb                         | 2   |
| ASSERTIONS                            | IS_DEFERRABLE                     | 4        | sql2003, hsqldb                         | 2   |
| ASSERTIONS                            | INITIALLY_DEFERRED                | 5        | sql2003, hsqldb                         | 2   |
| ATTRIBUTES                            | UDT_CATALOG                       | 1        | sql2003, pg                             | 2   |
| ATTRIBUTES                            | UDT_SCHEMA                        | 2        | sql2003, pg                             | 2   |
| ATTRIBUTES                            | UDT_NAME                          | 3        | sql2003, pg                             | 2   |
| ATTRIBUTES                            | ATTRIBUTE_NAME                    | 4        | sql2003, pg                             | 2   |
| ATTRIBUTES                            | ORDINAL_POSITION                  | 5        | sql2003, pg                             | 2   |
| ATTRIBUTES                            | ATTRIBUTE_DEFAULT                 | 6        | sql2003, pg                             | 2   |
| ATTRIBUTES                            | IS_NULLABLE                       | 7        | sql2003, pg                             | 2   |
| ATTRIBUTES                            | DATA_TYPE                         | 8        | sql2003, pg                             | 2   |
| ATTRIBUTES                            | CHARACTER_MAXIMUM_LENGTH          | 9        | sql2003, pg                             | 2   |
| ATTRIBUTES                            | CHARACTER_OCTET_LENGTH            | 10       | sql2003, pg                             | 2   |
| ATTRIBUTES                            | CHARACTER_SET_CATALOG             | 11       | sql2003, pg                             | 2   |
| ATTRIBUTES                            | CHARACTER_SET_SCHEMA              | 12       | sql2003, pg                             | 2   |
| ATTRIBUTES                            | CHARACTER_SET_NAME                | 13       | sql2003, pg                             | 2   |
| ATTRIBUTES                            | COLLATION_CATALOG                 | 14       | sql2003, pg                             | 2   |
| ATTRIBUTES                            | COLLATION_SCHEMA                  | 15       | sql2003, pg                             | 2   |
| ATTRIBUTES                            | COLLATION_NAME                    | 16       | sql2003, pg                             | 2   |
| ATTRIBUTES                            | NUMERIC_PRECISION                 | 17       | sql2003, pg                             | 2   |
| ATTRIBUTES                            | NUMERIC_PRECISION_RADIX           | 18       | sql2003, pg                             | 2   |
| ATTRIBUTES                            | NUMERIC_SCALE                     | 19       | sql2003, pg                             | 2   |
| ATTRIBUTES                            | DATETIME_PRECISION                | 20       | sql2003, pg                             | 2   |
| ATTRIBUTES                            | INTERVAL_TYPE                     | 21       | sql2003, pg                             | 2   |
| ATTRIBUTES                            | INTERVAL_PRECISION                | 22       | sql2003, pg                             | 2   |
| ATTRIBUTES                            | ATTRIBUTE_UDT_CATALOG             | 23       | sql2003, pg                             | 2   |
| ATTRIBUTES                            | ATTRIBUTE_UDT_SCHEMA              | 24       | sql2003, pg                             | 2   |
| ATTRIBUTES                            | ATTRIBUTE_UDT_NAME                | 25       | sql2003, pg                             | 2   |
| ATTRIBUTES                            | SCOPE_CATALOG                     | 26       | sql2003, pg                             | 2   |
| ATTRIBUTES                            | SCOPE_SCHEMA                      | 27       | sql2003, pg                             | 2   |
| ATTRIBUTES                            | SCOPE_NAME                        | 28       | sql2003, pg                             | 2   |
| ATTRIBUTES                            | MAXIMUM_CARDINALITY               | 29       | sql2003, pg                             | 2   |
| ATTRIBUTES                            | DTD_IDENTIFIER                    | 30       | sql2003, pg                             | 2   |
| ATTRIBUTES                            | IS_DERIVED_REFERENCE_ATTRIBUTE    | 31       | sql2003, pg                             | 2   |
| CHARACTER_SETS                        | CHARACTER_SET_CATALOG             | 1        | sql2003, pg, hsqldb                     | 3   |
| CHARACTER_SETS                        | CHARACTER_SET_SCHEMA              | 2        | sql2003, pg, hsqldb                     | 3   |
| CHARACTER_SETS                        | CHARACTER_SET_NAME                | 3        | sql2003, pg, mariadb, hsqldb            | 4   |
| CHARACTER_SETS                        | CHARACTER_REPERTOIRE              | 4        | sql2003, pg, hsqldb                     | 3   |
| CHARACTER_SETS                        | FORM_OF_USE                       | 5        | sql2003, pg, hsqldb                     | 3   |
| CHARACTER_SETS                        | DEFAULT_COLLATE_CATALOG           | 7        | sql2003, pg, hsqldb                     | 3   |
| CHARACTER_SETS                        | DEFAULT_COLLATE_SCHEMA            | 8        | sql2003, pg, hsqldb                     | 3   |
| CHARACTER_SETS                        | DEFAULT_COLLATE_NAME              | 9        | sql2003, pg, mariadb, hsqldb            | 4   |
| CHECK_CONSTRAINT_ROUTINE_USAGE        | CONSTRAINT_CATALOG                | 1        | sql2003, pg, hsqldb                     | 3   |
| CHECK_CONSTRAINT_ROUTINE_USAGE        | CONSTRAINT_SCHEMA                 | 2        | sql2003, pg, hsqldb                     | 3   |
| CHECK_CONSTRAINT_ROUTINE_USAGE        | CONSTRAINT_NAME                   | 3        | sql2003, pg, hsqldb                     | 3   |
| CHECK_CONSTRAINT_ROUTINE_USAGE        | SPECIFIC_CATALOG                  | 4        | sql2003, pg, hsqldb                     | 3   |
| CHECK_CONSTRAINT_ROUTINE_USAGE        | SPECIFIC_SCHEMA                   | 5        | sql2003, pg, hsqldb                     | 3   |
| CHECK_CONSTRAINT_ROUTINE_USAGE        | SPECIFIC_NAME                     | 6        | sql2003, pg, hsqldb                     | 3   |
| CHECK_CONSTRAINTS                     | CONSTRAINT_CATALOG                | 1        | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| CHECK_CONSTRAINTS                     | CONSTRAINT_SCHEMA                 | 2        | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| CHECK_CONSTRAINTS                     | CONSTRAINT_NAME                   | 3        | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| CHECK_CONSTRAINTS                     | CHECK_CLAUSE                      | 4        | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| COLLATION_CHARACTER_SET_APPLICABILITY | COLLATION_CATALOG                 | 1        | sql2003, pg                             | 2   |
| COLLATION_CHARACTER_SET_APPLICABILITY | COLLATION_SCHEMA                  | 2        | sql2003, pg                             | 2   |
| COLLATION_CHARACTER_SET_APPLICABILITY | COLLATION_NAME                    | 3        | sql2003, pg, mariadb                    | 3   |
| COLLATION_CHARACTER_SET_APPLICABILITY | CHARACTER_SET_CATALOG             | 4        | sql2003, pg                             | 2   |
| COLLATION_CHARACTER_SET_APPLICABILITY | CHARACTER_SET_SCHEMA              | 5        | sql2003, pg                             | 2   |
| COLLATION_CHARACTER_SET_APPLICABILITY | CHARACTER_SET_NAME                | 6        | sql2003, pg, mariadb                    | 3   |
| COLLATIONS                            | COLLATION_CATALOG                 | 1        | sql2003, pg, hsqldb                     | 3   |
| COLLATIONS                            | COLLATION_SCHEMA                  | 2        | sql2003, pg, hsqldb                     | 3   |
| COLLATIONS                            | COLLATION_NAME                    | 3        | sql2003, pg, mariadb, hsqldb            | 4   |
| COLLATIONS                            | PAD_ATTRIBUTE                     | 4        | sql2003, pg, hsqldb                     | 3   |
| COLUMN_COLUMN_USAGE                   | TABLE_CATALOG                     | 1        | sql2003, pg, hsqldb                     | 3   |
| COLUMN_COLUMN_USAGE                   | TABLE_SCHEMA                      | 2        | sql2003, pg, hsqldb                     | 3   |
| COLUMN_COLUMN_USAGE                   | TABLE_NAME                        | 3        | sql2003, pg, hsqldb                     | 3   |
| COLUMN_COLUMN_USAGE                   | COLUMN_NAME                       | 4        | sql2003, pg, hsqldb                     | 3   |
| COLUMN_COLUMN_USAGE                   | DEPENDENT_COLUMN                  | 5        | sql2003, pg, hsqldb                     | 3   |
| COLUMN_DOMAIN_USAGE                   | DOMAIN_CATALOG                    | 1        | sql2003, pg, mssql, hsqldb              | 4   |
| COLUMN_DOMAIN_USAGE                   | DOMAIN_SCHEMA                     | 2        | sql2003, pg, mssql, hsqldb              | 4   |
| COLUMN_DOMAIN_USAGE                   | DOMAIN_NAME                       | 3        | sql2003, pg, mssql, hsqldb              | 4   |
| COLUMN_DOMAIN_USAGE                   | TABLE_CATALOG                     | 4        | sql2003, pg, mssql, hsqldb              | 4   |
| COLUMN_DOMAIN_USAGE                   | TABLE_SCHEMA                      | 5        | sql2003, pg, mssql, hsqldb              | 4   |
| COLUMN_DOMAIN_USAGE                   | TABLE_NAME                        | 6        | sql2003, pg, mssql, hsqldb              | 4   |
| COLUMN_DOMAIN_USAGE                   | COLUMN_NAME                       | 7        | sql2003, pg, mssql, hsqldb              | 4   |
| COLUMN_PRIVILEGES                     | GRANTOR                           | 1        | sql2003, pg, mssql, hsqldb, h2          | 5   |
| COLUMN_PRIVILEGES                     | GRANTEE                           | 2        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| COLUMN_PRIVILEGES                     | TABLE_CATALOG                     | 3        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| COLUMN_PRIVILEGES                     | TABLE_SCHEMA                      | 4        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| COLUMN_PRIVILEGES                     | TABLE_NAME                        | 5        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| COLUMN_PRIVILEGES                     | COLUMN_NAME                       | 6        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| COLUMN_PRIVILEGES                     | PRIVILEGE_TYPE                    | 7        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| COLUMN_PRIVILEGES                     | IS_GRANTABLE                      | 8        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| COLUMNS                               | TABLE_CATALOG                     | 1        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| COLUMNS                               | TABLE_SCHEMA                      | 2        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| COLUMNS                               | TABLE_NAME                        | 3        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| COLUMNS                               | COLUMN_NAME                       | 4        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| COLUMNS                               | ORDINAL_POSITION                  | 5        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| COLUMNS                               | COLUMN_DEFAULT                    | 6        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| COLUMNS                               | IS_NULLABLE                       | 7        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| COLUMNS                               | DATA_TYPE                         | 8        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| COLUMNS                               | CHARACTER_MAXIMUM_LENGTH          | 9        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| COLUMNS                               | CHARACTER_OCTET_LENGTH            | 10       | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| COLUMNS                               | NUMERIC_PRECISION                 | 11       | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| COLUMNS                               | NUMERIC_PRECISION_RADIX           | 12       | sql2003, pg, mssql, hsqldb, h2          | 5   |
| COLUMNS                               | NUMERIC_SCALE                     | 13       | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| COLUMNS                               | DATETIME_PRECISION                | 14       | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| COLUMNS                               | INTERVAL_TYPE                     | 15       | sql2003, pg, hsqldb                     | 3   |
| COLUMNS                               | INTERVAL_PRECISION                | 16       | sql2003, pg, hsqldb                     | 3   |
| COLUMNS                               | CHARACTER_SET_CATALOG             | 17       | sql2003, pg, mssql, hsqldb              | 4   |
| COLUMNS                               | CHARACTER_SET_SCHEMA              | 18       | sql2003, pg, mssql, hsqldb              | 4   |
| COLUMNS                               | CHARACTER_SET_NAME                | 19       | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| COLUMNS                               | COLLATION_CATALOG                 | 20       | sql2003, pg, mssql, hsqldb              | 4   |
| COLUMNS                               | COLLATION_SCHEMA                  | 21       | sql2003, pg, mssql, hsqldb              | 4   |
| COLUMNS                               | COLLATION_NAME                    | 22       | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| COLUMNS                               | DOMAIN_CATALOG                    | 23       | sql2003, pg, mssql, hsqldb              | 4   |
| COLUMNS                               | DOMAIN_SCHEMA                     | 24       | sql2003, pg, mssql, hsqldb              | 4   |
| COLUMNS                               | DOMAIN_NAME                       | 25       | sql2003, pg, mssql, hsqldb              | 4   |
| COLUMNS                               | UDT_CATALOG                       | 26       | sql2003, pg, hsqldb                     | 3   |
| COLUMNS                               | UDT_SCHEMA                        | 27       | sql2003, pg, hsqldb                     | 3   |
| COLUMNS                               | UDT_NAME                          | 28       | sql2003, pg, hsqldb                     | 3   |
| COLUMNS                               | SCOPE_CATALOG                     | 29       | sql2003, pg, hsqldb                     | 3   |
| COLUMNS                               | SCOPE_SCHEMA                      | 30       | sql2003, pg, hsqldb                     | 3   |
| COLUMNS                               | SCOPE_NAME                        | 31       | sql2003, pg, hsqldb                     | 3   |
| COLUMNS                               | MAXIMUM_CARDINALITY               | 32       | sql2003, pg, hsqldb                     | 3   |
| COLUMNS                               | DTD_IDENTIFIER                    | 33       | sql2003, pg, hsqldb                     | 3   |
| COLUMNS                               | IS_SELF_REFERENCING               | 34       | sql2003, pg, hsqldb                     | 3   |
| COLUMNS                               | IS_IDENTITY                       | 35       | sql2003, pg, hsqldb                     | 3   |
| COLUMNS                               | IDENTITY_GENERATION               | 36       | sql2003, pg, hsqldb                     | 3   |
| COLUMNS                               | IDENTITY_START                    | 37       | sql2003, pg, hsqldb                     | 3   |
| COLUMNS                               | IDENTITY_INCREMENT                | 38       | sql2003, pg, hsqldb                     | 3   |
| COLUMNS                               | IDENTITY_MAXIMUM                  | 39       | sql2003, pg, hsqldb                     | 3   |
| COLUMNS                               | IDENTITY_MINIMUM                  | 40       | sql2003, pg, hsqldb                     | 3   |
| COLUMNS                               | IDENTITY_CYCLE                    | 41       | sql2003, pg, hsqldb                     | 3   |
| COLUMNS                               | IS_GENERATED                      | 42       | sql2003, pg, mariadb, hsqldb            | 4   |
| COLUMNS                               | GENERATION_EXPRESSION             | 43       | sql2003, pg, mariadb, hsqldb            | 4   |
| COLUMNS                               | IS_UPDATABLE                      | 44       | sql2003, pg, hsqldb                     | 3   |
| COLUMNS                               | COLUMN_TYPE                       |          | mariadb, h2                             | 2   |
| COLUMN_UDT_USAGE                      | UDT_CATALOG                       | 1        | sql2003, pg, hsqldb                     | 3   |
| COLUMN_UDT_USAGE                      | UDT_SCHEMA                        | 2        | sql2003, pg, hsqldb                     | 3   |
| COLUMN_UDT_USAGE                      | UDT_NAME                          | 3        | sql2003, pg, hsqldb                     | 3   |
| COLUMN_UDT_USAGE                      | TABLE_CATALOG                     | 4        | sql2003, pg, hsqldb                     | 3   |
| COLUMN_UDT_USAGE                      | TABLE_SCHEMA                      | 5        | sql2003, pg, hsqldb                     | 3   |
| COLUMN_UDT_USAGE                      | TABLE_NAME                        | 6        | sql2003, pg, hsqldb                     | 3   |
| COLUMN_UDT_USAGE                      | COLUMN_NAME                       | 7        | sql2003, pg, hsqldb                     | 3   |
| CONSTRAINT_COLUMN_USAGE               | TABLE_CATALOG                     | 1        | sql2003, pg, mssql, hsqldb              | 4   |
| CONSTRAINT_COLUMN_USAGE               | TABLE_SCHEMA                      | 2        | sql2003, pg, mssql, hsqldb              | 4   |
| CONSTRAINT_COLUMN_USAGE               | TABLE_NAME                        | 3        | sql2003, pg, mssql, hsqldb              | 4   |
| CONSTRAINT_COLUMN_USAGE               | COLUMN_NAME                       | 4        | sql2003, pg, mssql, hsqldb              | 4   |
| CONSTRAINT_COLUMN_USAGE               | CONSTRAINT_CATALOG                | 5        | sql2003, pg, mssql, hsqldb              | 4   |
| CONSTRAINT_COLUMN_USAGE               | CONSTRAINT_SCHEMA                 | 6        | sql2003, pg, mssql, hsqldb              | 4   |
| CONSTRAINT_COLUMN_USAGE               | CONSTRAINT_NAME                   | 7        | sql2003, pg, mssql, hsqldb              | 4   |
| CONSTRAINT_TABLE_USAGE                | TABLE_CATALOG                     | 1        | sql2003, pg, mssql, hsqldb              | 4   |
| CONSTRAINT_TABLE_USAGE                | TABLE_SCHEMA                      | 2        | sql2003, pg, mssql, hsqldb              | 4   |
| CONSTRAINT_TABLE_USAGE                | TABLE_NAME                        | 3        | sql2003, pg, mssql, hsqldb              | 4   |
| CONSTRAINT_TABLE_USAGE                | CONSTRAINT_CATALOG                | 4        | sql2003, pg, mssql, hsqldb              | 4   |
| CONSTRAINT_TABLE_USAGE                | CONSTRAINT_SCHEMA                 | 5        | sql2003, pg, mssql, hsqldb              | 4   |
| CONSTRAINT_TABLE_USAGE                | CONSTRAINT_NAME                   | 6        | sql2003, pg, mssql, hsqldb              | 4   |
| DATA_TYPE_PRIVILEGES                  | OBJECT_CATALOG                    | 1        | sql2003, pg, hsqldb                     | 3   |
| DATA_TYPE_PRIVILEGES                  | OBJECT_SCHEMA                     | 2        | sql2003, pg, hsqldb                     | 3   |
| DATA_TYPE_PRIVILEGES                  | OBJECT_NAME                       | 3        | sql2003, pg, hsqldb                     | 3   |
| DATA_TYPE_PRIVILEGES                  | OBJECT_TYPE                       | 4        | sql2003, pg, hsqldb                     | 3   |
| DATA_TYPE_PRIVILEGES                  | DTD_IDENTIFIER                    | 5        | sql2003, pg, hsqldb                     | 3   |
| DOMAIN_CONSTRAINTS                    | CONSTRAINT_CATALOG                | 1        | sql2003, pg, mssql, hsqldb              | 4   |
| DOMAIN_CONSTRAINTS                    | CONSTRAINT_SCHEMA                 | 2        | sql2003, pg, mssql, hsqldb              | 4   |
| DOMAIN_CONSTRAINTS                    | CONSTRAINT_NAME                   | 3        | sql2003, pg, mssql, hsqldb              | 4   |
| DOMAIN_CONSTRAINTS                    | DOMAIN_CATALOG                    | 4        | sql2003, pg, mssql, hsqldb              | 4   |
| DOMAIN_CONSTRAINTS                    | DOMAIN_SCHEMA                     | 5        | sql2003, pg, mssql, hsqldb              | 4   |
| DOMAIN_CONSTRAINTS                    | DOMAIN_NAME                       | 6        | sql2003, pg, mssql, hsqldb              | 4   |
| DOMAIN_CONSTRAINTS                    | IS_DEFERRABLE                     | 7        | sql2003, pg, mssql, hsqldb              | 4   |
| DOMAIN_CONSTRAINTS                    | INITIALLY_DEFERRED                | 8        | sql2003, pg, mssql, hsqldb              | 4   |
| DOMAINS                               | DOMAIN_CATALOG                    | 1        | sql2003, pg, mssql, hsqldb, h2          | 5   |
| DOMAINS                               | DOMAIN_SCHEMA                     | 2        | sql2003, pg, mssql, hsqldb, h2          | 5   |
| DOMAINS                               | DOMAIN_NAME                       | 3        | sql2003, pg, mssql, hsqldb, h2          | 5   |
| DOMAINS                               | DATA_TYPE                         | 4        | sql2003, pg, mssql, hsqldb, h2          | 5   |
| DOMAINS                               | CHARACTER_MAXIMUM_LENGTH          | 5        | sql2003, pg, mssql, hsqldb              | 4   |
| DOMAINS                               | CHARACTER_OCTET_LENGTH            | 6        | sql2003, pg, mssql, hsqldb              | 4   |
| DOMAINS                               | CHARACTER_SET_CATALOG             | 7        | sql2003, pg, mssql, hsqldb              | 4   |
| DOMAINS                               | CHARACTER_SET_SCHEMA              | 8        | sql2003, pg, mssql, hsqldb              | 4   |
| DOMAINS                               | CHARACTER_SET_NAME                | 9        | sql2003, pg, mssql, hsqldb              | 4   |
| DOMAINS                               | COLLATION_CATALOG                 | 10       | sql2003, pg, mssql, hsqldb              | 4   |
| DOMAINS                               | COLLATION_SCHEMA                  | 11       | sql2003, pg, mssql, hsqldb              | 4   |
| DOMAINS                               | COLLATION_NAME                    | 12       | sql2003, pg, mssql, hsqldb              | 4   |
| DOMAINS                               | NUMERIC_PRECISION                 | 13       | sql2003, pg, mssql, hsqldb              | 4   |
| DOMAINS                               | NUMERIC_PRECISION_RADIX           | 14       | sql2003, pg, mssql, hsqldb              | 4   |
| DOMAINS                               | NUMERIC_SCALE                     | 15       | sql2003, pg, mssql, hsqldb              | 4   |
| DOMAINS                               | DATETIME_PRECISION                | 16       | sql2003, pg, mssql, hsqldb              | 4   |
| DOMAINS                               | INTERVAL_TYPE                     | 17       | sql2003, pg, hsqldb                     | 3   |
| DOMAINS                               | INTERVAL_PRECISION                | 18       | sql2003, pg, hsqldb                     | 3   |
| DOMAINS                               | DOMAIN_DEFAULT                    | 19       | sql2003, pg, mssql, hsqldb              | 4   |
| DOMAINS                               | MAXIMUM_CARDINALITY               | 20       | sql2003, pg, hsqldb                     | 3   |
| DOMAINS                               | DTD_IDENTIFIER                    | 21       | sql2003, pg, hsqldb                     | 3   |
| ELEMENT_TYPES                         | OBJECT_CATALOG                    | 1        | sql2003, pg, hsqldb                     | 3   |
| ELEMENT_TYPES                         | OBJECT_SCHEMA                     | 2        | sql2003, pg, hsqldb                     | 3   |
| ELEMENT_TYPES                         | OBJECT_NAME                       | 3        | sql2003, pg, hsqldb                     | 3   |
| ELEMENT_TYPES                         | OBJECT_TYPE                       | 4        | sql2003, pg, hsqldb                     | 3   |
| ELEMENT_TYPES                         | COLLECTION_TYPE_IDENTIFIER        | 5        | sql2003, pg, hsqldb                     | 3   |
| ELEMENT_TYPES                         | DATA_TYPE                         | 6        | sql2003, pg, hsqldb                     | 3   |
| ELEMENT_TYPES                         | CHARACTER_MAXIMUM_LENGTH          | 7        | sql2003, pg, hsqldb                     | 3   |
| ELEMENT_TYPES                         | CHARACTER_OCTET_LENGTH            | 8        | sql2003, pg, hsqldb                     | 3   |
| ELEMENT_TYPES                         | CHARACTER_SET_CATALOG             | 9        | sql2003, pg, hsqldb                     | 3   |
| ELEMENT_TYPES                         | CHARACTER_SET_SCHEMA              | 10       | sql2003, pg, hsqldb                     | 3   |
| ELEMENT_TYPES                         | CHARACTER_SET_NAME                | 11       | sql2003, pg, hsqldb                     | 3   |
| ELEMENT_TYPES                         | COLLATION_CATALOG                 | 12       | sql2003, pg, hsqldb                     | 3   |
| ELEMENT_TYPES                         | COLLATION_SCHEMA                  | 13       | sql2003, pg, hsqldb                     | 3   |
| ELEMENT_TYPES                         | COLLATION_NAME                    | 14       | sql2003, pg, hsqldb                     | 3   |
| ELEMENT_TYPES                         | NUMERIC_PRECISION                 | 15       | sql2003, pg, hsqldb                     | 3   |
| ELEMENT_TYPES                         | NUMERIC_PRECISION_RADIX           | 16       | sql2003, pg, hsqldb                     | 3   |
| ELEMENT_TYPES                         | NUMERIC_SCALE                     | 17       | sql2003, pg, hsqldb                     | 3   |
| ELEMENT_TYPES                         | DATETIME_PRECISION                | 18       | sql2003, pg, hsqldb                     | 3   |
| ELEMENT_TYPES                         | INTERVAL_TYPE                     | 19       | sql2003, pg, hsqldb                     | 3   |
| ELEMENT_TYPES                         | INTERVAL_PRECISION                | 20       | sql2003, pg, hsqldb                     | 3   |
| ELEMENT_TYPES                         | UDT_CATALOG                       | 21       | sql2003, pg, hsqldb                     | 3   |
| ELEMENT_TYPES                         | UDT_SCHEMA                        | 22       | sql2003, pg, hsqldb                     | 3   |
| ELEMENT_TYPES                         | UDT_NAME                          | 23       | sql2003, pg, hsqldb                     | 3   |
| ELEMENT_TYPES                         | SCOPE_CATALOG                     | 24       | sql2003, pg, hsqldb                     | 3   |
| ELEMENT_TYPES                         | SCOPE_SCHEMA                      | 25       | sql2003, pg, hsqldb                     | 3   |
| ELEMENT_TYPES                         | SCOPE_NAME                        | 26       | sql2003, pg, hsqldb                     | 3   |
| ELEMENT_TYPES                         | MAXIMUM_CARDINALITY               | 27       | sql2003, pg, hsqldb                     | 3   |
| ELEMENT_TYPES                         | DTD_IDENTIFIER                    | 28       | sql2003, pg, hsqldb                     | 3   |
| ENABLED_ROLES                         | ROLE_NAME                         | 1        | sql2003, pg, mariadb, hsqldb            | 4   |
| INFORMATION_SCHEMA_CATALOG_NAME       | CATALOG_NAME                      | 1        | sql2003, pg, hsqldb                     | 3   |
| KEY_COLUMN_USAGE                      | CONSTRAINT_CATALOG                | 1        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| KEY_COLUMN_USAGE                      | CONSTRAINT_SCHEMA                 | 2        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| KEY_COLUMN_USAGE                      | CONSTRAINT_NAME                   | 3        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| KEY_COLUMN_USAGE                      | TABLE_CATALOG                     | 4        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| KEY_COLUMN_USAGE                      | TABLE_SCHEMA                      | 5        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| KEY_COLUMN_USAGE                      | TABLE_NAME                        | 6        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| KEY_COLUMN_USAGE                      | COLUMN_NAME                       | 7        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| KEY_COLUMN_USAGE                      | ORDINAL_POSITION                  | 8        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| KEY_COLUMN_USAGE                      | POSITION_IN_UNIQUE_CONSTRAINT     | 9        | sql2003, pg, mariadb, hsqldb, h2        | 5   |
| PARAMETERS                            | SPECIFIC_CATALOG                  | 1        | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| PARAMETERS                            | SPECIFIC_SCHEMA                   | 2        | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| PARAMETERS                            | SPECIFIC_NAME                     | 3        | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| PARAMETERS                            | ORDINAL_POSITION                  | 4        | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| PARAMETERS                            | PARAMETER_MODE                    | 5        | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| PARAMETERS                            | IS_RESULT                         | 6        | sql2003, pg, mssql, hsqldb              | 4   |
| PARAMETERS                            | AS_LOCATOR                        | 7        | sql2003, pg, mssql, hsqldb              | 4   |
| PARAMETERS                            | PARAMETER_NAME                    | 8        | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| PARAMETERS                            | FROM_SQL_SPECIFIC_CATALOG         | 9        | sql2003, hsqldb                         | 2   |
| PARAMETERS                            | FROM_SQL_SPECIFIC_SCHEMA          | 10       | sql2003, hsqldb                         | 2   |
| PARAMETERS                            | FROM_SQL_SPECIFIC_NAME            | 11       | sql2003, hsqldb                         | 2   |
| PARAMETERS                            | TO_SQL_SPECIFIC_CATALOG           | 12       | sql2003, hsqldb                         | 2   |
| PARAMETERS                            | TO_SQL_SPECIFIC_SCHEMA            | 13       | sql2003, hsqldb                         | 2   |
| PARAMETERS                            | TO_SQL_SPECIFIC_NAME              | 14       | sql2003, hsqldb                         | 2   |
| PARAMETERS                            | DATA_TYPE                         | 15       | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| PARAMETERS                            | CHARACTER_MAXIMUM_LENGTH          | 16       | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| PARAMETERS                            | CHARACTER_OCTET_LENGTH            | 17       | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| PARAMETERS                            | CHARACTER_SET_CATALOG             | 18       | sql2003, pg, mssql, hsqldb              | 4   |
| PARAMETERS                            | CHARACTER_SET_SCHEMA              | 19       | sql2003, pg, mssql, hsqldb              | 4   |
| PARAMETERS                            | CHARACTER_SET_NAME                | 20       | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| PARAMETERS                            | COLLATION_CATALOG                 | 21       | sql2003, pg, mssql, hsqldb              | 4   |
| PARAMETERS                            | COLLATION_SCHEMA                  | 22       | sql2003, pg, mssql, hsqldb              | 4   |
| PARAMETERS                            | COLLATION_NAME                    | 23       | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| PARAMETERS                            | NUMERIC_PRECISION                 | 24       | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| PARAMETERS                            | NUMERIC_PRECISION_RADIX           | 25       | sql2003, pg, mssql, hsqldb              | 4   |
| PARAMETERS                            | NUMERIC_SCALE                     | 26       | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| PARAMETERS                            | DATETIME_PRECISION                | 27       | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| PARAMETERS                            | INTERVAL_TYPE                     | 28       | sql2003, pg, mssql, hsqldb              | 4   |
| PARAMETERS                            | INTERVAL_PRECISION                | 29       | sql2003, pg, mssql, hsqldb              | 4   |
| PARAMETERS                            | UDT_CATALOG                       | 30       | sql2003, pg, hsqldb                     | 3   |
| PARAMETERS                            | UDT_SCHEMA                        | 31       | sql2003, pg, hsqldb                     | 3   |
| PARAMETERS                            | UDT_NAME                          | 32       | sql2003, pg, hsqldb                     | 3   |
| PARAMETERS                            | SCOPE_CATALOG                     | 33       | sql2003, pg, mssql, hsqldb              | 4   |
| PARAMETERS                            | SCOPE_SCHEMA                      | 34       | sql2003, pg, mssql, hsqldb              | 4   |
| PARAMETERS                            | SCOPE_NAME                        | 35       | sql2003, pg, mssql, hsqldb              | 4   |
| PARAMETERS                            | MAXIMUM_CARDINALITY               | 36       | sql2003, pg, hsqldb                     | 3   |
| PARAMETERS                            | DTD_IDENTIFIER                    | 37       | sql2003, pg, mariadb, hsqldb            | 4   |
| REFERENTIAL_CONSTRAINTS               | CONSTRAINT_CATALOG                | 1        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| REFERENTIAL_CONSTRAINTS               | CONSTRAINT_SCHEMA                 | 2        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| REFERENTIAL_CONSTRAINTS               | CONSTRAINT_NAME                   | 3        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| REFERENTIAL_CONSTRAINTS               | UNIQUE_CONSTRAINT_CATALOG         | 4        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| REFERENTIAL_CONSTRAINTS               | UNIQUE_CONSTRAINT_SCHEMA          | 5        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| REFERENTIAL_CONSTRAINTS               | UNIQUE_CONSTRAINT_NAME            | 6        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| REFERENTIAL_CONSTRAINTS               | MATCH_OPTION                      | 7        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| REFERENTIAL_CONSTRAINTS               | UPDATE_RULE                       | 8        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| REFERENTIAL_CONSTRAINTS               | DELETE_RULE                       | 9        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| ROLE_COLUMN_GRANTS                    | GRANTOR                           | 1        | sql2003, pg, hsqldb                     | 3   |
| ROLE_COLUMN_GRANTS                    | GRANTEE                           | 2        | sql2003, pg, hsqldb                     | 3   |
| ROLE_COLUMN_GRANTS                    | TABLE_CATALOG                     | 3        | sql2003, pg, hsqldb                     | 3   |
| ROLE_COLUMN_GRANTS                    | TABLE_SCHEMA                      | 4        | sql2003, pg, hsqldb                     | 3   |
| ROLE_COLUMN_GRANTS                    | TABLE_NAME                        | 5        | sql2003, pg, hsqldb                     | 3   |
| ROLE_COLUMN_GRANTS                    | COLUMN_NAME                       | 6        | sql2003, pg, hsqldb                     | 3   |
| ROLE_COLUMN_GRANTS                    | PRIVILEGE_TYPE                    | 7        | sql2003, pg, hsqldb                     | 3   |
| ROLE_COLUMN_GRANTS                    | IS_GRANTABLE                      | 8        | sql2003, pg, hsqldb                     | 3   |
| ROLE_ROUTINE_GRANTS                   | GRANTOR                           | 1        | sql2003, pg, hsqldb                     | 3   |
| ROLE_ROUTINE_GRANTS                   | GRANTEE                           | 2        | sql2003, pg, hsqldb                     | 3   |
| ROLE_ROUTINE_GRANTS                   | SPECIFIC_CATALOG                  | 3        | sql2003, pg, hsqldb                     | 3   |
| ROLE_ROUTINE_GRANTS                   | SPECIFIC_SCHEMA                   | 4        | sql2003, pg, hsqldb                     | 3   |
| ROLE_ROUTINE_GRANTS                   | SPECIFIC_NAME                     | 5        | sql2003, pg, hsqldb                     | 3   |
| ROLE_ROUTINE_GRANTS                   | ROUTINE_CATALOG                   | 6        | sql2003, pg, hsqldb                     | 3   |
| ROLE_ROUTINE_GRANTS                   | ROUTINE_SCHEMA                    | 7        | sql2003, pg, hsqldb                     | 3   |
| ROLE_ROUTINE_GRANTS                   | ROUTINE_NAME                      | 8        | sql2003, pg, hsqldb                     | 3   |
| ROLE_ROUTINE_GRANTS                   | PRIVILEGE_TYPE                    | 9        | sql2003, pg, hsqldb                     | 3   |
| ROLE_ROUTINE_GRANTS                   | IS_GRANTABLE                      | 10       | sql2003, pg, hsqldb                     | 3   |
| ROLE_TABLE_GRANTS                     | GRANTOR                           | 1        | sql2003, pg, hsqldb                     | 3   |
| ROLE_TABLE_GRANTS                     | GRANTEE                           | 2        | sql2003, pg, hsqldb                     | 3   |
| ROLE_TABLE_GRANTS                     | TABLE_CATALOG                     | 3        | sql2003, pg, hsqldb                     | 3   |
| ROLE_TABLE_GRANTS                     | TABLE_SCHEMA                      | 4        | sql2003, pg, hsqldb                     | 3   |
| ROLE_TABLE_GRANTS                     | TABLE_NAME                        | 5        | sql2003, pg, hsqldb                     | 3   |
| ROLE_TABLE_GRANTS                     | PRIVILEGE_TYPE                    | 6        | sql2003, pg, hsqldb                     | 3   |
| ROLE_TABLE_GRANTS                     | IS_GRANTABLE                      | 7        | sql2003, pg, hsqldb                     | 3   |
| ROLE_TABLE_GRANTS                     | WITH_HIERARCHY                    | 8        | sql2003, pg, hsqldb                     | 3   |
| ROLE_UDT_GRANTS                       | GRANTOR                           | 1        | sql2003, pg, hsqldb                     | 3   |
| ROLE_UDT_GRANTS                       | GRANTEE                           | 2        | sql2003, pg, hsqldb                     | 3   |
| ROLE_UDT_GRANTS                       | UDT_CATALOG                       | 3        | sql2003, pg, hsqldb                     | 3   |
| ROLE_UDT_GRANTS                       | UDT_SCHEMA                        | 4        | sql2003, pg, hsqldb                     | 3   |
| ROLE_UDT_GRANTS                       | UDT_NAME                          | 5        | sql2003, pg, hsqldb                     | 3   |
| ROLE_UDT_GRANTS                       | PRIVILEGE_TYPE                    | 6        | sql2003, pg, hsqldb                     | 3   |
| ROLE_UDT_GRANTS                       | IS_GRANTABLE                      | 7        | sql2003, pg, hsqldb                     | 3   |
| ROLE_USAGE_GRANTS                     | GRANTOR                           | 1        | sql2003, pg, hsqldb                     | 3   |
| ROLE_USAGE_GRANTS                     | GRANTEE                           | 2        | sql2003, pg, hsqldb                     | 3   |
| ROLE_USAGE_GRANTS                     | OBJECT_CATALOG                    | 3        | sql2003, pg, hsqldb                     | 3   |
| ROLE_USAGE_GRANTS                     | OBJECT_SCHEMA                     | 4        | sql2003, pg, hsqldb                     | 3   |
| ROLE_USAGE_GRANTS                     | OBJECT_NAME                       | 5        | sql2003, pg, hsqldb                     | 3   |
| ROLE_USAGE_GRANTS                     | OBJECT_TYPE                       | 6        | sql2003, pg, hsqldb                     | 3   |
| ROLE_USAGE_GRANTS                     | PRIVILEGE_TYPE                    | 7        | sql2003, pg, hsqldb                     | 3   |
| ROLE_USAGE_GRANTS                     | IS_GRANTABLE                      | 8        | sql2003, pg, hsqldb                     | 3   |
| ROUTINE_COLUMN_USAGE                  | SPECIFIC_CATALOG                  | 1        | sql2003, hsqldb                         | 2   |
| ROUTINE_COLUMN_USAGE                  | SPECIFIC_SCHEMA                   | 2        | sql2003, hsqldb                         | 2   |
| ROUTINE_COLUMN_USAGE                  | SPECIFIC_NAME                     | 3        | sql2003, hsqldb                         | 2   |
| ROUTINE_COLUMN_USAGE                  | ROUTINE_CATALOG                   | 4        | sql2003, hsqldb                         | 2   |
| ROUTINE_COLUMN_USAGE                  | ROUTINE_SCHEMA                    | 5        | sql2003, hsqldb                         | 2   |
| ROUTINE_COLUMN_USAGE                  | ROUTINE_NAME                      | 6        | sql2003, hsqldb                         | 2   |
| ROUTINE_COLUMN_USAGE                  | TABLE_CATALOG                     | 7        | sql2003, hsqldb                         | 2   |
| ROUTINE_COLUMN_USAGE                  | TABLE_SCHEMA                      | 8        | sql2003, hsqldb                         | 2   |
| ROUTINE_COLUMN_USAGE                  | TABLE_NAME                        | 9        | sql2003, hsqldb                         | 2   |
| ROUTINE_COLUMN_USAGE                  | COLUMN_NAME                       | 10       | sql2003, hsqldb                         | 2   |
| ROUTINE_PRIVILEGES                    | GRANTOR                           | 1        | sql2003, pg, hsqldb                     | 3   |
| ROUTINE_PRIVILEGES                    | GRANTEE                           | 2        | sql2003, pg, hsqldb                     | 3   |
| ROUTINE_PRIVILEGES                    | SPECIFIC_CATALOG                  | 3        | sql2003, pg, hsqldb                     | 3   |
| ROUTINE_PRIVILEGES                    | SPECIFIC_SCHEMA                   | 4        | sql2003, pg, hsqldb                     | 3   |
| ROUTINE_PRIVILEGES                    | SPECIFIC_NAME                     | 5        | sql2003, pg, hsqldb                     | 3   |
| ROUTINE_PRIVILEGES                    | ROUTINE_CATALOG                   | 6        | sql2003, pg, hsqldb                     | 3   |
| ROUTINE_PRIVILEGES                    | ROUTINE_SCHEMA                    | 7        | sql2003, pg, hsqldb                     | 3   |
| ROUTINE_PRIVILEGES                    | ROUTINE_NAME                      | 8        | sql2003, pg, hsqldb                     | 3   |
| ROUTINE_PRIVILEGES                    | PRIVILEGE_TYPE                    | 9        | sql2003, pg, hsqldb                     | 3   |
| ROUTINE_PRIVILEGES                    | IS_GRANTABLE                      | 10       | sql2003, pg, hsqldb                     | 3   |
| ROUTINE_ROUTINE_USAGE                 | SPECIFIC_CATALOG                  | 1        | sql2003, hsqldb                         | 2   |
| ROUTINE_ROUTINE_USAGE                 | SPECIFIC_SCHEMA                   | 2        | sql2003, hsqldb                         | 2   |
| ROUTINE_ROUTINE_USAGE                 | SPECIFIC_NAME                     | 3        | sql2003, hsqldb                         | 2   |
| ROUTINE_ROUTINE_USAGE                 | ROUTINE_CATALOG                   | 4        | sql2003, hsqldb                         | 2   |
| ROUTINE_ROUTINE_USAGE                 | ROUTINE_SCHEMA                    | 5        | sql2003, hsqldb                         | 2   |
| ROUTINE_ROUTINE_USAGE                 | ROUTINE_NAME                      | 6        | sql2003, hsqldb                         | 2   |
| ROUTINES                              | SPECIFIC_CATALOG                  | 1        | sql2003, pg, mssql, hsqldb              | 4   |
| ROUTINES                              | SPECIFIC_SCHEMA                   | 2        | sql2003, pg, mssql, hsqldb              | 4   |
| ROUTINES                              | SPECIFIC_NAME                     | 3        | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| ROUTINES                              | ROUTINE_CATALOG                   | 4        | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| ROUTINES                              | ROUTINE_SCHEMA                    | 5        | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| ROUTINES                              | ROUTINE_NAME                      | 6        | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| ROUTINES                              | ROUTINE_TYPE                      | 7        | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| ROUTINES                              | MODULE_CATALOG                    | 8        | sql2003, pg, mssql, hsqldb              | 4   |
| ROUTINES                              | MODULE_SCHEMA                     | 9        | sql2003, pg, mssql, hsqldb              | 4   |
| ROUTINES                              | MODULE_NAME                       | 10       | sql2003, pg, mssql, hsqldb              | 4   |
| ROUTINES                              | UDT_CATALOG                       | 11       | sql2003, pg, mssql, hsqldb              | 4   |
| ROUTINES                              | UDT_SCHEMA                        | 12       | sql2003, pg, mssql, hsqldb              | 4   |
| ROUTINES                              | UDT_NAME                          | 13       | sql2003, pg, mssql, hsqldb              | 4   |
| ROUTINES                              | DATA_TYPE                         | 14       | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| ROUTINES                              | CHARACTER_MAXIMUM_LENGTH          | 15       | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| ROUTINES                              | CHARACTER_OCTET_LENGTH            | 16       | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| ROUTINES                              | CHARACTER_SET_CATALOG             | 17       | sql2003, pg, mssql, hsqldb              | 4   |
| ROUTINES                              | CHARACTER_SET_SCHEMA              | 18       | sql2003, pg, mssql, hsqldb              | 4   |
| ROUTINES                              | CHARACTER_SET_NAME                | 19       | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| ROUTINES                              | COLLATION_CATALOG                 | 20       | sql2003, pg, mssql, hsqldb              | 4   |
| ROUTINES                              | COLLATION_SCHEMA                  | 21       | sql2003, pg, mssql, hsqldb              | 4   |
| ROUTINES                              | COLLATION_NAME                    | 22       | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| ROUTINES                              | NUMERIC_PRECISION                 | 23       | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| ROUTINES                              | NUMERIC_PRECISION_RADIX           | 24       | sql2003, pg, mssql, hsqldb              | 4   |
| ROUTINES                              | NUMERIC_SCALE                     | 25       | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| ROUTINES                              | DATETIME_PRECISION                | 26       | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| ROUTINES                              | INTERVAL_TYPE                     | 27       | sql2003, pg, mssql, hsqldb              | 4   |
| ROUTINES                              | INTERVAL_PRECISION                | 28       | sql2003, pg, mssql, hsqldb              | 4   |
| ROUTINES                              | TYPE_UDT_CATALOG                  | 29       | sql2003, pg, mssql, hsqldb              | 4   |
| ROUTINES                              | TYPE_UDT_SCHEMA                   | 30       | sql2003, pg, mssql, hsqldb              | 4   |
| ROUTINES                              | TYPE_UDT_NAME                     | 31       | sql2003, pg, mssql, hsqldb              | 4   |
| ROUTINES                              | SCOPE_CATALOG                     | 32       | sql2003, pg, mssql, hsqldb              | 4   |
| ROUTINES                              | SCOPE_SCHEMA                      | 33       | sql2003, pg, mssql, hsqldb              | 4   |
| ROUTINES                              | SCOPE_NAME                        | 34       | sql2003, pg, mssql, hsqldb              | 4   |
| ROUTINES                              | MAXIMUM_CARDINALITY               | 35       | sql2003, pg, mssql, hsqldb              | 4   |
| ROUTINES                              | DTD_IDENTIFIER                    | 36       | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| ROUTINES                              | ROUTINE_BODY                      | 37       | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| ROUTINES                              | ROUTINE_DEFINITION                | 38       | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| ROUTINES                              | EXTERNAL_NAME                     | 39       | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| ROUTINES                              | EXTERNAL_LANGUAGE                 | 40       | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| ROUTINES                              | PARAMETER_STYLE                   | 41       | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| ROUTINES                              | IS_DETERMINISTIC                  | 42       | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| ROUTINES                              | SQL_DATA_ACCESS                   | 43       | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| ROUTINES                              | IS_NULL_CALL                      | 44       | sql2003, pg, mssql, hsqldb              | 4   |
| ROUTINES                              | SQL_PATH                          | 45       | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| ROUTINES                              | SCHEMA_LEVEL_ROUTINE              | 46       | sql2003, pg, mssql, hsqldb              | 4   |
| ROUTINES                              | MAX_DYNAMIC_RESULT_SETS           | 47       | sql2003, pg, mssql, hsqldb              | 4   |
| ROUTINES                              | IS_USER_DEFINED_CAST              | 48       | sql2003, pg, mssql, hsqldb              | 4   |
| ROUTINES                              | IS_IMPLICITLY_INVOCABLE           | 49       | sql2003, pg, mssql, hsqldb              | 4   |
| ROUTINES                              | SECURITY_TYPE                     | 50       | sql2003, pg, mariadb, hsqldb            | 4   |
| ROUTINES                              | TO_SQL_SPECIFIC_CATALOG           | 51       | sql2003, pg, hsqldb                     | 3   |
| ROUTINES                              | TO_SQL_SPECIFIC_SCHEMA            | 52       | sql2003, pg, hsqldb                     | 3   |
| ROUTINES                              | TO_SQL_SPECIFIC_NAME              | 53       | sql2003, pg, hsqldb                     | 3   |
| ROUTINES                              | AS_LOCATOR                        | 54       | sql2003, pg, hsqldb                     | 3   |
| ROUTINES                              | CREATED                           | 55       | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| ROUTINES                              | LAST_ALTERED                      | 56       | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| ROUTINES                              | NEW_SAVEPOINT_LEVEL               | 57       | sql2003, pg, hsqldb                     | 3   |
| ROUTINES                              | IS_UDT_DEPENDENT                  | 58       | sql2003, pg, hsqldb                     | 3   |
| ROUTINES                              | RESULT_CAST_FROM_DATA_TYPE        | 59       | sql2003, pg, hsqldb                     | 3   |
| ROUTINES                              | RESULT_CAST_AS_LOCATOR            | 60       | sql2003, pg, hsqldb                     | 3   |
| ROUTINES                              | RESULT_CAST_CHAR_MAX_LENGTH       | 61       | sql2003, pg, hsqldb                     | 3   |
| ROUTINES                              | RESULT_CAST_CHAR_OCTET_LENGTH     | 62       | sql2003, pg, hsqldb                     | 3   |
| ROUTINES                              | RESULT_CAST_CHAR_SET_CATALOG      | 63       | sql2003, pg, hsqldb                     | 3   |
| ROUTINES                              | RESULT_CAST_CHAR_SET_SCHEMA       | 64       | sql2003, pg, hsqldb                     | 3   |
| ROUTINES                              | RESULT_CAST_CHARACTER_SET_NAME    | 65       | sql2003, hsqldb                         | 2   |
| ROUTINES                              | RESULT_CAST_COLLATION_CATALOG     | 66       | sql2003, pg, hsqldb                     | 3   |
| ROUTINES                              | RESULT_CAST_COLLATION_SCHEMA      | 67       | sql2003, pg, hsqldb                     | 3   |
| ROUTINES                              | RESULT_CAST_COLLATION_NAME        | 68       | sql2003, pg, hsqldb                     | 3   |
| ROUTINES                              | RESULT_CAST_NUMERIC_PRECISION     | 69       | sql2003, pg, hsqldb                     | 3   |
| ROUTINES                              | RESULT_CAST_NUMERIC_RADIX         | 70       | sql2003, hsqldb                         | 2   |
| ROUTINES                              | RESULT_CAST_NUMERIC_SCALE         | 71       | sql2003, pg, hsqldb                     | 3   |
| ROUTINES                              | RESULT_CAST_DATETIME_PRECISION    | 72       | sql2003, pg, hsqldb                     | 3   |
| ROUTINES                              | RESULT_CAST_INTERVAL_TYPE         | 73       | sql2003, pg, hsqldb                     | 3   |
| ROUTINES                              | RESULT_CAST_INTERVAL_PRECISION    | 74       | sql2003, pg, hsqldb                     | 3   |
| ROUTINES                              | RESULT_CAST_TYPE_UDT_CATALOG      | 75       | sql2003, pg, hsqldb                     | 3   |
| ROUTINES                              | RESULT_CAST_TYPE_UDT_SCHEMA       | 76       | sql2003, pg, hsqldb                     | 3   |
| ROUTINES                              | RESULT_CAST_TYPE_UDT_NAME         | 77       | sql2003, pg, hsqldb                     | 3   |
| ROUTINES                              | RESULT_CAST_SCOPE_CATALOG         | 78       | sql2003, pg, hsqldb                     | 3   |
| ROUTINES                              | RESULT_CAST_SCOPE_SCHEMA          | 79       | sql2003, pg, hsqldb                     | 3   |
| ROUTINES                              | RESULT_CAST_SCOPE_NAME            | 80       | sql2003, pg, hsqldb                     | 3   |
| ROUTINES                              | RESULT_CAST_MAX_CARDINALITY       | 81       | sql2003, hsqldb                         | 2   |
| ROUTINES                              | RESULT_CAST_DTD_IDENTIFIER        | 82       | sql2003, pg, hsqldb                     | 3   |
| ROUTINE_SEQUENCE_USAGE                | SPECIFIC_CATALOG                  | 1        | sql2003, hsqldb                         | 2   |
| ROUTINE_SEQUENCE_USAGE                | SPECIFIC_SCHEMA                   | 2        | sql2003, hsqldb                         | 2   |
| ROUTINE_SEQUENCE_USAGE                | SPECIFIC_NAME                     | 3        | sql2003, hsqldb                         | 2   |
| ROUTINE_SEQUENCE_USAGE                | SEQUENCE_CATALOG                  | 4        | sql2003, hsqldb                         | 2   |
| ROUTINE_SEQUENCE_USAGE                | SEQUENCE_SCHEMA                   | 5        | sql2003, hsqldb                         | 2   |
| ROUTINE_SEQUENCE_USAGE                | SEQUENCE_NAME                     | 6        | sql2003, hsqldb                         | 2   |
| ROUTINE_TABLE_USAGE                   | SPECIFIC_CATALOG                  | 1        | sql2003, hsqldb                         | 2   |
| ROUTINE_TABLE_USAGE                   | SPECIFIC_SCHEMA                   | 2        | sql2003, hsqldb                         | 2   |
| ROUTINE_TABLE_USAGE                   | SPECIFIC_NAME                     | 3        | sql2003, hsqldb                         | 2   |
| ROUTINE_TABLE_USAGE                   | ROUTINE_CATALOG                   | 4        | sql2003, hsqldb                         | 2   |
| ROUTINE_TABLE_USAGE                   | ROUTINE_SCHEMA                    | 5        | sql2003, hsqldb                         | 2   |
| ROUTINE_TABLE_USAGE                   | ROUTINE_NAME                      | 6        | sql2003, hsqldb                         | 2   |
| ROUTINE_TABLE_USAGE                   | TABLE_CATALOG                     | 7        | sql2003, hsqldb                         | 2   |
| ROUTINE_TABLE_USAGE                   | TABLE_SCHEMA                      | 8        | sql2003, hsqldb                         | 2   |
| ROUTINE_TABLE_USAGE                   | TABLE_NAME                        | 9        | sql2003, hsqldb                         | 2   |
| SCHEMATA                              | CATALOG_NAME                      | 1        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| SCHEMATA                              | SCHEMA_NAME                       | 2        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| SCHEMATA                              | SCHEMA_OWNER                      | 3        | sql2003, pg, mssql, hsqldb, h2          | 5   |
| SCHEMATA                              | DEFAULT_CHARACTER_SET_CATALOG     | 4        | sql2003, pg, mssql, hsqldb              | 4   |
| SCHEMATA                              | DEFAULT_CHARACTER_SET_SCHEMA      | 5        | sql2003, pg, mssql, hsqldb              | 4   |
| SCHEMATA                              | DEFAULT_CHARACTER_SET_NAME        | 6        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| SCHEMATA                              | SQL_PATH                          | 7        | sql2003, pg, mariadb, hsqldb            | 4   |
| SCHEMATA                              | DEFAULT_COLLATION_NAME            |          | mariadb, h2                             | 2   |
| SEQUENCES                             | SEQUENCE_CATALOG                  | 1        | sql2003, pg, mssql, hsqldb, h2          | 5   |
| SEQUENCES                             | SEQUENCE_SCHEMA                   | 2        | sql2003, pg, mssql, hsqldb, h2          | 5   |
| SEQUENCES                             | SEQUENCE_NAME                     | 3        | sql2003, pg, mssql, hsqldb, h2          | 5   |
| SEQUENCES                             | DATA_TYPE                         | 4        | sql2003, pg, mssql, hsqldb              | 4   |
| SEQUENCES                             | NUMERIC_PRECISION                 | 5        | sql2003, pg, mssql, hsqldb              | 4   |
| SEQUENCES                             | NUMERIC_PRECISION_RADIX           | 6        | sql2003, pg, mssql, hsqldb              | 4   |
| SEQUENCES                             | NUMERIC_SCALE                     | 7        | sql2003, pg, mssql, hsqldb              | 4   |
| SEQUENCES                             | MAXIMUM_VALUE                     | 8        | sql2003, pg, mssql, hsqldb              | 4   |
| SEQUENCES                             | MINIMUM_VALUE                     | 9        | sql2003, pg, mssql, hsqldb              | 4   |
| SEQUENCES                             | INCREMENT                         | 10       | sql2003, pg, mssql, hsqldb, h2          | 5   |
| SEQUENCES                             | CYCLE_OPTION                      | 11       | sql2003, pg, mssql, hsqldb              | 4   |
| SEQUENCES                             | DECLARED_DATA_TYPE                |          | mssql, hsqldb                           | 2   |
| SEQUENCES                             | DECLARED_NUMERIC_PRECISION        |          | mssql, hsqldb                           | 2   |
| SEQUENCES                             | DECLARED_NUMERIC_SCALE            |          | mssql, hsqldb                           | 2   |
| SEQUENCES                             | START_VALUE                       |          | pg, mssql, hsqldb                       | 3   |
| SQL_FEATURES                          | FEATURE_ID                        | 1        | sql2003, pg, hsqldb                     | 3   |
| SQL_FEATURES                          | FEATURE_NAME                      | 2        | sql2003, pg, hsqldb                     | 3   |
| SQL_FEATURES                          | SUB_FEATURE_ID                    | 3        | sql2003, pg, hsqldb                     | 3   |
| SQL_FEATURES                          | SUB_FEATURE_NAME                  | 4        | sql2003, pg, hsqldb                     | 3   |
| SQL_FEATURES                          | IS_SUPPORTED                      | 5        | sql2003, pg, hsqldb                     | 3   |
| SQL_FEATURES                          | IS_VERIFIED_BY                    | 6        | sql2003, pg, hsqldb                     | 3   |
| SQL_FEATURES                          | COMMENTS                          | 7        | sql2003, pg, hsqldb                     | 3   |
| SQL_IMPLEMENTATION_INFO               | IMPLEMENTATION_INFO_ID            | 1        | sql2003, pg, hsqldb                     | 3   |
| SQL_IMPLEMENTATION_INFO               | IMPLEMENTATION_INFO_NAME          | 2        | sql2003, pg, hsqldb                     | 3   |
| SQL_IMPLEMENTATION_INFO               | INTEGER_VALUE                     | 3        | sql2003, pg, hsqldb                     | 3   |
| SQL_IMPLEMENTATION_INFO               | CHARACTER_VALUE                   | 4        | sql2003, pg, hsqldb                     | 3   |
| SQL_IMPLEMENTATION_INFO               | COMMENTS                          | 5        | sql2003, pg, hsqldb                     | 3   |
| SQL_LANGUAGES                         | SQL_LANGUAGE_SOURCE               | 1        | sql2003, pg                             | 2   |
| SQL_LANGUAGES                         | SQL_LANGUAGE_YEAR                 | 2        | sql2003, pg                             | 2   |
| SQL_LANGUAGES                         | SQL_LANGUAGE_CONFORMANCE          | 3        | sql2003, pg                             | 2   |
| SQL_LANGUAGES                         | SQL_LANGUAGE_INTEGRITY            | 4        | sql2003, pg                             | 2   |
| SQL_LANGUAGES                         | SQL_LANGUAGE_IMPLEMENTATION       | 5        | sql2003, pg                             | 2   |
| SQL_LANGUAGES                         | SQL_LANGUAGE_BINDING_STYLE        | 6        | sql2003, pg                             | 2   |
| SQL_LANGUAGES                         | SQL_LANGUAGE_PROGRAMMING_LANGUAGE | 7        | sql2003, pg                             | 2   |
| SQL_PACKAGES                          | FEATURE_ID                        | 1        | sql2003, pg                             | 2   |
| SQL_PACKAGES                          | FEATURE_NAME                      | 2        | sql2003, pg                             | 2   |
| SQL_PACKAGES                          | IS_SUPPORTED                      | 3        | sql2003, pg, hsqldb                     | 3   |
| SQL_PACKAGES                          | IS_VERIFIED_BY                    | 4        | sql2003, pg, hsqldb                     | 3   |
| SQL_PACKAGES                          | COMMENTS                          | 5        | sql2003, pg, hsqldb                     | 3   |
| SQL_PARTS                             | FEATURE_ID                        | 1        | sql2003, pg                             | 2   |
| SQL_PARTS                             | FEATURE_NAME                      | 2        | sql2003, pg                             | 2   |
| SQL_PARTS                             | IS_SUPPORTED                      | 3        | sql2003, pg, hsqldb                     | 3   |
| SQL_PARTS                             | IS_VERIFIED_BY                    | 4        | sql2003, pg, hsqldb                     | 3   |
| SQL_PARTS                             | COMMENTS                          | 5        | sql2003, pg, hsqldb                     | 3   |
| SQL_SIZING                            | SIZING_ID                         | 1        | sql2003, pg, hsqldb                     | 3   |
| SQL_SIZING                            | SIZING_NAME                       | 2        | sql2003, pg, hsqldb                     | 3   |
| SQL_SIZING                            | SUPPORTED_VALUE                   | 3        | sql2003, pg, hsqldb                     | 3   |
| SQL_SIZING                            | COMMENTS                          | 4        | sql2003, pg, hsqldb                     | 3   |
| SQL_SIZING_PROFILES                   | SIZING_ID                         | 1        | sql2003, pg, hsqldb                     | 3   |
| SQL_SIZING_PROFILES                   | SIZING_NAME                       | 2        | sql2003, pg, hsqldb                     | 3   |
| SQL_SIZING_PROFILES                   | PROFILE_ID                        | 3        | sql2003, pg, hsqldb                     | 3   |
| SQL_SIZING_PROFILES                   | PROFILE_NAME                      | 4        | sql2003, hsqldb                         | 2   |
| SQL_SIZING_PROFILES                   | REQUIRED_VALUE                    | 5        | sql2003, pg, hsqldb                     | 3   |
| SQL_SIZING_PROFILES                   | COMMENTS                          | 6        | sql2003, pg, hsqldb                     | 3   |
| TABLE_CONSTRAINTS                     | CONSTRAINT_CATALOG                | 1        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| TABLE_CONSTRAINTS                     | CONSTRAINT_SCHEMA                 | 2        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| TABLE_CONSTRAINTS                     | CONSTRAINT_NAME                   | 3        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| TABLE_CONSTRAINTS                     | TABLE_CATALOG                     | 4        | sql2003, pg, mssql, hsqldb, h2          | 5   |
| TABLE_CONSTRAINTS                     | TABLE_SCHEMA                      | 5        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| TABLE_CONSTRAINTS                     | TABLE_NAME                        | 6        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| TABLE_CONSTRAINTS                     | CONSTRAINT_TYPE                   | 7        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| TABLE_CONSTRAINTS                     | IS_DEFERRABLE                     | 8        | sql2003, pg, mssql, hsqldb, h2          | 5   |
| TABLE_CONSTRAINTS                     | INITIALLY_DEFERRED                | 9        | sql2003, pg, mssql, hsqldb, h2          | 5   |
| TABLE_PRIVILEGES                      | GRANTOR                           | 1        | sql2003, pg, mssql, hsqldb, h2          | 5   |
| TABLE_PRIVILEGES                      | GRANTEE                           | 2        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| TABLE_PRIVILEGES                      | TABLE_CATALOG                     | 3        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| TABLE_PRIVILEGES                      | TABLE_SCHEMA                      | 4        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| TABLE_PRIVILEGES                      | TABLE_NAME                        | 5        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| TABLE_PRIVILEGES                      | PRIVILEGE_TYPE                    | 6        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| TABLE_PRIVILEGES                      | IS_GRANTABLE                      | 7        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| TABLE_PRIVILEGES                      | WITH_HIERARCHY                    | 8        | sql2003, pg, hsqldb                     | 3   |
| TABLES                                | TABLE_CATALOG                     | 1        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| TABLES                                | TABLE_SCHEMA                      | 2        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| TABLES                                | TABLE_NAME                        | 3        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| TABLES                                | TABLE_TYPE                        | 4        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| TABLES                                | SELF_REFERENCING_COLUMN_NAME      | 5        | sql2003, pg, hsqldb                     | 3   |
| TABLES                                | REFERENCE_GENERATION              | 6        | sql2003, pg, hsqldb                     | 3   |
| TABLES                                | USER_DEFINED_TYPE_CATALOG         | 7        | sql2003, pg, hsqldb                     | 3   |
| TABLES                                | USER_DEFINED_TYPE_SCHEMA          | 8        | sql2003, pg, hsqldb                     | 3   |
| TABLES                                | USER_DEFINED_TYPE_NAME            | 9        | sql2003, pg, hsqldb                     | 3   |
| TABLES                                | IS_INSERTABLE_INTO                | 10       | sql2003, pg, hsqldb                     | 3   |
| TABLES                                | IS_TYPED                          | 11       | sql2003, pg, hsqldb                     | 3   |
| TABLES                                | COMMIT_ACTION                     | 12       | sql2003, pg, hsqldb                     | 3   |
| TRANSFORMS                            | UDT_CATALOG                       | 1        | sql2003, pg                             | 2   |
| TRANSFORMS                            | UDT_SCHEMA                        | 2        | sql2003, pg                             | 2   |
| TRANSFORMS                            | UDT_NAME                          | 3        | sql2003, pg                             | 2   |
| TRANSFORMS                            | SPECIFIC_CATALOG                  | 4        | sql2003, pg                             | 2   |
| TRANSFORMS                            | SPECIFIC_SCHEMA                   | 5        | sql2003, pg                             | 2   |
| TRANSFORMS                            | SPECIFIC_NAME                     | 6        | sql2003, pg                             | 2   |
| TRANSFORMS                            | GROUP_NAME                        | 7        | sql2003, pg                             | 2   |
| TRANSFORMS                            | TRANSFORM_TYPE                    | 8        | sql2003, pg                             | 2   |
| TRANSLATIONS                          | TRANSLATION_CATALOG               | 1        | sql2003, hsqldb                         | 2   |
| TRANSLATIONS                          | TRANSLATION_SCHEMA                | 2        | sql2003, hsqldb                         | 2   |
| TRANSLATIONS                          | TRANSLATION_NAME                  | 3        | sql2003, hsqldb                         | 2   |
| TRANSLATIONS                          | SOURCE_CHARACTER_SET_CATALOG      | 4        | sql2003, hsqldb                         | 2   |
| TRANSLATIONS                          | SOURCE_CHARACTER_SET_SCHEMA       | 5        | sql2003, hsqldb                         | 2   |
| TRANSLATIONS                          | SOURCE_CHARACTER_SET_NAME         | 6        | sql2003, hsqldb                         | 2   |
| TRANSLATIONS                          | TARGET_CHARACTER_SET_CATALOG      | 7        | sql2003, hsqldb                         | 2   |
| TRANSLATIONS                          | TARGET_CHARACTER_SET_SCHEMA       | 8        | sql2003, hsqldb                         | 2   |
| TRANSLATIONS                          | TARGET_CHARACTER_SET_NAME         | 9        | sql2003, hsqldb                         | 2   |
| TRANSLATIONS                          | TRANSLATION_SOURCE_CATALOG        | 10       | sql2003, hsqldb                         | 2   |
| TRANSLATIONS                          | TRANSLATION_SOURCE_SCHEMA         | 11       | sql2003, hsqldb                         | 2   |
| TRANSLATIONS                          | TRANSLATION_SOURCE_NAME           | 12       | sql2003, hsqldb                         | 2   |
| TRIGGER_COLUMN_USAGE                  | TRIGGER_CATALOG                   | 1        | sql2003, hsqldb                         | 2   |
| TRIGGER_COLUMN_USAGE                  | TRIGGER_SCHEMA                    | 2        | sql2003, hsqldb                         | 2   |
| TRIGGER_COLUMN_USAGE                  | TRIGGER_NAME                      | 3        | sql2003, hsqldb                         | 2   |
| TRIGGER_COLUMN_USAGE                  | TABLE_CATALOG                     | 4        | sql2003, hsqldb                         | 2   |
| TRIGGER_COLUMN_USAGE                  | TABLE_SCHEMA                      | 5        | sql2003, hsqldb                         | 2   |
| TRIGGER_COLUMN_USAGE                  | TABLE_NAME                        | 6        | sql2003, hsqldb                         | 2   |
| TRIGGER_COLUMN_USAGE                  | COLUMN_NAME                       | 7        | sql2003, hsqldb                         | 2   |
| TRIGGERED_UPDATE_COLUMNS              | TRIGGER_CATALOG                   | 1        | sql2003, pg, hsqldb                     | 3   |
| TRIGGERED_UPDATE_COLUMNS              | TRIGGER_SCHEMA                    | 2        | sql2003, pg, hsqldb                     | 3   |
| TRIGGERED_UPDATE_COLUMNS              | TRIGGER_NAME                      | 3        | sql2003, pg, hsqldb                     | 3   |
| TRIGGERED_UPDATE_COLUMNS              | EVENT_OBJECT_CATALOG              | 4        | sql2003, pg, hsqldb                     | 3   |
| TRIGGERED_UPDATE_COLUMNS              | EVENT_OBJECT_SCHEMA               | 5        | sql2003, pg, hsqldb                     | 3   |
| TRIGGERED_UPDATE_COLUMNS              | EVENT_OBJECT_TABLE                | 6        | sql2003, pg, hsqldb                     | 3   |
| TRIGGERED_UPDATE_COLUMNS              | EVENT_OBJECT_COLUMN               | 7        | sql2003, pg, hsqldb                     | 3   |
| TRIGGER_ROUTINE_USAGE                 | TRIGGER_CATALOG                   | 1        | sql2003, hsqldb                         | 2   |
| TRIGGER_ROUTINE_USAGE                 | TRIGGER_SCHEMA                    | 2        | sql2003, hsqldb                         | 2   |
| TRIGGER_ROUTINE_USAGE                 | TRIGGER_NAME                      | 3        | sql2003, hsqldb                         | 2   |
| TRIGGER_ROUTINE_USAGE                 | SPECIFIC_CATALOG                  | 4        | sql2003, hsqldb                         | 2   |
| TRIGGER_ROUTINE_USAGE                 | SPECIFIC_SCHEMA                   | 5        | sql2003, hsqldb                         | 2   |
| TRIGGER_ROUTINE_USAGE                 | SPECIFIC_NAME                     | 6        | sql2003, hsqldb                         | 2   |
| TRIGGERS                              | TRIGGER_CATALOG                   | 1        | sql2003, pg, mariadb, hsqldb, h2        | 5   |
| TRIGGERS                              | TRIGGER_SCHEMA                    | 2        | sql2003, pg, mariadb, hsqldb, h2        | 5   |
| TRIGGERS                              | TRIGGER_NAME                      | 3        | sql2003, pg, mariadb, hsqldb, h2        | 5   |
| TRIGGERS                              | EVENT_MANIPULATION                | 4        | sql2003, pg, mariadb, hsqldb            | 4   |
| TRIGGERS                              | EVENT_OBJECT_CATALOG              | 5        | sql2003, pg, mariadb, hsqldb            | 4   |
| TRIGGERS                              | EVENT_OBJECT_SCHEMA               | 6        | sql2003, pg, mariadb, hsqldb            | 4   |
| TRIGGERS                              | EVENT_OBJECT_TABLE                | 7        | sql2003, pg, mariadb, hsqldb            | 4   |
| TRIGGERS                              | ACTION_ORDER                      | 8        | sql2003, pg, mariadb, hsqldb            | 4   |
| TRIGGERS                              | ACTION_CONDITION                  | 9        | sql2003, pg, mariadb, hsqldb            | 4   |
| TRIGGERS                              | ACTION_STATEMENT                  | 10       | sql2003, pg, mariadb, hsqldb            | 4   |
| TRIGGERS                              | ACTION_ORIENTATION                | 11       | sql2003, pg, mariadb, hsqldb            | 4   |
| TRIGGERS                              | ACTION_TIMING                     | 12       | sql2003, pg, mariadb, hsqldb            | 4   |
| TRIGGERS                              | ACTION_REFERENCE_OLD_TABLE        | 13       | sql2003, pg, mariadb, hsqldb            | 4   |
| TRIGGERS                              | ACTION_REFERENCE_NEW_TABLE        | 14       | sql2003, pg, mariadb, hsqldb            | 4   |
| TRIGGERS                              | ACTION_REFERENCE_OLD_ROW          | 15       | sql2003, pg, mariadb, hsqldb            | 4   |
| TRIGGERS                              | ACTION_REFERENCE_NEW_ROW          | 16       | sql2003, pg, mariadb, hsqldb            | 4   |
| TRIGGERS                              | CREATED                           | 17       | sql2003, pg, mariadb, hsqldb            | 4   |
| TRIGGER_SEQUENCE_USAGE                | TRIGGER_CATALOG                   | 1        | sql2003, hsqldb                         | 2   |
| TRIGGER_SEQUENCE_USAGE                | TRIGGER_SCHEMA                    | 2        | sql2003, hsqldb                         | 2   |
| TRIGGER_SEQUENCE_USAGE                | TRIGGER_NAME                      | 3        | sql2003, hsqldb                         | 2   |
| TRIGGER_SEQUENCE_USAGE                | SEQUENCE_CATALOG                  | 4        | sql2003, hsqldb                         | 2   |
| TRIGGER_SEQUENCE_USAGE                | SEQUENCE_SCHEMA                   | 5        | sql2003, hsqldb                         | 2   |
| TRIGGER_SEQUENCE_USAGE                | SEQUENCE_NAME                     | 6        | sql2003, hsqldb                         | 2   |
| TRIGGER_TABLE_USAGE                   | TRIGGER_CATALOG                   | 1        | sql2003, hsqldb                         | 2   |
| TRIGGER_TABLE_USAGE                   | TRIGGER_SCHEMA                    | 2        | sql2003, hsqldb                         | 2   |
| TRIGGER_TABLE_USAGE                   | TRIGGER_NAME                      | 3        | sql2003, hsqldb                         | 2   |
| TRIGGER_TABLE_USAGE                   | TABLE_CATALOG                     | 4        | sql2003, hsqldb                         | 2   |
| TRIGGER_TABLE_USAGE                   | TABLE_SCHEMA                      | 5        | sql2003, hsqldb                         | 2   |
| TRIGGER_TABLE_USAGE                   | TABLE_NAME                        | 6        | sql2003, hsqldb                         | 2   |
| UDT_PRIVILEGES                        | GRANTOR                           | 1        | sql2003, pg, hsqldb                     | 3   |
| UDT_PRIVILEGES                        | GRANTEE                           | 2        | sql2003, pg, hsqldb                     | 3   |
| UDT_PRIVILEGES                        | UDT_CATALOG                       | 3        | sql2003, pg, hsqldb                     | 3   |
| UDT_PRIVILEGES                        | UDT_SCHEMA                        | 4        | sql2003, pg, hsqldb                     | 3   |
| UDT_PRIVILEGES                        | UDT_NAME                          | 5        | sql2003, pg, hsqldb                     | 3   |
| UDT_PRIVILEGES                        | PRIVILEGE_TYPE                    | 6        | sql2003, pg, hsqldb                     | 3   |
| UDT_PRIVILEGES                        | IS_GRANTABLE                      | 7        | sql2003, pg, hsqldb                     | 3   |
| USAGE_PRIVILEGES                      | GRANTOR                           | 1        | sql2003, pg, hsqldb                     | 3   |
| USAGE_PRIVILEGES                      | GRANTEE                           | 2        | sql2003, pg, hsqldb                     | 3   |
| USAGE_PRIVILEGES                      | OBJECT_CATALOG                    | 3        | sql2003, pg, hsqldb                     | 3   |
| USAGE_PRIVILEGES                      | OBJECT_SCHEMA                     | 4        | sql2003, pg, hsqldb                     | 3   |
| USAGE_PRIVILEGES                      | OBJECT_NAME                       | 5        | sql2003, pg, hsqldb                     | 3   |
| USAGE_PRIVILEGES                      | OBJECT_TYPE                       | 6        | sql2003, pg, hsqldb                     | 3   |
| USAGE_PRIVILEGES                      | PRIVILEGE_TYPE                    | 7        | sql2003, pg, hsqldb                     | 3   |
| USAGE_PRIVILEGES                      | IS_GRANTABLE                      | 8        | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | USER_DEFINED_TYPE_CATALOG         | 1        | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | USER_DEFINED_TYPE_SCHEMA          | 2        | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | USER_DEFINED_TYPE_NAME            | 3        | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | USER_DEFINED_TYPE_CATEGORY        | 4        | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | IS_INSTANTIABLE                   | 5        | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | IS_FINAL                          | 6        | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | ORDERING_FORM                     | 7        | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | ORDERING_CATEGORY                 | 8        | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | ORDERING_ROUTINE_CATALOG          | 9        | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | ORDERING_ROUTINE_SCHEMA           | 10       | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | ORDERING_ROUTINE_NAME             | 11       | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | REFERENCE_TYPE                    | 12       | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | DATA_TYPE                         | 13       | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | CHARACTER_MAXIMUM_LENGTH          | 14       | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | CHARACTER_OCTET_LENGTH            | 15       | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | CHARACTER_SET_CATALOG             | 16       | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | CHARACTER_SET_SCHEMA              | 17       | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | CHARACTER_SET_NAME                | 18       | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | COLLATION_CATALOG                 | 19       | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | COLLATION_SCHEMA                  | 20       | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | COLLATION_NAME                    | 21       | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | NUMERIC_PRECISION                 | 22       | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | NUMERIC_PRECISION_RADIX           | 23       | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | NUMERIC_SCALE                     | 24       | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | DATETIME_PRECISION                | 25       | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | INTERVAL_TYPE                     | 26       | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | INTERVAL_PRECISION                | 27       | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | SOURCE_DTD_IDENTIFIER             | 28       | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | REF_DTD_IDENTIFIER                | 29       | sql2003, pg, hsqldb                     | 3   |
| VIEW_COLUMN_USAGE                     | VIEW_CATALOG                      | 1        | sql2003, pg, mssql, hsqldb              | 4   |
| VIEW_COLUMN_USAGE                     | VIEW_SCHEMA                       | 2        | sql2003, pg, mssql, hsqldb              | 4   |
| VIEW_COLUMN_USAGE                     | VIEW_NAME                         | 3        | sql2003, pg, mssql, hsqldb              | 4   |
| VIEW_COLUMN_USAGE                     | TABLE_CATALOG                     | 4        | sql2003, pg, mssql, hsqldb              | 4   |
| VIEW_COLUMN_USAGE                     | TABLE_SCHEMA                      | 5        | sql2003, pg, mssql, hsqldb              | 4   |
| VIEW_COLUMN_USAGE                     | TABLE_NAME                        | 6        | sql2003, pg, mssql, hsqldb              | 4   |
| VIEW_COLUMN_USAGE                     | COLUMN_NAME                       | 7        | sql2003, pg, mssql, hsqldb              | 4   |
| VIEW_ROUTINE_USAGE                    | TABLE_CATALOG                     | 1        | sql2003, pg                             | 2   |
| VIEW_ROUTINE_USAGE                    | TABLE_SCHEMA                      | 2        | sql2003, pg                             | 2   |
| VIEW_ROUTINE_USAGE                    | TABLE_NAME                        | 3        | sql2003, pg                             | 2   |
| VIEW_ROUTINE_USAGE                    | SPECIFIC_CATALOG                  | 4        | sql2003, pg, hsqldb                     | 3   |
| VIEW_ROUTINE_USAGE                    | SPECIFIC_SCHEMA                   | 5        | sql2003, pg, hsqldb                     | 3   |
| VIEW_ROUTINE_USAGE                    | SPECIFIC_NAME                     | 6        | sql2003, pg, hsqldb                     | 3   |
| VIEWS                                 | TABLE_CATALOG                     | 1        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| VIEWS                                 | TABLE_SCHEMA                      | 2        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| VIEWS                                 | TABLE_NAME                        | 3        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| VIEWS                                 | VIEW_DEFINITION                   | 4        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| VIEWS                                 | CHECK_OPTION                      | 5        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| VIEWS                                 | IS_UPDATABLE                      | 6        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| VIEWS                                 | INSERTABLE_INTO                   | 7        | sql2003, hsqldb                         | 2   |
| VIEWS                                 | IS_TRIGGER_DELETABLE              |          | pg, hsqldb                              | 2   |
| VIEWS                                 | IS_TRIGGER_INSERTABLE_INTO        |          | pg, hsqldb                              | 2   |
| VIEWS                                 | IS_TRIGGER_UPDATABLE              |          | pg, hsqldb                              | 2   |
| VIEW_TABLE_USAGE                      | VIEW_CATALOG                      | 1        | sql2003, pg, mssql, hsqldb              | 4   |
| VIEW_TABLE_USAGE                      | VIEW_SCHEMA                       | 2        | sql2003, pg, mssql, hsqldb              | 4   |
| VIEW_TABLE_USAGE                      | VIEW_NAME                         | 3        | sql2003, pg, mssql, hsqldb              | 4   |
| VIEW_TABLE_USAGE                      | TABLE_CATALOG                     | 4        | sql2003, pg, mssql, hsqldb              | 4   |
| VIEW_TABLE_USAGE                      | TABLE_SCHEMA                      | 5        | sql2003, pg, mssql, hsqldb              | 4   |
| VIEW_TABLE_USAGE                      | TABLE_NAME                        | 6        | sql2003, pg, mssql, hsqldb              | 4   |
