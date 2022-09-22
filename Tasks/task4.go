package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

type Student struct{
	rollno int
	name string
	address string
	subjects []string
}

func AddStudent(rollno int,name string,address string,subjects []string)(*Student){
	s:=new(Student)
	s.rollno=rollno
	s.name=name
	s.address=address
	s.subjects=subjects
	return s
}
type Student_list struct{
	list []*Student
}
func (ls *Student_list)CreateStudent(rollno int,name string,address string,subjects []string)(*Student){
	student:=new(Student)
	student=AddStudent(rollno,name,address,subjects)
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
		fmt.Printf("student subjects\t%s\n",students.list[i].subjects)
	}
}

func (students *Student_list)Hash()([][32]uint8){
	str:=""
	hashes:=make([][32]uint8,len(students.list))
	fmt.Println("Block Data\t\tBlock Hash")
	for i:=0;i<len(students.list);i++{
		str=(students.list[i].name)+strconv.Itoa(students.list[i].rollno)+(students.list[i].address)
		for subject:=0;subject<len(students.list[i].subjects);subject++{
			str=str+students.list[i].subjects[subject]
		}
		sum := sha256.Sum256([]byte(str))
		fmt.Printf("%s\t\t%x\n",str,sum)		
		hashes[i]=sum
	}
	return hashes

}
func main(){
student_list:=new(Student_list)
student_list.CreateStudent(1,"Ali","XYZ Town", []string{"SSD","PEH","NPT"})
student_list.CreateStudent(2,"Abdullah","ABC Town",[]string{"Blockchain","PDC","NTC"})
student_list.CreateStudent(3,"Awais","--- Town",[]string{"Information Security","OSINT","Malware Analysis"})



student_list.Print()
hashes:=student_list.Hash()
fmt.Printf("\nHashes of all bloacks are:%x",hashes)

//fmt.Println(student_list.list[0])
//fmt.Printf("%T\n",student_list.list[0])

//sum = sha256.Sum256(student_list.list[0])
//fmt.Printf("%x\n", sum) //hexadecimal


}