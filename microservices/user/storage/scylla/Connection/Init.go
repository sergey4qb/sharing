package Connection

import (
	"fmt"
	"github.com/gocql/gocql"
	"log"
)

//type Connection struct {
//
//}
func New(config *Config) (*gocql.Session, error) {
	conf := NewConfig(config.Hosts, config.Keyspace, config.Strategy, config.RF, config.ProtoVersion, config.Timeout, config.CL, config.NumConns )
	connector := NewConnector(conf)
	session, err := connector.Connect()
	if err != nil {
		log.Println(err)
	}
	createKeyspace := fmt.Sprintf(`
		CREATE KEYSPACE IF NOT EXISTS %s
			WITH REPLICATION = {
				'class' : '%s',
				'replication_factor' : %d
			}`,
		connector.Cfg.Keyspace,
		connector.Cfg.Strategy,
		connector.Cfg.RF,
	)

	if err := session.Query(createKeyspace).Exec(); err != nil {
		return &gocql.Session{}, err
	}
	if err != nil {
		log.Println(err)
	}
	return session, nil
}
