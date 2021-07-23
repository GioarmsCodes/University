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
    fmt.Println("xased01--> ",nomi[x])
    x=rand.Intn(169)
    fmt.Println("Gioarm--> ",nomi[x])
    x=rand.Intn(169)
    fmt.Println("Alexynewgen--> ",nomi[x])
    x=rand.Intn(169)
    fmt.Println("Pippofendi--> ",nomi[x])
    x=rand.Intn(169)
    fmt.Println("Edofendi--> ",nomi[x])
    x=rand.Intn(169)
    fmt.Println("Simone--> ",nomi[x])




}
