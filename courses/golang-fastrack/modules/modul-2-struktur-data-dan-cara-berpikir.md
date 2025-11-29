# MODUL 2: STRUKTUR DATA & CARA BERPIKIR DI GO

## 2.1 Array

Array adalah koleksi data dengan ukuran tetap.

```go
package main

import "fmt"

func main() {
    // Deklarasi array dengan ukuran 3
    var angka [3]int
    angka[0] = 10
    angka[1] = 20
    angka[2] = 30
    fmt.Println(angka)  // [10 20 30]

    // Inisialisasi array
    buah := [3]string{"Apel", "Mangga", "Pisang"}
    fmt.Println(buah)

    // Akses elemen
    fmt.Println(buah[0])  // Apel
    fmt.Println(buah[1])  // Mangga

    // Looping array
    for i, val := range buah {
        fmt.Printf("Index %d: %s\n", i, val)
    }
}
```

## 2.2 Slice

Slice adalah wrapper dinamis di atas array. Bisa bertambah/berkurang ukurannya.

```go
package main

import "fmt"

func main() {
    // Buat slice
    var warna []string
    warna = append(warna, "Merah")
    warna = append(warna, "Hijau")
    warna = append(warna, "Biru")
    fmt.Println(warna) // Output: [Merah Hijau Biru]

    // Slice dengan make
    angka := make([]int, 5)  // slice dengan kapasitas 5
    angka[0] = 10
    angka[1] = 20
    fmt.Println(angka)  // [10 20 0 0 0]

    // Slice dari array
    arr := [5]int{1, 2, 3, 4, 5}
    slice1 := arr[1:4]  // index 1 sampai 3
    fmt.Println(slice1) // [2 3 4]

    // Length dan capacity
    slice2 := make([]int, 3, 5)
    fmt.Println(len(slice2))  // 3 (jumlah elemen saat ini)
    fmt.Println(cap(slice2))  // 5 (kapasitas maksimal)
}
```

## 2.3 Map

Map adalah struktur key-value, mirip dictionary atau object.

```go
package main

import "fmt"

func main() {
    // Deklarasi map
    var stok map[string]int      // deklarasi, stok == nil (belum bisa dipakai)
    stok = make(map[string]int)  // alokasikan map kosong, siap diisi data

    // Tambah data
    stok["Apel"] = 50
    stok["Mangga"] = 30
    stok["Pisang"] = 20

    fmt.Println(stok)  // map[Apel:50 Mangga:30 Pisang:20]
    fmt.Println(stok["Apel"])  // 50

    // Map literal
    user := map[string]string{
        "nama": "Ahmad",
        "email": "ahmad@mail.com",
        "kota": "Jakarta",
    }
    fmt.Println(user["nama"])

    // Cek keberadaan key
    val, ok := stok["Jeruk"]
    if ok {
        fmt.Println("Jeruk ada dengan stok:", val)
    } else {
        fmt.Println("Jeruk tidak ada")
    }

    // Looping map
    // resultnya:
    // Mangga: 30
    // Apel: 50
    // Pisang: 20
    for key, val := range stok {
        fmt.Printf("%s: %d\n", key, val)
    }

    // Delete
    delete(stok, "Apel")
    fmt.Println(stok)  // Apel sudah dihapus
}
```

## 2.4 Struct

Struct digunakan untuk merepresentasikan object/entity dengan beberapa field.

```go
package main

import "fmt"

// Definisi struct
type User struct {
    ID    int
    Nama  string
    Email string
    Umur  int
}

type Task struct {
    ID        int
    Title     string
    Completed bool
}

func main() {
    // Buat instance struct
    user1 := User{
        ID:    1,
        Nama:  "Ahmad",
        Email: "ahmad@mail.com",
        Umur:  25,
    }
    fmt.Println(user1)
    fmt.Println(user1.Nama)  // Ahmad

    // Struct zero value
    var user2 User
    fmt.Println(user2)  // {0  0} - semua field kosong

    // Slice of struct
    users := []User{
        {ID: 1, Nama: "Ahmad", Email: "ahmad@mail.com", Umur: 25},
        {ID: 2, Nama: "Budi", Email: "budi@mail.com", Umur: 30},
    }

    for _, user := range users {
        fmt.Printf("ID: %d, Nama: %s\n", user.ID, user.Nama)
    }

    // Task
    task := Task{
        ID:        1,
        Title:     "Belajar Go",
        Completed: false,
    }
    fmt.Println(task)
}
```

## 2.5 Pointer

Pointer menyimpan alamat memori dari suatu variabel. Berguna untuk mengubah data dari dalam fungsi lain atau efisiensi memory.

```go
package main

import "fmt"

type User struct {
    Nama string
    Umur int
}

// Fungsi dengan pointer receiver
func ubahNama(user *User, namaBaru string) {
    user.Nama = namaBaru  // mengubah nama langsung
}

func main() {
    user := User{Nama: "Ahmad", Umur: 25}
    fmt.Println(user)  // {Ahmad 25}

    // Buat pointer ke user
    pointerUser := &user
    fmt.Println(pointerUser)      // &{Ahmad 25}
    fmt.Println(pointerUser.Nama) // Ahmad

    // Ubah via pointer
    pointerUser.Nama = "Budi"
    fmt.Println(user)  // {Budi 25} - user asli juga berubah

    // Atau gunakan fungsi
    ubahNama(&user, "Chandra")
    fmt.Println(user)  // {Chandra 25}
}
```

## 2.6 Error Handling

Go menggunakan pola error yang sederhana: mengembalikan error sebagai return value terakhir.

```go
package main

import (
    "fmt"
    "errors"
)

// Fungsi yang mengembalikan error
func validasiNama(nama string) error {
    if nama == "" {
        return errors.New("nama tidak boleh kosong")
    }
    if len(nama) < 3 {
        return errors.New("nama minimal 3 karakter")
    }
    return nil
}

func konversiKePersen(nilai float64) (float64, error) {
    if nilai < 0 || nilai > 100 {
        return 0, fmt.Errorf("nilai harus antara 0-100, mendapat %.2f", nilai)
    }
    return nilai, nil
}

func main() {
    // Contoh 1
    err := validasiNama("Al")
    if err != nil {
        fmt.Println("Error:", err)
        // nama minimal 3 karakter
    }

    err = validasiNama("Ahmad")
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Nama valid!")
    }

    // Contoh 2
    persen, err := konversiKePersen(150)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("Persen: %.2f\n", persen)
    }
}
```

## 2.7 Latihan Modul 2

**Latihan 1: Manajemen User Sederhana**

Buat program yang:
- Memiliki struct `User` dengan field: ID, Nama, Email
- Slice berisi daftar user
- Fungsi untuk menambah user baru
- Fungsi untuk mencari user berdasarkan ID
- Fungsi untuk menampilkan semua user

**Latihan 2: Inventory Produk**

Buat program yang:
- Memiliki struct `Product` dengan field: ID, Name, Stok
- Map untuk menyimpan produk
- Fungsi untuk menambah stok
- Fungsi untuk mengurangi stok (dengan error checking)
- Fungsi untuk menampilkan semua produk
