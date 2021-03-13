package main
import "fmt"
/*il programma dato una data di nascita ti dice se puoi votare ,votare camera,votare camera e senato */
func main(){
    var anno, annoCorrente int
    annoCorrente = 2020
    fmt.Print("Inserisci anno di nascita: ")
    fmt.Scan(&anno)
    if annoCorrente - anno < 18 {
    }else {
        if annoCorrente - anno <25 {
            fmt.Println("puoi votare solo camera")
        }else{
            fmt.Println("puoi votare senato e camera")
        }
     }
}
