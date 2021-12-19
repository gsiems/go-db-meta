package sqlite

import (
	m "github.com/gsiems/go-db-meta/model"
)

// UniqueConstraints defines the query for obtaining a list of unique
// constraints for the (tableSchema, tableName) parameters and returns the
// results of executing the query
func UniqueConstraints(db *m.DB, tableSchema, tableName string) ([]m.UniqueConstraint, error) {

	//q := ``
	//return db.UniqueConstraints(q, tableSchema, tableName)

	var d []m.UniqueConstraint

	catName, err := catalogName(db)
	if err != nil {
		return d, err
	}

	q := `
SELECT m.name AS table_name,
        m.sql
    FROM sqlite_master m
    CROSS JOIN (
        SELECT coalesce ( $2, '' ) AS table_name
        ) AS args
    WHERE m.type = 'table'
        AND m.tbl_name NOT LIKE '%s'
        AND ( args.table_name = '' OR args.table_name = m.name )
`

	rows, err := db.Query(q, "sqlite_%", tableName)
	if err != nil {
		return d, err
	}
	defer func() {
		if cerr := rows.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	for rows.Next() {

		var name string
		var sql string

		err = rows.Scan(&name,
			&sql,
		)
		if err != nil {
			return d, err
		}

		for _, v := range strings.Split(sql, "\n") {

			tv := strings.Trim(strings.RTrim(v, "\n\r ,"))
			if strings.HasPrefix(tv, "CONSTRAINT ") {
				if strings.Contains(tv, " UNIQUE ") {

					u := m.UniqueConstraint{
						TableCatalog: catName,
						TableSchema:  tableSchema,
						TableName:    name,
					}

					pts := SplitN(tv, "(", 2)

					u.ConstraintName = strings.Trim(strings.TrimPrefix(pts[0], "CONSTRAINT"), " ")
					u.ConstraintColumns = strings.TrimSuffix(pts[1], ")")

					d = append(d, u)

				}
			}
		}
	}
	return d, err
}

/*
type UniqueConstraint struct {
	TableCatalog      sql.NullString `json:"tableCatalog"`
	TableSchema       sql.NullString `json:"tableSchema"`
	TableName         sql.NullString `json:"tableName"`
	ConstraintName    sql.NullString `json:"constraintName"`
	ConstraintColumns sql.NullString `json:"constraintColumns"`
	Status            sql.NullString `json:"status"`
	Comment           sql.NullString `json:"comment"`
}



type sqliteTabCons struct {
tableName string
constraints []string
}


SELECT '%s' AS table_catalog,
        coalesce ( $1, '' ) AS table_schema,
        name AS table_name,
        NULL AS table_owner,
        sql
    FROM sqlite_master
    WHERE type = 'table'
        AND tbl_name NOT LIKE '%s'



func (db *DB) QSingleString(q string) (sql.NullString, error) {

	var v sql.NullString
	rows, err := db.Query(q)
	if err != nil {
		return v, err
	}
	defer func() {
		if cerr := rows.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	if rows.Next() {
		err = rows.Scan(&v)
	}







	return v, err
}



for _, v := range strings.Split ( sql, "\n" ) {
tv := strings.Trim ( v )
    if  strings.HasPrefix(tv, "CONSTRAINT ") {







    }
}



sql column
CREATE TABLE newTable4_nk
    (
        id INTEGER,
        col2 INTEGER,
        col3 TEXT,
        CONSTRAINT newTable4_ck1 CHECK (col2 > 10),
        CONSTRAINT newTable4_pk PRIMARY KEY (id),
        CONSTRAINT newTable4_ix1 UNIQUE (col3)
    )





sub get_unique_constraints {
    my ( $self, $filters ) = @_;
    $self->{logger}->log_info("Retrieving unique constraint information...");
    my %return;
    my $schema = $self->{schema};
    my %tables = $self->get_tables($filters);
    my $name_idx = 1;    # "un-named" constraint name index-- SQLite doesn't require constraint names

    foreach my $table_name ( keys %tables ) {
        next unless ( 'TABLE' eq $tables{$table_name}{table_type} );
        my %pt = $self->_parse_table_ddl( $schema, $table_name );

        foreach my $constraint ( @{ $pt{table_constraints} } ) {
            next unless ( $constraint =~ m/\b(UNIQUE)\b]/i );
            my @ary = grep { $_ } split /\b(UNIQUE)\b/i, $constraint;

            my ( $constraint_name, $column_names );

            #    [CONSTRAINT [constraint_name]] UNIQUE(column[,column[,...]])
            while (@ary) {
                my $token = shift @ary;
                next unless ($token);
                if ( uc $token eq 'CONSTRAINT' ) {
                    $constraint_name = shift @ary;
                    $constraint_name =~ s/^\s+//;
                    $constraint_name =~ s/\s+$//;
                }
                elsif ( uc $token eq 'UNIQUE' ) {
                    $column_names = shift @ary;
                    $column_names =~ s/^[\s(]+//;
                    $column_names =~ s/[\s)]+$//;
                }
            }
            $constraint_name ||= 'un_named_unique_' . $name_idx++;

            $return{$table_name}{$constraint_name}{table_schema}    = $schema;
            $return{$table_name}{$constraint_name}{table_name}      = $table_name;
            $return{$table_name}{$constraint_name}{constraint_name} = $constraint_name;
            @{ $return{$table_name}{$constraint_name}{column_names} } = split /\s*,\s* /, $column_names;
            $return{$table_name}{$constraint_name}{status}   = 'Enabled';
            $return{$table_name}{$constraint_name}{comments} = '';
        }

        # Column:
        foreach my $column_name ( @{ $pt{column_names} } ) {
            foreach my $constraint ( @{ $pt{columns}{$column_name}{constraints} } ) {
                next unless ( $constraint =~ m/\bUNIQUE[\s\n(]/i );
                my @ary = grep { $_ } split /\b(UNIQUE)\b/i, $constraint;
                my ($constraint_name);

                while (@ary) {
                    my $token = shift @ary;
                    next unless ($token);
                    if ( uc $token eq 'CONSTRAINT' ) {
                        $constraint_name = shift @ary;
                        $constraint_name =~ s/^\s+//;
                        $constraint_name =~ s/\s+$//;
                    }
                }
                $constraint_name ||= 'un_named_unique_' . $name_idx++;

                $return{$table_name}{$constraint_name}{table_schema}    = $schema;
                $return{$table_name}{$constraint_name}{table_name}      = $table_name;
                $return{$table_name}{$constraint_name}{constraint_name} = $constraint_name;
                @{ $return{$table_name}{$constraint_name}{column_names} } = ($column_name);
                $return{$table_name}{$constraint_name}{status}   = 'Enabled';
                $return{$table_name}{$constraint_name}{comments} = '';
            }
        }
    }
    return wantarray ? %return : \%return;
}

*/
