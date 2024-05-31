package models

import (
	"github.com/volatiletech/null"
	"time"
)

type Task struct {
	ID          string    `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	PendingAt   time.Time `json:"pendingAt" db:"pending_at"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	ArchivedAt  null.Time `json:"archivedAt" db:"archived_at"`
}
