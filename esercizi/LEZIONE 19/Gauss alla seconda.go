package main
import "fmt"
/* Calcola la somma 1^2+2^2+3^2+4^2 */
func main(){
    var numero, s, i int
    s = 0
    fmt.Print("Inserisci il numero:  ")
    fmt.Scan(&numero)
    for i=1 ; i<=numero ; i++ {
        s += i*i
    }
    fmt.Print(s)
}
