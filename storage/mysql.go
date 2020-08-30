package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	mysqlUsername = "root"
	mysqlPassword = "password"
	mysqlHostPort = "localhost:3306"
	mysqlDB       = "db"
)

type MySQL struct {
	client *sql.DB
}

func NewMySQL() (*MySQL, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		mysqlUsername, mysqlPassword, mysqlHostPort, mysqlDB,
	)
	Client, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = Client.Ping(); err != nil {
		return nil, err
	}
	return &MySQL{
		client: Client,
	}, err
}
