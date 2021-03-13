package main

import(
    "fmt"
    "persona"
)

type hasDate interface {
	Datadaoggi() int
}


func Contasotto(i int,l []hasDate) int{
    var cont int
    for _,x := range l {
        if x.Datadaoggi() <= i {
            cont++
        }
    }
    return cont
}
func main(){
    var lista []hasDate

    lista = append(lista,persona.Utente{"Armani","Islam",24,11,2001})
    lista = append(lista,persona.Utente{"Armani","Islam",31,12,2003})
    lista = append(lista,persona.Utente{"Armani","Islam",12,03,2005})
    lista = append(lista,persona.Utente{"Armani","Islam",16,04,2012})
    lista = append(lista,persona.Utente{"Armani","Islam",27,06,1998})
    lista = append(lista,persona.Utente{"Armani","Islam",01,02,2007})
    lista = append(lista,persona.Utente{"Armani","Islam",01,02,1970})

    fmt.Println(Contasotto(19,lista))
}
