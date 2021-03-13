package main

import(
    "fmt"
    "os"
    "strconv"
)

func main(){
    s := os.Args[1]
    var t string
    var temp2 int

    for i:=0; i<len(s); {
        temp1,_ := strconv.Atoi(string(s[i]))
        if i < len(s)-1{
        temp2,_ = strconv.Atoi(string(s[i+1]))
}
        if temp1 <= temp2{
            t = t + string(s[i])
            i++
        }else{
            t += string(s[i])
            i++
            fmt.Println(t)
            t= ""
        }

    }
    fmt.Println(t)

}
