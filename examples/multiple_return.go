package main
//multiple returns.
import (
	"fmt"
	"math"
	"errors"
)
//Use of named variabled
func SumProd(i, j int)(a int,b int,c int){
	a=i*j
	b=i/j
	c=i-j
	return
}
/*
Return multiple values from SQRT.
*/
func Sqrt(v float64)(float64,error){
	if(v<0){
		return float64(math.NaN()), errors.New("Not possible here")
	}
	return math.Sqrt(v),nil
}
func main(){
	a,b,c:=SumProd(56,25)
	fmt.Println("Prouct:",a)
	fmt.Println("Division:",b)
	fmt.Println("Subtraction: ",c)
	v:=56.7
	sq,err:=Sqrt(v)
	if(err!=nil){
		fmt.Println("Error:",err)
	}else{
		fmt.Println("Sqrt of ", v, " is ",sq)
	}
}