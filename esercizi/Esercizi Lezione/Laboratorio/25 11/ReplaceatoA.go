package main
/* Replace the a rune with another rune inside a string, the
 runes are and the strings comes from the outside*/
import "fmt"
func Replace(s string,o string,n string)  string{
    c := []rune(s)
    for i:=0;i<len(c);i++{
        if c[i] == rune(o[0]) {
            c[i] = rune(n[0])
        }
    }
    return string(c)
}
func main(){
    var s string
    var vecchio string
    var nuovo string

    fmt.Print("inserisci stringa: ")
    fmt.Scan(&s)
    fmt.Print("inserisci vecchia variabile: ")
    fmt.Scan(&vecchio)
    fmt.Print("inserisci nuova variabile: ")
    fmt.Scan(&nuovo)
    c := Replace(s,vecchio,nuovo)
    fmt.Println(string(c))
}
