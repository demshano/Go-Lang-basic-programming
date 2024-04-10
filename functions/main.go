package main
import ("fmt")

func main(){

	greetings();
	birthdayGreetings("Deshan", "Manodya")   // calling the function with two arguments 
	fmt.Print(addition(12,34))
	fmt.Print(multiplication(2,4))
	fmt.Println();
	fmt.Print(multiReturn(12, "deshan"))

}

//==========================normal function declaration
func greetings(){
	fmt.Print("hello everyone");
}

//======================parameterize function declaration

func birthdayGreetings(firstName string, lastName string){
	fmt.Print("Happy Birthday ", firstName," "+lastName);
}

//here firstName and LastName are parameters of the function

//==============function return

/*
If you want the function to return a value, you need to define the data type of the return value 
(such as int, string, etc), and also use the return keyword inside the function:

func FunctionName(param1 type, param2 type) type {
  // code to be executed
  return output
}
*/

func addition(x int, y int) int{
	return x + y;
}
//=============named return values

/*
Here, we name the return value as result (of type int), and return the value with a naked return
(means that we use the return statement without specifying the variable name):
*/

func multiplication(number1 float32, number2 float32) (result float32){
	result = number1 * number2;
	return
}

//==============multiple return values

/*
Here, myFunction() returns one integer (result) and one string (txt1):

func myFunction(x int, y string) (result int, txt1 string) {
  result = x + x
  txt1 = y + " World!"
  return
}

func main() {
  fmt.Println(myFunction(5, "Hello"))
}
*/

func multiReturn(number int, name string) (numberResult int, massege string){
	numberResult = number * number;
	massege = name + ", you are awesome!";
	return

}