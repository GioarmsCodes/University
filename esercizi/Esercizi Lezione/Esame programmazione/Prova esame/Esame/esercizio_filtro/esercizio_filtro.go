package main

import (
    "fmt"
    "os"
    "strconv"
)

func main(){
    for i:=1 ; i<len(os.Args) ; i+=2{
        lettera := os.Args[i]
        numero ,_ := strconv.Atoi(os.Args[i+1])

        for j:=0 ; j<numero ; j++{
            fmt.Print(lettera)
        }
    }
}
