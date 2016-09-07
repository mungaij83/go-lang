package main
/*
Bubble sort algorithm
Runnint time: O(n^2)
*/
import (
	"fmt"
)
func swap(a *int,b *int){
	tmp:=*a
	*a=*b
	*b=tmp
}
func bubble_sort(A *[] int){
	var N int=len(*A)-1
	for i:=0;i<N;i++{ //bubble from 0 to N
		for j:=N;j>=i+1;j--{ //Down to i
			if((*A)[j]<(*A)[j-1]){
				tmp:=(*A)[j-1]
				(*A)[j-1]=(*A)[j]
				(*A)[j]=tmp
			}
		}

	}
}
func main(){
	array:=[] int{5,7,2,1,-7,4,40,-30}
	fmt.Println("Array: ",array)
	bubble_sort(&array)
	fmt.Println("Sorted Array:",array)
}