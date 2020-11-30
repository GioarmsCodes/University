package main
import "fmt"
/*epoch   01/01/1970    */
// si legge una data e si calcola quanti giorni sono trascorsi dall epoch
func main(){
    var g, m, a, anno, mese,gg int
    anno = 1970
    mese = 1
    gg = 0
    fmt.Print(" g  m  a : ")
    fmt.Scan(&g)
    fmt.Scan(&m)
    fmt.Scan(&a)
    if a >= anno{
        anno=1970
        gg+=g
    }else {
        a,anno=anno,a
        mese=m
        m=12
        if mese ==  11 || mese ==  4 || mese ==  6 || mese ==  9 {
            giorni += 31 - g
        } else if mese==2{
            if anno%400 == 0{
                giorni += 30 - g
            }else if anno % 4 == 0 && anno % 100 == 1{
                giorni += 29 - g
                }

        }
    }

    for ;anno<a; anno++ {
        gg += 365
        if anno%400 == 0{
            gg++
        }else if anno % 4 == 0 && anno % 100 == 1{
            gg++
            }
    }
    for ; mese < m; mese++{
        if mese==1 || mese==4 || mese ==6 || mese==9{
            gg+=30
        }else if mese==2{
            gg+=28
            if a%400 == 0{
                gg++
            }else if a % 4 == 0 && a % 100 == 1{
                gg++
                }
        }else {
            g += 31
        }
    }
    gg+=g

    fmt.Print("i giorni trascorsi da epoch sono: ",gg,"\n")
    }
