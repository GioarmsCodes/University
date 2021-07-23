package main

import (
    "fmt"
    "time"
)

func Primo(n int) bool{
    for i:=2 ; i<n/2+1 ; i++{
        if n%i == 0{
            return false
        }
    }
    return true
}

func main(){
    for i:=2 ; ;i++{
        if Primo(i){
            fmt.Println(i)
            time.Sleep(99999)
        }
    }
}
