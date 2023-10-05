package task14

import (
	"fmt"
	"reflect"
)

func Task14() {
	var x interface{} = 42
	var y interface{} = "Vibe"
	var z interface{} = true
	var o interface{} = make(chan int)
	var a interface{} = make(chan string)
	checkType(x)
	checkType(y)
	checkType(z)
	checkType(o)
	checkType(a)

}

func checkType(i interface{}) {
	switch reflect.TypeOf(i).Kind() {
	case reflect.Int:
		fmt.Println(i, " Is integer")
	case reflect.String:
		fmt.Println(i, " Is string")
	case reflect.Bool:
		fmt.Println(i, " Is boolean")
	case reflect.Chan:
		channelType := reflect.TypeOf(i)
		fmt.Printf("Chan type: %v\n", channelType)
	default:
		fmt.Println("Unknown type")
	}
}
