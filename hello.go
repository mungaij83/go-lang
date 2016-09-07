package main

import ("fmt"
		"math"
)

type Circle struct{
	radius float64
}
//Go methods
func(circle Circle) area() float64{
	return math.Pi*circle.radius*circle.radius
}
/* Function*/
//Also
func main(){
	fmt.Println("Hello Go");
	var age=56
	var s=true
	fmt.Println(age)
	fmt.Println(s)
	values:=func(num1 float64) float64{
		return math.Sqrt(num1)
	}
	sequence:=getSequence()
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(values(8.5))
	circle:=Circle{radius:7}
	fmt.Println("Area: ",circle.area())
	var sl=make([] int,5,5)
	sl=append(sl,9)
	sl=append(sl,1200)
	fmt.Printf("Array of size=%d, sliece=%v\n",len(sl),sl)
	fmt.Printf("Factorial: %d!=%d\n",5,factorial(5))
	fmt.Println("Fibonacci: ",Fibonacci(12))
	fmt.Println("Fibonacci: ",Fibonacci(25))
	pascal(5)
}
//function as value
func getSequence() func() int{
	var i=0;
	return func() int{
		i+=1
		return i
	}
}
func factorial(num int) int{
	if(num<=1){
		return 1
	}
	return num*factorial(num-1)
}
func pascal(n int){
	var triangle=make([]int,n,n)
	var j,i int;
	triangle[0][0]=1
	triangle[1][0]=1
	triangle[1][1]=1;
	for i=0;i<n;i++{
		for j=0;j<i;j++{
			triangle[i]=append(triangle[i],triangle[i-1][j]+triangle[i-1][j+1])
			fmt.Println(triangle[i][j])
		}
		fmt.Printf("\n")
	}
}
func Fibonacci(num int) [] int{
	var nums [] int;
	var n int;
	for n=1;n<num;n++{
		nums=append(nums,fibonacci(n))
	}
	return nums
}
func fibonacci(num int) int{
	if(num==0){
		return 0
	}
	if(num==1){
		return 1
	}
	return fibonacci(num-1)+fibonacci(num-2)
}