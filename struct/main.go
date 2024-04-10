package main
import (
	"fmt"
	// "reflect"
)

type Person struct {
	name string
	age int
	married bool
	salary float32
}

func main(){

	

	//create an instance
	deshan := Person {
		name: "deshan",
		age: 25,
		married: false,
		salary: 24500.00,
	}

	// fmt.Print(reflect.TypeOf(deshan))		//main.person
	// fmt.Println(deshan.age)
	personDetails(deshan)
	

}

/* 

A struct (short for structure) is used to create a collection of members of different data types, into a single variable.

While arrays are used to store multiple values of the same data type into a single variable, structs are used to store multiple values of different data types into a single variable.

A struct can be useful for grouping data together to create records.

to declare a struct go use type and struct keywords

*/

// pass struct as function argument

func personDetails(deshan Person){
	fmt.Println("Name: ",deshan.name);
	fmt.Println("Age: ",deshan.age);
	fmt.Println("Married: ",deshan.married);
	fmt.Println("Salary: ",deshan.salary);

	if deshan.married{
		fmt.Println("deshan is  married")
	} else{
		fmt.Println("deshan is not married married")
	}
		

}