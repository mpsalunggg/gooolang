package main

import "fmt"

type Tabungan struct {
	Saldo float64
}

type Barang struct {
	ID    int
	Nama  string
	Harga float64
}

func main() {
	tabungan := &Tabungan{Saldo: 100000}

	toko := []Barang{
		{ID: 1, Nama: "Laptop", Harga: 150000},
		{ID: 2, Nama: "Mouse", Harga: 50000},
		{ID: 3, Nama: "Keyboard", Harga: 75000},
	}

	fmt.Println(*tabungan)
	fmt.Println(toko)
}
