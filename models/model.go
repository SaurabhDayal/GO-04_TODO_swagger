package models

import "time"

type Todo struct {
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	PendingAt   time.Time `json:"pendingAt" db:"pending_at"`
}

type TodoResponse struct {
	UserId      int       `json:"userId" db:"user_id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Status      bool      `json:"status" db:"status"`
	PendingAt   time.Time `json:"pendingAt" db:"pending_at"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
}
