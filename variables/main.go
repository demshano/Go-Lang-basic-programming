package main
import ("fmt")

func main(){
	
	//=========decalring with var keyword
	// var firstName string = "deshan";
	// var lastName string;
	// lastName = "manodya"
	// fmt.Print("first name: "+ firstName, "\nlast name: "+ lastName)

	//========decalring with := sign

	// firstName := "deshan";
	// lastName := "manodya";
	// fmt.Print("first Name: "+ firstName +"\n"+ "last name: "+ lastName);

	// NOTED: if we decalring with := sign , we must assign a value
	//NOTED: only we can use inside of the function, can not use outside of the function

	// var numberInt int =12;
	// var numberFloat float32 =  12.22;
	// var flag bool = true;

	// fmt.Print(numberInt,"\n");
	// fmt.Println(numberFloat);
	// fmt.Println(flag);

	/*
	var a string;  // default value is: ""
	var b int;     // default value is: 0
	var c bool;		// default value is: false

	*/

	//========declaring multiple  variables in one line
	var i,j int =5,6;
	fmt.Printf("%d %d",i,j);

	//here we go use fmt.Printf() function that specify the formatted string %d is about integer values

	// ======fmt package provide servaral formatting verbs
	/*
		%d: Print an integer in decimal format.
		%f: Print a floating-point number (float64) in decimal notation.
		%s: Print a string.
		%t: Print a boolean value.
		%c: Print a single character (byte).
		%v: Print the value in its default format (general-purpose verb).
		%b: Print an integer in binary format.
		%o: Print an integer in octal format.
		%x: Print an integer in hexadecimal format, with lowercase letters.
		%X: Print an integer in hexadecimal format, with uppercase letters.
		%p: Print a pointer address.
		%T: Print the type of the value.
	*/

	//===============type conversion==================
	var x int = 47;
	var y float32 = float32(x);
	fmt.Println(y);

	//===============constants==================
	const pi float32 = 3.14;
	fmt.Println(pi);

	// =========go output functions
	//in go there are three main type of output functyion
	
	/*  
		Print()
		Println()
		Printf()
	*/
	

}