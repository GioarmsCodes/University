package main

import(
    "fmt"
    "math/rand"
    "time"
    "bufio"
    "os"
)

func Newteam(all []string) ([]string, []string){
    var team1 []string
    var team2 []string
    var temp1 string
    var rep bool
    var m,l int
for i:=0 ; i<5 ; i++ {
    x:=rand.Intn(10)
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
    for m=0 ; m<10 ; m++{
        for l=0 ; l<5 ; l++{
            if all[m]==team1[l]{
                break
            }
        }
        if !(l<5){
            team2 = append(team2,all[m])
        }
    }
return team1,team2
}
func main() {
    rand.Seed(time.Now().UTC().UnixNano())
    var s string
    var nomi []string
    b := bufio.NewScanner(os.Stdin)
for b.Scan(){
    s = b.Text()
    nomi = append(nomi,s)
}
    t1,t2 := Newteam(nomi)
    fmt.Println("[TOP  ,JUNGLE  ,MID ,ADC , SUPP]")
    fmt.Println(t1)
    fmt.Println("[TOP  ,JUNGLE  ,MID ,ADC  , SUPP]")
    fmt.Println(t2)
    rand.Seed(time.Now().UTC().UnixNano())
    c := rand.Intn(2)
    if c==0 {
        fmt.Println("Il team 1 ha il First Pick")
    }else{
        fmt.Println("\n\nIl team 2 ha il First Pick")
    }


}
