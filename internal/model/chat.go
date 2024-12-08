package model

import "time"

type Message struct {
	ChatId    int64
	From      string
	Text      string
	Timestamp *time.Time
}
