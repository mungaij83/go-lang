package main
/*
Go implementation of merge sort algorithm

*/
import (
	"fmt"
	//"math"
)
type Comparable interface{
	less(c Comparable) bool
}

func merge(A *[]int, p int, q int,r int){
	n1:=q-p+1
	n2:=r-q
	var L [] int
	var R [] int
	for i:=0;i<n1;i++{
		L=append(L,(*A)[p+i])
	}
	for i:=0;i<n2;i++{
		R=append(R,(*A)[q+i+1])
	}
	var i int=0
	var j int=0
	var k int=p
	for(k<r && i<n1 && j<n2){
		if L[i]<=R[j] {
			(*A)[k]=L[i]
			i+=1
		}else{
			(*A)[k]=R[j]
			j+=1
		}
		k++
	}
	//Copy remaining items in L and R back to A
	for(i<n1){
		(*A)[k]=L[i]
		i+=1
		k+=1
	}
	for(j<n2){
		(*A)[k]=R[j]
		j+=1
		k+=1
	}
	//fmt.Println("R=",R,"L=",L)
	//fmt.Println("A=> ",A)
}
func merge_sort(array *[] int, p int ,r int) {
	if(p<r){
		//Same as (p+r)/2, Avoid overflow for large integers
		q:=(int)(p+(r-p)/2)
		//fmt.Println("p=",p,"q=",q,"r=",r)
		merge_sort(array,p,q)
		merge_sort(array,q+1,r)
		merge(array,p,q,r)
	}
}
func main() {
	array:=[] int{3,6,2,5,1,6,8,9,0,-4,34,34,23,56,33,50,4,-1,-100}
	fmt.Printf("Array:%v\n",array)
	merge_sort(&array,0,len(array)-1)
	fmt.Printf("Sorted Array:%v\n",array)
}