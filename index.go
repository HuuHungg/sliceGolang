package main

import "fmt"

type Student struct {
	name string
}

// Define method
func (s Student) getName() string {
	return s.name
}

// 1.Value Reciever
func (s Student) changeName() {
	fmt.Printf("p2 = %p", &s)
	s.name = "Robin"
}

// 2.Pointer receiver
func (s *Student) changeName2() {
	fmt.Printf("p2 = %p", s)
	fmt.Println()

	s.name = "Robin2"
}

// 3. Non-struct

type String string

func (s String) append(str string) string {
	return fmt.Sprintf("%s%s", s, str)
}

func main() {
	student := &Student{name: "Ryan"}

	fmt.Printf("p1 = %p", student)
	fmt.Println()

	student.changeName2()
	fmt.Println(student.name)

	s1 := String("a")
	newStr := s1.append("c")

	fmt.Println(newStr)

}
