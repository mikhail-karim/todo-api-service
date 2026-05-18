package main

import (
	"log"
	"net/http"
)

func main() {
	store := NewTodoStore() // membuat tempat penyimpanan todo di memory
	app := NewApp(store)    // membuat app utama dan memasukkan store ke dalam app

	mux := http.NewServeMux() // membuat router bawaan dari package net/http

	mux.HandleFunc("/health", app.healthHandler)   // route untuk mengecek apakah api berjalan
	mux.HandleFunc("/todos", app.todosHandler)     // route untuk list todo dan create todo
	mux.HandleFunc("/todos/", app.todoByIDHandler) // route untuk todo berdasarkan id

	log.Println("Server running at http://localhost:8080") // menampilkan pesan saat server mulai berjalan
	log.Fatal(http.ListenAndServe(":8080", mux))           // menjalankan server di port 8080
}
