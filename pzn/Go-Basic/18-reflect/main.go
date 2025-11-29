package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func main() {
	number := 23
	reflectValue := reflect.ValueOf(number)

	fmt.Println("type variable : ", reflectValue.Type())
	fmt.Println(reflectValue.Kind())
	fmt.Println(reflectValue.Int())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variabel :", reflectValue.Int())
	}

	// Access with interface casting
	fmt.Println("interface : ", reflectValue.Interface().(int))

	fmt.Println("===========================")
	// Access with object
	var s1 = &Student{Name: "wick", Age: 2}
	s1.getProperty()

	fmt.Println("current name : ", s1.Name)
	s1.SetName("Putra")

	reflectValue2 := reflect.ValueOf(s1)
	method := reflectValue2.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("putra"),
	})
	fmt.Println("updated name : ", s1.Name)
}

func (s *Student) getProperty() {
	var reflectValue = reflect.ValueOf(s)
	fmt.Println(reflectValue)
	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
		fmt.Println("elem : ", reflectValue)
	}

	var reflectType = reflectValue.Type()
	fmt.Println("type = ", reflectType)
	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("nama      :", reflectType.Field(i).Name)
		fmt.Println("tipe data :", reflectType.Field(i).Type)
		fmt.Println("nilai     :", reflectValue.Field(i).Interface())
		fmt.Println("")
	}
}

func (s *Student) SetName(name string) {
	s.Name = name
}
