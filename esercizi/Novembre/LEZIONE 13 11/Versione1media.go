package main
import (
    "fmt"
    "math"
)
func main() {
    var persone,sommaM,sommaV,media,varianza int
    var altezza [100]int

    fmt.Print("Inserisci numero persone: ")
    fmt.Scan(&persone)

    for i:=0 ; i<persone ; i++{
        fmt.Print("Inserisci altezza: ")
        fmt.Scan(&altezza[i])
        sommaM+=altezza[i]
    }
    media = sommaM / persone
    fmt.Println("la media e' :",media)
    for j:=0;j<persone;j++{
        sommaV+=(altezza[j]-media) * (altezza[j]-media)
        }
    varianza= sommaV /persone

     fmt.Println("la Varianza e' :",varianza)
     fmt.Println("lo scarto quadratico e' :",math.Sqrt(float64(varianza)))
}
