package main

import "fmt"


type StudentRecord struct{
	rollnumber int
	name string
	address string

}
func (s *StudentRecord) AddStudent(rollno int, name string,address string) *StudentRecord{
	s.rollnumber=rollno
	s.name=name
	s.address=address
	return s
}

func main(){
student1:=new(StudentRecord)
student1.AddStudent(1,"Ali","XYZ Town")
fmt.Println(&student1,student1)

}