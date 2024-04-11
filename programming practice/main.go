package main
import ("fmt")

func main(){

	numbers := [...]int{23,45,7,8,43,65,88,91,3,55}

	fmt.Print(bubbleSorting(numbers))



}

func bubbleSorting(numbers [10]int)([10]int){
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers)-1; j++ {
			if numbers[j] > numbers[j+1] {

				temp := numbers[j];
				numbers[j] = numbers[j+1]
				numbers[j+1] = temp
			

		}
	}
}
	return numbers
}

