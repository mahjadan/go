package main

import (
	"fmt"
	"reflect"
)

type Data struct {
	Msg string
}
type New struct {
	Msg string
}

func main() {

	// var i int
	// i = 10
	// ref(i)
	// ref(&i, 20)

	d := Data{Msg: "Hello"}
	fmt.Println("value before is :  ", d)

	change(&d, Data{Msg: "new Hello"})
	fmt.Println("value after is :  ", d)
	// m := make(map[string]int)
	// s := make([]string, 7)

	// change(nil, m)
}

func change(v interface{}, newVal interface{}) {
	vRefVal := reflect.ValueOf(v)
	if vRefVal.Kind() != reflect.Ptr {
		panic(fmt.Sprintf("the value passed into the interface is not a pointer, got: %v", vRefVal.Kind()))
	}
	newValRefVal := reflect.ValueOf(newVal)
	fmt.Println("kind of interface: ", vRefVal.Kind())
	fmt.Println("kind of newVal: ", newValRefVal.Kind())
	if newValRefVal.Kind() == reflect.Invalid {
		panic("invalid kind pass in newVal or nil.")
	}

	vObjValueType := vRefVal.Elem().Type()
	newValType := reflect.TypeOf(newVal)
	fmt.Println("type of the value referenced in the interface: ", vObjValueType)
	fmt.Println("type of newVal: ", newValType)
	if vObjValueType != newValType {
		panic(fmt.Sprintf("the type of the value passed does not match with the type of interface being passed, %v != %v", newValType, vObjValueType))
	}
	fmt.Println("------changing the value------")
	// reflect.ValueOf(v).Elem().Set(reflect.ValueOf(newVal))
	vRefVal.Elem().Set(newValRefVal)
}
