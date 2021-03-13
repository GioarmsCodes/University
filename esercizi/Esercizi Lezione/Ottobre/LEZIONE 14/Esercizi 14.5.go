package main
import "fmt"
/*il programma dato un numero mi dice se finisce con 3 */
func main(){
    var n int
    fmt.Println("inserisci numero: ")
    fmt.Scan(&n)
    if (n<10){
        if (n==3){
            fmt.Println("il numero finisce con 3")
        }else {
            fmt.Println("il numero non finisce con 3")

        }
    }else {
        n=n-3
        if(n % 10 == 0){
            fmt.Println("il numero finisce con 3")
        }else {
            fmt.Println("il numero non finisce con 3")
        }
    }
}
