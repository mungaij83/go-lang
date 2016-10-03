package main
/*
Matrix multiplication
*/
import(
	"fmt"
)
type Struct Matrix{
	M [][] int
}
func createMatrix(n int,m int,val) Matrix{
	matrix:=Matrix
	matrix.M=[n][m] int
	return matrix
}
func size(mat *Matrix){
	r:=len(mat)
	c:=len(mat[0])
	return r,c
}
func multiply(A *Matrix,B *Matrix) Matrix{
	r_a, c_a=A.size()
	r_b,c_b=B.size()
	if(c_a==r_b){
		matrix:=createMatrix(r_a,c_b)
		for(i:=0;i<r_a;i++){
			for(j:=0;j<c_b;j++){
				c:=0
				for(k:=0;k<r_a;k++){
					c:=c+A.M[i][k]*B[j][k]
				}
				matrix.M[i][j]=c
			}
		}
		return matrix
	}
	else{
		panic("Matrix size cannot be multiplied")
	}
}
func strasens(A *Matrix,B Matrix) Matrix{
	return nil
}

func main(){
	matrix:=Matrix
	matrix2:=Matrix
	matrix.M=[][]int {{9,5}{7,7}}
	matrix2.M=[][]int{{5,5}{4,3}}
	m:=multiply(matrix,matrix2)
	Println(m)
}