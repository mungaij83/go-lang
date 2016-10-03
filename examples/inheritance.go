package main
/*
Inheritance using anonymous classes: subclassing and inheritance
Multiple inheritance can also be done using the same technique
*/
import (
	"fmt"
)
type Car struct{
	numberOfWheels int
	name string
}
func (c Car) Wheels() int{
	return c.numberOfWheels
}
type Ferrari struct{ //Inherit
	Car //Anonymous class
	max_speed int
}
func (f Ferrari) Moving() int{
	return f.max_speed
}
func main(){
	f:=Ferrari{Car{numberOfWheels:4,name:"Ferrari"},233}
	c:=Car{numberOfWheels:4,name:"Car"}
	fmt.Println("Number of Wheels: ",f.Wheels()) //Method wheels on super class
	fmt.Println("Max Speed: ",f.Moving(),"Mph")
	// fmt.Println("Car max speed: ",c.Moving()) error here: c.Moving undefined (type Car has no field or method Moving)
	fmt.Println("Car: ",c) 
}