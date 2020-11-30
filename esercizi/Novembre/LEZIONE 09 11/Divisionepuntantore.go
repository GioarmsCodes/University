package main
import "fmt"
func Divisione (divisore int,dividendo int,quoz *int, resto *int){
    *quoz,*resto = dividendo/divisore,dividendo%divisore
}
func main() {
    var n1,n2,q,resto int

    fmt.Print("Inserisci dividendo: ")
    fmt.Scan(&n1)
    fmt.Print("Inserisci divisore: ")
    fmt.Scan(&n2)
     Divisione(n1,n2,&q,&resto)
    fmt.Println("Risultato: ",q,"\nResto: ",resto)
}
