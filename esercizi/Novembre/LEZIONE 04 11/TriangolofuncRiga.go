package main
import "fmt"
/* Triangolo con funzione stampa riga*/
func Spazio(n int){
}
func Stampariga(n int, s int) {
    var i int
    for i=0; i<n ; i++ { //stampa gli spazi
        fmt.Print(" ")
    }
    for j:=0; j<s ; j++ { //stampa gli asterischi
        fmt.Print("*")
    }
    fmt.Println("")
}
func main(){
    var numero,c,a,i int
    fmt.Print("Inserisci numero: ")
    fmt.Scan(&numero)
    for i=0 ; i<numero ; i++ { //le righe del triangolo
        c=0
        a=0
        for m:=0; m<i ; m++ { //conta gli spazi
            c++
            }
        for j:=0 ; j<numero-i ; j++ { // conta gli asterischi
            a++
        }
        Stampariga(c, a)
    }
}
