package main

import (
    "fmt"
    "unicode"
    "os"
    "bufio"
    "strings"
    "strconv"
)

func cancella(lista []string) []string{
    var copia []string
    for i := 0 ; i < len(lista) ; i++{
        if unicode.IsNumber(rune(lista[i][0])){
            temp,_ := strconv.Atoi(string(lista[i]))
            i += temp
            continue
        }else{
            copia = append(copia,lista[i])
        }
    }
    return copia
}

func main(){
    var parole []string
    var l string
    if len(os.Args) != 2{
        fmt.Println("Fornire 1 nome di file")
        return
    }

    s := os.Args[1]
    f,err := os.Open(s)
    if err != nil{
        fmt.Println("File non accessibile")
        return
    }

    b:= bufio.NewScanner(f)
        for b.Scan(){
            l += b.Text() + " "

        }

        parole = strings.Split(l[:len(l)-1]," ")
        fmt.Println(parole)
        dummy := cancella(parole)
        fmt.Println(dummy)
}
