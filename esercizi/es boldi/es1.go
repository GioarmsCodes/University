package main

import (
    "fmt"
    "os"
)

func Find(c byte ,s string ) int{
    if len(s) == 0{
        return -1
    }
    if s[0] == c{
        return 0
    }
    x := Find(c,s[1:])
    if x == -1{
        return -1
    }
    return 1+x
}

func main(){
    c := os.Args[1][0]

    s:= os.Args[2]


    dummy  := Find(c,s)

        fmt.Println(dummy)


}
