package main

import "fmt"

func main(){
    var n int
    var s string
    fmt.Print("Inserisci numero : ")
    fmt.Scan(&n)



    for i:=0 ; i<(n*2)+1 ;i++{

        for j:= 0 ; j<(n*2)+1 ; j++{
            if j == n || i == n {
                s+="* "
            }else{
                s+="  "
            }
        }
        fmt.Println(s)
        s = ""
    }

}
