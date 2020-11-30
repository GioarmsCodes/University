package main
import "fmt"
        func main(){
            var n int
            var s string
            fmt.Print("Inserisci parola : ") //  inserisci parola
            fmt.Scan(&s)
            fmt.Print("Inserisci traslazione : ") // inserisci numero traslazione
            fmt.Scan(&n)
            fmt.Print(s[n:],s[:n])  // stranda la stringa da n fino alla fine e da inizio fino a n
            fmt.Println("") // a capo
        }
