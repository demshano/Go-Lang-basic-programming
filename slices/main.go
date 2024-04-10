package main
import (
	"fmt"
	// "reflect"
)

func main(){
	//slices are similar to array but more powerfull and flexible
	// Like arrays, slices are also used to store multiple values of the same type in a single variable.
	// However, unlike arrays, the length of a slice can grow and shrink as you see fit.

	/* 
	In Go, there are several ways to create a slice:

		-> Using the []datatype{values} format
		-> Create a slice from an array
		-> Using the make() function
 */

 //========================creating a slice with []datatype{values}

//  sliceName := []datatype{};

// sliceArray := [3]int{1,2,3};
// name := "deshan"
// fmt.Println(sliceArray)
// fmt.Print(reflect.TypeOf(sliceArray))
// fmt.Print(reflect.TypeOf(name))

// =========len() and cap()
// len() function - returns the length of the slice (the number of elements in the slice)
// cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)

// mySclice := []int{};
// fmt.Print(len(mySclice),cap(mySclice)); //output -> 0 0

// mySlice2 := []int{1,2,3,4,5};
// fmt.Println("Length : ",len(mySlice2), "\nCapacity :", cap(mySlice2)); // output -> Length :  5 Capacity :   5

//========================creating slice from an array

// var myArray = [length]dataType{elements};
// mySclice := myArray[start:end]  a slice made from the array

// var myArray = [5]int{1,2,3,4,5};
// mySlice := myArray[1:3];   // output -> {2 3}
// fmt.Println("My Slice is : ",mySlice);

//===============================craeting a slice with make function

// mySlice := make([]type, length, capacity)
// NOTED: if the capacity  isn't given then it will be set to length

// var mySlice = make([]int, 5, 10);
// fmt.Printf("The type of MySlice is %T\n",mySlice);                    //Outputs The type of MySlice is []int
// fmt.Println("Length of MySlice is : ",len(mySlice));                // Outputs Length of MySlice is : 5
// fmt.Println("Capacity of MySlice is : ",cap(mySlice));              // Outputs Capacity of MySlice is : 10

// =======================================modify the slice==========================

// =========================Access Elements of a Slice

//accesing the elements same like arrays in go

// mySlice := []int{12,34,45,56,67,89};

// var firstElement int =  mySlice[0];                      // Accessing The First Element Of The Slice

// fmt.Print(firstElement);

// //==========change the elements of a slice

// //changing the elements same as array

// mySlice[0] = 23;

// fmt.Print(mySlice)		the first index element has been changed

//=======================append the elements to a slice

// sliceName = append(sliceName, element1, element2,...);

// var mySlice = []string{"deshan", "sahan"};

// mySlice = append(mySlice,  "Dasun");
// fmt.Println("The new slice is : ",mySlice);

//=============================append one slice to another slice

//newSlice := append(slice1, slice2...);
//Note: The '...' after slice2 is necessary when appending the elements of one slice to another.
// slice1 := make([]int, 5, 20);
// slice2 := make([]int, 10,30);

// slice3 := append(slice1, slice2...);

// fmt.Print(slice3)

//=======================copy()
// copy(destination, source)

 source := []int{1,2,3,4,5,6,7,8,9};
 destination := make([]int, 3);
  sliceCopy := copy(destination,source);			///3 elements copied output of silceCopy is 3

  fmt.Print(sliceCopy) // 3
  fmt.Print(destination)   ///output [1 2  3]

}