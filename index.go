package main

import (
	"fmt"
)

// Struct

type StudentInfo struct {
	age   int
	class string
}

type Student struct {
	id    int
	name  string
	email string
	info  StudentInfo
}

type Teacher struct {
	id   int
	name string
}

type Childrent struct {
	id   int
	name string
}

func main() {
	st1 := Student{
		id:    123,
		name:  "Huu Hung Nguyen",
		email: "nghuuhung0810@gmail.com",
	}

	fmt.Println(st1)

	fmt.Println(st1.id)
	fmt.Println(st1.name)
	fmt.Println(st1.email)

	st2 := Student{id: 456, name: "Robin"}

	fmt.Println(st2)
	fmt.Println(st2.id)
	fmt.Println(st2.name)

	var st3 Teacher = struct {
		id   int
		name string
	}{
		id:   777,
		name: "Bill",
	}

	fmt.Println(st3)

	// anonymouse struct
	var anonymous = struct {
		age   int
		email string
	}{
		age:   21,
		email: "nghuuhung0810@gmail.com",
	}

	fmt.Println(anonymous)

	// pointer tro toi struct

	pointer := &Student{
		id:   000,
		name: "Pointer",
	}

	fmt.Println(&pointer)
	fmt.Println((*pointer).id)
	fmt.Println((*pointer).name)

	// anonymous fields

	type NoName struct {
		string
		int
	}
	n := NoName{"Huu Hung Nguyen", 9999}
	fmt.Println(n)
	fmt.Println("=============================")
	// Struct long struct ---
	student := Student{
		id:    2,
		name:  "Guenfufun",
		email: "nghuuhung0810@gmail.com",
		info: StudentInfo{
			age:   21,
			class: "NB1102",
		},
	}

	fmt.Println(student)

	// Compare 2 struct
	type struct1 struct {
		id   int
		name string
	}

	s1 := struct1{1, "A"}
	s2 := struct1{1, "A"}

	if s1 == s2 {
		fmt.Println("s1 == s2")
	} else {
		fmt.Println("s1 != s2")
	}

	var chidrent Childrent

	chidrent.name = "ryan"
	fmt.Println(chidrent)

}
