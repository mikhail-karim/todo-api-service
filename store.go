package main

import (
	"sort"
	"strings"
	"sync"
	"time"
)

// struct untuk menyimpan data todo
type TodoStore struct {
	mu     sync.Mutex   // mutex digunakan agar data aman saat diakses bersamaan
	todos  map[int]Todo // map digunakan untuk menyimpan todo berdasarkan id
	nextID int          // nextid digunakan untuk membuat id otomatis
}

// function untuk membuat todo store baru
func NewTodoStore() *TodoStore {
	return &TodoStore{
		todos:  make(map[int]Todo), // membuat map kosong untuk menyimpan todo
		nextID: 1,                  // id pertama dimulai dari 1
	}
}

// method untuk mengambil semua todo
func (s *TodoStore) List() []Todo {
	s.mu.Lock()         // mengunci store agar aman saat dibaca
	defer s.mu.Unlock() // membuka kunci setelah function selesai

	result := make([]Todo, 0, len(s.todos)) // membuat slice kosong untuk menampung hasil todo

	// memasukkan semua todo dari map ke slice
	for _, todo := range s.todos {
		result = append(result, todo)
	}

	// mengurutkan todo berdasarkan id
	sort.Slice(result, func(i, j int) bool {
		return result[i].ID < result[j].ID
	})

	return result
}

// method untuk membuat todo baru
func (s *TodoStore) Create(title string) Todo {
	s.mu.Lock()         // mengunci store karena data akan diubah
	defer s.mu.Unlock() // membuka kunci setelah function selesai

	now := time.Now() // mengambil waktu saat ini

	// membuat object todo baru
	todo := Todo{
		ID:        s.nextID,
		Title:     title,
		Done:      false,
		CreatedAt: now,
		UpdatedAt: now,
	}

	s.todos[todo.ID] = todo // menyimpan todo ke dalam map berdasarkan id
	s.nextID++              // menaikkan id untuk todo berikutnya

	return todo
}

// method untuk mengambil todo berdasarkan id
func (s *TodoStore) Get(id int) (Todo, bool) {
	s.mu.Lock()         // mengunci store agar aman saat dibaca
	defer s.mu.Unlock() // membuka kunci setelah function selesai

	todo, exists := s.todos[id] // mencari todo berdasarkan id
	return todo, exists
}

// method untuk update todo berdasarkan id
func (s *TodoStore) Update(id int, req UpdateTodoRequest) (Todo, bool) {
	s.mu.Lock()         // mengunci store karena data akan diubah
	defer s.mu.Unlock() // membuka kunci setelah function selesai

	todo, exists := s.todos[id] // mencari todo berdasarkan id
	// jika todo tidak ditemukan, kembalikan false
	if !exists {
		return Todo{}, false
	}

	// jika title dikirim, update title
	if req.Title != nil {
		todo.Title = strings.TrimSpace(*req.Title)
	}

	// jika done dikirim, update status done
	if req.Done != nil {
		todo.Done = *req.Done
	}

	todo.UpdatedAt = time.Now() // update waktu perubahan terakhir
	s.todos[id] = todo          // simpan kembali todo yang sudah diubah

	return todo, true
}

// method untuk menghapus todo berdasarkan id
func (s *TodoStore) Delete(id int) bool {
	s.mu.Lock()         // mengunci store karena data akan dihapus
	defer s.mu.Unlock() // membuka kunci setelah function selesai

	// cek apakah todo ada
	if _, exists := s.todos[id]; !exists {
		return false
	}

	delete(s.todos, id) // hapus todo dari map
	return true
}
