package main
import "fmt"
func main(){
    var numero int
    for {
        fmt.Print("Inserisci valore: ")
        fmt.Scan(&numero)
        if numero <= 0 {
            return}
            for i:=1;i<=numero;i*=2{
                if numero == i{
                    fmt.Println("e' una potenza di 2")
                    break
                }

            }
        }
    }
