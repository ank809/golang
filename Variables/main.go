package main 

import "fmt"


func main(){
	
	// A variable is just a name give to the memory location where the data is stored

	// Declaring variables 
	// 1. Using var keyword ---> var var_name type

	var name string="Ankita"
	var age int=19

	fmt.Println("Name is", name)
	fmt.Println("Age is", age)

	// 2. Variable declaration using var but without specifying type
	var marks = 12
	fmt.Println("Marks are", marks)


	//3. Using := expression assignment operator , var_name := value

	university := "LPU"
	fmt.Println("University is", university)

	//	What if no initial value is assigned to a variable? --> default value will be considered then
	var rollnum int
	fmt.Println("Roll number is",rollnum)

	length, breadth :=12, 13
	fmt.Println(length, breadth)

	var (
		isPresent =false
		lovesCoding=true
	)

	fmt.Println(isPresent, lovesCoding)

	// TO print the type of variabele use Printf with %T. 
	// We can't use %T with Println to print the type of variable as Println is used to print without
	// format specifiers

	fmt.Printf("Type of isPresent is %T", isPresent)

	a := 24234324.21432435345345
	fmt.Printf("%T \n", a)
	
	fmt.Printf("%.3f \n", a)
	
	fmt.Printf("%18f", a)

	/* Important points
	
	1. := operator can not be used outside a function.
	2. You need to initialize and declare in same line using := operator.
	3. You can initialize later on if you declare a variable using var keyword.
	4. Variable can be declared using var keyword outside the function.

	*/

	
}