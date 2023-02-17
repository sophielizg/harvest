package mysql

import (
	"database/sql"
)

type CookieService struct {
	Db *sql.DB
}

func (c *CookieService) GetCookies(runId int, host string) (string, error) {
	rows, err := c.Db.Query("CALL getCookiesForHost(?, ?);", runId, host)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	for rows.Next() {
		var value string
		err := rows.Scan(&value)
		if err != nil {
			return "", err
		}
		return value, nil
	}

	return "", rows.Err()
}

func (c *CookieService) SetCookies(runId int, host string, value string) error {
	_, err := c.Db.Exec("CALL addOrUpdateCookies(?, ?, ?);", runId, host, value)
	return err
}
