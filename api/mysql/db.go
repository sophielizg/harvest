package mysql

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sophielizg/harvest/api/harvest"
)

var db *sql.DB

type MysqlConfig struct {
	user     string
	password string
	protocol string
	host     string
	port     int
	dbname   string
}

func (c *MysqlConfig) DSNString() string {
	return fmt.Sprintf("%s:%s@%s(%s:%d)/%s",
		c.user, c.password, c.protocol, c.host, c.port, c.dbname)
}

func Open(configService *harvest.ConfigService) error {
	mysqlConfigString, err := (*configService).Value("mysql")
	if err != nil {
		return err
	}

	mysqlConfig := MysqlConfig{}
	err = json.Unmarshal([]byte(mysqlConfigString), &mysqlConfig)
	if err != nil {
		return err
	}

	db, err := sql.Open("mysql", mysqlConfig.DSNString())
	if err != nil {
		return err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db.Ping()
}

func Close() {
	db.Close()
}
