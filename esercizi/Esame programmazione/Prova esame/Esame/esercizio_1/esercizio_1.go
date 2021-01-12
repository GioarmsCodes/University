package main

import (
    "fmt"
    "os"
    "unicode"
)

func main(){
    var flag bool
    for i:=1 ;i<len(os.Args) ; i++{
        if i%2 == 1{
            flag = true
        }else{
            flag  = false
        }
        for _,x := range os.Args[i]{
            if flag{
                fmt.Print(string(unicode.ToUpper(x)))
                flag = false
            }else{
                fmt.Print(string(unicode.ToLower(x)))
                flag = true
            }
        }
        fmt.Print(" ")

    }
}
