package main

import (

    "fmt"
    "os"
    "strconv"
)

func GeneraNumeri (N,k int) []int{
    var variazioni []int
    s:= strconv.Itoa(N)

    for i:=0 ;i<len(s) ;i++{
        for j:=i+1 ;j<len(s) ; j++{
            if len(s[i:j+1]) == k {
                temp,_ := strconv.Atoi(s[i:j+1])
                variazioni = append(variazioni,temp)
            }
        }
    }
    return variazioni
}
func FiltraNumeri (num []int) []int{
    var Filtro []int
    for i:=0 ;i<len(num) ; i++{
        if num[i] % 2 ==-0 {
            Filtro = append(Filtro,num[i])
        }
    }
    return Filtro
}
func main(){
    s1 := os.Args[1]
    s2 := os.Args[2]
     N,_ :=strconv.Atoi(s1)
    n,_ :=strconv.Atoi(s2)

    test := GeneraNumeri(N,n)
    test = FiltraNumeri(test)
    fmt.Println(test)

}
