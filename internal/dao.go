package internal

type Dao interface {
	AddFeed(url string) error
	GetFeeds() ([]Feed, error)
	Close() error
}
