package main

import "fmt"


func clessidra(n, spaz int) {
    if n == 1 {
        for i := 0 ; i< spaz ; i++{
            fmt.Print(" ")
        }
            fmt.Print("*")
        fmt.Print("\n")
        return
    }

    for i := 0; i < spaz; i++ {
        fmt.Print(" ")
    }
    for i := 0; i < (2 * n) - 1; i++ {
        fmt.Print("*")
    }
    fmt.Print("\n")

    clessidra(n - 1, spaz + 1)

    for i := 0; i < spaz; i++ {
        fmt.Print(" ")
    }
    for i := 0; i < (2 * n) - 1; i++ {
        fmt.Print("*")
    }
    fmt.Print("\n")
}



func main() {
    var n int
    fmt.Println("Inserisci numero: ")
    fmt.Scan(&n)

    clessidra(n,0)

}
