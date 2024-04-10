package main
import ("fmt")

// Define an interface named Shape
type Shape interface {
    Area() float64
}

// Define a struct named Circle
type Circle struct {
    Radius float64
}

//Define a struct named Rectangle
type Rectangle struct {
    Width, Height float64
}

// Implement the Shape interface for the Circle type
func (c Circle) Area() float64 {
    return 3.14159 * c.Radius * c.Radius
}

// Implement the Shape interface for the Rectangle type
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}



func main(){

	//create instane for circle and rectangle
	circle := Circle{
		Radius:  5,
	}

	rectangle := Rectangle{
		Width: 8,
		Height: 7,
	}
	fmt.Println(getArea(circle));
	fmt.Println(getArea(rectangle))

}

// A function that accepts a Shape interface and calls its Area method
func getArea(shape Shape) float64 {
    return shape.Area()
}