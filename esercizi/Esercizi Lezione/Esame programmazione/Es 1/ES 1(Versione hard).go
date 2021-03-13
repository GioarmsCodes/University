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
    var n,resto,resto2,resto3 string
    var variazioni map[string]bool
    variazioni = make(map[string]bool)
    fmt.Println("Inserisci numero: ")
    fmt.Scan(&n)
        for i:=0;i<len(n);i++{
            resto=n[:i]+n[i+1:]
            variazioni[resto] = true
            for j:=0;j<len(resto);j++{
                resto2=resto[:j]+resto[j+1:]
                variazioni[resto2] = true
                for m:=0;m<len(resto2);m++{
                    resto3=resto2[:m]+resto2[m+1:]
                    variazioni[resto3] = true
                }
            }
        }
        fmt.Println(variazioni)
        for k,_ := range variazioni{
            num,_:=strconv.Atoi(k)
            if Eprimo(num){
                fmt.Println(num)
            }
        }

}
