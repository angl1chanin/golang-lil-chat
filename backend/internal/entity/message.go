package entity

import "time"

type Message struct {
	ID        int       `json:"id"`
	Author    string    `json:"author"`
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
}
