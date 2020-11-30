package main
import "fmt"
func Divisione (dividendo int, divisore int) (d int ,r int) {
    d = dividendo / divisore
    r = dividendo % divisore
    return
}

func main() {
    var n1,n2,d,r int
    fmt.Print("Inserisci primo numero: ")
    fmt.Scan(&n1)
    fmt.Print("Inserisci secondo numero: ")
    fmt.Scan(&n2)
    d, r = Divisione(n1,n2)
    fmt.Println("Risultato: ",d,"\nResto: ",r)
}
