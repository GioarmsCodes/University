package main

import (
    "fmt"
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
    var n,l,h,serp,cont int
    var spazi string
    fmt.Println("Inserisci numero: ")
    fmt.Scan(&n)
    fmt.Scan(&l)
    fmt.Scan(&h)
    
    for i:=0;i<l-1;i++{
        spazi +=" "
    }

for {
    for i:=0;i<l;i++{
        fmt.Print(serp)
        if serp == 9{
            serp =0
        }else{
        serp++
    }
    }

    fmt.Println()
    cont++
    if cont == n+1{
        break
    }
    for i:=0;i<h-2;i++{
        fmt.Print(spazi,serp,"\n")
        if serp == 9{
            serp =0
        }else{
        serp++
    }
    }
    temp :=serp -1
    for i:=l;i>0;i--{
        fmt.Print(temp+i)
        if serp == 9{
            serp =0
        }else{
        serp++
    }
    }
    fmt.Println()
    cont++
    if cont == n+1{
        break
    }
    for i:=0;i<h-2;i++{
        fmt.Print(serp,spazi,"\n")
        if serp == 9{
            serp =0
        }else{
        serp++
    }
    }


}


}
