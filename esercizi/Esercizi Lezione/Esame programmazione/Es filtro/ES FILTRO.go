package main

import (
    "fmt"
    "strconv"
)

func main(){
    var s string
    var numeri []int
    var flag bool
    fmt.Println("Inserisci stringa: ")
    fmt.Scan(&s)

    for _,x:= range s{
    temp,err := strconv.Atoi(string(x))
    if err != nil{
        continue
    }
    numeri =append(numeri,temp)
    }

    for i:=1 ;i<len(numeri);i++{
        if numeri[i]>numeri[i-1]{
            flag = true
            break
        }
    }
    if flag {
        fmt.Println("Sequenza nascosta non ordinata.")
    }else{
        fmt.Println("Sequenza nascosta ordinata.")
    }
}
