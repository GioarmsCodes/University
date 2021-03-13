package main
import "fmt"
/*il programma mi dice se nelle ultime 3 cifre c'e' uno zero*/
func main(){
    var numero,c int
    c=0
    fmt.Print("Inserisci numero Coglione: ")
    fmt.Scan(&numero)

    if numero%10==0{
        fmt.Println("il numero contiene uno Zero(0) nell' ultima cifra")
        c+=1
    }
    numero/=10
    if numero%10==0{
        fmt.Println("il numero contiene uno Zero(0) nella penultima cifra")
            c+=1
    }
    numero/=10
    if numero%10==0 {
        fmt.Println("il numero contiene uno Zero(0) nella terzultima cifra")
     }
    if numero%10!=0 && c==0{
         fmt.Println("il numero non contiene uno Zero(0) nelle ultime 3 cifre")
      }
}
