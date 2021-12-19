package sqlite

import (
	m "github.com/gsiems/go-db-meta/model"
)

// CheckConstraints returns an empty set it is unclear how to extract
// check constraints from SQLite other than parsing the sqlite_master.sql
// column
func CheckConstraints(db *m.DB, tableSchema, tableName string) ([]m.CheckConstraint, error) {

	q := ``
	return db.CheckConstraints(q, tableSchema, tableName)

	var d []m.CheckConstraint

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
				if strings.Contains(tv, " CHECK ") {

					u := m.CheckConstraint{
						TableCatalog: catName,
						TableSchema:  tableSchema,
						TableName:    name,
					}

					pts := SplitN(tv, "(", 2)

					u.ConstraintName = strings.Trim(strings.TrimPrefix(pts[0], "CONSTRAINT"), " ")
					u.CheckClause = strings.TrimSuffix(pts[1], ")")

					d = append(d, u)

				}
			}
		}
	}
	return d, err

}

// Column constraints: column_name [type_name [(p[,s]]] [constraint [constraint]]
//  [CONSTRAINT constraint_name]
//     ( PRIMARY KEY [ASC|DESC] [ON CONFLICT (ROLLBACK|ABORT|FAIL|IGNORE|REPLACE)][AUTOINCREMENT]
//     | NOT NULL [ON CONFLICT (ROLLBACK|ABORT|FAIL|IGNORE|REPLACE)]
//     | UNIQUE [ON CONFLICT (ROLLBACK|ABORT|FAIL|IGNORE|REPLACE)]
//     | CHECK (expr)
//     | DEFAULT (NUMBER|STRING|(expr))
//     | COLLATE collate_name
//     | FOREIGN KEY REFERENCES table_name ...
//     )

/*
sub get_check_constraints {
    my ( $self, $filters ) = @_;
    $self->{logger}->log_info("Extracting check_constraint information...");
    my %return;
    my $schema = $self->{schema};
    my %tables = $self->get_tables($filters);
    my $name_idx = 1;    # "un-named" constraint name index-- SQLite doesn't require constraint names

    foreach my $table_name ( keys %tables ) {
        next unless ( 'TABLE' eq $tables{$table_name}{table_type} );
        my %pt = $self->_parse_table_ddl( $schema, $table_name );

        foreach my $constraint ( @{ $pt{table_constraints} } ) {
            next unless ( $constraint =~ m/\b(CHECK)\b/i );
            my @ary = grep { $_ } split /\b(CONSTRAINT|CHECK)\b/i, $constraint;
            my ( $constraint_name, $search_condition );

            while (@ary) {
                my $token = shift @ary;
                next unless ($token);
                if ( uc $token eq 'CONSTRAINT' ) {
                    $constraint_name = shift @ary;
                    $constraint_name =~ s/^\s+//;
                    $constraint_name =~ s/\s+$//;
                }
                elsif ( uc $token eq 'CHECK' ) {
                    $search_condition = shift @ary;
                    $search_condition =~ s/^[\s]+//;
                    $search_condition =~ s/[\s]+$//;
                }
            }
            $constraint_name ||= 'unnamed_check_' . $name_idx++;

            $return{$table_name}{$constraint_name}{table_schema}     = $schema;
            $return{$table_name}{$constraint_name}{table_name}       = $table_name;
            $return{$table_name}{$constraint_name}{constraint_name}  = $constraint_name;
            $return{$table_name}{$constraint_name}{search_condition} = $search_condition;
            $return{$table_name}{$constraint_name}{status}           = 'Enabled';
            $return{$table_name}{$constraint_name}{comments}         = '';
        }

        # Column:
        foreach my $column_name ( @{ $pt{column_names} } ) {
            foreach my $constraint ( @{ $pt{columns}{$column_name}{constraints} } ) {
                next unless ( $constraint =~ m/\b(CHECK)\b/i );
                my @ary = grep { $_ } split /\b(CONSTRAINT|CHECK)\b/i, $constraint;
                my ( $constraint_name, $search_condition );
                while (@ary) {
                    my $token = shift @ary;
                    next unless ($token);

                    if ( uc $token eq 'CONSTRAINT' ) {
                        $constraint_name = shift @ary;
                        $constraint_name =~ s/^\s+//;
                        $constraint_name =~ s/\s+$//;
                    }
                    elsif ( uc $token eq 'CHECK' ) {
                        $search_condition = shift @ary;
                        $search_condition =~ s/^[\s]+//;
                        $search_condition =~ s/[\s]+$//;
                    }
                }
                $constraint_name ||= 'unnamed_check_' . $name_idx++;

                $return{$table_name}{$constraint_name}{table_schema}     = $schema;
                $return{$table_name}{$constraint_name}{table_name}       = $table_name;
                $return{$table_name}{$constraint_name}{constraint_name}  = $constraint_name;
                $return{$table_name}{$constraint_name}{search_condition} = $search_condition;
                $return{$table_name}{$constraint_name}{status}           = 'Enabled';
                $return{$table_name}{$constraint_name}{comments}         = '';
            }
        }
    }
    return wantarray ? %return : \%return;
}
*/
