package main
import ("fmt")

func main(){

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//range, range i used to iterate elements easy

	array := []int{1,2,3,4,5};

	for index, value := range array {
		fmt.Println("index", index, "value", value);
	}
}