package models

import (
    "database/sql"
)

var DB *sql.DB

type User struct {
    UserId int
	Email  string
}

type Feed struct {
    FeedId int
	Url    string
}