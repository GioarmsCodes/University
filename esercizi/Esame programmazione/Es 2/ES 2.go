package main

import (
    "fmt"
    "os"
    "strconv"
)

func main(){
    var s string
    n,_:=strconv.Atoi(os.Args[1])

    for i:=0 ; i<n ; i++{
            for j:=i ; j<n ; j++{
                s +="0"
            }
            if i==0{
            s += "+"
        }else{
            s += " "
            for m:=0 ;m<i*2;m++{
                s += " "
            }
        }
        for j:=i ; j<n ; j++{
            s +="0"
        }
        fmt.Println(s)
        s = ""
    }
}
