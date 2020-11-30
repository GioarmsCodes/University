package main
import "fmt"
/* Stampa potenze di 2 ^ n  (n inserito dall' utente)*/
func main(){
    var numero, contatore, i int
    contatore = 0
    fmt.Print("dammi numero:  ")
    fmt.Scan(&numero)
    for i=1; contatore<=numero; i*=2 {
        fmt.Println(i,"2 alla ",contatore)
        contatore++
        }
    }
