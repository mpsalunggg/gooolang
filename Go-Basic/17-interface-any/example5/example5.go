package main

import (
	"errors"
	"fmt"
)

type Pembayaran interface {
	KurangiSaldo(jumlah float64) error
	CekSaldo() float64
}

type Tabungan struct {
	Saldo float64
}

type Barang struct {
	ID    int
	Nama  string
	Harga float64
}

func (t *Tabungan) CekSaldo() float64 {
	return t.Saldo
}

func (t *Tabungan) KurangiSaldo(jumlah float64) error {
	if jumlah > t.Saldo {
		return errors.New("saldo tidak mencukupi")
	}
	t.Saldo -= jumlah
	return nil
}

func CariBarangByID(barangList []Barang, id int) (*Barang, error) {
	for _, barang := range barangList {
		if barang.ID == id {
			return &barang, nil
		}
	}
	return nil, errors.New("barang dengan ID tersebut tidak ditemukan")
}

func BeliBarang(saldo Pembayaran, barang *Barang) error {
	fmt.Printf("Ingin membeli: %s (Harga: %.2f)\n", barang.Nama, barang.Harga)
	if err := saldo.KurangiSaldo(barang.Harga); err != nil {
		return err
	}
	
	fmt.Println("Pembelian berhasil!")
	return nil
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
		if err := BeliBarang(tabungan, barang); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("Saldo sekarang: %.2f\n", tabungan.Saldo)
		}
	}
}
