package storage

import "time"

type Storage interface {
	Close() error

	// CreateShortUrl() error
	GetUrl(url string) (URL, error)
	CreateShortUrl(short URL) error
}

type URL struct {
	ID        int64
	URL       string
	Slug      string
	CreatedAt time.Time
}
