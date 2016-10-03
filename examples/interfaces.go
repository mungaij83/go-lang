package main
/*
Interfaces in GoLang
Use "er" convention to indicate an interface
Polymophism: Different behaviours on same type of object.
*/
import (
	"fmt"
	"math"
)
type Shaper interface{
	Area() float64
	Perimeter() float64
}
type Rectangle struct{
	length,width int
}
func(r Rectangle) Area() float64{
	return float64(r.length*r.width)
}
func(r Rectangle) Perimeter() (t float64){
	t=float64(r.length+r.width)*2
	return
}
type Circle struct{
	radius float64
}
func (c Circle) Area() float64{
	return math.Pi*c.radius*c.radius
}
func (c Circle)Perimeter() float64{
	return 2*math.Pi*c.radius
}
func main(){
	r:=Rectangle{width:50,length:30}
	fmt.Println("Rectangle: ",r)
	fmt.Println("Area:	\t",r.Area())
	fmt.Println("Perimeter:\t",r.Perimeter())
	s:=Shaper(r) //Type casting to Shaper, not necessary just demostrating
	fmt.Println("Rectangle: ",s)
	fmt.Println("Area:	\t",s.Area())
	fmt.Println("Perimeter:\t",s.Perimeter())
	circle:=Circle{radius:float64(7.0)}
	fmt.Println("Circle: ",circle)
	fmt.Println("Area:	\t",circle.Area())
	fmt.Println("Perimeter:\t",circle.Perimeter())
}