package Connection

import (
	"fmt"
	"github.com/gocql/gocql"
	"log"
)

type Connector struct {
	Cfg    *Config
	Client *gocql.Session
}

func NewConnector(cfg *Config) *Connector {
	return &Connector{Cfg: cfg}
}


func (connector *Connector) Connect() (*gocql.Session, error) {
	cluster := gocql.NewCluster(connector.Cfg.Hosts...)
	cluster.Hosts = connector.Cfg.Hosts
	cluster.Timeout = connector.Cfg.Timeout
	cluster.NumConns = connector.Cfg.NumConns
	cluster.ProtoVersion = connector.Cfg.ProtoVersion
	session, err := cluster.CreateSession()
	if err != nil {
		log.Println(err)
	}
	connector.Client = session
	return session, nil
}

func (connector *Connector) CreateKeyspace() error {

	queryString := fmt.Sprintf(`
		CREATE KEYSPACE IF NOT EXISTS %s
			WITH REPLICATION = {
				'class' : '%s',
				'replication_factor' : %d
			}`,
		connector.Cfg.Keyspace,
		connector.Cfg.Strategy,
		connector.Cfg.RF,
	)

	if err := connector.Client.Query(queryString).Exec(); err != nil {
		return err
	}
	return nil
}