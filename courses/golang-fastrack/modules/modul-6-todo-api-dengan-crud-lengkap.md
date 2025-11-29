# MODUL 6: STUDI KASUS - TO-DO API DENGAN CRUD LENGKAP

## 6.1 Desain Project

Kita akan membuat REST API lengkap untuk manajemen to-do list dengan CRUD operations.

**Endpoint yang akan dibuat:**
- `GET /tasks` → Ambil semua task
- `GET /tasks/:id` → Ambil task spesifik
- `POST /tasks` → Buat task baru
- `PUT /tasks/:id` → Update task
- `DELETE /tasks/:id` → Hapus task

**Data Model:**
```
Task {
  id: integer
  title: string
  description: string
  completed: boolean
}
```

## 6.2 Struktur Project

```
todo-api/
├── main.go
├── go.mod
└── internal/
    ├── models/
    │   └── task.go
    ├── handlers/
    │   └── task_handler.go
    └── storage/
        └── task_storage.go
```

## 6.3 Step 1: Buat go.mod dan Struktur Folder

```bash
mkdir todo-api
cd todo-api
go mod init github.com/username/todo-api
mkdir -p internal/models
mkdir -p internal/handlers
mkdir -p internal/storage
```

## 6.4 Step 2: Buat Models

File: `internal/models/task.go`

```go
package models

type Task struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Completed   bool   `json:"completed"`
}
```

## 6.5 Step 3: Buat Storage Layer

File: `internal/storage/task_storage.go`

```go
package storage

import (
    "fmt"
    "github.com/username/todo-api/internal/models"
)

type TaskStorage struct {
    tasks  []models.Task
    nextID int
}

func NewTaskStorage() *TaskStorage {
    return &TaskStorage{
        tasks:  []models.Task{},
        nextID: 1,
    }
}

// GetAll mengembalikan semua task
func (ts *TaskStorage) GetAll() []models.Task {
    return ts.tasks
}

// GetByID mencari task berdasarkan ID
func (ts *TaskStorage) GetByID(id int) (*models.Task, error) {
    for i := range ts.tasks {
        if ts.tasks[i].ID == id {
            return &ts.tasks[i], nil
        }
    }
    return nil, fmt.Errorf("task dengan ID %d tidak ditemukan", id)
}

// Create membuat task baru
func (ts *TaskStorage) Create(title, description string) models.Task {
    task := models.Task{
        ID:          ts.nextID,
        Title:       title,
        Description: description,
        Completed:   false,
    }
    ts.tasks = append(ts.tasks, task)
    ts.nextID++
    return task
}

// Update mengupdate task
func (ts *TaskStorage) Update(id int, title, description string, completed bool) (*models.Task, error) {
    for i := range ts.tasks {
        if ts.tasks[i].ID == id {
            ts.tasks[i].Title = title
            ts.tasks[i].Description = description
            ts.tasks[i].Completed = completed
            return &ts.tasks[i], nil
        }
    }
    return nil, fmt.Errorf("task dengan ID %d tidak ditemukan", id)
}

// Delete menghapus task
func (ts *TaskStorage) Delete(id int) error {
    for i := range ts.tasks {
        if ts.tasks[i].ID == id {
            ts.tasks = append(ts.tasks[:i], ts.tasks[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("task dengan ID %d tidak ditemukan", id)
}
```

## 6.6 Step 4: Buat Handler

File: `internal/handlers/task_handler.go`

```go
package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "strings"
    "github.com/username/todo-api/internal/models"
    "github.com/username/todo-api/internal/storage"
)

type TaskHandler struct {
    storage *storage.TaskStorage
}

func NewTaskHandler(storage *storage.TaskStorage) *TaskHandler {
    return &TaskHandler{storage: storage}
}

// HandleTasks menangani GET dan POST /tasks
func (th *TaskHandler) HandleTasks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    switch r.Method {
    case "GET":
        th.handleGetTasks(w, r)
    case "POST":
        th.handleCreateTask(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

// handleGetTasks menangani GET /tasks
func (th *TaskHandler) handleGetTasks(w http.ResponseWriter, r *http.Request) {
    tasks := th.storage.GetAll()
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(tasks)
}

// handleCreateTask menangani POST /tasks
func (th *TaskHandler) handleCreateTask(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Title       string `json:"title"`
        Description string `json:"description"`
    }

    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    if req.Title == "" {
        http.Error(w, "Title is required", http.StatusBadRequest)
        return
    }

    task := th.storage.Create(req.Title, req.Description)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(task)
}

// HandleTaskByID menangani GET, PUT, DELETE /tasks/:id
func (th *TaskHandler) HandleTaskByID(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // Extract ID dari path
    parts := strings.Split(r.URL.Path, "/")
    idStr := parts[len(parts)-1]
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid task ID", http.StatusBadRequest)
        return
    }

    switch r.Method {
    case "GET":
        th.handleGetTaskByID(w, id)
    case "PUT":
        th.handleUpdateTask(w, r, id)
    case "DELETE":
        th.handleDeleteTask(w, id)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

// handleGetTaskByID menangani GET /tasks/:id
func (th *TaskHandler) handleGetTaskByID(w http.ResponseWriter, id int) {
    task, err := th.storage.GetByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(task)
}

// handleUpdateTask menangani PUT /tasks/:id
func (th *TaskHandler) handleUpdateTask(w http.ResponseWriter, r *http.Request, id int) {
    var req struct {
        Title       string `json:"title"`
        Description string `json:"description"`
        Completed   bool   `json:"completed"`
    }

    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    if req.Title == "" {
        http.Error(w, "Title is required", http.StatusBadRequest)
        return
    }

    task, err := th.storage.Update(id, req.Title, req.Description, req.Completed)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(task)
}

// handleDeleteTask menangani DELETE /tasks/:id
func (th *TaskHandler) handleDeleteTask(w http.ResponseWriter, id int) {
    err := th.storage.Delete(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}
```

