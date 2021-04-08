package storage

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Kamzs/Kamil-Ambroziak/logger"

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
	for {
		if err = Client.Ping(); err != nil {
			logger.Info("mysql not reachable")
			time.Sleep(time.Second * 5)
		} else {
			return &MySQL{
				client: Client,
			}, err
		}
	}

}
