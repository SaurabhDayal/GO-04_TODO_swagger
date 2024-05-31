package models

import "time"

type UserClient struct {
	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type UserRole struct {
	Id        int       `json:"id" db:"id"`
	UserId    int       `json:"userId" db:"user_id"`
	Role      string    `json:"role" db:"role"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type UserSession struct {
	UserId int    `db:"user_id"`
	Token  string `db:"session_token"`
}

type Todo struct {
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	PendingAt   time.Time `json:"pendingAt" db:"pending_at"`
}

type TodoClient struct {
	UserId      int       `json:"userId" db:"user_id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Status      bool      `json:"status" db:"status"`
	PendingAt   time.Time `json:"pendingAt" db:"pending_at"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
}
