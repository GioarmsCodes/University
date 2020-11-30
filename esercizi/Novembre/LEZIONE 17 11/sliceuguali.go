package main
import "fmt"

func Sliceuguali (s1 []int, s2 []int) (c bool) {
    var v int
    v=0
    if len(s1) == len(s2) {
        for i:=0 ; i<len(s1) ; i++ {
            if s1[i] == s2[i]{
                v++
            }
        }
    }
        if v == len(s1){
            c=true
        }else{
            c=false
        }
        return
}

func main() {

    array:= [6]int{1,2,3,32,3,1236}

    //test
    fmt.Println(Sliceuguali(array [3:5],array [3:5]))
    fmt.Println(Sliceuguali(array [4:5],array [3:5]))
    fmt.Println(Sliceuguali(array [2:5],array [3:5]))
    fmt.Println(Sliceuguali(array [3:5],array [1:5]))
    //test
}
