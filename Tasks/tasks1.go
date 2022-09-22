package main

import "fmt"

type Person struct
{
	name string
	rolno string
	semester int	
}
func Show_Person_Info(obj Person){
	fmt.Println("Name is:",obj.name)
	fmt.Println("Rol No is:",obj.rolno)
	fmt.Println("Semester is:",obj.semester)

} 
func main(){
	fmt.Println("Hello World")

	var person1 Person

	person1.name="Umer"
	person1.rolno="19I1650"
	person1.semester=7

	Show_Person_Info(person1)

}
