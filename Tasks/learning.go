package main

import "fmt"
func Learning_Functions(){
	fmt.Println("Executing Function")
}
func Print_MyBiodata(f_name string,l_name string,age int){
	fmt.Println("First Name: ",f_name,"\nLast Name: ",l_name,"\nAge: ",age)
}

func sum(x int,y int)(int){
	return x+y
}
func main(){
/*	fmt.Print("Hello WOrld")

	//Variable Declation
	//1
	var name="Umer"
	fmt.Print(name)
	//2
	var name1 string="Umer"
	fmt.Print(name1)
	//3

	name2:="Javed"
	fmt.Print(name2)
	//VAriable Declaration

	var a string
	var b int
	var c bool
	var d float32
	fmt.Print("\n************Break************\n")
	fmt.Print(a)
	fmt.Print(b)
	fmt.Print(c)
	fmt.Print(d)

	fmt.Print("\n************Break************\n")
*/
	//Multiple Variable Declaration 
	//1
/*	var a,b,c,d=1,2,3,4
	var fname,lname="Umer","Javed"
	fmt.Print(a,b,c,d)
	var result=fname+lname
	fmt.Print("\n",result)
*/
/*	//2
	var(
		a int
		b=1
		c="Umer"
	)
	fmt.Print("\n",a,b,c)
	//Const

	const readable ="can't be changed"
	fmt.Print("\n",readable,"\n")
	//Multiple Const Declaration

	const(
		A=1
		B int =2
		C string="Learning"
	)
	fmt.Print(A,B,C,"\n")

	fmt.Println(A,B,C)
	fmt.Println("Test")

	fmt.Printf("Test1")
	fmt.Printf("Test2")

	i:="Umer"
	j:=22	

	fmt.Printf("Name is:%v and its type is:%T \n",i,i)
	fmt.Printf("Age is:%v and its type is:%T \n",j,j)

	//Formatting Verbs

	fmt.Printf("Name Default Format is:%v \t Name GO Format is:%#v",i,i)
	fmt.Printf("\n100%%")

	var number=15
	fmt.Printf("\n")
	fmt.Printf("%d\n",number)
	fmt.Printf("%b\n",number)
	fmt.Printf("%+d\n",number)
	fmt.Printf("%o\n",number)
	fmt.Printf("%O\n",number)
	fmt.Printf("%x\n",number)
	fmt.Printf("%X\n",number)

	var txt="Hello"
	fmt.Printf("%s\n",txt)
	fmt.Printf("%q\n",txt)

	fmt.Printf("%8s\n",txt)
	fmt.Printf("%-8s\n",txt)

	var decision=true
	fmt.Printf("%t\n",decision)
	decision=false
	fmt.Printf("%t\n",decision)
*/
	//int,int8,int16,int32,int64
/*	var x int=15
	var y int=-15
	fmt.Printf("\nValue of X is:%v and it's Type is:%T",x,x)
	fmt.Printf("\nValue of X is:%v and it's Type is:%T",y,y)

	var positive uint16=15777
	fmt.Printf("\nValue of X is:%v and it's Type is:%T",positive,positive)

	//Floating Numbers
	var decimal_num1 float32=2.2
	var decimal_num2 float64=2.2222

	fmt.Printf("\nValue:%v \t Type:%T",decimal_num1,decimal_num1)
	fmt.Printf("\nValue:%v \t Type:%T",decimal_num2,decimal_num2)
*/
	//Arrays

	var array1=[4] int {1,2,3,4}
	fmt.Println(array1)

	array2:=[6]int{3,4,5,6,7,8}
	fmt.Println(array2)

	array3 :=[6] int{3:11,4:12}
	fmt.Println(array3)

	fmt.Println(len(array3))

	//GO Slices
	
	myslice:=[4]int{1,2,3}
	fmt.Println(myslice)
	fmt.Println(cap(myslice))
	fmt.Println(len(myslice))

	//Slicing on array
	var myarray=[8]int{1,2,3,4}
	slice:=myarray[3:len(myarray)]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	//Slicing Using Make function

	slice1:=make([]int,5,10)
	fmt.Println(slice1,"\t",len(slice1),"\t",cap(slice1))

	//Slice Modification
	slice1[0]=2
	fmt.Println(slice1,"\t",len(slice1),"\t",cap(slice1))

	slice1=append(slice1,6,7,8)
	fmt.Println(slice1,"\t",len(slice1),"\t",cap(slice1))

	slice2:=append(slice,slice1...)
	fmt.Println(slice2,"\t",len(slice2),"\t",cap(slice2))
	
	//Conditional Structure
	if 2==2{
		fmt.Println("Equal")
	}else
	{
		fmt.Println("Not Equal")
	}
	//Single and Multi Switch is same as C++ switch

	//For loop same as C++

	fruits:=[5]string{"Apple","Mango","Banana","Pine Apple","Melon"}

	for index,value:= range fruits{
		fmt.Println(index,"\t",value)
	}

	//Functions

	Learning_Functions()
	Print_MyBiodata("Umer","Javed",22)
	fmt.Println(sum(2,3))


}
