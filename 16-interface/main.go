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

type Describer interface {
	Describe() string
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) Describe() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

func PrintDescription(d Describer) {
	fmt.Println(d.Describe())
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

	fmt.Println("================================")
	// Example 2
	p := &Person{Name: "Putra", Age: 23}
	PrintDescription(p)

	p.Name = "Test"
	p.Age = 24

	PrintDescription(p)
}
