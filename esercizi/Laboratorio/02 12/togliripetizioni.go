package main


import  "fmt"

func main(){
    s:= []string{"ciao","ciao","come","stai","io","tutto","bene","bene"}

    test:= removeDuplicate(s)

    fmt.Println(test)
}
//Funzione che rimuove eventuali duplicati all interno di una slice di stringhe
func removeDuplicate(s []string) []string{
    var temp bool
    var test []string
    for i:=0;i<len(s);i++{
        for j:=0;j<i;j++{
            if i==j{
                continue
            }
            if s[i]==s[j]{
                temp = true
            }
        }
    if !temp{
        test = append(test,s[i])
    }
    temp = false
    }
    return test
}
