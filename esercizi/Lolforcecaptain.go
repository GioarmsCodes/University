package main

import(
    "fmt"
    "math/rand"
    "time"
)

func Newteam(all []string) ([]string, []string){
    var team1 []string
    var team2 []string
    var temp1 string
    var rep bool
    var m,l int
for i:=0 ; i<4 ; i++ {
    x:=rand.Intn(8)
    temp1 = all[x]
    for j:=0 ; j<len(team1);j++{
        if team1[j]==temp1 {
            rep=true
            i--
        }
    }
    if !rep{
        team1 = append(team1,all[x])
    }
    rep=false
}
    for m=0 ; m<8 ; m++{
        for l=0 ; l<4 ; l++{
            if all[m]==team1[l]{
                break
            }
        }
        if !(l<4){
            team2 = append(team2,all[m])
        }
    }
return team1,team2
}
func main() {
    rand.Seed(time.Now().UTC().UnixNano())

    var s,cpt1,cpt2 string
    var nomi []string
    fmt.Print("Inserisci nome del primo capitano: ")
    fmt.Scan(&cpt1)
    fmt.Print("Inserisci nome del secondo capitano: ")
    fmt.Scan(&cpt2)
for i:= 0 ; i<8;i++{
    fmt.Print("Inserisci nome: ")
    fmt.Scan(&s)
    nomi = append(nomi,s)

}
 t1,t2 := Newteam(nomi)
    fmt.Print("TEAM ",cpt1,": ",t1,"\nTEAM ",cpt2,": ",t2,"\n")

    c := rand.Intn(2)
    if c==0 {
        fmt.Println("Il team ",cpt1 ," ha il First Pick")
    }else{
        fmt.Println("Il team ",cpt2 ," ha il First Pick")
    }

}
