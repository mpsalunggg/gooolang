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

func checkType(data any) {
	switch temp := data.(type) {
	case *User:
		fmt.Printf("Ini User %s , umurnya %d\n", temp.Name, temp.Age)
	case map[string]string:
		for key, value := range temp {
			fmt.Println(key, value)
		}
	case []int:
		res := 0
		for _, val := range temp {
			res += val
		}
		fmt.Println(res)
	}
}

func main() {
	user := &User{"mps", 12}
	fmt.Println(*user)

	updateUser(user, "Putra", 32)

	checkType(user)

	data2 := map[string]string{
		"Role":  "Software engineer",
		"Focus": "Frontend engineer",
	}

	checkType(data2)

	data3 := []int{4, 2, 1, 3}

	checkType(data3)

	data4 := &User{Name: "test", Age: 20}
	if d, ok := any(data4).(*User); ok {
		fmt.Printf("Successfully casted: Name=%s, Age=%d\n", d.Name, d.Age)
	} else {
		fmt.Println("Failed to cast data to *User")
	}
}
