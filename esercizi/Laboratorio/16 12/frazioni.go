package main

import (
	"fmt"
	"os"
	"strconv"
)
type Frazione struct {num, den int}

//definire le seguenti funzioni
 //       ---------Fatto-------------
func CreaFrazione(n,d int) Frazione {
   var x Frazione
	x.num =n
	x.den =d
	return x
}
// ----------------Fatto-----------------------
func CreaFrazioneParticolare (n int) Frazione {
	return CreaFrazione(n, 1)
}

//                              ---------Fatto---------
//CreaFrazioneDaStringa crea una Frazione da una stringa come ad es "5/7"
//ritorna la coppia ("",false) se la stringa s non è nel formato attesa
func CreaFrazioneDaStringa(s string) (Frazione,bool) {
	var salvaindice int
	var x Frazione
	var ok bool
	var err,err2 error
  for i,x := range s {
	  if x == '/'{
		  salvaindice = i
		  break
	  }
  }
  x.num,err = strconv.Atoi(s[:salvaindice])
  x.den,err2 = strconv.Atoi(s[salvaindice+1:])
  if salvaindice == 0 || err != nil || err2 != nil{
	  x.num = 0
	  x.den = 0
	  ok = false
	  return x,ok
  }
  ok = true
  return x, ok
}

// ----------------Fatto------------------------
func (f1 Frazione) Somma (f2 Frazione) Frazione {
	var somma Frazione
	somma.den = f1.den * f2.den
	somma.num = f1.num*f2.den + f2.num*f1.den
	return somma
}

//Sommatoria calcola (ricorsivamente) la Frazione che risulta sommando
//le frazioni passate come argomento

func Sommatoria (sommafrazioni []Frazione) Frazione {
	if len(sommafrazioni) == 0 {
		 return CreaFrazione(0,0) //ritorniamo una frazione non definita
	}
	if len(sommafrazioni) == 1 {
		return sommafrazioni[0]
	}
	//la lunghezza della slice è > 1: caso ricorsivo
	return sommafrazioni[0].Somma(Sommatoria(sommafrazioni[1:]))

}

//main legge da linea di comando una serie di frazioni e ne
//stampa la somma (se possibile, ma non necessariamente, in forma normalizzata)
//nel caso la serie sia vuota produrrà il messaggio "nessun argomento"
//sullo standard error
func main() {
	var totF []Frazione
	for i:=1 ;i<len(os.Args);i++{
		t,err := CreaFrazioneDaStringa(os.Args[i])
		if !err{
			fmt.Fprintf(os.Stderr,"%v\n",err)
		}
		totF = append(totF,t)
	}
	fmt.Println(totF)

	test := CreaFrazione(3,2)
	test2 := CreaFrazione(2,3)
	risultato := test.Somma(test2)
	fmt.Println(risultato)

	
	prova := Sommatoria(totF)
	floatprova := float64(prova.num)/float64(prova.den)
	fmt.Printf("%E",floatprova)

}
