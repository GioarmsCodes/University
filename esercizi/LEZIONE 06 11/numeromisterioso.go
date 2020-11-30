package main
import (
    "bufio"
    "os"
    "fmt"
    "strconv"
)
func misterioso(s string) (n int,b bool) {
    var numero string
    numero=""
    for _,x:= range s {
        switch true {
        case x=='#':
        case  x> 47 &&  x< 58:
            numero = numero + string(x)
            b = true
        default:
            b = false
        }
    }
    n , _ = strconv.Atoi(numero)
    return n*2 , b
}

func main(){
    var n int
    var error bool
    b := bufio.NewScanner(os.Stdin)
    b.Scan()
    s := b.Text()
    fmt.Println(s)
    n , error = misterioso(s)
    if !error {
        fmt.Println("Hai inserito un valore errato")
    }else{
        fmt.Println(n)
    }
}
