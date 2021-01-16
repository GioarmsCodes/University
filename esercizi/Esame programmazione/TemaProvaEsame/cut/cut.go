package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
)

func main(){

    riga,_ := strconv.Atoi(os.Args[1])
    colonna,_ := strconv.Atoi(os.Args[2])
    nomefile := os.Args[3]

    f,err := os.Open(nomefile)
    if err != nil {
        fmt.Println("Errore durante l'apertura del file")
        os.Exit(1)
    }
    b := bufio.NewScanner(f)
    for b.Scan(){
        s := b.Text()

        for i:= 0 ; i<len(s) ; i++{
            if i>=riga-1 && i<=colonna-1{
                fmt.Print(string(s[i]))
            }

        }
        fmt.Println("")
    }
}
