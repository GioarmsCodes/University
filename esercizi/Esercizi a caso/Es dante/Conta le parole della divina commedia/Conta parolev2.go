//Programma che conta l' Occorrenza delle lettere dell'alfabeto
package main
import (
    "fmt"
    "os"
    "bufio"
)


func Leggifile(s string)(lettere map[string]int,err error){
    var temp string
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
        for _,x:= range l{
            if x!= ' '{
                temp +=string(x)
                }else{
                    lettere[temp]++
                    temp = ""
                }
            }
            lettere[temp]++
            temp = ""
        }
        return
    }



    func main() {
        var somma int
        parole,err := Leggifile(os.Args[1])
        if err != nil {
            fmt.Println("Errore durante l apertura del file")
        }
        for k,v := range parole{
            fmt.Println(k,": ",v)
            somma+=v
        }
        fmt.Println("Parole distinte: ",len(parole),"Parole totali: ",somma)
    }
