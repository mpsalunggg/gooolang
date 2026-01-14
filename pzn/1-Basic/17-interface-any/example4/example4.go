package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	var data interface{}

	user := &User{Name: "Putra", Age: 13}
	data = user

	castedUser, ok := data.(*User)
	if ok {
		fmt.Printf("Nama: %s, Umur: %d\n", castedUser.Name, castedUser.Age)
	} else {
		fmt.Println("Gagal melakukan casting")
	}

	var wrongData interface{} = "ini string, bukan User"
    if _, ok := wrongData.(*User); !ok {
        fmt.Println("Casting gagal karena tipe tidak sesuai.")
    }
}
