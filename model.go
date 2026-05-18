package main

import "time"

// struct utama untuk data todo
type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// struct untuk request saat membuat todo baru
type CreateTodoRequest struct {
	Title string `json:"title"`
}

// struct untuk request saat update todo
type UpdateTodoRequest struct {
	Title *string `json:"title"`
	Done  *bool   `json:"done"`
}

// struct untuk response list todo
type ListTodosResponse struct {
	Data  []Todo `json:"data"`
	Count int    `json:"count"`
}

// struct untuk response error
type ErrorResponse struct {
	Error string `json:"error"`
}
