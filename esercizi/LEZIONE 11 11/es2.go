package main
import (
    "fmt"
    "strconv"
)
type data struct{
    g,m,a int
}
func Modificadata(data *data, s string, x int) {
    if s == "g"{
        (*data).g=x
    }else if s == "m"{
    (*data).g=x
}else if s == "a"{
        (*data).a=x
        }
}
func main(){
    var d data
    var x int
    var s string
    d.g , d.m , d.a = 24,11,2001
    fmt.Print("Inserisci stringa: ")
    fmt.Scan(&s)
    fmt.Print("Inserisci numero: ")
    fmt.Scan(&x)
    Modificadata(&d,s,x)
    fmt.Print(strconv.Itoa(d.g),"-",strconv.Itoa(d.m),"-",strconv.Itoa(d.a) )
    fmt.Println()
}
