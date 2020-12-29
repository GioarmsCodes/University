package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
)

func main(){
    var flag bool
    var banca map[string]int
    banca = make(map[string]int)
    f,err := os.Open(os.Args[1])
    var anno,mese,giorno,all,tipo string
    var indextemp,importo int

    if err != nil {
        fmt.Println("Errore durante l'apertura del file")
        os.Exit(1)
    }
    b := bufio.NewScanner(f)
    flag = false
    for b.Scan(){
        s:= b.Text()

        for i,x:= range s{

            if x=='_' && flag == false{
                indextemp = i + 1
                anno = s[:i]
                flag = true
                continue
            }

            if x=='_' && flag == true{
                mese = s[indextemp:i]
                if len(mese) == 1{
                    mese = "0" + mese
                }
                indextemp = i +1
                flag = false
                continue
            }
            if x==';' && flag == false{
                giorno = s[indextemp:i]
                if len(giorno) == 1{
                    giorno = "0" + giorno
                }
                indextemp = i +1
                flag = true
                continue
            }
            if x==';' && flag == true {
                tipo = s[indextemp:i]
                indextemp = i +1
                continue
        }
        importo,_= strconv.Atoi(s[indextemp:])
        }

        all = anno + mese + giorno

        if tipo == "V" {
            banca[all] += importo
        }else if tipo == "P"{
            banca[all] -= importo
        }
    indextemp = 0
    flag = false
    }
    for k,v := range banca{
        fmt.Println(k[:3],"/",k[4:6],"/",k[6:8],"-->",v)
    }
}
