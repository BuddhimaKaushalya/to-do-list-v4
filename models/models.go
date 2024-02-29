package models

import "time"

type Task struct {
	Description string `json:"Task"`
	CreatedAt   time.Time
}

var Tasks []Task
