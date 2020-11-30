package main
import "fmt"
/* conto alla rovescia */
func main(){
    var numero, i  int
    fmt.Print("Inserisci numero :  ")
    fmt.Scan(&numero)
    for i=0 ; i<= numero ; i++ {
        fmt.Println(numero-i)
    }
}
