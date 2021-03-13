//Programma che conta l' Occorrenza delle lettere dell'alfabeto 
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
            for i:='a';i<='z';i++{
            if i==x{
                lettere[x]++
                break
            }
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
    for i:='a';i<='z';i++{
        fmt.Println(string(i),": ",parole[i])
    }
}
