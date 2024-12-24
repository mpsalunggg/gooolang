package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

func main() {
	names := []string{"Putra", "Satria"}
	callName("Haiii, ", names)
	addNum := add(2, 3)
	fmt.Println(addNum)
	randomNew()
}

func callName(message string, names []string) string {
	nameString := strings.Join(names, " ")
	fmt.Println(message, nameString)
	return nameString
}

// function with callback return value
func add(a int, b int) int {
	return a + b
}

// function rand.New()
func randomNew() {
	num := rand.New(rand.NewSource(time.Now().Unix()))
	fmt.Println(num)
}

// function multiple return
func calculate(d float64) (float64, float64) {
	var area = math.Pi * math.Pow(d/2, 2)
	var circumference = math.Pi * d
	return circumference, area
}

// function predefined value
func calculate2(d float64) (area float64, circumference float64) {
	area = math.Pi * math.Pow(d/2, 2)
	circumference = math.Pi * d
	return
}
