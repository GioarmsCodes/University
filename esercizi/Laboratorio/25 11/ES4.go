package main
// funzione che data una slice di un tipo x ed un indice rimuove  l'elemento dalla slice
import "fmt"

func Insert(c []string,s string,i int) []string {
    var x []string
    for j:=0;j<len(c);j++{
        if j==i{
            x = append(x,s)
        }
        x = append(x,c[j])
    }
    return x
}
func main(){
    var s string
    var i int
    fmt.Print("Inserisci stringa: ")
    fmt.Scan(&s)
    fmt.Print("Inserisci indice: ")
    fmt.Scan(&i)

    c := []string{"ciao","come","stai","tutto","bene?"}
    fmt.Println(c)

    bigc := Insert(c,s,i)
    fmt.Println(bigc)
}
