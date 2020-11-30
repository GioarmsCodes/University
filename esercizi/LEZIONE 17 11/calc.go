package main
import (
    "fmt"
    "os"
    "strconv"
)

func Calcolo (s1 string,op string,s2 string) float64 {
    var t int
    var n1,n2 float64
    t,_=strconv.Atoi(s1)
    n1=float64(t)
    t,_=strconv.Atoi(s2)
    n2=float64(t)

    switch op {
    case "+":
        return n1+n2
    case "-":
        return n1-n2
    case "/":
        return n1/n2
    case "x":
        return n1*n2
    default:
        return 0
    }
}
func main() {
    fmt.Println(Calcolo(os.Args[1],os.Args[2],os.Args[3]))
    }
