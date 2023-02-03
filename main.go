package main

import "fmt"

type School struct{
	name string
	address string
	students_num int
	rating int
}

func (s *School) changeStudNum(newNum int) {
    s.students_num = newNum
}
func (s *School) changeRating(newNum int) {
    s.rating = newNum
}

func main(){
	school1 := School{"Zerde","Markova 47B",250, 5}
	school2 := School{"Leader","Orbita 4",360, 4}

	school1.changeRating(10)
	school2.changeStudNum(843)

	fmt.Printf("School 1 name: %s\nAddress: %s\nStudents number: %d\nRating: %d\n", school1.name, school1.address, school1.students_num, school1.rating)
	fmt.Println("==========================================================================")
	fmt.Println("==========================================================================")
	fmt.Printf("School 2 name: %s\nAddress: %s\nStudents number: %d\nRating: %d\n", school2.name, school2.address, school2.students_num, school2.rating)
}	