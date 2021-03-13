package main
import "fmt"
/* Stampa potenze di 2 */
func main(){
    var numero, i  int
    fmt.Print("Inserisci numero :  ")
    fmt.Scan(&numero)
    if numero < 2 {
        fmt.Println("inserisci un numero maggiore uguale a 2")
    }else {
    for i=2 ;i <= numero ; i*=2{
        fmt.Println(i)
        }
     }
    }
