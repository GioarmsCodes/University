package main

import (
    "fmt"
    "math/rand"
    "time"
)

type CartellaTombola struct{
    n [15]int;
    v [15]bool;
}

func String(c CartellaTombola) string{

    return fmt.Sprintf("%v",c.n)
}

func Creacartella() CartellaTombola{
    var c CartellaTombola
    var temp int
    for i:=0 ;i<5;i++{
        temp = rand.Intn(30) + 1
        for j:=0;j<i;j++{
            if temp == c.n[j]{
                i--
                continue
            }
        }
        c.n[i] = temp
    }
    for i:=5 ;i<10;i++{
        temp = 30+ rand.Intn(30) + 1
        for j:=5;j<i;j++{
            if temp == c.n[j]{
                i--
                continue
            }
        }
        c.n[i] = temp
    }
    for i:=10 ;i<15;i++{
        temp = 60 + rand.Intn(30) + 1
        for j:=10;j<i;j++{
            if temp == c.n[j]{
                i--
                continue
            }
        }
        c.n[i] = temp
    }
 return c
}

func (c CartellaTombola) Presente(n int) bool{
    for i:=0 ;i<len(c.n);i++{
        if n==c.n[i]{
            return true
        }
    }
    return false
}

func Copri(c *CartellaTombola,n int) bool{
    for i:=0 ;i<15;i++{
        if n==c.n[i]{
            c.v[i] = true
            return false
        }
    }
    return true //non ha trovato il numero -->Error
}

func Scopri(c *CartellaTombola,n int) bool{
    for i:=0 ;i<len(c.n);i++{
        if n==c.n[i] && c.v[i] == true{
            c.v[i] = false
            return false
        }

    }
    return true //non ha trovato il numero o la carta era gia' scoperta --> Error
}

func (c CartellaTombola) Coperto(n int) bool{
    for i:=0 ;i<len(c.n);i++{
        if n==c.n[i] && c.v[i] == true{
            return true
        }
    }
    return false //non e' coperto quindi false
}
func main(){
    rand.Seed(time.Now().UnixNano())
    var s CartellaTombola
    s = Creacartella()
    //test := s.Presente(s.n[3]) //Funziona
    test1 := Copri(&s,s.n[3])
    //fmt.Println(s.v)
    test2 := s.Coperto(s.n[3])
    test3 := Scopri(&s,s.n[3])
    fmt.Println(s.n[3],"-->",s.n,"\n",s.v,test1,test2,test3)

}
