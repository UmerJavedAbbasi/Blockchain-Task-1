package main

import (
	"fmt"
	"strings"
)

type Student struct{
	rollno int
	name string
	address string
}
func AddStudent(rollno int,name string,address string)(*Student){
	s:=new(Student)
	s.rollno=rollno
	s.name=name
	s.address=address
	return s
}
type Student_list struct{
	list []*Student
}
func (ls *Student_list)CreateStudent(rollno int,name string,address string)(*Student){
	student:=new(Student)
	student=AddStudent(rollno,name,address)
	ls.list=append(ls.list,student)
	return student
}

func (students *Student_list)Print(){
	fmt.Println("Students Data")
	
	for i:=0;i<len(students.list);i++{
		fmt.Printf("%s List %d %s \n",strings.Repeat("=",25),i,strings.Repeat("=",25))
		fmt.Printf("student rollno\t%d\n",students.list[i].rollno)
		fmt.Printf("student name\t%s\n",students.list[i].name)
		fmt.Printf("student address\t%s\n",students.list[i].address)
	}
}
func main(){
student_list:=new(Student_list)
student_list.CreateStudent(1,"Ali","XYZ Town")
student_list.CreateStudent(2,"Abdullah","ABC Town")
student_list.CreateStudent(3,"Awais","--- Town")

fmt.Println(student_list.list[2])

student_list.Print()
}