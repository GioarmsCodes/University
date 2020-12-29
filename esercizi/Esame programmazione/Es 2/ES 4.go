
package main

import (
    "fmt"
    "os"
)

func main(){

    s := os.Args[1]
    var variazioni map[string]int

    var key []string
    var value []int
    variazioni = make(map[string]int)

    for i:=0 ; i<len(s) ; i++{
        for j:=i+1 ; j<len(s) ; j++{
            if s[i] == s[j]{
                if len(s[i:j+1]) >= 3{
                variazioni[s[i:j+1]]++
            }
            }
        }
    }

    for k,v := range variazioni {
        key = append(key,k)
        value = append(value,v)
    }

    for i := 0 ; i < len(key) ; i++{
        imin := i
        min := len(key[i])
        for j:=i+1 ; j<len(key) ; j++{
            if len(key[j]) < min {
                imin = j
                min = len(key[j])
            }
        }
        key[i],key[imin] =key[imin],key[i]
        value[i],value[imin] =value[imin],value[i]
    }
    for i:=len(value) -1 ;i>=0 ; i--{
        fmt.Println(key [i],"->",value[i])
}
}
