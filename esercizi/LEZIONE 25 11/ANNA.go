package main

import (
    "fmt"
    "math/rand"
    "time"
)
const max = 5
func Anna(max int) []rune{
    var parola []rune
    for i:=0 ; i<max ; i++ {
        x:=rand.Intn(26)
        parola = append(parola,rune(x+'a'))
    }
    return parola
}

func Check (s string,s2 string) (c bool){
    cp :=[]rune(s2)
    var i,j,q,pg,pb int
    pl := []rune(s)
    for q=0 ; q<len(cp) ;q++{
        if pl[q]!=cp[q]{
            break
        }
    }
    if !(q<len(cp)){
        return true
    }

    for i=0 ;i<len(cp) ; i++ {
        for j=0 ; j<len(cp);j++{
            if cp[i]==pl[j] && i==j{
                    pg++
                    cp[j]='*'
                    pl[j]='+'
                    break
                }
            }
        }
    fmt.Println("Lettere nel posto giusto: ",pg)
    for i=0 ;i<len(cp) ; i++ {
        for j=0 ; j<len(cp);j++{
            if cp[i]==pl[j]{
                    pb++
                    cp[i]='*'
                    break
                    }
                }
            }
            fmt.Println("Lettere giuste nel posto sbagliato: ",pb)
            pg=0
            pb=0
        return false
    }
    func main() {
        var s string
        var winorlose bool
        rand.Seed(time.Now().UTC().UnixNano())
        a:= Anna(max)
        for i:=1;;i++{
            fmt.Println("\nTENTATIVO [",i,"]: \n")
            fmt.Print("Inserisci stringa ",max," [caratteri]:  ")
            fmt.Scan(&s)
            if len(s)!=max{
                fmt.Println("hai inserito un stringa di numero diverso da ",max)
                continue
            }
            s2 := string(a)
            winorlose = Check(s,s2)
            if winorlose {
                fmt.Println("Hai indovinato!")
                break
            }

        }
    }
