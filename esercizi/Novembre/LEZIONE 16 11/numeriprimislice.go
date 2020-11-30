package main
import "fmt"
//Predicta la capacita' di un array avendo la lunghezza delle celle occupate
func Potenza2predict(n int) int {
    var i int
    for i=1;i<=n;i*=2{
        if n <= i {
            return i
        }
    }
    return i
}
// Funzione che controlla se un numero inserito e' primo
func Isprimo(numero int) bool{
    var primo bool
    var i int
        for i=2;i<=numero/2;i++{
            if numero%i==0{
                break
            }
        }
    if i>numero/2 {
        primo = true
        return primo
    }else{
        primo = false
        return primo
    }
}
//Funzione che crea una slice solo con i numeri primi fi a n [Inserito dell utente]
func Sliceprimi(nmax int) []int{
    var x []int
    for i:=2;i<=nmax;i++{
        if Isprimo(i){
            x=append(x,i)
        }
    }
    return x
}


func main() {
    var n,predict int
    var primi []int
    fmt.Print("Inserisci numero: ")
    fmt.Scan(&n)
    primi = Sliceprimi(n)
    predict = Potenza2predict(len(primi))
    fmt.Print(primi,"\nLunghezza: ",len(primi),"\nPredict: ",predict,"\nCapacita': ",cap(primi),"\n")
}
