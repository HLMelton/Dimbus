package database

import (
	"database/sql"
)

type Link struct {
	Code string
	URL  string
}

func (db *DB) CreateLink(code, url string) error {
	_, err := db.Exec(`INSERT INTO short_links (code, url) VALUES ($1, $2)`, code, url)
	return err
}

func (db *DB) GetLink(code string) (string, error) {
	var url string
	err := db.QueryRow(`SELECT url FROM short_links WHERE code = $1`, code).Scan(&url)
	if err == sql.ErrNoRows {
		return "", nil
	}
	return url, err
}

func (db *DB) CodeExists(code string) (bool, error) {
	var exists bool
	err := db.QueryRow(`SELECT EXISTS (SELECT 1 FROM short_links WHERE code = $1)`, code).Scan(&exists)
	return exists, err
} 