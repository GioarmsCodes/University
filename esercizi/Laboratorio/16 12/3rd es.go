package main

import (
	"fmt"
)
func Binomiale(n int ) [][]int {
triangolo:=make([][]int,n+1)
for i:=0 ;i<n;i++{
triangolo[i]=make([]int,n+1)
}
triangolo[0][0] = 1
for i:=1;i<n;i++{
	triangolo[0][i]=0
}
for j:=1;j<n;j++{
	triangolo[j][0]=1
	for m:=1;m<n;m++{
		triangolo[j][m] = triangolo[j-1][m-1]+triangolo[j-1][m]
	}
}
	return triangolo
}
func main() {
	var n int


	fmt.Println("Inserisci Numero")
	fmt.Scan(&n)
	bin := Binomiale(n)
	fmt.Println(bin)
	for i:=0 ; i<n ;i++{
		for j:=0 ;j<n ;j++{
			if bin[i][j]!= 0 {
				fmt.Print(bin[i][j])
			}
		}
		fmt.Println()
	}
}
