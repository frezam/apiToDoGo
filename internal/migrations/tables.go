package migrations

import _ "embed"

//go:embed sql/DDL_tables.sql
var TablesScript string
