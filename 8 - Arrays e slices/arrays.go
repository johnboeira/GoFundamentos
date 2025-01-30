package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"P1", "p2", "p3", "p4"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3}
	fmt.Println(array3)

	//Slice
	slice := []int{1, 2, 3}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))

	//add
	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

}
