package main

import (
	"net/http"
	"strconv"
	"strings"
)

// struct app menyimpan dependency yang dibutuhkan handler
type App struct {
	store *TodoStore
}

// function untuk membuat app baru
func NewApp(store *TodoStore) *App {
	return &App{
		store: store,
	}
}

// handler untuk mengecek apakah api aktif
func (app *App) healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"status": "ok",
	})
}

// handler untuk endpoint /todos
func (app *App) todosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		app.getTodos(w, r)
	case http.MethodPost:
		app.createTodo(w, r)
	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

// handler untuk endpoint /todos/{id}
func (app *App) todoByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, ok := getIDFromPath(r.URL.Path)
	if !ok {
		writeError(w, http.StatusBadRequest, "invalid todo id")
		return
	}

	switch r.Method {
	case http.MethodGet:
		app.getTodoByID(w, r, id)
	case http.MethodPatch:
		app.updateTodo(w, r, id)
	case http.MethodDelete:
		app.deleteTodo(w, r, id)
	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

// function untuk mengambil semua todo
func (app *App) getTodos(w http.ResponseWriter, r *http.Request) {
	todos := app.store.List()

	writeJSON(w, http.StatusOK, ListTodosResponse{
		Data:  todos,
		Count: len(todos),
	})
}

// function untuk membuat todo baru
func (app *App) createTodo(w http.ResponseWriter, r *http.Request) {
	var req CreateTodoRequest

	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	req.Title = strings.TrimSpace(req.Title)

	if req.Title == "" {
		writeError(w, http.StatusBadRequest, "title is required")
		return
	}

	todo := app.store.Create(req.Title)

	writeJSON(w, http.StatusCreated, todo)
}

// function untuk mengambil todo berdasarkan id
func (app *App) getTodoByID(w http.ResponseWriter, r *http.Request, id int) {
	todo, exists := app.store.Get(id)
	if !exists {
		writeError(w, http.StatusNotFound, "todo not found")
		return
	}

	writeJSON(w, http.StatusOK, todo)
}

// function untuk update todo berdasarkan id
func (app *App) updateTodo(w http.ResponseWriter, r *http.Request, id int) {
	var req UpdateTodoRequest

	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	if req.Title == nil && req.Done == nil {
		writeError(w, http.StatusBadRequest, "at least one field is required")
		return
	}

	if req.Title != nil && strings.TrimSpace(*req.Title) == "" {
		writeError(w, http.StatusBadRequest, "title cannot be empty")
		return
	}

	todo, exists := app.store.Update(id, req)
	if !exists {
		writeError(w, http.StatusNotFound, "todo not found")
		return
	}

	writeJSON(w, http.StatusOK, todo)
}

// function untuk menghapus todo berdasarkan id
func (app *App) deleteTodo(w http.ResponseWriter, r *http.Request, id int) {
	deleted := app.store.Delete(id)
	if !deleted {
		writeError(w, http.StatusNotFound, "todo not found")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// function untuk mengambil id dari url path
func getIDFromPath(path string) (int, bool) {
	idText := strings.TrimPrefix(path, "/todos/")

	if idText == "" || strings.Contains(idText, "/") {
		return 0, false
	}

	id, err := strconv.Atoi(idText)
	if err != nil {
		return 0, false
	}

	return id, true
}
