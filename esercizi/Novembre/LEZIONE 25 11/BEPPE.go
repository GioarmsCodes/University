package main
import "fmt"


func predict(i int,lunghezza int,s string) string{
    var correct bool
    var t rune
    if i==lunghezza-1{
        for c:='a'; c<='z';c++{
            t=c
            correct = Check(s,c,i)
            if correct{
                break
            }
        }
        return string(t)
    }
    for c:='a'; c<='z';c++{
        t=c
        correct = Check(s,c,i)
        if correct{
            break
        }
    }
    return string(t)+predict(i+1,lunghezza,s)
}

func Check (s string,r2 rune,cont int) (c bool){
    sr := []rune(s)
    if sr[cont] == r2 {
        return true
    }else{
        return false
    }
}

func main(){
    var s string
    fmt.Print("Inserisci stringa: ")
    fmt.Scan(&s)
    test := predict(0,len(s),s)
    fmt.Println(test)
    }
