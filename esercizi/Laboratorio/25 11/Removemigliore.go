package main
// funzione che data una slice di un tipo x ed una slice di indici rimuove  l'elemento dall indice
import (
    "fmt"
    "bufio"
    "os"
)

func Remove(c []rune,i []int) []rune {
    var cont int
    var l []rune
    for j:=0 ;j<len(c);j++{
        if j==i[cont]{
            if j<=len(i)-1 {
                cont++
            }
        }else{
            l = append(l,c[j])
         }
        }
        return l
    }
    func main(){
        var s string
        var q,e int
        var listi []int
        b:= bufio.NewScanner(os.Stdin)

        fmt.Print("Inserisci stringa: ")
        b.Scan()
        s = b.Text()

        fmt.Print("Quanti elementi vuoi togliere")
        fmt.Scan(&e)
        for i:=0;i<e;i++{
            fmt.Print("Inserisci indice:")
            fmt.Scan(&q)
            listi = append(listi,q)
        }

        c := []rune(s)
        smallc := Remove(c,listi)
        fmt.Println(string(smallc))
    }
