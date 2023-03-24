package main

import (
	"fmt"
)

func addItem(item int, list ...int) {
	list = append(list, item, 11, 22, 33, 44, 55, 66)
	fmt.Println(list)
}

func change(list ...int) {
	list[2] = 999
}

func main() {
	addItem(1, 200, 300, 400, 600)

	var list = []int{100, 200, 300}
	addItem(100, list...)

	change(list...)
	fmt.Println(list)

}
