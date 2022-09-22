package main

import "fmt"

type employee struct {
	name     string
	salary   int
	position string
}

type company struct {
	companyName string
	employees   []employee
}


func main(){

	emp1:=employee{"Amir",8000,"Full-Stack Developer"}
	emp2:=employee{"Ali",10000,"Mean Stack Developer"}
	emp3:=employee{"Abdullah",12000,"Mern Stack Developer"}

	fmt.Println(emp1)
	fmt.Println(emp2)
	fmt.Println(emp3)

	emplys:=[]employee{emp1,emp2,emp3}
	fmt.Println(emplys)

	cmpny:=company{"Software House",emplys}
	fmt.Println("Company Data is")
	fmt.Println(cmpny)

}
