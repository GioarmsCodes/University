//Programma che conta l' Occorrenza delle lettere dell'alfabeto
package main
import (
    "fmt"
    "os"
    "bufio"
)


func Leggifile(s string)(lettere map[string]int,err error){
    var prima int
    lettere = make(map[string]int)
    p,e := os.Open(s)
    if e != nil{
        err = e
        return
    }
    defer p.Close()
    b := bufio.NewScanner(p)

    for b.Scan(){
        l := b.Text()
        if len(l) == 0{
            continue
        }
        prima = 0
        for i,x:= range l{
            if x == ' ' {
                temp:=l[prima:i]
                lettere[string(temp)]++
                prima = i+1
            }
        }
    }
    return
}



func main() {
    parole,err := Leggifile(os.Args[1])
    if err != nil {
        fmt.Println("Errore durante l apertura del file")
    }
    for k,v := range parole{
        fmt.Println(k,": ",v)
    }
    fmt.Println(len(parole))
}
