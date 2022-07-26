package main

import (
	"errors"
	"fmt"
	"reflect"
)

type stubMapping map[string]interface{}

var StrubStorage = stubMapping{}

func setup() {

	StrubStorage = map[string]interface{}{
		"longRunningTest1": longRunningTest1,
		"longRunningTest2": longRunningTest2,
	}
}
func main() {
	setup()

	/*
		Check for the enable of the tests.
		For each enable status, call the function - Open/Close design principle (not checked)

		Keep a counter for enable tests

		If error is found, return and cancel using Context

	*/

	result1, _ := Call("longRunningTest1", "max", "mustermann")

	outputStr1 := result1.(string)
	fmt.Println(outputStr1)

	result2, _ := Call("longRunningTest2", 100)

	outputStr2 := result2.(int)
	fmt.Println(outputStr2)
}

func Call(funcName string, params ...interface{}) (result interface{}, err error) {
	f := reflect.ValueOf(StrubStorage[funcName])

	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is out of index")
		return
	}

	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}

	res := f.Call(in)
	result = res[0].Interface()
	return
}

func longRunningTest1(a, b string) string {
	return fmt.Sprintf("Done for %v, %v", a, b)
}

func longRunningTest2(val int) int {
	return 10 * val
}
