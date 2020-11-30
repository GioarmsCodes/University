package main
import "fmt"
func main(){
    var numero, m, i, c int
    c=0
    fmt.Print("Inserisci numero : ")
    fmt.Scan(&numero)
    for{
        for i=2; ; i++ {

            for m=2; m < i ; m++{

                if i % m == 0 {
                    break
                }
            }
            if !(m < i ) {
                c++
                fmt.Println(i)
            }
            if c == numero {
                return
            }
        }

    }
}
