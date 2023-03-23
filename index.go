package main

import "fmt"

func main() {
	// Khai báo slice
	var slice []int
	fmt.Println(slice)

	// Khai báo và khởi tạo slice
	var slice1 = []int{1, 2, 3, 4}
	fmt.Println(slice1)

	// Tạo một slice tử 1 array
	var array = [7]int{1, 2, 3, 4, 5, 6, 7} // array1
	slice2 := array[1:3]
	fmt.Println(slice2)

	slice3 := array[:] // Lấy tấy cả giá trị
	fmt.Println(slice3)

	slice4 := array[2:] // lấy từ 2 đổ đi
	fmt.Println(slice4)

	slice5 := array[:3] // lấy từ 0 đến 3
	fmt.Println(slice5)

	// Tạo một slice từ một slice khác
	var slice6 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice7 := slice6
	fmt.Println(slice7)

	slice8 := slice6[1:]
	fmt.Println(slice8)

	var array1 = [5]int{1, 2, 3, 4, 5}
	slice9 := array1[:]
	slice9[0] = 777

	fmt.Println(slice9)
	fmt.Println(array1)

	// length và capacity cua slice
	countries := [...]string{"PHP", "Ruby", "JAVA", "C++", "C#", "NodeJS", "Golang"}
	slice10 := countries[3:4]
	fmt.Println(slice10)

	fmt.Println(len(slice10))
	fmt.Println(cap(slice10))

	// len : Số lượng phần tử của slice
	// cap : Số lượng phần tử của underlying array bắt đầu từ vị trí start khi mà slice được tạo
	// cap là từ vịt trí đấy đổ ngược lại xem có bao nhiêu phần tử

	// make, copy, append

	// make, copy, append
	// make khai báo slice cung cấp một giá trị cụ thể

	fmt.Println("===================")
	slice11 := make([]int, 3, 6)
	fmt.Println(slice11)
	fmt.Println(len(slice11))
	fmt.Println(cap(slice11))

	fmt.Println("===================")
	slice12 := make([]int, 2)
	fmt.Println(slice12)
	fmt.Println(len(slice12))
	fmt.Println(cap(slice12))

	fmt.Println("===================")
	// append thêm một giá trị và slice
	var slice13 []string
	slice13 = append(slice13, "毎日ITを勉強してる頑張りましょう")
	fmt.Println(slice13)

	fmt.Println("===================")
	// 3.Copy
	src := []string{"A", "B", "C", "D"}
	dest := make([]string, 3)

	number := copy(dest, src)
	fmt.Println(dest)
	fmt.Println(number)

	fmt.Println("===================")
	// delete item width index = 1
	src = append(src[:1], src[2:]...) // slice - slice = append(slice1, slice2)
	fmt.Println(src)
}
