package main

import (
  "fmt"
  "list"
)

func main() {
  var x List
  x = AddInOrder("pippo", x)
  x = AddInOrder("pluto", x)
  x = list.AddInOrder("topolino", x)
  x = list.AddInOrder("ultimo", x)
  x = list.AddInOrder("abate", x)
  x = list.AddInOrder("ragno", x)

  //Prove delle print
  list.PrintWithArrows(x)
  list.PrintAllElements(x)

 fmt.Println("Test di armani : --")
  //devo aggiungere delle funzioni di stampa identiche ma ricorsive
  fmt.Println("Stampa all in modo ricorsivo ")
  list.PrintAllRic(x) // checked funziona

  //stampa sempre delle lista ma in modo ricorsivo
  fmt.Println("Stampa Frecce  in modo ricorsivo ")
  list.PrintArrowRic(x)

  //Aggiugere un elemento ad un determinato IndexList
  list.IndexList("ciao",x,2) //checked funziona
  list.PrintArrowRic(x)
//Rimuove un elemento da un determinato index
  list.IndexRemove(x,3) //checked funziona
  list.IndexRemove(x,3)
  list.PrintArrowRic(x)

  fmt.Printf("Lunghezza della lista: %d\n", list.Length(x))
}
