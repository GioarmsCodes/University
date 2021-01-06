package main



import (
    "fmt"
    "os"
    "bufio"
    "strconv"
)

type giocatore struct {
    nome string
    n int
}

func Sortgiocatori (p []giocatore){

    for i:=0 ;i<len(p);i++{
        min := p[i].n
        mini := i
        for m:=i+1;m<len(p);m++{
            if p[m].n < min{
                min = p[m].n
                mini = m
            }
        }
        p[i],p[mini] = p[mini],p[i]
    }
}

func main() {
    var play giocatore
    var giocatori []giocatore
    b := bufio.NewScanner(os.Stdin)

    for b.Scan(){

        s:= b.Text()

        for i,x:= range s{
            if x==' '{
                play.nome = s[:i]
                temp,_ := strconv.Atoi(s[i+1:])
                play.n = temp
                break
            }
        }
                    giocatori = append(giocatori,play)
    }
    fmt.Println(giocatori)

    Sortgiocatori(giocatori)
    for i:=len(giocatori) - 1;i>=0 ;i--{
        fmt.Println(giocatori[i].nome ,"-",giocatori[i].n)
    }
}
