package main

import (
    "fmt"
    "os"
)

func main(){

    s := os.Args[1]
    var prec rune
    for i,x := range s{
        if i!=0{
            if ((x>=64 && x<= 90) || (x>=97 && x<= 122)) && (prec>=48 && prec<=57){
                fmt.Printf("%s\n",string(x))
                }
        }
        prec = x
    }
    fmt.Println()
}
