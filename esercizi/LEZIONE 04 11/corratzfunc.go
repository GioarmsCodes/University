package main
import "fmt"
/* Successione di Corratz*/
func Corratz(n int) {
    for ; n!=1 ; {
        if n%2 == 0{
            n/=2
        }else{
            n=3*n+1
        }
        fmt.Println(n)
        }
}
func main(){
    var numero int
    fmt.Print("Inserisci numero: ")
    fmt.Scan(&numero)
    Corratz(numero)
}
