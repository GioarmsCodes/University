package main

import (
    "fmt"
    "os"
    "strconv"
)

func Find(s string,n int) string{
    var index int
    var min int
    if n==1 {
        min,_= strconv.Atoi(string(s[0]))
        for i:=0 ; i<len(s); i++{
            temp,_ := strconv.Atoi(string(s[i]))
            if temp < min{
                min = temp
            }
            }
            return strconv.Itoa(min)
    }
    min,_= strconv.Atoi(string(s[0]))
    for i:=0 ; i<len(s)-n+1; i++{
        temp,_ := strconv.Atoi(string(s[i]))
        if temp < min{
            min = temp
            index = i
        }
        }
    return strconv.Itoa(min) + Find(s[index+1:],n-1)
    }

func main(){
    n1 := os.Args[1]
    n2,_:= strconv.Atoi(os.Args[2])
    lunghezza := len(n1) - n2
    test := Find(n1,lunghezza)
    fmt.Println("numero migliore:",test)

}
