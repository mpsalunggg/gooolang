package main

import (
	"fmt"
	"math"
)

type Hitung interface {
	luas() float64
	keliling() float64
}

type Lingkaran struct {
	diameter float64
}

func (l Lingkaran) jariJari() float64 {
	return l.diameter / 2
}

func (l Lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l Lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

type Persegi struct {
	sisi float64
}

func (p Persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p Persegi) keliling() float64 {
	return p.sisi * 4
}

func main() {
	// INTERFACE

	var bangunDatar Hitung

	bangunDatar = Persegi{10.0}
	fmt.Println("===== persegi")
	fmt.Println("luas      :", bangunDatar.luas())
	fmt.Println("keliling  :", bangunDatar.keliling())

	bangunDatar = Lingkaran{14.0}
	fmt.Println("===== lingkaran")
	fmt.Println("luas      :", bangunDatar.luas())
	fmt.Println("keliling  :", bangunDatar.keliling())
	fmt.Println("jari-jari :", bangunDatar.(Lingkaran).jariJari())
}
