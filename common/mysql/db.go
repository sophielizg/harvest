package mysql

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	harvest "github.com/sophielizg/harvest/common"
)

type MysqlConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Protocol string `json:"protocol"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Dbname   string `json:"dbname"`
}

func (c *MysqlConfig) DSNString() string {
	return fmt.Sprintf("%s:%s@%s(%s:%d)/%s?parseTime=true",
		c.User, c.Password, c.Protocol, c.Host, c.Port, c.Dbname)
}

func OpenDb(configService harvest.ConfigService) (*sql.DB, error) {
	mysqlConfigBytes, err := configService.Value("mysql")
	if err != nil {
		return nil, err
	}

	mysqlConfig := MysqlConfig{}
	err = json.Unmarshal(mysqlConfigBytes, &mysqlConfig)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("mysql", mysqlConfig.DSNString())
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, db.Ping()
}

func CloseDb(db *sql.DB) {
	db.Close()
}
