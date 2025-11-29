# MODUL 4: PENGENALAN HTTP & net/http

## 4.1 Dasar HTTP Singkat

HTTP adalah protokol untuk komunikasi client-server di web. Metode HTTP yang umum:
- **GET**: Mengambil data
- **POST**: Mengirim/membuat data
- **PUT**: Mengupdate seluruh resource
- **PATCH**: Mengupdate sebagian resource
- **DELETE**: Menghapus data

Struktur endpoint:
- `GET /users` → Ambil semua user
- `GET /users/1` → Ambil user dengan ID 1
- `POST /users` → Buat user baru (body berisi data)
- `PUT /users/1` → Update user ID 1
- `DELETE /users/1` → Hapus user ID 1

## 4.2 net/http: HTTP Server Minimal

Go memiliki package `net/http` di standard library untuk membuat HTTP server.

**Konsep dasar:**
- Handler: Fungsi yang menangani request, signature: `func(w http.ResponseWriter, r *http.Request)`
- ResponseWriter: Digunakan untuk menulis response
- Request: Berisi informasi tentang request dari client

**Contoh Paling Sederhana:**

```go
package main

import (
    "fmt"
    "net/http"
)

// Handler function
func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, Go!")
}

func main() {
    // Register handler
    http.HandleFunc("/hello", helloHandler)

    // Jalankan server di port 8080
    fmt.Println("Server berjalan di http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
```

**Cara test:**
- Buka browser dan ketik: `http://localhost:8080/hello`
- Atau gunakan curl: `curl http://localhost:8080/hello`

## 4.3 Multiple Handlers & Routing Sederhana

```go
package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Halaman Hello")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Halaman About")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Halaman Contact")
}

func main() {
    http.HandleFunc("/hello", helloHandler)
    http.HandleFunc("/about", aboutHandler)
    http.HandleFunc("/contact", contactHandler)

    fmt.Println("Server berjalan di http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
```

## 4.4 Membaca Request Data

```go
package main

import (
    "fmt"
    "net/http"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
    // Cek method
    fmt.Fprintf(w, "Method: %s\n", r.Method)

    // Cek path
    fmt.Fprintf(w, "Path: %s\n", r.URL.Path)

    // Query parameter
    fmt.Fprintf(w, "Query String: %s\n", r.URL.RawQuery)

    // Ambil query parameter spesifik
    nama := r.URL.Query().Get("nama")
    fmt.Fprintf(w, "Nama: %s\n", nama)
}

func main() {
    http.HandleFunc("/test", testHandler)
    http.ListenAndServe(":8080", nil)
}
```

**Test dengan curl:**
```bash
curl "http://localhost:8080/test?nama=Ahmad"
```

## 4.5 Response Headers & Status Code

```go
package main

import (
    "net/http"
)

func okHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)  // 200
    w.Write([]byte("OK"))
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusNotFound)  // 404
    w.Write([]byte("Resource tidak ditemukan"))
}

func badRequestHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusBadRequest)  // 400
    w.Write([]byte("Bad request"))
}

func main() {
    http.HandleFunc("/ok", okHandler)
    http.HandleFunc("/notfound", notFoundHandler)
    http.HandleFunc("/badrequest", badRequestHandler)

    http.ListenAndServe(":8080", nil)
}
```

## 4.6 Latihan Modul 4

**Latihan 1: Server Sederhana dengan 3 Endpoint**

Buat server dengan endpoint:
- `GET /` → Tampilkan "Selamat datang di API"
- `GET /status` → Tampilkan "Status: OK"
- `GET /info?name=Ahmad` → Tampilkan "Hello, Ahmad!"

**Latihan 2: Method Checking**

Buat handler yang menampilkan:
- Status code 200 jika method GET
- Status code 405 (Method Not Allowed) jika method lain
