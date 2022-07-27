package Connection

import (
	"github.com/gocql/gocql"
	"time"
)

type Config struct {
	Hosts        []string
	Keyspace     string
	Strategy     string
	RF           int
	ProtoVersion int
	Timeout      time.Duration
	CL           gocql.Consistency
	NumConns     int
}

func NewConfig(hosts []string, keyspace string, strategy string, RF int, protoVersion int, timeout time.Duration, CL gocql.Consistency, numConns int) *Config {
	return &Config{Hosts: hosts, Keyspace: keyspace, Strategy: strategy, RF: RF, ProtoVersion: protoVersion, Timeout: timeout, CL: CL, NumConns: numConns}
}
