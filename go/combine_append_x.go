package demo

import "fmt"

func CombineAppend() {
	slice := Int{1, 2, 3, 4}
	slice = append(slice, 5, 6, 7)

	fmt.Println(slice)
}
