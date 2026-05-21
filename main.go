package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// ambil jwt secret dari environment variable
	// kalau kosong, pakai default yang sama dengan auth-service
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "secret-key-demo"
	}

	// membuat tempat penyimpanan todo di memory
	store := NewTodoStore()

	// membuat app utama dan memasukkan store + jwt secret
	app := NewApp(store, jwtSecret)

	// membuat router bawaan dari package net/http
	mux := http.NewServeMux()

	// route ini tidak perlu token, cuma untuk cek service hidup
	mux.HandleFunc("/health", app.healthHandler)

	// route todo wajib pakai token
	mux.HandleFunc("/todos", app.authMiddleware(app.todosHandler))
	mux.HandleFunc("/todos/", app.authMiddleware(app.todoByIDHandler))

	// menampilkan pesan saat server mulai berjalan
	log.Println("server running at http://localhost:8080")

	// menjalankan server di port 8080
	log.Fatal(http.ListenAndServe(":8080", mux))
}
