// Funzione che grea un istogramma con le ricorrenze dei caratteri presi in stdin 

package main

import (
    "fmt"
    "bufio"
    "os"
    "unicode"
)


//Legge il testo da stdin e lo mette tutto in una sola stringa
func Leggitesto() (s string){
    b := bufio.NewScanner(os.Stdin)
for b.Scan() {
        s+= b.Text()
    }
    return s
}

//Func che Prende un input una mappa rune:int e stampa il carattere in stringa della runa a lui data a n * Rispettivi a value di quella runa
func StampaOccorrenze(Occorrenze map[rune]int) {
    for k,v := range Occorrenze{
        fmt.Print(string(k),": ")
        for i:= 0 ; i<v ; i++{
            fmt.Print("*")
        }
      fmt.Println("")
    }
}
//Func che crea una mappa dove per ogni Rune calcola quante volte lo trova nella stringa presa in INPUT
func Occorrenze(s string) map[rune]int{
    o := make(map[rune]int)
    for _,x:= range s{
        if unicode.IsLetter(x){
            o[x]++
        }
        }
    return o
}

func main(){
var c map[rune]int
c = make(map[rune]int)
s := Leggitesto()
c = Occorrenze(s)
StampaOccorrenze(c)
}
