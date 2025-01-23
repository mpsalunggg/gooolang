package main

import (
	"errors"
	"fmt"
)

type Tabungan struct {
	Saldo float64
}

type Barang struct {
	ID    int
	Nama  string
	Harga float64
}

func CariBarangByID(barangList []Barang, id int) (*Barang, error) {
	for _, barang := range barangList {
		if barang.ID == id {
			return &barang, nil
		}
	}
	return nil, errors.New("barang dengan ID tersebut tidak ditemukan")
}

func main() {
	tabungan := &Tabungan{Saldo: 100000}

	toko := []Barang{
		{ID: 1, Nama: "Laptop", Harga: 150000},
		{ID: 2, Nama: "Mouse", Harga: 50000},
		{ID: 3, Nama: "Keyboard", Harga: 75000},
	}

	fmt.Println(tabungan.Saldo)
	fmt.Println(toko)
	fmt.Println("========================")
	fmt.Printf("Saldo awal: Rp.%.2f\n", tabungan.Saldo)

	for _, barang := range toko {
		fmt.Printf("ID: %d, Nama: %s, Harga: %.2f\n", barang.ID, barang.Nama, barang.Harga)
	}

	var id int
	fmt.Print("Masukkan ID barang yang ingin dibeli: ")
	fmt.Scan(&id)

	barang, err := CariBarangByID(toko, id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(barang)
	}
}
