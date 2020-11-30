package main
import "fmt"
func traslazione(s string,n int) string{
    var traslato string
    traslato = ""
    for i:=n ; i < len(s) ; i++ {
            traslato = traslato + string(s[i])
    }
    for i:=0 ; i < n ; i++ {
            traslato = traslato + string(s[i])
    }
    return traslato
}
func main(){
    var n int
    var s, traslato string
    fmt.Print("Inserisci parola : ")
    fmt.Scan(&s)
    fmt.Print("Inserisci traslazione : ")
    fmt.Scan(&n)
    traslato=traslazione(s, n)
    fmt.Println(traslato)

}
