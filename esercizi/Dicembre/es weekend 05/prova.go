 package main

import (
	"amazon"
	"fmt"
	"mlist"
)

// HasLength is any type with an associated Length() method that returns an int
type hasLength interface {
	Length() int
}

type hasVolume interface {
	Volume() float64
}

func averageLength(x []hasLength) float64 {
	var sum int
	for _, b := range x {
		sum += b.Length()
	}
	return float64(sum) / float64(len(x))
}


func averageVolume (listv []hasVolume) float64 {
	var contatore int
	var somma float64
	for _,x := range listv {
		contatore++
		somma +=x.Volume()
	}
	return somma/float64(contatore)
}
/*
func cerca(x []interface{}, y interface{}) int {
	for i, t := range x {
		if t == y {
			return i
		}
	}
	return -1
}*/

func main() {
	var x mlist.List
	var pacco amazon.Box
	var sfera amazon.Sphere
	x = x.AddInOrder("pippo")
	x = x.AddInOrder("pluto")
	x = x.AddInOrder("topolino")
	x = x.AddInOrder("ultimo")
	x = x.AddInOrder("abate")
	x = x.AddInOrder("ragno")
	x.PrintWithArrows()
	fmt.Printf("Lunghezza della lista: %d\n", x.Length())

	fmt.Println(x)

	var scatola amazon.Box
	scatola = amazon.Box{30, 20, 15, 10}

	fmt.Printf("Volume della scatola: %d\n", scatola.Volume())
	fmt.Printf("Lunghezza della scatola: %d\n", scatola.Length())

	var y []hasLength
	y = append(y, amazon.Box{30, 20, 10, 20})
	y = append(y, amazon.Box{50, 25, 15, 30})
	y = append(y, x)
	y = append(y, amazon.Box{32, 20, 15, 10})

	fmt.Printf("Lunghezza media: %.2f\n", averageLength(y))
	fmt.Println(scatola)

//Aggiunta armani, ho creato 2 metodi per il peso del pacco e la sua densita'
	pacco = amazon.Box{30,20,10,15}
	paccovolume := pacco.Volume()
	paccopeso := pacco.Weight ()
	paccodensity := pacco.Density()
	fmt.Println("Volume: ",paccovolume,"\nPeso: ",paccopeso,"\nDensita'",paccodensity)

//esercizio parte2 devo creare un nuovo tipo di pacco che e' una sfera e voglio che calcoli la mia densita'
	sfera = amazon.Sphere{10}

	sferavolume := sfera.Volume()
	fmt.Println("volume della sfera: ",sferavolume)

	var test []hasVolume

	test = append(test, amazon.Box{30, 10, 10, 20})
	test = append(test, amazon.Sphere{10})
	test = append(test, amazon.Box{32, 20, 31, 20})
	test = append(test, amazon.Box{21, 31, 32, 20})
	test = append(test, amazon.Sphere{5})
	test = append(test, amazon.Box{11, 12, 10, 22})
	test = append(test, amazon.Box{30, 20, 31, 20})
	test = append(test, amazon.Sphere{2})
	test = append(test, amazon.Box{33, 20, 10, 11})
	test = append(test, amazon.Sphere{8})

	fmt.Println(test)
	fmt.Println(averageVolume(test))


}
