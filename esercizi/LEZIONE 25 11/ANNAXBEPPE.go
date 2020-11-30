package main
import (
    "fmt"
    "math/rand"
    "time"
)

var volte = 0

func predict(i int,lunghezza int,s string) string{
    var correct bool
    var t rune
    if i==lunghezza-1{
        for c:='a'; c<='z';c++{
            t=c
            volte++
            correct = Check(s,c,i)
            if correct{
                break
            }
        }
        return string(t)
    }
    for c:='a'; c<='z';c++{
        t=c
        volte++
        correct = Check(s,c,i)
        if correct{
            break
        }
    }
    return string(t)+predict(i+1,lunghezza,s)
}

func Check(seq []rune, parola []rune) (int, int) {
    var giuste int = 0
    var pos_sbagliati int = 0
    var temp1 []rune
    var temp2 []rune
    for x := 0; x < len(seq); x++ {
        if seq[x] == parola[x] {
            giuste++
        } else {
            temp1 = append(temp1, seq[x])
            temp2 = append(temp2, parola[x])
        }
    }
    for x := 0; x < len(temp1); x++ {
        for y := 0; y < len(temp2); y++ {
            if temp1[x] == temp2[y] {
                pos_sbagliati++
                temp2[y] = 999
                break
            }
        }
    }
    return giuste, pos_sbagliati
}

const max = 5

func Anna(max int) string{
    var s string
    var parola []rune
    for i:=0 ; i<max ; i++ {
        x:=rand.Intn(26)
        parola = append(parola,rune(x+'a'))
    }
    for i:=0;i<len(parola);i++{
        s=s+string(parola[i])
    }
    return s
}

    func main() {
        var s string
        rand.Seed(time.Now().UTC().UnixNano())
        //s = Anna(max)
        s="zzzzz"
        fmt.Println("A: ho creato la parola --> ",s)
        fmt.Println("B: La indovino subito!")

        test := predict(0,max,s)
        fmt.Println(test,volte)

    }
