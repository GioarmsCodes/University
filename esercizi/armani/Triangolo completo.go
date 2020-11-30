package main
import "fmt"
func main(){
    var numero int
    fmt.Print("Inserisci numero : ")
    fmt.Scan(&numero)

    for i:=0; i<numero; i++{
        for m:=0; m<i+1; m++{
            fmt.Print(" ")
        }
        for j:=0; j<numero-i; j++{
            fmt.Print("**")
        }
        fmt.Println("")
    }

}
