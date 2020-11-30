package main
import "fmt"
/*il programma dati 2 date dico quale dei due e' piu vecchio */
func main(){
    var giorno1, mese1, anno1, giorno2, mese2, anno2 int
    fmt.Println("1)Inserisci giorno mese anno (es: 3 10 2001)")
fmt.Scan(&giorno1)
fmt.Scan(&mese1)
fmt.Scan(&anno1)

fmt.Println("2)Inserisci giorno mese anno (es: 3 10 2001)")
fmt.Scan(&giorno2)
fmt.Scan(&mese2)
fmt.Scan(&anno2)

if(anno1<anno2){
    fmt.Println("il primo e' piu vecchio del secondo")
}
if(anno1>anno2){
    fmt.Println("il secondo e' piu vecchio del primo")
}
if(anno1==anno2){
    if(mese1<mese2){
        fmt.Println("il primo e' piu vecchio del secondo")
    }
    if(mese1>mese2){
        fmt.Println("il secondo e' piu vecchio del primo")
    }
    if(mese1==mese2){
        if(giorno1<giorno2){
            fmt.Println("il primo e' piu vecchio del secondo")
        }
        if(giorno1>giorno2){
            fmt.Println("il secondo e' piu vecchio del primo")
        }
        if(giorno1==giorno2){
            fmt.Println("SIETE NATI NELLO STESSO GIORNO")
        }
    }
}
}
