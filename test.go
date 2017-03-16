package main

import (
	"fmt"
	"reflect"
)

type testStruct struct {
	stuff string
	other int
}

func main() {
	fmt.Println("Hello World")
	m := testStruct{}
	fmt.Println(m.stuff == "")
	fmt.Println(&(m.other))

}

func array(obj interface{}) []interface{} {
	m := make([]interface{}, 0)
	m = append(m, obj)
	m = append(m, obj)
	return m
}

func makeObj(obj interface{}) interface{} {
	m := reflect.Zero(reflect.TypeOf(obj))
	fmt.Println(m)
	return m
}
