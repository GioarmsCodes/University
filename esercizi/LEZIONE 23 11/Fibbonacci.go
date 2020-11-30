package main

import "fmt"

func fibbonacci(n int) int{
    if cache[n-1]!= 0{
        return cache[n-1]
    }
    if n==1 || n==2{
        return 1
    }
    cache[n-1]= fibbonacci(n-1)+fibbonacci(n-2)
    return cache[n-1]

}

var cache []int
func main() {
    var n int
    fmt.Print("Inserisci numero :")
    fmt.Scan(&n)
    cache=make([]int,n)
    for i:=1 ; i<=n ; i++{
        x:= fibbonacci(i)
        fmt.Println(x)
    }
    for i:=1 ; i<n ; i++ {
        if i>2{
        c:= float64(cache[i-1])/float64(cache[i])
        fmt.Println(c)
        fmt.Println(cache[i])
    }
    }
}
