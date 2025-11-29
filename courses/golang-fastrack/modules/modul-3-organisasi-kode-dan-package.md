# MODUL 3: ORGANISASI KODE & PACKAGE

## 3.1 Konsep Package

Setiap folder berisi file-file Go dengan deklarasi `package` yang sama. Package adalah cara Go mengorganisir kode menjadi modul yang bisa dipake ulang.

**Aturan package:**
- Satu folder = satu package
- Semua file `.go` dalam folder harus punya package yang sama
- `main` adalah package khusus untuk entry point program
- Fungsi/variable yang dimulai huruf capital adalah PUBLIC (bisa diakses dari package lain)
- Fungsi/variable yang dimulai huruf lowercase adalah PRIVATE (hanya dalam package itu sendiri)

## 3.2 Struktur Project Recommended

Untuk project kecil tanpa framework, struktur yang bagus adalah:

```
belajar-go-basic/
├── main.go
├── internal/
│   ├── models/
│   │   └── user.go
│   └── services/
│       └── userservice.go
└── go.mod
```

Penjelasan:
- `main.go`: Entry point program
- `internal/`: Folder untuk paket internal project (bukan untuk di-export)
- `models/`: Struct dan model data
- `services/`: Logika bisnis / operasi
- `go.mod`: File modul Go (mencatat dependensi)

## 3.3 Membuat Project Bertahap

**Langkah 1: Inisialisasi Project**

```bash
mkdir belajar-go-api
cd belajar-go-api
go mod init github.com/username/belajar-go-api
```

Perintah `go mod init` membuat file `go.mod` yang berisi metadata project.

**Langkah 2: Buat Struktur Folder**

```bash
mkdir -p internal/models
mkdir -p internal/services
```

**Langkah 3: Buat File Models**

File: `internal/models/user.go`

```go
package models

type User struct {
    ID    int
    Name  string
    Email string
}
```

**Langkah 4: Buat File Services**

File: `internal/services/userservice.go`

```go
package services

import "github.com/username/belajar-go-api/internal/models"

var users []models.User
var nextID = 1

func CreateUser(name, email string) models.User {
    user := models.User{
        ID:    nextID,
        Name:  name,
        Email: email,
    }
    users = append(users, user)
    nextID++
    return user
}

func GetAllUsers() []models.User {
    return users
}

func GetUserByID(id int) *models.User {
    for i := range users {
        if users[i].ID == id {
            return &users[i]
        }
    }
    return nil
}
```

**Langkah 5: Buat Main.go**

File: `main.go`

```go
package main

import (
    "fmt"
    "github.com/username/belajar-go-api/internal/models"
    "github.com/username/belajar-go-api/internal/services"
)

func main() {
    // Buat user
    user1 := services.CreateUser("Ahmad", "ahmad@mail.com")
    user2 := services.CreateUser("Budi", "budi@mail.com")

    fmt.Println("User baru:", user1)
    fmt.Println("User baru:", user2)

    // Ambil semua user
    allUsers := services.GetAllUsers()
    for _, user := range allUsers {
        fmt.Printf("ID: %d, Nama: %s, Email: %s\n", user.ID, user.Name, user.Email)
    }

    // Ambil user by ID
    user := services.GetUserByID(1)
    if user != nil {
        fmt.Println("Ditemukan:", user.Name)
    }
}
```

**Langkah 6: Jalankan**

```bash
go run main.go
```

## 3.4 Latihan Modul 3

**Latihan: Refactor Inventory Project ke Multiple Packages**

Ambil latihan dari Modul 2 tentang inventory produk, dan refactor menjadi:
- `internal/models/product.go` - struct Product
- `internal/services/productservice.go` - semua fungsi CRUD
- `main.go` - test/demo program
