# MODUL 5: JSON & REST API TANPA FRAMEWORK

## 5.1 Encoding JSON

Package `encoding/json` mengubah struct Go menjadi JSON dan sebaliknya.

```go
package main

import (
    "encoding/json"
    "fmt"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

func main() {
    // Marshal: struct → JSON
    user := User{ID: 1, Name: "Ahmad", Email: "ahmad@mail.com"}
    jsonData, err := json.Marshal(user)
    if err != nil {
        fmt.Println("Error:", err)
    }
    fmt.Println(string(jsonData))  // {"id":1,"name":"Ahmad","email":"ahmad@mail.com"}

    // Unmarshal: JSON → struct
    jsonString := `{"id":2,"name":"Budi","email":"budi@mail.com"}`
    var user2 User
    err = json.Unmarshal([]byte(jsonString), &user2)
    if err != nil {
        fmt.Println("Error:", err)
    }
    fmt.Println(user2)  // {2 Budi budi@mail.com}
}
```

**JSON Tags:**
- `json:"fieldname"` - nama field di JSON
- `json:"fieldname,omitempty"` - exclude jika kosong
- `json:"-"` - exclude field ini dari JSON

## 5.2 Response JSON di HTTP Handler

```go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
    users := []User{
        {ID: 1, Name: "Ahmad", Email: "ahmad@mail.com"},
        {ID: 2, Name: "Budi", Email: "budi@mail.com"},
    }

    // Set header
    w.Header().Set("Content-Type", "application/json")

    // Marshal dan tulis
    json.NewEncoder(w).Encode(users)
}

func main() {
    http.HandleFunc("/users", getUsersHandler)
    fmt.Println("Server di http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
```

**Test dengan curl:**
```bash
curl http://localhost:8080/users
```

## 5.3 Membaca Request Body JSON

```go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var user User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)  // 201
    json.NewEncoder(w).Encode(user)
}

func main() {
    http.HandleFunc("/users", createUserHandler)
    http.ListenAndServe(":8080", nil)
}
```

**Test dengan curl:**
```bash
curl -X POST -H "Content-Type: application/json" \
  -d '{"name":"Ahmad","email":"ahmad@mail.com"}' \
  http://localhost:8080/users
```

## 5.4 Routing & Handling Multiple Methods

Untuk API yang lebih terstruktur, sering perlu menangani multiple methods di satu path:

```go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strings"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

var users []User
var nextID = 1

func usersHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        handleGetUsers(w, r)
    case "POST":
        handleCreateUser(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func handleGetUsers(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func handleCreateUser(w http.ResponseWriter, r *http.Request) {
    var user User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    user.ID = nextID
    nextID++
    users = append(users, user)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

func main() {
    http.HandleFunc("/users", usersHandler)
    fmt.Println("Server di http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
```

## 5.5 Latihan Modul 5

**Latihan: Simple API dengan GET dan POST**

Buat API dengan endpoint:
- `GET /tasks` → Tampilkan semua task (dalam JSON array)
- `POST /tasks` → Buat task baru (terima JSON dengan field "title")

Gunakan in-memory storage (slice).
