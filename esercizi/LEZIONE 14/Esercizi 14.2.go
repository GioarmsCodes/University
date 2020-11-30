package main
import "fmt"
/*Programma che dato un numero float64 restistuisce un numero intero troncato*/
func main(){
    var numero float64
    fmt.Println("inserisci numero con la virgola")
    fmt.Scan(&numero)
    numero = float64(int(numero + 0.5))
    fmt.Println(numero)
}
