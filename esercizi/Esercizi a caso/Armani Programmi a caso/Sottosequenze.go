package main
import (
    "fmt"
    "os"
)
func main(){
    var s string
    s = os.Args[1]

    for i:=0 ; i<len(s) ; i++ {
        for j:=i+1 ; j<len(s) ; j++ {
            if s[i] == s[j]{
                subslice := s[i:j+1]
                fmt.Println(subslice)
            }
        }
    }
}
