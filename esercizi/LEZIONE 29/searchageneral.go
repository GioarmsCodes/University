package main
import (
    "fmt"
)
/* search a on a string */
func main(){
    var parola string
    var c int

    fmt.Print("inserisci parola: ")
    fmt.Scan(&parola)
    for _,x:= range parola{
        if x=='a' {
            c++
        }
        
    }
    fmt.Println("ci sono: ",c, "a")
}
