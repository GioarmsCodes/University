package main

import (
    "fmt"
   "strconv"
)

func Eprimo(n int) bool {
    for i:=2 ;i<n;i++{
        if n%i == 0{
            return false
        }
    }
    return true
}

func main(){
    var n,resto string
    var variazioni []string
    fmt.Println("Inserisci numero: ")
    fmt.Scan(&n)
    variazioni=append(variazioni,n)
        for i:=0;i<len(n);i++{
            resto=n[:i]+n[i+1:]
            variazioni=append(variazioni,resto)
        }
        for i:=0;i<len(n)-1;i++{
            resto=n[:i]+n[i+2:]
            variazioni=append(variazioni,resto)
        }
        for i:=0;i<len(n)-2;i++{
            resto=n[:i]+n[i+3:]
            variazioni=append(variazioni,resto)
        }

        fmt.Println(variazioni)
        for _,x := range variazioni{
            num,_:=strconv.Atoi(x)
            if Eprimo(num){
                fmt.Println(num)
            }
        }

}
