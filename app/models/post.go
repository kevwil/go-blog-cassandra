package models

import (
	"time"
)

type Post struct {
	Id      uint
	Title   string
	Tags    string
	Content string
	Date    time.Time
}