## 6.7 Step 5: Buat Main.go

File: `main.go`

```go
package main

import (
    "fmt"
    "net/http"
    "github.com/username/todo-api/internal/handlers"
    "github.com/username/todo-api/internal/storage"
)

func main() {
    // Inisialisasi storage dan handler
    taskStorage := storage.NewTaskStorage()
    taskHandler := handlers.NewTaskHandler(taskStorage)

    // Register routes
    http.HandleFunc("/tasks", taskHandler.HandleTasks)
    http.HandleFunc("/tasks/", taskHandler.HandleTaskByID)

    fmt.Println("API Server berjalan di http://localhost:8080")
    fmt.Println("Endpoints:")
    fmt.Println("  GET    /tasks")
    fmt.Println("  POST   /tasks")
    fmt.Println("  GET    /tasks/:id")
    fmt.Println("  PUT    /tasks/:id")
    fmt.Println("  DELETE /tasks/:id")

    http.ListenAndServe(":8080", nil)
}
```

## 6.8 Menjalankan Project

```bash
go run main.go
```

## 6.9 Testing API

Gunakan curl atau Postman untuk test:

**1. GET semua task (awalnya kosong):**
```bash
curl http://localhost:8080/tasks
```
Output: `null` atau `[]`

**2. CREATE task baru:**
```bash
curl -X POST -H "Content-Type: application/json" \
  -d '{"title":"Belajar Go","description":"Memahami dasar Go"}' \
  http://localhost:8080/tasks
```
Output:
```json
{"id":1,"title":"Belajar Go","description":"Memahami dasar Go","completed":false}
```

**3. CREATE task lagi:**
```bash
curl -X POST -H "Content-Type: application/json" \
  -d '{"title":"Membuat API","description":"REST API tanpa framework"}' \
  http://localhost:8080/tasks
```

**4. GET semua task:**
```bash
curl http://localhost:8080/tasks
```

**5. GET task spesifik:**
```bash
curl http://localhost:8080/tasks/1
```

**6. UPDATE task:**
```bash
curl -X PUT -H "Content-Type: application/json" \
  -d '{"title":"Belajar Go dan Framework","description":"Pelajari Go lebih dalam","completed":true}' \
  http://localhost:8080/tasks/1
```

**7. DELETE task:**
```bash
curl -X DELETE http://localhost:8080/tasks/1
```

**8. Verify penghapusan:**
```bash
curl http://localhost:8080/tasks
```

## 6.10 Penambahan Fitur (Opsional)

**Fitur 1: Validasi Input**

Tambahkan validasi di handler untuk memastikan title tidak kosong dan panjangnya minimal 3 karakter.

**Fitur 2: Logging**

Tambahkan logging sederhana untuk setiap request:

```go
func logRequest(r *http.Request) {
    fmt.Printf("[%s] %s %s\n", r.Method, r.URL.Path, time.Now().Format("2006-01-02 15:04:05"))
}
```

**Fitur 3: Search/Filter**

Tambahkan query parameter untuk filter task yang completed atau belum.

## 6.11 Latihan Modul 6

**Challenge 1: Tambahan Endpoint**

Tambahkan endpoint baru:
- `GET /tasks/completed` → Tampilkan hanya task yang sudah complete

**Challenge 2: Pagination**

Tambahkan query parameter `?page=1&limit=5` untuk pagination di endpoint `GET /tasks`.

**Challenge 3: Sorting**

Tambahkan query parameter `?sort=created` atau `?sort=title` untuk sort hasil.
