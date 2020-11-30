package main
import "fmt"
/*il programma dato un numero ti dice se finisce per OOO */
func main(){
    var numero int
    fmt.Print("Inserisci numero: ")
    fmt.Scan(&numero)
    if numero <1000 {
        fmt.Println("Il numero inserito non finisce per OOO")
    }else {
        if numero % 1000 == 0 {
            fmt.Println("Il numero finisce per 000")
        }else{
            fmt.Println("Il numero inserito non finisce per OOO")
         }

    }

}
