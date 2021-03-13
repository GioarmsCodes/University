package main
import (
    "fmt"

)
/* search a on a string */
func main(){
    var parola string
    var i,c int

    fmt.Print("inserisci parola: ")
    fmt.Scan(&parola)
    x:=len(parola)
    for i=0;i<x; i++{
        if parola[i] == 'a' {
            c++
        }
    }
    fmt.Println("contiene ",c," a")
}
