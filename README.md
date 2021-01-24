# go-db-meta

The goal with this project is to provide a tool for extracting metadata
from various RDBMS engines and presenting that data in a consistent
fashion across engines. The initial/primary focus is to extract
information about the tables/views in reasonably sufficient (though
perhaps not exact) detail to permit generating the DDL needed for
re-creating the tables/views and the relationships between them.

It should be noted that, while the SQL standard information schema is
used to inform the design of this project, this is not intended to
exactly reflect the structure of the standard information schema. See
[information_schema](information_schema/readme.md) for a comparison of
different information_schema implementations.

# Roadmap

 1. Describe the structure of tables/views.
 1. Describe any user defined domains/types used in tables.
 1. Describe the relationsips between tables/views.
 1. Describe the constraints on tables.
 1. ???
