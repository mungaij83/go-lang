package main
/*
Insertion sort algorithm in GoLang

Sorts integers only
*/
import ("fmt")
/*
Base case: Only one element or number
General Case: More than one elements in the array, apply insertion sort algorithm.
*/
func sort(numbers [] int) [] int{
	if(len(numbers)==1){
		return numbers
	}
	var i,j,tmp int
	for i=1;i<len(numbers);i++{
		tmp=numbers[i]
		j=i-1
		for(j>=0 && numbers[j]>tmp) {
			numbers[j+1]=numbers[j]
			j-=1
		}
		numbers[j+1]=tmp
		fmt.Printf("Iterations[%d]: State %v\n",i,numbers)
	}
	return numbers
}
func main(){
	numbers:=[] int{4,12,2,7,8,3,45,9,10}
	others:=[] int {-5,6,3,8,12,78,12,45,2,1,-90}
	var sorted[] int;
	sorted=sort(numbers)
	fmt.Printf("Numbers: %v\n",numbers)
	fmt.Printf("Sorted= %v\n",sorted)
	fmt.Println("-----------------------------")
	fmt.Printf("Numbers: %v\n",others)
	fmt.Printf("Sorted= %v\n",sort(others))
}