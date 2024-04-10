package main
import ("fmt")

func main(){
	i, j := 42,43; //here i and j variables get memory address of the memory
	fmt.Println(i,j)
	fmt.Println(&i, &j) //printing the memory addresses of i and j here"&"" means --address of--
}