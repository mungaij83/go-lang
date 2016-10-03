package main
/*
Go object oriented. Using struct
1. Protected aspect- If first leter is capital, it is exported outside the package 
2. Private variable- Start with small leters, only visible in current package
Contain methods or behaviour
When defining a function or method associated with a type, it is given as a named variable e.g (c Cube) funcName() retType;
You will be able to add methods for a type only if the type is defined in the <<same package>>.Anonymous fields, can be used to extend a struct in another package
*/
import(
	"fmt"
)

type Cube struct{
	width int
	height int
	length int
	name string
}
/*
Method to return area
*/
func (c Cube) Area() int{
	return 2*(c.width*c.height+c.height*c.length+c.width*c.length)
}
/*
Method to return volume
*/
func (c Cube) Volume() int{
	return c.width*c.length*c.height
}
func main(){
	cube:=Cube{width:30,height:50,length:50,name:"First"}
	pr:=new (Cube) //Pointer to an instance
	pr.width=50
	pr.height=34
	pr.length=5
	pr.name="Cube 3"
	fmt.Println("Cube Pointer",(*pr))
	fmt.Println("Cube volume:",pr.Volume(),"cm3 Area: ",pr.Area())
	fmt.Println("Cube: ",cube)
	fmt.Println("Cube volume:",cube.Volume(),"cm3 Area: ",cube.Area())
}
