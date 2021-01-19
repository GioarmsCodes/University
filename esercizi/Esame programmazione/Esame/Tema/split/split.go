package main

import (
    "bufio"
    "os"
    "fmt"
    "strconv"
)

func main(){
    num,_:= strconv.Atoi(os.Args[1])
    nomefile := os.Args[2]

    f,_ := os.Open(nomefile)

    defer f.Close()

    b:= bufio.NewScanner(f)

    cont := 0
    taglio := 0
    for b.Scan(){

        if cont == 0 || cont%num == 0{
            fmt.Printf("::::::::::::::\ntaglio-%02d\n::::::::::::::\n",taglio)
            taglio++
        }
        s := b.Text()
        fmt.Println(s)
        cont++
    }
}
