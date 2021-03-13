package main

import "fmt"

func Poligono(x []int,y []int) bool{
    var numero int
    j := x
    q := y
    for i:=0 ; i<len(X) ; i++{
        for m:=0 ; m<len(X) ; m++{
            if x[i] == j[m] && y[i] == q[m]{
                return false
            }
        }
    }
    return true
}

func main() {
        var numero int
        var test bool

        fmt.Println("Inserisci numero di punti")
        fmt.Scan(&numero)

        for i:=0;i<numero;i++{
            var x []int
            var y []int
            fmt.Println("inserisci Punto n-",i+1,": [X Y]")
            fmt.Scan(&x[i])
            fmt.Scan(&y[i])
        }

        test = Poligono(x,y)
}
