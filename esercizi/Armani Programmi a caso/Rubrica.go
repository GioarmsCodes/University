package main
import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)
func main(){
    var rubrica map[string]int
    rubrica = make(map[string]int)
    //var rubricabyte []byte
    var nome string
    var numero int
    b := bufio.NewScanner(os.Stdin)
    b1 :=bufio.NewScanner(os.Stdin)
    fmt.Println("Inserire Nome[INVIO]/Inserire Numero[INVIO]")
    for b.Scan() && b1.Scan(){
        nome= b.Text()
        numero,_ = strconv.Atoi(b1.Text())
        fmt.Println(nome,numero)
        rubrica[nome] = numero
    }
    fmt.Println(rubrica)
    f,_ := os.Create("Rubrica.txt")
    defer f.Close()
    fmt.Fprintf(f,"Ciao")
    temp := fmt.Sprintf("%v",rubrica)
    f.Write([]byte(temp))
    fmt.Println("ciao")



}
