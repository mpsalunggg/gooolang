package main

import "fmt"

func printData(data interface{}) {
	switch v := data.(type) {
	case int:
		fmt.Println("Data bertipe int:", v)
	case string:
		fmt.Println("Data bertipe string:", v)
	case bool:
		fmt.Println("Data bertipe bool:", v)
	default:
		fmt.Println("Data bertipe lain:", v)
	}
}

func main(){
	var data interface{}

	data = 21
	printData(data)

	data = "Halo, dunia!"
	printData(data)

}