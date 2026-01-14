package main

import (
	"fmt"
	"reflect"
)

type Greeting struct{}

func (g Greeting) SayHello(name string) {
	fmt.Printf("Helloooo gesssss, my name is %s\n", name)
}

func (g Greeting) RepeatHi(count int, message string) {
	for i := 0; i < count; i++ {
		fmt.Println("Hellooo " + message)
	}
}

func main() {
	greeting := Greeting{}

	val := reflect.ValueOf(greeting)
	typ := reflect.TypeOf(greeting)

	method := val.MethodByName("SayHello")
	if method.IsValid() {
		method.Call([]reflect.Value{
			reflect.ValueOf("mps"),
		})
	}

	methodRepeat := val.MethodByName("RepeatHi")
	if methodRepeat.IsValid() {
		methodRepeat.Call([]reflect.Value{
			reflect.ValueOf(4),
			reflect.ValueOf("brohh"),
		})
	}

	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)
		fmt.Printf("Method %d: %s\n", i+1, method.Name)
	}   
}
