package main
import "fmt"
/*Dato un numero float64 restistuisce solo la parte frazionaria*/
func main(){
    var numero float64
    fmt.Println("inserisci numero con la virgola")
    fmt.Scan(&numero)
    numero -=float64(int(numero))
    fmt.Println(numero)
}
