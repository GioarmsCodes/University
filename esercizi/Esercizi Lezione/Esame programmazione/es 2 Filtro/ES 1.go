package main

import (

    "fmt"
    "os"
    "unicode"
)

func main(){
var notAlph bool
var c int
    for i := 0 ; i<len(os.Args) ; i++ {
        for _,x := range os.Args[i] {
            if !unicode.IsLetter(x){
                notAlph = true
                break
        }

    }
    if notAlph {
        c++
        }else{
            fmt.Print(os.Args[i]," ")
        }
        notAlph = false

    }
    fmt.Print(c)

}
