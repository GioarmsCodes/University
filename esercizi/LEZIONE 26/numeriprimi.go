package main
import (
    "math"
    "fmt"
)
func main(){
    var numero, m, i, c int
    c=0
    fmt.Print("Inserisci numero : ")
    fmt.Scan(&numero)

    for i=2; i <= numero; i++ {

        for m=2; m < i ; m++{

            if i % m == 0 {
                break
            }

        }

        if !(m < i ) {
            c++
            fmt.Println(i)

        }
    }
    aprossimazione:= float64(numero)/(math.Log(float64(numero)))
    fmt.Print("ci sono ",c," numeri primi \n")
    fmt.Print(aprossimazione)
}
