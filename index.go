package main

import (
	"fmt"
)

func main() {
	// Tạo một map với key là string và value là int
	ages := make(map[string]string)

	// Thêm các phần tử vào map
	ages["Alice"] = "ITを勉強してる"
	ages["Bob"] = "難しいです"
	ages["Charlie"] = "頑張りましょう"

	// Lấy giá trị của phần tử có key là "Alice"
	aliceAge := ages["Alice"]
	fmt.Println("Alice is", aliceAge, "years old")

	// Kiểm tra xem một key có tồn tại trong map hay không
	if age, ok := ages["David"]; ok {
		fmt.Println("David is", age, "years old")
	} else {
		fmt.Println("David is not in the map")
	}

}
