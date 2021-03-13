package main
import "fmt"
func main(){
    var numero int
    fmt.Print("Inserisci numero : ")
    fmt.Scan(&numero)
    for numero>1 {
        if numero%2 == 0 {
            numero /= 2
        }else {
            numero = 3 * numero + 1
        }

        fmt.Println(numero)
    }

}
