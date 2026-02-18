package main

import "time"

type Todo struct {
	Title       string
	Completed   bool
	CompletedAt time.Time
	CreatedAt   *time.Time
}


