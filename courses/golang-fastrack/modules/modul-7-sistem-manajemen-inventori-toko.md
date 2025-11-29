# MODUL 7: STUDI KASUS - SISTEM MANAJEMEN INVENTORI TOKO

## 7.1 Desain Project

Kita akan membuat REST API lengkap untuk sistem manajemen inventori toko dengan fitur CRUD produk dan tracking stok.

**Fitur yang akan dibuat:**
- Menambah produk baru ke inventori
- Melihat daftar produk dengan stok
- Mencari produk berdasarkan ID atau nama
- Update harga dan stok produk
- Mengurangi stok saat ada penjualan
- Menambah stok saat ada pembelian
- Menghapus produk dari inventori

**Endpoint yang akan dibuat:**
- `GET /products` → Ambil semua produk
- `GET /products/:id` → Ambil produk spesifik
- `POST /products` → Tambah produk baru
- `PUT /products/:id` → Update produk
- `DELETE /products/:id` → Hapus produk
- `POST /products/:id/sell` → Kurangi stok penjualan
- `POST /products/:id/restock` → Tambah stok pembelian
- `GET /products/low-stock` → Tampilkan produk dengan stok rendah

**Data Model:**
```
Product {
  id: integer
  name: string
  description: string
  price: float64
  stock: integer
  sku: string
  category: string
  created_at: string
}
```

## 7.2 Struktur Project

```
inventory-api/
├── main.go
├── go.mod
└── internal/
    ├── models/
    │   └── product.go
    ├── handlers/
    │   └── product_handler.go
    ├── storage/
    │   └── product_storage.go
    └── utils/
        └── validator.go
```

## 7.3 Step 1: Buat go.mod dan Struktur Folder

```bash
mkdir inventory-api
cd inventory-api
go mod init github.com/username/inventory-api
mkdir -p internal/models
mkdir -p internal/handlers
mkdir -p internal/storage
mkdir -p internal/utils
```

## 7.4 Step 2: Buat Models

File: `internal/models/product.go`

```go
package models

import "time"

type Product struct {
    ID          int       `json:"id"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    Price       float64   `json:"price"`
    Stock       int       `json:"stock"`
    SKU         string    `json:"sku"`
    Category    string    `json:"category"`
    CreatedAt   time.Time `json:"created_at"`
}

type SellRequest struct {
    Quantity int `json:"quantity"`
}

type RestockRequest struct {
    Quantity int `json:"quantity"`
}

type CreateProductRequest struct {
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Price       float64 `json:"price"`
    Stock       int     `json:"stock"`
    SKU         string  `json:"sku"`
    Category    string  `json:"category"`
}

type UpdateProductRequest struct {
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Price       float64 `json:"price"`
    SKU         string  `json:"sku"`
    Category    string  `json:"category"`
}
```

## 7.5 Step 3: Buat Validator Utility

File: `internal/utils/validator.go`

```go
package utils

import "errors"

func ValidateProduct(name string, price float64, stock int, sku string) error {
    if name == "" {
        return errors.New("nama produk tidak boleh kosong")
    }
    if len(name) < 3 {
        return errors.New("nama produk minimal 3 karakter")
    }
    if price < 0 {
        return errors.New("harga tidak boleh negatif")
    }
    if stock < 0 {
        return errors.New("stok tidak boleh negatif")
    }
    if sku == "" {
        return errors.New("SKU tidak boleh kosong")
    }
    return nil
}

func ValidateSellQuantity(currentStock int, sellQuantity int) error {
    if sellQuantity <= 0 {
        return errors.New("jumlah penjualan harus lebih dari 0")
    }
    if sellQuantity > currentStock {
        return errors.New("stok tidak cukup untuk penjualan ini")
    }
    return nil
}

func ValidateRestockQuantity(restockQuantity int) error {
    if restockQuantity <= 0 {
        return errors.New("jumlah restok harus lebih dari 0")
    }
    return nil
}
```

## 7.6 Step 4: Buat Storage Layer

File: `internal/storage/product_storage.go`

```go
package storage

import (
    "fmt"
    "time"
    "github.com/username/inventory-api/internal/models"
)

type ProductStorage struct {
    products map[int]models.Product
    nextID   int
}

func NewProductStorage() *ProductStorage {
    return &ProductStorage{
        products: make(map[int]models.Product),
        nextID:   1,
    }
}

// GetAll mengembalikan semua produk
func (ps *ProductStorage) GetAll() []models.Product {
    var products []models.Product
    for _, product := range ps.products {
        products = append(products, product)
    }
    return products
}

// GetByID mencari produk berdasarkan ID
func (ps *ProductStorage) GetByID(id int) (*models.Product, error) {
    product, exists := ps.products[id]
    if !exists {
        return nil, fmt.Errorf("produk dengan ID %d tidak ditemukan", id)
    }
    return &product, nil
}

