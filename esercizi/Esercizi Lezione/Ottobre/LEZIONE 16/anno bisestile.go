package main
import "fmt"
/*il programma dato una data.anno calcola se 'e un anno bisestile*/
func main(){
    var anno int
    fmt.Print("Inserisci anno corrente:  ")
    fmt.Scan(&anno)
    if anno%400 == 0{
        fmt.Println("L anno e' bisestile")
    }else if anno % 4 == 0 && anno % 100 != 0{
        fmt.Println("L anno e' bisestile")
        }else {
            fmt.Println("L anno non e' bisestile")
        }
}
