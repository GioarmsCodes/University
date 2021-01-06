package main

import (
    "fmt"
    "os"
     )

func main() {
    var n int
    var s string
    var r rune
    var flag bool
    fmt.Println("Inserisci numero : ")
    fmt.Scan(&n)

    if n<0 {
        fmt.Println("Altezza del rombo minore di zero.")
        os.Exit(1)
    }
    for i:=0 ; i < n; i++{
        r = 'A'
        r += rune(i)
        for m:=i ;m<n-1 ;m++{
            s += "  "
        }

        for j:=0 ; j<2*i+1 ; j++{
                    s+= " " +string(r)
            if r == 'A' || flag{
                r++
                flag = true
            }else{
                r--
            }
        }
        fmt.Println(s)
        s = ""
        flag = false

}
for i:=n-2 ; i>=0; i--{
    r = 'A'
    r += rune(i)
    for m:=i ;m<n-1 ;m++{
        s += "  "
    }

    for j:=2*i+1  ; j>0 ; j--{
        s+= " " +string(r)
        if r == 'A' || flag{
            r++
            flag = true
        }else{
            r--
        }
    }
    fmt.Println(s)
    s = ""
    flag = false

}
}
