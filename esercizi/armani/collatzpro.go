package main
import "fmt"
func main(){
    var n int
    fmt.Print("Inserisci numero : ")
    fmt.Scan(&n)
    for i:=1; i<=n; i++{
        numero := i
        fmt.Println(i,")")
        for numero > 1 {
            if numero%2 == 0 {
                numero /= 2
            }else {
                numero = 3 * numero + 1
            }
            fmt.Print(i, " ")
            fmt.Print(numero, " ")

        }
        fmt.Println(" ")
    }

}
