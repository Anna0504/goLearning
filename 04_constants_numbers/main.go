// **************** NUMBERS ****************
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var e int
	fmt.Println(e)
	var f float32 = 58888888889999
	g := 999999999999999999999i
	h := 5.78
	var i float64 = 5.7
	fmt.Println(f)
	fmt.Printf("%T", f)
	fmt.Println(reflect.TypeOf(g))
	fmt.Println(reflect.TypeOf(h))
	fmt.Println(reflect.TypeOf(i))
	// Using 8-bit unsigned int
	var X uint8 = 225
	fmt.Println(X, X-3)
	// Using 16-bit signed int
	var Y int16 = 32767
	fmt.Println(Y+2, Y-2)
	// 225 222
	// -32767 32765
}
