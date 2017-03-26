package crabgo

import (
	"fmt"
	"reflect"
)

func Dump(i interface{}) {
	fmt.Println("-----------------begin--------------------")
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(i)
	fmt.Println("-----------------end----------------------")
}
