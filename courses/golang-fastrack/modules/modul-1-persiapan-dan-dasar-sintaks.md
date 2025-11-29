# MODUL 1: PERSIAPAN & DASAR SINTAKS GO

## 1.1 Instalasi Go

**Langkah-langkah instalasi:**

1. Kunjungi situs resmi: https://go.dev/dl
2. Pilih versi untuk sistem operasi kamu (Windows, macOS, atau Linux)
3. Download dan jalankan installer
4. Ikuti langkah-langkah installer hingga selesai

**Verifikasi instalasi:**

Buka terminal/command prompt dan jalankan:

```bash
go version
```

Jika berhasil, kamu akan melihat output seperti:

```
go version go1.21.x darwin/amd64
```

**Setup Editor:**

Rekomendasi untuk pemula: Visual Studio Code

- Download dari https://code.visualstudio.com
- Install ekstensi "Go" dari Go Team at Google
- Buka folder project di VS Code

## 1.2 Struktur Program Go Paling Sederhana

Setiap program Go yang bisa dijalankan harus memiliki:

1. **Package**: `package main` (entry point program)
2. **Fungsi main**: `func main() { ... }` (tempat program dimulai)

**Contoh file pertama - hello.go:**

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

**Cara menjalankan:**

```bash
go run hello.go
```

Output:

```
Hello, Go!
```

## 1.3 Variabel & Konstanta

**Deklarasi Variabel:**

Ada 3 cara mendeklarasikan variabel di Go:

```go
package main

import "fmt"

func main() {
    // Cara 1: var dengan tipe eksplisit
    var nama string = "Budi"
    fmt.Println(nama)

    // Cara 2: var tanpa tipe (tipe inferensi)
    var umur = 25
    fmt.Println(umur)

    // Cara 3: short declaration (paling sering dipakai)
    kota := "Jakarta"
    fmt.Println(kota)

    // Variabel kosong (zero value)
    var harga float64  // akan bernilai 0.0
    fmt.Println(harga)
}
```

**Konstanta:**

Konstanta adalah nilai yang tidak bisa diubah setelah dideklarasikan:

```go
package main

import "fmt"

func main() {
    const pi = 3.14
    const namaApp = "To-Do App"

    fmt.Println(pi)
    fmt.Println(namaApp)

    // Berikut akan ERROR:
    // pi = 3.15  // ERROR: cannot assign to pi
}
```

## 1.4 Tipe Data Dasar

Go memiliki tipe data berikut:

```go
package main

import "fmt"

func main() {
    // String
    var nama string = "Ahmad"

    // Integer (berbagai ukuran)
    var umur int = 25
    var bilangan8 int8 = 127
    var bilangan64 int64 = 9223372036854775807

    // Unsigned integer (hanya positif)
    var positif uint = 100

    // Float
    var tinggi float64 = 175.5
    var berat float32 = 75.2

    // Boolean
    var aktif bool = true
    var nonaktif bool = false

    fmt.Println(nama)
    fmt.Println(umur)
    fmt.Println(tinggi)
    fmt.Println(aktif)
}
```

## 1.5 Operator Aritmatika & Perbandingan

```go
package main

import "fmt"

func main() {
    a := 10
    b := 3

    // Operator Aritmatika
    fmt.Println("Tambah:", a+b)      // 13
    fmt.Println("Kurang:", a-b)      // 7
    fmt.Println("Kali:", a*b)        // 30
    fmt.Println("Bagi:", a/b)        // 3
    fmt.Println("Modulo:", a%b)      // 1

    // Operator Perbandingan (hasilnya bool)
    fmt.Println("a > b:", a > b)     // true
    fmt.Println("a < b:", a < b)     // false
    fmt.Println("a == b:", a == b)   // false
    fmt.Println("a != b:", a != b)   // true
}
```

## 1.6 Struktur Kontrol: if/else & switch

**if/else:**

```go
package main

import "fmt"

func main() {
    nilai := 85

    if nilai >= 80 {
        fmt.Println("Nilai A - Excellent!")
    } else if nilai >= 70 {
        fmt.Println("Nilai B - Good")
    } else if nilai >= 60 {
        fmt.Println("Nilai C - Acceptable")
    } else {
        fmt.Println("Nilai D - Fail")
    }
}
```

**switch:**

```go
package main

import "fmt"

func main() {
    hari := "Senin"

    switch hari {
    case "Senin":
        fmt.Println("Awal minggu")
    case "Sabtu", "Minggu":
        fmt.Println("Weekend!")
    default:
        fmt.Println("Hari biasa")
    }
}
```

## 1.7 Perulangan: for

Go hanya memiliki satu keyword perulangan: `for`. Tetapi bisa digunakan untuk berbagai pola:

```go
package main

import "fmt"

func main() {
    // Pola 1: for tradisional
    for i := 0; i < 5; i++ {
        fmt.Println("Iterasi", i)
    }

    // Pola 2: for dengan kondisi (seperti while)
    counter := 0
    for counter < 3 {
        fmt.Println("Counter:", counter)
        counter++
    }

    // Pola 3: for infinite (harus ada break)
    count := 0
    for {
        if count >= 2 {
            break
        }
        fmt.Println("Infinite loop", count)
        count++
    }

    // Pola 4: for range (untuk slice, map, string)
    buah := []string{"Apel", "Mangga", "Pisang"}
    for i, fruit := range buah {
        fmt.Printf("Index %d: %s\n", i, fruit)
    }
}
```

## 1.8 Fungsi & Multiple Return Values

**Fungsi Dasar:**

```go
package main

import "fmt"

// Fungsi tanpa parameter dan return
func salam() {
    fmt.Println("Assalamualaikum!")
}

// Fungsi dengan parameter
func tambah(a int, b int) int {
    return a + b
}

// Fungsi dengan multiple return values
func pembagian(a int, b int) (int, int) {
    hasil := a / b
    sisa := a % b
    return hasil, sisa
}

func main() {
    salam()

    hasil := tambah(5, 3)
    fmt.Println("5 + 3 =", hasil)

    bagi, modulo := pembagian(10, 3)
    fmt.Println("10 / 3 = ", bagi, "sisa", modulo)
}
```

**Fungsi dengan Error Handling:**

```go
package main

import "fmt"

// Fungsi yang mengembalikan error
func bagi(a int, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("error: pembagi tidak boleh 0")
    }
    return a / b, nil
}

func main() {
    hasil, err := bagi(10, 2)
    if err != nil {
        fmt.Println("Ada error:", err)
    } else {
        fmt.Println("Hasil:", hasil)
    }

    // Coba dengan pembagi 0
    hasil, err = bagi(10, 0)
    if err != nil {
        fmt.Println("Ada error:", err)  // Output: Ada error: error: pembagi tidak boleh 0
    }
}
```

## 1.9 Latihan Modul 1

**Latihan 1: Program Kalkulator Sederhana**

Buat program yang menggunakan fungsi untuk:

- Penjumlahan
- Pengurangan
- Perkalian
- Pembagian

Contoh output:

```
5 + 3 = 8
5 - 3 = 2
5 * 3 = 15
5 / 3 = 1
```

**Latihan 2: Program Cek Bilangan Ganjil/Genap**

Buat fungsi `cekGanjilGenap(n int) string` yang mengembalikan "Ganjil" atau "Genap".

**Latihan 3: Program Grade Nilai**

Buat fungsi yang menerima nilai numerik (0-100) dan mengembalikan grade (A, B, C, D, E).
