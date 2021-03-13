package main

import (
	"fmt"
)

//definire un tipo Stack (di interi), due funzioni Pop e Push, e una funzione String
type Stack []int


func Push (s Stack,x int) Stack {
	s = append(s,x)
	return s
}

func Pop (s Stack) (ns Stack,top int,ok bool){
	if len(s) == 0{
		ok = false
		return
	}
	top = s[len(s)-1]
	ns = s[:len(s)-1]
	ok = true
	return
}

func String(s Stack) string {
	var c string
	for i:=0 ;i<len(s); i++{
	c += fmt.Sprintf("%v",s[i])
 }
	return c
}

func main() {
	var stack Stack
	stack = Push(stack,1)
	stack = Push(stack,2)
	stack = Push(stack,3)
	stack = Push(stack,4)
	stack = Push(stack,5)
	fmt.Println(String(stack))
	stack, top, ok := Pop(stack) //la funzione pop ritorna una tripla di valori
	fmt.Println("rimosso", top, "dallo stack")
	fmt.Println(String(stack))
	stack, top, ok = Pop(stack)
	fmt.Println("rimosso", top, "dallo stack")
	fmt.Println(String(stack))
	stack, top, ok = Pop(stack)
	fmt.Println("rimosso", top, "dallo stack")
	fmt.Println(String(stack))
	stack, top, ok = Pop(stack)
	if !ok {
		fmt.Println("stack vuoto")
	}
}
