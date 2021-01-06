package main

import (
    "fmt"
    "strings"
    "bufio"
    "os"
)

func Inverti(s string) (ret string){
    for i:=len(s)-1 ;i>=0 ; i--{
        ret +=string(s[i])
    }
    return
}

func main(){
    var basi,final,invl []string
    var s string
    var cambiamento bool

    b := bufio.NewReader(os.Stdin)
    q := bufio.NewReader(os.Stdin)
fmt.Println("Inserisci stringa base: ")
    base,_ := b.ReadString('\n')

    basi = strings.Split(base[:len(base)-2]," ")
fmt.Println("Inserisci stringa da girare: ")
    inv,_ := q.ReadString('\n')

    invl = strings.Split(inv[:len(inv)-2]," ")

for i:=0 ;i<len(basi) ;i++{
    for j:=0;j<len(invl) ;j++{
        if basi[i] == invl[j] {
            s = Inverti(basi[i])
            cambiamento = true
        }
    }
    if cambiamento{
        final = append(final,s)
        cambiamento = false
    }else{
        final = append(final,basi[i])
    }
}

    fmt.Println(final)
}
