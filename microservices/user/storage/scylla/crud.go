package scylla

import (
	"encoding/json"
	"fmt"
	"github.com/gocql/gocql"
	"log"
	"time"
	"user/models"
	"user/storage/scylla/Connection"
)

type TableNames struct {
	User string
}

type Scylla struct {
	*gocql.Session
	Keyspace string
	UserTableName *TableNames
}

func NewScylla(conf *Connection.Config, names *TableNames) (*Scylla, error) {
	client, err := Connection.New(conf)
	if err != nil {
		log.Fatal(err)
	}
	return &Scylla{client, conf.Keyspace, names}, err
}

func (s Scylla) Create(user *models.User) error {
	stmt := fmt.Sprintf(`INSERT INTO %s.%s(id, name, address, created_at, updated_at) VALUES (?,?,?,?,?)`, s.Keyspace, s.UserTableName.User)
	err := s.Query(stmt, user.ID, user.Name, user.Address, user.CreatedAt, user.UpdatedAt).Exec()
	if err != nil {
		return err
	}
	return err
}

func (s Scylla) GetUser(id uint32) ([]byte, error) {
	var user models.User
	stmt := fmt.Sprintf(`SELECT id, name, address, created_at, updated_at FROM %s.%s WHERE id = ?`, s.Keyspace, s.UserTableName.User)
	err := s.Query(stmt, id).Scan(&user.ID, &user.Name, &user.Address, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(user)
}

func (s Scylla) Update(userID uint32, data []byte) error {
	var user models.User
	var userTmpForCreatedAt models.User
	err := json.Unmarshal(data, &user)
	if err != nil {
		log.Println(err)
	}
	user.UpdatedAt = time.Now().Unix()
	updateData := prepareUpdateStruct(&user)
	fmt.Println(updateData)
	if err != nil {
		return err
	}
	userCreatedAt,err  := s.GetUser(userID)
	if err != nil {
		return err
	}
	err = json.Unmarshal(userCreatedAt, &userTmpForCreatedAt)
	if err != nil {
		log.Println(err)
	}

	stmt := fmt.Sprintf(`UPDATE %s.%s SET %s WHERE id = ? and created_at = ?`, s.Keyspace, s.UserTableName.User, updateData)
	err = s.Query(stmt, int(userID), userTmpForCreatedAt.CreatedAt).Exec()
	if err != nil {
		return err
	}
	return nil
}

func (s Scylla) Delete(id uint32) error {
	stmt := fmt.Sprintf(`DELETE FROM %s.%s WHERE id =?`, s.Keyspace, s.UserTableName.User)
	err := s.Query(stmt, id).Exec()
	if err != nil {
		return err
	}
	return nil
}
