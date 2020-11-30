package main
import "fmt"

func main(){
    var n , somma int
    fmt.Print("Inserisci numero: ")
    fmt.Scan(&n)

    for i:=1; i<n+1 ; i++ {
        somma=0
        for j:=1 ; j<i ; j++ {
            if i % j==0 {
                somma+=j
            }
        }
        if somma>i{
            fmt.Println(i)
        }
    }
}
