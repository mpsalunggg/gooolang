package main

import (
	"fmt"
	"reflect"
)

func main(){
	number := 23
	reflectValue := reflect.ValueOf(number)

	fmt.Println("type variable : ", reflectValue.Type())
	fmt.Println(reflectValue.Kind())
	fmt.Println(reflectValue.Int)

	if reflectValue.Kind() == reflect.Int {
        fmt.Println("nilai variabel :", reflectValue.Int())
    }
}