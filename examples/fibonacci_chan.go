package main
/*
Fibonaci sequence using Go Chanel
*/
import ("fmt")
func fibonacci(n int,c chan int){
	x,y:=0,1
	i:=0
	for i<=n{
		c<-x //Write to chanel
		x,y=y,x+y
		i+=1
	}
	close(c)
}
func main(){
	ch:=make(chan int,3)
	N:=30
	go func(){
		i:=0
		for(i<N){
			v,ok:=<-ch
			if(ok){
				fmt.Print(v," ")
				i+=1
			}
		}
		fmt.Println("")
	}()
	fibonacci(N,ch)

}