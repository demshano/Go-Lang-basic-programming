package main
import ("fmt");
func main(){
	
	//mainly there are servaral datya types in go language
	/*
	int, float32, float64, bool and string

	

	*/

		var numberInt int  = 32;
		var numberFloat32 float32 = 32.4323;
		var numberFloat64 float64 = 32.24346256744;
		var flag bool = true;
		var name string = "cloud parallax";

		
		fmt.Println("Integer:", numberInt)
		fmt.Println("Float32:", numberFloat32)
		fmt.Println("Float64:", numberFloat64)
		fmt.Println("Boolean:", flag)
		fmt.Println("String :", name)

		fmt.Printf("%s", "------using Printf function--------\n");

		fmt.Printf("Integer %d, float32 %f, float64 %f, boolean %b, string %s", numberInt, numberFloat32, numberFloat64, flag, name);

	


}