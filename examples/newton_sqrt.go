package main
/*
Newtons Square root aproximation
Z(n)=x/2
Z(n+1)=Zn-(Z(n)^2-x)/2Z(n)
*/
import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z:=x/2
	for i:=0;i<100;i++{
		z=z-((z*z-x)/(2*z))
	}
	return z
}

func main() {
	fmt.Println(Sqrt(8))
}
