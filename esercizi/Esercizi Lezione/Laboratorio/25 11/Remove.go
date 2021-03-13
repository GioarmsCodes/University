package main
// funzione che data una slice di un tipo x ed un indice rimuove  l'elemento dalla slice
import "fmt"

func Remove(c []rune,i int) []rune {
    var l []rune
    for j:=0 ;j<len(c);j++{
        if j==i{
            continue
        }else{
            l = append(l,c[j])
        }
    }
    return l
}
func main(){
    var s string
    var i int
    fmt.Print("Inserisci stringa: ")
    fmt.Scan(&s)
    fmt.Print("Inserisci indice: ")
    fmt.Scan(&i)

    c := []rune(s)

    smallc := Remove(c,i)
    fmt.Println(string(smallc))
}
