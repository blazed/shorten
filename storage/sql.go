package storage

import (
	"fmt"

	"os"

	"github.com/jmoiron/sqlx"
	// Third party drivers
	_ "github.com/lib/pq"
)

var schema = `
CREATE TABLE IF NOT EXISTS shorten_urls (
	id BIGSERIAL NOT NULL PRIMARY KEY,
	slug VARCHAR NOT NULL UNIQUE,
	url VARCHAR NOT NULL,
	created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
)`

type conn struct {
	db *sqlx.DB
}

// Open makes sure we can connect the database
func Open() (Storage, error) {
	conn, err := open()
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func open() (*conn, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME")))
	if err != nil {
		return nil, err
	}
	c := &conn{db}
	c.db.MustExec(schema)
	return c, nil
}

func (c *conn) Close() error {
	return c.db.Close()
}

func (c *conn) GetURL(slug string) (url URL, err error) {
	err = c.db.Get(&url, "SELECT * FROM shorten_urls WHERE slug = $1", slug)
	if err != nil {
		return url, err
	}
	return url, nil

}

func (c *conn) CreateShortURL(short URL) error {
	_, err := c.db.NamedExec(`INSERT INTO shorten_urls (slug, url, created_at) values (:slug, :url, :created_at);`, short)
	if err != nil {
		return err
	}
	return nil
}
