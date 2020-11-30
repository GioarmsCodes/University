package main

import "fmt"

func main() {
    var n int
    fmt.Print("n :")
    fmt.Scan(&n)
    if n>1 {
        fmt.Println("2")
        for candidato := 3; candidato <= n; candidato+=2 {
            var divisore int
            for divisore = 3; divisore * divisore <= candidato ; divisore+=2 {
                if candidato  % divisore == 0 {
                    break
                }
            }
            if divisore * divisore > candidato{
                fmt.Println(candidato)
            }
        }
    }
}