// GetBySKU mencari produk berdasarkan SKU
func (ps *ProductStorage) GetBySKU(sku string) (*models.Product, error) {
    for _, product := range ps.products {
        if product.SKU == sku {
            return &product, nil
        }
    }
    return nil, fmt.Errorf("produk dengan SKU %s tidak ditemukan", sku)
}

// Create membuat produk baru
func (ps *ProductStorage) Create(name, description string, price float64, stock int, sku, category string) models.Product {
    product := models.Product{
        ID:          ps.nextID,
        Name:        name,
        Description: description,
        Price:       price,
        Stock:       stock,
        SKU:         sku,
        Category:    category,
        CreatedAt:   time.Now(),
    }
    ps.products[ps.nextID] = product
    ps.nextID++
    return product
}

// Update mengupdate produk
func (ps *ProductStorage) Update(id int, name, description string, price float64, sku, category string) (*models.Product, error) {
    product, exists := ps.products[id]
    if !exists {
        return nil, fmt.Errorf("produk dengan ID %d tidak ditemukan", id)
    }

    product.Name = name
    product.Description = description
    product.Price = price
    product.SKU = sku
    product.Category = category

    ps.products[id] = product
    return &product, nil
}

// Delete menghapus produk
func (ps *ProductStorage) Delete(id int) error {
    _, exists := ps.products[id]
    if !exists {
        return fmt.Errorf("produk dengan ID %d tidak ditemukan", id)
    }
    delete(ps.products, id)
    return nil
}

// Sell mengurangi stok penjualan
func (ps *ProductStorage) Sell(id int, quantity int) (*models.Product, error) {
    product, exists := ps.products[id]
    if !exists {
        return nil, fmt.Errorf("produk dengan ID %d tidak ditemukan", id)
    }

    if quantity > product.Stock {
        return nil, fmt.Errorf("stok tidak cukup. stok saat ini: %d, diminta: %d", product.Stock, quantity)
    }

    product.Stock -= quantity
    ps.products[id] = product
    return &product, nil
}

// Restock menambah stok pembelian
func (ps *ProductStorage) Restock(id int, quantity int) (*models.Product, error) {
    product, exists := ps.products[id]
    if !exists {
        return nil, fmt.Errorf("produk dengan ID %d tidak ditemukan", id)
    }

    product.Stock += quantity
    ps.products[id] = product
    return &product, nil
}

// GetLowStock menampilkan produk dengan stok rendah (< threshold)
func (ps *ProductStorage) GetLowStock(threshold int) []models.Product {
    var lowStockProducts []models.Product
    for _, product := range ps.products {
        if product.Stock < threshold {
            lowStockProducts = append(lowStockProducts, product)
        }
    }
    return lowStockProducts
}
```

## 7.7 Step 5: Buat Handler

File: `internal/handlers/product_handler.go`

```go
package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "strings"
    "github.com/username/inventory-api/internal/models"
    "github.com/username/inventory-api/internal/storage"
    "github.com/username/inventory-api/internal/utils"
)

type ProductHandler struct {
    storage *storage.ProductStorage
}

func NewProductHandler(storage *storage.ProductStorage) *ProductHandler {
    return &ProductHandler{storage: storage}
}

