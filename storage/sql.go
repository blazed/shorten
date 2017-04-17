package storage

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var schema = `
CREATE TABLE IF NOT EXISTS shorten_urls (
	id BIGSERIAL NOT NULL PRIMARY KEY,
	slug VARCHAR NOT NULL,
	url VARCHAR NOT NULL
)`

type conn struct {
	db *sqlx.DB
}

func Open() (Storage, error) {
	conn, err := open()
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func open() (*conn, error) {
	db, err := sqlx.Connect("postgres", "user=blazed dbname=shorten sslmode=disable")
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

func (c *conn) GetUrl(slug string) (url URL, err error) {
	err = c.db.Get(&url, "SELECT * FROM shorten_urls WHERE slug = $1", slug)
	if err != nil {
		return url, err
	}
	return url, nil

}

func (c *conn) CreateShortUrl(short URL) error {
	_, err := c.db.NamedExec(`INSERT INTO shorten_urls (slug, url) values (:slug, :url);`, short)
	if err != nil {
		return err
	}
	return nil
}
