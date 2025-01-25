package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func updateUser(user *User, name string, age int) {
	user.Name = name
	user.Age = age
	fmt.Println("Success change user : ", user.Name)
}

func checkType(data any){
	switch temp := data.(type) {
	case *User :
		fmt.Printf("Ini User %s , umurnya %d\n", temp.Name, temp.Age)
	case map[string]string:
		for key, value := range temp{
			fmt.Println(key, value)
		}
	}
}

func main() {
	user := &User{"mps", 12}
	fmt.Println(*user)

	updateUser(user, "Putra", 32)

	checkType(user)

	data2 := map[string]string{
		"Role": "Software engineer",
		"Focus": "Frontend engineer",
	}

	checkType(data2)
}
