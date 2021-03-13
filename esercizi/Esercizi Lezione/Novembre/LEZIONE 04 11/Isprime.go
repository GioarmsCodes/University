package main
import "fmt"
 // Funzione che Mi restituisce 1 se e' primo 0 se non e' primo

func IsPrime(n int)  {
    var i int
    for i=2;i<=n;i++{
        if n%i==0{
            break
        }
    }
    if i==n{
        fmt.Println("E' un numero primo")
    }else{
        fmt.Println("non e' un numero primo")
    }
}

func main(){
    var numero int
    fmt.Print("Inserisci numero: ")
    fmt.Scan(&numero)
    IsPrime(numero)
}
