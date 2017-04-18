package storage

import "time"

type Storage interface {
	Close() error

	// CreateShortUrl() error
	GetURL(url string) (URL, error)
	CreateShortURL(short URL) error
}

type URL struct {
	ID        int64
	URL       string
	Slug      string
	CreatedAt time.Time `db:"created_at"`
}
