package main
import "fmt"
/* numero 3 search */
func main(){
    var numero, i int
    fmt.Print("inserisci numero:  ")
    fmt.Scan(&numero)

    for {
        if numero % 10 == 3 {
            break
        }
        numero/=10
        }
    if i!=numero{
        fmt.Println("contiene un 3")
    }else {
        fmt.Println("non contiene un 3")
    }

}
