package main
import ("fmt")
func main() {

	//there are two ways to declare an array, with var keyword and := sign

	//var names_withVar1 = [3]string{"deshan", "kosala", "harsha"}; //using var keyword 
	// var names_withVar2 = [...]int{12,234,45,67};
	// fmt.Print(names_withVar1);
	// fmt.Print(names_withVar2);

	// fmt.Println();

	// for x := 0; x < len(names_withVar1); x++ {
	// 	fmt.Println("The element at index ",(x+1)," is : ",names_withVar1[x]);
	// }

	// ====accesing array elements

	//fmt.Print(names_withVar1[0]);

	//=============array initializzation

	/*

	//not initialize
	array1 := [5]int{}                   //all zeros	[0,0,0,0,0]
	//partially initialize
	array2 := [5]int{1,2}					// [1,2,0,0,0]
	//fully initialize
	array3 := [5]int{1,2,3,4,5};			//[1,2,3,4,5]

	*/

	//=============initialize only specific elements

	var arrayNumbers = [5]int{1:12,3:30};
	fmt.Print(arrayNumbers);

	//=====find the length of the array
	// go used len() function to find the array length
	lengthOfArray := len(arrayNumbers)
	fmt.Print(lengthOfArray);


}