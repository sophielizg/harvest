package mysql

import (
	"database/sql"
)

type CookieService struct {
	Db *sql.DB
}

func (c *CookieService) GetCookies(runId int, host string) (string, error) {
	rows, err := c.Db.Query("CALL addOrUpdateCookies(?, ?);", runId, host)
	if err != nil {
		return "", err
	}

	for rows.Next() {
		var value string
		err := rows.Scan(&value)
		if err != nil {
			return "", err
		}
		return value, nil
	}

	return "", nil
}

func (c *CookieService) SetCookies(runId int, host string, value string) error {
	stmt, err := c.Db.Prepare("CALL addOrUpdateCookies(?, ?, ?);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(runId, host, value)
	return err
}
