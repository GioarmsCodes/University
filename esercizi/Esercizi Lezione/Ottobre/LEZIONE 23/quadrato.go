package main
import "fmt"
/* numero 3 search */
func main(){
    var numero  int
    fmt.Print("inserisci numero:  ")
    fmt.Scan(&numero)

    for i:=0 ; i <numero ; i++ {
        for j:=0 ; j <numero ; j++ {
            fmt.Print("*")
            }
        fmt.Println("")
        }

}
