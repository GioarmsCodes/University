package main
import "fmt"
func main(){
    var numero int
    fmt.Print("Inserisci numero : ")
    fmt.Scan(&numero)

    for i:=0; i<numero; i++{
        for m:=0; m<numero-i; m++{
            fmt.Print(" ")
        }
        for j:=0; j<i+1; j++{
            if i%2==0{
                fmt.Print("**")
            }else {
                fmt.Print("°°")
            }
        }
        fmt.Println("")
    }
    for i:=0; i<numero; i++{
        for m:=0; m<i+1; m++{
            fmt.Print(" ")
        }
        for j:=0; j<numero-i; j++{
            if i%2==0{
                fmt.Print("**")
            }else {
                fmt.Print("°°")
            }
        }
        fmt.Println("")
    }

}
