package main

import "fmt"

// type triangle struct {
// 	width, height int32
// }

// func (s triangle) setArea(width int32, height int32) {
// 	s.width = width
// 	s.height = height
// }

// func (s triangle) getArea() (int32, int32) {
// 	return s.width, s.height
// }

// func (s *triangle) setAreaPtr(width int32, height int32) {
// 	s.width = width
// 	s.height = height
// }

// func main() {
// 	tri1 := new(triangle)

// 	x, y := tri1.getArea()

// 	fmt.Printf("-- %v, %d/%d\n", tri1, x, y)

// 	tri1.setArea(10, 10)
// 	x, y = tri1.getArea()

// 	fmt.Printf("-- %v, %d/%d\n", tri1, x, y)

// 	tri1.setAreaPtr(11, 11)
// 	x, y = tri1.getArea()

// 	fmt.Printf("-- %v, %d/%d\n", tri1, x, y)

// }

type Person struct {
	name string
	age  int
}

func (p *Person) greeting() {
	fmt.Println("Hello~")
}

type Student struct {
	Person // 필드명 없이 타입만 선언하면 포함(Is-a) 관계가 됨
	school string
	grade  int
}

func main() {
	// s := Student{name: "PSY", age: 29, school: "school", grade: 1} // 컴파일 오류.
	// s := Student{"PSY", 29, "school", 1} // 컴파일 오류.
	var s = Student{Person{"PSY", 29}, "school", 1}

	s.Person.greeting() // Hello~
	s.greeting()        // Hello~

	fmt.Printf("---\n")
}



{}