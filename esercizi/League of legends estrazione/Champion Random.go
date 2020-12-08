package main

import(
    "fmt"
    "math/rand"
    "time"
    "bufio"
    "os"
)


func main() {
    rand.Seed(time.Now().UTC().UnixNano())
    var s string
    var nomi []string
    b := bufio.NewScanner(os.Stdin)
for b.Scan(){
    s = b.Text()
    nomi = append(nomi,s)
}

    x:=rand.Intn(169)
    fmt.Println("Davilk--> ",nomi[x])
    x=rand.Intn(169)
    fmt.Println("Gioarm--> ",nomi[x])
    x=rand.Intn(169)
    fmt.Println("Sajjad--> ",nomi[x])
    x=rand.Intn(169)
    fmt.Println("Raiden--> ",nomi[x])
    x=rand.Intn(169)
    fmt.Println("Tokaji--> ",nomi[x])
    x=rand.Intn(169)
    fmt.Println("Jacopon--> ",nomi[x])
    x=rand.Intn(169)
    fmt.Println("Fahim--> ",nomi[x])
    x=rand.Intn(169)
    fmt.Println("Hiro--> ",nomi[x])
    x=rand.Intn(169)
    fmt.Println("Alphyie--> ",nomi[x])
    x=rand.Intn(169)
    fmt.Println("Simonoggia--> ",nomi[x])




}
