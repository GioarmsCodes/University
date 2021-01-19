package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
)

func main(){
    var s string
    lunghezza,_ := strconv.Atoi(os.Args[1])
    nomefile := os.Args[2]
    f,_ := os.Open(nomefile)
    defer f.Close()
    b := bufio.NewScanner(f)

    for b.Scan(){
        s+=b.Text()
        if string(s[len(s)-1]) != " "{
            s += " "
            }
    }

    for i := 0 ; i<len(s) ; {
        for j:=0 ; j<lunghezza ; j++{
            fmt.Print(string(s[i]))
            i++
            if i == len(s){
                break
            }
        }
        fmt.Println("")
    }
}
