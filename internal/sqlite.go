package internal

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type SqliteDao struct {
	db *sqlx.DB
}

func NewSqliteDao(dsn string) (*SqliteDao, error) {
	s, err := sqlx.Connect("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	driver, err := sqlite3.WithInstance(s.DB, &sqlite3.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://scripts/migrations",
		"sqlite3", driver)
	if err != nil {
		return nil, err
	}

	err = m.Up()
	if err != nil && err.Error() != "no change" {
		s.Close()
		return nil, err
	}

	sd := &SqliteDao{
		db: s,
	}
	return sd, nil
}

func (sd SqliteDao) AddFeed(url string) error {
	_, err := sd.db.Exec("INSERT INTO feed (url) VALUES($1)", url)
	return err
}

func (sd SqliteDao) GetFeeds() ([]Feed, error) {
	var feeds []Feed
	if err := sd.db.Select(&feeds, "SELECT id, url FROM feed"); err != nil {
		return nil, err
	}
	return feeds, nil
}

func (sd SqliteDao) Close() error {
	return sd.db.Close()
}
