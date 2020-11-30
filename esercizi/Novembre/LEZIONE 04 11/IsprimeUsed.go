package main
import "fmt"
//Utilizzo la mia funzione creata precedentemente per un programma assegnatomi in precedenza
func IsPrime(n int) int{
    var i int
    for i=2;i<=n;i++{
        if n%i==0{
            break
        }
    }
    if i==n{
        return 1
    }else{
        return 0
    }
}

func main(){
    var numero int
    fmt.Print("Inserisci numero: ")
    fmt.Scan(&numero)
    for i:=0;i<=numero;i++{
        if IsPrime(i) == 1{
            fmt.Println(i)
        }
    }
}
