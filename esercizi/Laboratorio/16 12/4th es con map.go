package main

import (
    "fmt"
    "math/rand"
    "time"
)

type CartellaTombola map[int]bool

func String(c CartellaTombola) string{

    return fmt.Sprintf("%v",c)
}

func Creacartella() CartellaTombola{
    var c CartellaTombola
    c = make(CartellaTombola)
    var temp int
    for i:=0 ;i<5;i++{
        temp = rand.Intn(30) + 1
        for k,_ := range c{
            if temp == k{
                i--
                continue
            }
        }
        c[temp]= false
    }
    for i:=5 ;i<10;i++{
        temp = 30+ rand.Intn(30) + 1
        for k,_ := range c{
            if temp == k{
                i--
                continue
            }
        }
        c[temp] = false
    }
    for i:=10 ;i<15;i++{
        temp = 60 + rand.Intn(30) + 1
        for k,_ := range c{
            if temp == k{
                i--
                continue
            }
        }
        c[temp] = false
    }
 return c
}

func (c CartellaTombola) Presente(n int) bool{
    for k,_ := range c{
        if n==k{
            return true
        }
    }
    return false
}

func Copri(c CartellaTombola,n int) bool{
    for k,v := range c{
        if n==k && v == false{
            c[k] = true
            return false
        }
    }
    return true //non ha trovato il numero -->Error
}

func Scopri(c CartellaTombola,n int) bool{
    for k,v := range c{
        if n==k && v == true{
            c[k] = false
            return false
        }

    }
    return true //non ha trovato il numero o la carta era gia' scoperta --> Error
}

func (c CartellaTombola) Coperto(n int) bool{
    for k,v := range c{
        if n==k && v == true{
            return true
        }
    }
    return false //non e' coperto quindi false o non lo ha trovato
}

func main(){
    rand.Seed(time.Now().UnixNano())
    var s CartellaTombola
    s = Creacartella()
    test := s.Presente(2) //Funziona
    test1 := Copri(s,2)
    test2 := s.Coperto(2)
    test3 := Scopri(s,2)
    fmt.Println(s,test,test1,test2,test3)

}
