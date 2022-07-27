package Tables

import "github.com/gocql/gocql"

type Tables struct {
	Session *gocql.Session
}

func NewTables(session *gocql.Session) *Tables {
	return &Tables{Session: session}
}
