# go-db-meta

The goal with this project is to provide a tool for extracting metadata
from various RDBMS engines and presenting that data in a consistent
fashion across engines.

It should be noted that, while the SQL standard information schema is
used to inform the design of this project, this is not intended to
exactly reflect the structure of the standard information schema. It
should also be noted that no two database systems appear to implement
the information_schema the same. See
[information_schema](information_schema/readme.md) for a comparison of
different information_schema implementations.

# Intended uses

 * Generate data dictionaries
 * Basic DDL extraction/generation
 * Database/schema comparisons
 * ???

# Supported databases

 * PostgreSQL
 * SQLite
 * MariaDB (work in progress)
 * MS-SQL (work in progress)
 * MySQL (work in progress)
 * Oracle

# Available queries

 * CheckConstraints
 * Columns
 * CurrentCatalog
 * Dependencies
 * Domains
 * Indexes
 * PrimaryKeys
 * ReferentialConstraints
 * Schemata
 * Tables
 * Types
 * UniqueConstraints

# Examples

See the ```_example``` directory.
