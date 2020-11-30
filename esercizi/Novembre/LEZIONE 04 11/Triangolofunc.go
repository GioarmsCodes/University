package main
import "fmt"
/* Programma che stampa triangolo con 2 funzione*/
func Spazio(n int){
    for i:=0 ; i<n ; i++ {
        fmt.Print(" ")
    }
}
func Triangolo(n int) {
    for i:=0 ; i<n ; i++ {
        Spazio(i)
        for j:=0 ; j<n-i ; j++ {
            fmt.Print("*")
        }
        fmt.Println("")
    }
}
func main(){
    var numero int
    fmt.Print("Inserisci numero: ")
    fmt.Scan(&numero)
    Triangolo(numero)
}
