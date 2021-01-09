//Programma che cerca di creare 2 team bilanciati dato 10 nomi:value da un file

package main

import (

    "fmt"
    "bufio"
    "os"
    "strconv"
)

func Ord(n []int) {
    var min,mini int

    for i:=0 ; i<len(n) ; i++{
    min = n[i]
    mini = i
        for j:=i ; j<len(n); j++{
            if n[j]<min{

                mini = j
                min = n[j]
            }
        }
                        n[i] , n[mini] = n[mini] , n[i]
}
}

func main() {
    var numeri []int
    var team1 []string
    var team2 []string

    var giocatori map[string]int

    giocatori = make(map[string]int)
    f,_ := os.Open("giocatori.txt")
    b := bufio.NewScanner(f)

    for b.Scan(){

        s:= b.Text()


        for i,x :=  range s {
            if x == ':' {
                temp ,_  := strconv.Atoi(s[i+1:])
                giocatori[s[:i]] = temp
            }
        }
    }

    for _,v := range giocatori {
        numeri = append(numeri,v)
    }

    Ord(numeri)

    for i:= 0 ;i<len(numeri) ;i++{
        if  i == 0 || i == 3 || i == 4 || i == 7 || i == 8{
            for k,v:= range giocatori {
                if v == numeri[i] {
                    team1 = append(team1,k)
                    giocatori[k] = 0
                    break
                }
            }
        }
    }
    for k,v := range giocatori {
        if v!= 0 {
        team2 = append(team2,k)
    }
    }
    fmt.Println(team1)
    fmt.Println(team2)
}
