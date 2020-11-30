package main
import "fmt"
/*Il programma dato 2 o 3 numeri restistuisce la meda fra di loro*/
func main(){
    var nNumeri, primon, secondon, terzon float64
        fmt.Println("con quanti numeri vuoi lavorare? (2/3)")
        fmt.Scan(&nNumeri)
        if (nNumeri == 2){
            fmt.Println("Inserisci il primo numero")
            fmt.Scan(&primon)
            fmt.Println("Inserisci il secondo numero")
            fmt.Scan(&secondon)
            somma := primon + secondon
            media := somma / 2
            fmt.Println("la media fa ", media)
        }
        if (nNumeri == 3){
            fmt.Println("Inserisci il primo numero")
            fmt.Scan(&primon)
            fmt.Println("Inserisci il secondo numero")
            fmt.Scan(&secondon)
            fmt.Println("Inserisci il terzo numero")
            fmt.Scan(&terzon)
            somma := primon + secondon + terzon
            media := somma/3
            fmt.Println("la media fa ", media)
        }
        if (nNumeri != 2 && nNumeri != 3){
            fmt.Println("kittemuort metti il 2 o il 3")
        }
}
