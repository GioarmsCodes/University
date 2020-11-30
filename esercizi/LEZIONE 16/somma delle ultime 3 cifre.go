package main
import "fmt"
/*il programma dato un numero da 3 cifre mi dice se la somma delle 3 cifre e' maggiore o minore di 10*/
func main(){
    var numero,somma int
    somma = 0
    fmt.Print("Inserisci numero da 3 cifre (es: 123):  ")
    fmt.Scan(&numero)
    if numero > 99 {
        if numero % 10 == 0 {
            somma+=0
        }
        if numero % 10 == 1 {
            somma+=1
        }
        if numero % 10 == 2 {
            somma+=2
        }
        if numero % 10 == 3 {
            somma+=3
        }
        if numero % 10 == 4 {
            somma+=4
        }
        if numero % 10 == 5 {
            somma+=5
        }
        if numero % 10 == 6 {
            somma+=6
        }
        if numero % 10 == 7 {
            somma+=7
        }
        if numero % 10 == 8 {
            somma+=8
        }
        if numero % 10 == 9 {
            somma+=9
        }

        numero = (numero -(numero % 10)) / 10

        if numero % 10 == 0 {
            somma+=0
        }
        if numero % 10 == 1 {
            somma+=1
        }
        if numero % 10 == 2 {
            somma+=2
        }
        if numero % 10 == 3 {
            somma+=3
        }
        if numero % 10 == 4 {
            somma+=4
        }
        if numero % 10 == 5 {
            somma+=5
        }
        if numero % 10 == 6 {
            somma+=6
        }
        if numero % 10 == 7 {
            somma+=7
        }
        if numero % 10 == 8 {
            somma+=8
        }
        if numero % 10 == 9 {
            somma+=9
        }

        numero = (numero -(numero % 10)) / 10

        if numero % 10 == 0 {
            somma+=0
        }
        if numero % 10 == 1 {
            somma+=1
        }
        if numero % 10 == 2 {
            somma+=2
        }
        if numero % 10 == 3 {
            somma+=3
        }
        if numero % 10 == 4 {
            somma+=4
        }
        if numero % 10 == 5 {
            somma+=5
        }
        if numero % 10 == 6 {
            somma+=6
        }
        if numero % 10 == 7 {
            somma+=7
        }
        if numero % 10 == 8 {
            somma+=8
        }
        if numero % 10 == 9 {
            somma+=9
        }
    }
    fmt.Println(somma)
    if somma > 10 {
        fmt.Println("La somma e' maggiore di 10")
    }else if somma < 10 {
        fmt.Println("La somma e' minore di 10")
     }else {
         fmt.Println("La somma e' 10")
     }
}