// HandleProducts menangani GET dan POST /products
func (ph *ProductHandler) HandleProducts(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    switch r.Method {
    case "GET":
        ph.handleGetProducts(w, r)
    case "POST":
        ph.handleCreateProduct(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

// handleGetProducts menangani GET /products
func (ph *ProductHandler) handleGetProducts(w http.ResponseWriter, r *http.Request) {
    products := ph.storage.GetAll()
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(products)
}

// handleCreateProduct menangani POST /products
func (ph *ProductHandler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
    var req models.CreateProductRequest

    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    // Validasi input
    err = utils.ValidateProduct(req.Name, req.Price, req.Stock, req.SKU)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Cek SKU sudah ada atau belum
    _, err = ph.storage.GetBySKU(req.SKU)
    if err == nil {
        http.Error(w, "SKU sudah terdaftar", http.StatusConflict)
        return
    }

    product := ph.storage.Create(req.Name, req.Description, req.Price, req.Stock, req.SKU, req.Category)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(product)
}

// HandleProductByID menangani GET, PUT, DELETE /products/:id
func (ph *ProductHandler) HandleProductByID(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // Extract ID dari path
    id := ph.extractID(r.URL.Path)
    if id == -1 {
        http.Error(w, "Invalid product ID", http.StatusBadRequest)
        return
    }

    switch r.Method {
    case "GET":
        ph.handleGetProductByID(w, id)
    case "PUT":
        ph.handleUpdateProduct(w, r, id)
    case "DELETE":
        ph.handleDeleteProduct(w, id)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

// handleGetProductByID menangani GET /products/:id
func (ph *ProductHandler) handleGetProductByID(w http.ResponseWriter, id int) {
    product, err := ph.storage.GetByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(product)
}

// handleUpdateProduct menangani PUT /products/:id
func (ph *ProductHandler) handleUpdateProduct(w http.ResponseWriter, r *http.Request, id int) {
    var req models.UpdateProductRequest

    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    if req.Name == "" || req.Price < 0 || req.SKU == "" {
        http.Error(w, "Invalid product data", http.StatusBadRequest)
        return
    }

    product, err := ph.storage.Update(id, req.Name, req.Description, req.Price, req.SKU, req.Category)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(product)
}

// handleDeleteProduct menangani DELETE /products/:id
func (ph *ProductHandler) handleDeleteProduct(w http.ResponseWriter, id int) {
    err := ph.storage.Delete(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}

// HandleSellProduct menangani POST /products/:id/sell
func (ph *ProductHandler) HandleSellProduct(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    if r.Method != "POST" {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    id := ph.extractID(r.URL.Path)
    if id == -1 {
        http.Error(w, "Invalid product ID", http.StatusBadRequest)
        return
    }

    var req models.SellRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    err = utils.ValidateSellQuantity(0, req.Quantity)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    product, err := ph.storage.Sell(id, req.Quantity)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(product)
}

// HandleRestockProduct menangani POST /products/:id/restock
func (ph *ProductHandler) HandleRestockProduct(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    if r.Method != "POST" {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    id := ph.extractID(r.URL.Path)
    if id == -1 {
        http.Error(w, "Invalid product ID", http.StatusBadRequest)
        return
    }

    var req models.RestockRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    err = utils.ValidateRestockQuantity(req.Quantity)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    product, err := ph.storage.Restock(id, req.Quantity)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(product)
}

// HandleLowStock menangani GET /products/low-stock
func (ph *ProductHandler) HandleLowStock(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    if r.Method != "GET" {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    threshold := 10 // Default threshold
    thresholdParam := r.URL.Query().Get("threshold")
    if thresholdParam != "" {
        t, err := strconv.Atoi(thresholdParam)
        if err == nil && t > 0 {
            threshold = t
        }
    }

    products := ph.storage.GetLowStock(threshold)
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(products)
}

// Helper function untuk extract ID dari path
func (ph *ProductHandler) extractID(path string) int {
    parts := strings.Split(path, "/")
    if len(parts) < 3 {
        return -1
    }
    idStr := parts[2]
    id, err := strconv.Atoi(idStr)
    if err != nil {
        return -1
    }
    return id
}
```

## 7.8 Step 6: Buat Main.go

File: `main.go`

```go
package main

import (
    "fmt"
    "net/http"
    "github.com/username/inventory-api/internal/handlers"
    "github.com/username/inventory-api/internal/storage"
)

func main() {
    // Inisialisasi storage dan handler
    productStorage := storage.NewProductStorage()
    productHandler := handlers.NewProductHandler(productStorage)

    // Register routes
    http.HandleFunc("/products", productHandler.HandleProducts)
    http.HandleFunc("/products/low-stock", productHandler.HandleLowStock)
    http.HandleFunc("/products/", handleProductRoutes(productHandler))

    fmt.Println("=== Inventory Management API ===")
    fmt.Println("Server berjalan di http://localhost:8080")
    fmt.Println("\nEndpoints:")
    fmt.Println("  GET    /products                  - Ambil semua produk")
    fmt.Println("  POST   /products                  - Tambah produk baru")
    fmt.Println("  GET    /products/:id              - Ambil produk spesifik")
    fmt.Println("  PUT    /products/:id              - Update produk")
    fmt.Println("  DELETE /products/:id              - Hapus produk")
    fmt.Println("  POST   /products/:id/sell         - Kurangi stok penjualan")
    fmt.Println("  POST   /products/:id/restock      - Tambah stok pembelian")
    fmt.Println("  GET    /products/low-stock        - Tampilkan stok rendah")

    http.ListenAndServe(":8080", nil)
}

func handleProductRoutes(ph *handlers.ProductHandler) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        path := r.URL.Path

        // Cek apakah ada suffix /sell atau /restock
        if len(path) > 10 && path[len(path)-5:] == "/sell" {
            ph.HandleSellProduct(w, r)
            return
        }

        if len(path) > 10 && path[len(path)-8:] == "/restock" {
            ph.HandleRestockProduct(w, r)
            return
        }

        // Jika tidak ada suffix, handle normal product by ID
        ph.HandleProductByID(w, r)
    }
}
```

## 7.9 Menjalankan Project

```bash
go run main.go
```

## 7.10 Testing API

Gunakan curl untuk test semua endpoint:

**1. CREATE produk baru:**
```bash
curl -X POST -H "Content-Type: application/json" \
  -d '{
    "name":"Laptop Asus",
    "description":"Laptop Gaming Performance",
    "price":12000000,
    "stock":5,
    "sku":"ASUS-001",
    "category":"Elektronik"
  }' \
  http://localhost:8080/products
```

**2. CREATE produk lagi:**
```bash
curl -X POST -H "Content-Type: application/json" \
  -d '{
    "name":"Mouse Logitech",
    "description":"Mouse Gaming RGB",
    "price":500000,
    "stock":20,
    "sku":"LOG-001",
    "category":"Aksesoris"
  }' \
  http://localhost:8080/products
```

**3. CREATE produk ketiga:**
```bash
curl -X POST -H "Content-Type: application/json" \
  -d '{
    "name":"Keyboard Mechanical",
    "description":"Keyboard RGB Mechanical",
    "price":1200000,
    "stock":3,
    "sku":"KEY-001",
    "category":"Aksesoris"
  }' \
  http://localhost:8080/products
```

**4. GET semua produk:**
```bash
curl http://localhost:8080/products
```

**5. GET produk spesifik:**
```bash
curl http://localhost:8080/products/1
```

**6. UPDATE produk (ubah harga):**
```bash
curl -X PUT -H "Content-Type: application/json" \
  -d '{
    "name":"Laptop Asus TUF",
    "description":"Laptop Gaming Performance - Updated",
    "price":13000000,
    "sku":"ASUS-001",
    "category":"Elektronik"
  }' \
  http://localhost:8080/products/1
```

**7. SELL produk (kurangi stok):**
```bash
curl -X POST -H "Content-Type: application/json" \
  -d '{"quantity":2}' \
  http://localhost:8080/products/1/sell
```

**8. RESTOCK produk (tambah stok):**
```bash
curl -X POST -H "Content-Type: application/json" \
  -d '{"quantity":10}' \
  http://localhost:8080/products/2/restock
```

**9. GET low stock products (stok rendah):**
```bash
curl http://localhost:8080/products/low-stock
```

**10. GET low stock dengan threshold custom:**
```bash
curl "http://localhost:8080/products/low-stock?threshold=5"
```

**11. DELETE produk:**
```bash
curl -X DELETE http://localhost:8080/products/3
```

## 7.11 Penambahan Fitur (Opsional)

**Fitur 1: Search Produk berdasarkan Nama**

Tambahkan endpoint:
- `GET /products/search?q=laptop` → Cari produk dengan nama mengandung "laptop"

**Fitur 2: Filter berdasarkan Kategori**

Tambahkan endpoint:
- `GET /products?category=Elektronik` → Tampilkan hanya produk kategori tertentu

**Fitur 3: Laporan Penjualan Harian**

Tambahkan tracking untuk setiap penjualan dan buat endpoint:
- `GET /sales/report` → Tampilkan laporan penjualan

**Fitur 4: Export ke Format Lain**

Tambahkan endpoint:
- `GET /products/export?format=csv` → Export data produk ke CSV

## 7.12 Latihan Modul 7

**Challenge 1: Tracking Riwayat Penjualan**

Tambahkan:
- Struct `SalesHistory` untuk tracking setiap transaksi
- Endpoint `GET /products/:id/sales-history` untuk melihat riwayat penjualan produk tertentu

**Challenge 2: Multi-Toko**

Modifikasi sistem untuk mendukung multiple toko:
- Tambahkan field `store_id` di Product struct
- Endpoint `/stores/:store_id/products` untuk produk per toko
- Endpoint untuk transfer stok antar toko

**Challenge 3: Notifikasi Stok Minimal**

Implementasikan sistem notifikasi:
- Endpoint `POST /settings/low-stock-threshold` untuk set threshold per user
- Endpoint `GET /notifications` untuk melihat notifikasi stok rendah

## 7.13 Kesimpulan

Selamat! Kamu telah menyelesaikan Modul 7 dan mempelajari cara membuat sistem manajemen inventori toko yang lengkap dengan Go.

**Skill yang sudah dikuasai di modul ini:**
- Membuat REST API kompleks dengan multiple endpoints
- Implementasi validasi input yang baik
- Manajemen state dengan in-memory storage
- Error handling yang comprehensive
- Struktur project yang terorganisir dengan baik

**Next Steps:**
1. **Database Integration**: Integrasikan dengan database (PostgreSQL, MySQL)
2. **Authentication**: Tambahkan sistem autentikasi dan autorisasi
3. **Testing**: Tulis unit tests dan integration tests
4. **Deployment**: Deploy aplikasi ke production
5. **Monitoring**: Tambahkan logging dan monitoring
