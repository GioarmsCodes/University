package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	var nomefile string
	var file []string
	nomefile = os.Args[1]
	f,err := os.Open(nomefile)

	if err != nil {
		fmt.Fprint(os.Stderr,"%v\n",err)
	}
	defer f.Close()





	b:= bufio.NewScanner(f)



	for b.Scan(){
		s := b.Text()
		file = append(file,s)
	}






	for i:=0 ; i<len(file); i++{
		if file[i] =="" && file[i+1] == ""{
			continue
		}
		fmt.Println(file[i])
	}
}
