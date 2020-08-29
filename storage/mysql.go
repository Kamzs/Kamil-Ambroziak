package storage

import (
	"Kamil-Ambroziak/logger"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

const (
	mysqlUsername = "root"
	mysqlPassword = "password"
	mysqlHostPort = "localhost:3306"
	mysqlDB       = "sys"
)

type MySQL struct {
	client *sql.DB
	tables []string
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
	mysql.SetLogger(logger.GetLogger())

	log.Println("database successfully configured")
	return &MySQL{
		client: Client,
		tables: []string{"fetchers","history"},
	}, err
}
