package main

import (
    "fmt"
    "strconv"
    "math"
    "os"
)
func main(){
    eq := os.Args[1][2:]

    var elemento string
    var elementi []string


    for i,char := range eq{

        if (char=='+' || char=='-') && i != 0{
            elementi = append(elementi,elemento)
            elemento = ""
        }
            elemento += string(char)

    }
    elementi = append(elementi,elemento)
    a,_ := strconv.Atoi(elementi[2][:len(elementi[2])-3])
    b,_ := strconv.Atoi(elementi[1][:len(elementi[1])-1])
    c,_ := strconv.Atoi(elementi[0])

    fmt.Println(a,b,c)
    delta := b*b - 4*a*c

    if delta < 0{
        fmt.Printf("non ci sono soluzioni reali\n")
    }else if delta == 0{
        soluzione := -float64(b)/ float64(2*a)
        fmt.Printf("Esiste un'unica soluzione reale: %.3f\n",soluzione)
    }else{
        x1 := (-float64(b) + math.Sqrt(float64(b*b - 4*a*c))) / float64(2*a)

        x2 := (-float64(b) - math.Sqrt(float64(b*b - 4*a*c))) / float64(2*a)

        fmt.Printf("Ci sono due soluzioni reali: %.3f e %.3f\n",x1,x2)
    }


}
