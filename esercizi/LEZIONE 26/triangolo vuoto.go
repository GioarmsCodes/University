package main
import "fmt"
/* numero 3 search */
func main(){
    var numero  int
    fmt.Print("inserisci numero:  ")
    fmt.Scan(&numero)

    for i:=0 ; i <numero ; i++ {
        for j:=0 ; j <numero-i ; j++ {
            if i==0 || j==0 || j==numero-i-1{
            fmt.Print("*")
        }else{
            fmt.Print(" ")
        }

            }
        fmt.Println("")
        }

}
