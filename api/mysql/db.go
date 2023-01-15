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
	User     string `json:"user"`
	Password string `json:"password"`
	Protocol string `json:"protocol"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Dbname   string `json:"dbname"`
}

func (c *MysqlConfig) DSNString() string {
	return fmt.Sprintf("%s:%s@%s(%s:%d)/%s",
		c.User, c.Password, c.Protocol, c.Host, c.Port, c.Dbname)
}

func Open(app *harvest.App) error {
	mysqlConfigString, err := app.ConfigService.Value("mysql")
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
