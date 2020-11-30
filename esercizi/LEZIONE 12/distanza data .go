package main

import (
    "fmt"
)

func main() {
    var mese, giorno, totGiorni int

    fmt.Print("Inserisci il numero della data odierna: ")
    fmt.Scan(&giorno)

    fmt.Print("Inserisci il numero del mese della data odierna: ")
    fmt.Scan(&mese)

    totGiorni = ((mese - 1) * 30) + giorno
    fmt.Println("sono passati", totGiorni, " giorni dal primo gennaio ")
    totGiorni =355 - ((mese - 1) * 30 + giorno)

    if (totGiorni < 0 ) {
        fmt.Println("sono passati", - totGiorni, "giorni da natale")
        totGiorni = 360 + totGiorni
    }
    fmt.Println("mancano ", totGiorni, "giorni per natale ")
}
