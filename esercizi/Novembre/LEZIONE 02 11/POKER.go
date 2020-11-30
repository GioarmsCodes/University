package main
import "fmt"
/* Carte poker */
func main(){
    var numero,x int
    fmt.Println("INSERISCI NUMERO: ")
    fmt.Scan(&numero)
    if numero < 13{
        x=numero
        }else if numero%13 == 0{
            x=0
            }else{
                x=numero%13
            }
            switch x{
            case 1:
                fmt.Print("A")
            case 10:
                fmt.Print("10")
            case 11:
                fmt.Print("J")
            case 12:
                fmt.Print("Q")
            case 0:
                fmt.Print("K")
            default:
                fmt.Print(x)
            }

            if numero>0 && numero <14 {
                fmt.Print("♡")
            }
            if numero>13 && numero <26 {
                fmt.Print("♢")
            }
            if numero>25 && numero <39 {
                fmt.Print("♣")
            }
            if numero>38 && numero <53 {
                fmt.Print("♤")
            }
        }
