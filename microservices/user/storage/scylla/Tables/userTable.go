package Tables

import "fmt"

func (tables *Tables) MakeUserTables(keyspace, tableName string) error {
	stmt := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %v.%v (id int, name text, address text, created_at int, updated_at int, PRIMARY KEY (id, created_at)) WITH CLUSTERING ORDER BY (created_at DESC)", keyspace, tableName)
	return tables.Session.Query(stmt).Exec()
}
func (tables *Tables) DropUserTables(keyspace, tableName string) error {
	stmt := fmt.Sprintf("DROP TABLE IF EXISTS %v.%v", keyspace, tableName)
	return tables.Session.Query(stmt).Exec()
}
