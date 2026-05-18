# todo api golang

todo api sederhana yang dibuat menggunakan golang dan package bawaan `net/http`.

project ini dibuat sebagai latihan dasar backend api dengan golang. fitur utama yang tersedia adalah membuat, melihat, mengubah, dan menghapus todo.

## fitur

- health check api
- membuat todo baru
- melihat semua todo
- melihat todo berdasarkan id
- mengubah todo
- menghapus todo
- response dalam format json
- validasi request sederhana
- penyimpanan data sementara di memory

## teknologi

- golang
- net/http
- json
- curl / postman untuk testing

## struktur project

```txt
todo_api/
├── main.go
├── model.go
├── store.go
├── handler.go
└── response.go
```

## endpoint api

```txt
GET     /health
GET     /todos
POST    /todos
GET     /todos/{id}
PATCH   /todos/{id}
DELETE  /todos/{id}
```

## cara menjalankan

jalankan perintah berikut di terminal:

```bash
go run .
```

server akan berjalan di:

```txt
http://localhost:8080
```

## contoh request

### health check

```bash
curl http://localhost:8080/health
```

### membuat todo

untuk windows cmd:

```bat
curl -X POST http://localhost:8080/todos -H "Content-Type: application/json" -d "{\"title\":\"belajar golang api\"}"
```

untuk powershell:

```powershell
curl.exe -X POST http://localhost:8080/todos -H "Content-Type: application/json" -d '{"title":"belajar golang api"}'
```

### melihat semua todo

```bash
curl http://localhost:8080/todos
```

### melihat todo berdasarkan id

```bash
curl http://localhost:8080/todos/1
```

### update todo

untuk windows cmd:

```bat
curl -X PATCH http://localhost:8080/todos/1 -H "Content-Type: application/json" -d "{\"done\":true}"
```

untuk powershell:

```powershell
curl.exe -X PATCH http://localhost:8080/todos/1 -H "Content-Type: application/json" -d '{"done":true}'
```

### hapus todo

```bash
curl -X DELETE http://localhost:8080/todos/1
```

## catatan

data todo masih disimpan di memory, jadi data akan hilang saat server dimatikan.

project ini dibuat untuk memahami dasar backend api di golang, seperti routing, handler, struct, json, pointer, dan method http.