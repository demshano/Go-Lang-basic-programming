package main

import (
	"fmt"
	
	
)

func main() {
	//  fmt.Print("Go Lang\nprogramming")
	
	// variables()
	// assignmentOp()
	// arrays()
	// forLoops()
	// whileLoops()
	numbers := []int{1, 2, 3, 4, 5}
	minMax(numbers)
	var num  = []int{23,45,7}
	fmt.Print(minMax(num))
}

func variables() {
	var firstName string = "deshan"
	var lastName string = "manodya"
	//implicitly type inference, no need to declare data type of variable
	fullname := firstName + " " + lastName //concatenation operator in Go is '+'
	age := 21                              //integer
	//boolean or bool

	fmt.Print(fullname, "\n", "age is", age) //boolean or bool
}
func assignmentOp() {
	number1 := 12
	number2 := 23

	// or we can use var number1 = 12; and so on...

	fmt.Print(number1 + number2)

}

func minMax(arr []int) (int, int) {
	minVal := arr[0]
	maxVal := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > maxVal {
			maxVal = arr[i]
		} else if arr[i] < minVal {
			minVal = arr[i]
		}
	}
	return minVal, maxVal
}




func arrays() {
	var names = [3]string{"John", "Jane", "Doe"} //arrays are static and fixed size unlike
	prices := [3]int{10, 20, 30}

	fmt.Println(prices[0])
	fmt.Println(prices[2])
	fmt.Print(names)

	prices[0] = 100
	fmt.Println(prices)
	
}

func forLoops() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func whileLoops(){
	i := 0
	for(i < 5){
		fmt.Println(i);
		i++;
	}
}
