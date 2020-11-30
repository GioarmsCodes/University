package main
import "fmt"

func Quantevolte (s1 []int, s2 []int) (c int) {
        for i:=0 ; i<len(s1) ; i++ {
            for j:=0 ; j<len(s2) ; j++{
                if s1[i] == s2 [j]{
                    c++
                }
            }

}
return
}
func main() {
    array:= [10]int{0,1,2,3,4,5,6,7,8,9}
    v:= Quantevolte(array [3:7],array [2:8])
    fmt.Println(v)
}
