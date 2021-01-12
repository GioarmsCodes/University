package main

import (
    "fmt"
    "os"
)


func main(){
    stringa := os.Args[1]
    var occorrenze map[string]int
    var nomi []string
    var numeri []int
    var imin,min int
    occorrenze = make(map[string]int)
    for i:=0 ; i<len(stringa) ; i++{
        for j:= i+1 ; j<len(stringa) ; j++{
            if stringa[i] == stringa[j] && len(stringa[i:j+1])>=3{
                occorrenze[stringa[i:j+1]]++
            }
        }
    }

    // metto i valori della mappa in 2 stringhe
    for k,v := range occorrenze{
        nomi = append(nomi,k)
        numeri = append(numeri,v)
    }

    //sortiing nomi in base alla granddezza
    for i:=0 ;i<len(nomi) ; i++{
        imin = i
        min = len(nomi[i])
        for j:=i+1 ; j<len(nomi) ; j++{
            if len(nomi[j])<min{
                min = len(nomi[j])
                imin = j
            }
        }
        nomi[i] , nomi[imin] = nomi[imin] , nomi[i]
        numeri[i] , numeri[imin] = numeri[imin] , numeri[i]

    }

    for i:=len(nomi)-1 ; i>=0 ; i--{
        fmt.Println(nomi[i],"->","Occorrenze:",numeri[i])
    }
}
