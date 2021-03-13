package main

import "fmt"

func Ricors(s string) string {
    var indexmax int
    if len(s) == 1{
        return s
    }
    max := rune(s[0])
    for i,x := range s{
        if x<max {
            max = x
            indexmax = i
        }
    }
    resto := s[:indexmax]+s[indexmax+1:]
    return string(max) + Ricors(resto)
}


func main(){

    var s string
    fmt.Println("Inserisci stringa")
    fmt.Scan(&s)

    l := Ricors(s)

    fmt.Println(l)

}
