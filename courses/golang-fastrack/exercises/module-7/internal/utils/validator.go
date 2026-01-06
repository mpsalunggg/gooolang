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
