//Programma che ti conta le lettere nel programma contando pure i punti le virgole

package main
import (
    "fmt"
    "os"
    "bufio"
)


func Leggifile(s string)(lettere map[rune]int,err error){
    lettere = make(map[rune]int)
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
        for _,x:= range l{
            lettere[x]++
        }
    }
    return
}



func main() {
    parole,err := Leggifile(os.Args[1])
    if err != nil {
        fmt.Println("Errore durante l apertura del file")
    }
    fmt.Println(parole)
    for k,v := range parole {
        fmt.Println(string(k),": ",v)
    }
}
