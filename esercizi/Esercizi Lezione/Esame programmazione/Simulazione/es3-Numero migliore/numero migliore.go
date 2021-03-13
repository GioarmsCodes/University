package main

import (
    "fmt"
    "strconv"
    "math"
    "os"
)

func RicMigliornumero (s string,dec int) string{
    var index int
    min,_ := strconv.Atoi(string(s[0]))

    if dec == 1 {

        for i:=0; i<len(s);i++{
            temp,_ := strconv.Atoi(string(s[i]))
            if temp<min{
                min = temp
            }
        }
        stest := strconv.Itoa(min)
        return stest
    }

    for i:=0; i<len(s)-dec+1;i++{
        temp,_ := strconv.Atoi(string(s[i]))
        if temp<min{
            min = temp
            index = i
        }
    }
    stest := strconv.Itoa(min)
    copyof := s[index+1:]
    return stest + RicMigliornumero(copyof,dec-1)
}

func main(){

    var n ,dec int

    fmt.Println("Inserisci numero: ")
    fmt.Scan(&n)
    fmt.Println("INSERISCI cifre da togliere: ")
    fmt.Scan(&dec)
    temp := strconv.Itoa(n)
    fmt.Println(temp)
    if len(temp)<dec {
        fmt.Println(math.NaN())
        os.Exit(1)
    }

    t := RicMigliornumero(temp,len(temp)-dec)
    fmt.Println(t)
}
