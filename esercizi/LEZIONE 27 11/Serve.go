package main

import (
    "fmt"
    "image"
    "image/png"
    "github.com/holizz/terrapin"
    "math"
    "net/http"
)

//////////////////////////CREATORI DEL PNG////////////////////////////////////////////////////////
func pageImmagine(w http.ResponseWriter, r *http.Request) {
    i := image.NewRGBA(image.Rect(0, 0, 1000, 1000))
    t := terrapin.NewTerrapin(i, terrapin.Position{200.0, 250.0})
    Fiore(t, 40.0, 5)
    png.Encode(w, i)
}

func Disegnokoch(w http.ResponseWriter, r *http.Request) {
    i := image.NewRGBA(image.Rect(0, 0, 1000, 1000))
    t := terrapin.NewTerrapin(i, terrapin.Position{200.0, 600.0})
    kochSnowFlake(t, 20.0, 3)
    png.Encode(w, i)
}

func DisegnoQuadrato(w http.ResponseWriter, r *http.Request) {
    i := image.NewRGBA(image.Rect(0, 0, 1000, 1000))
    t := terrapin.NewTerrapin(i, terrapin.Position{200.0, 250.0})
    Quadrato(t,200,100)
    png.Encode(w, i)
}

func DisegnoStra(w http.ResponseWriter, r *http.Request) {
    i := image.NewRGBA(image.Rect(0, 0, 1000, 1000))
    t := terrapin.NewTerrapin(i, terrapin.Position{200.0, 250.0})
    Stra(t,10,150)
    png.Encode(w, i)
}
//////////////////////////CREATORI DEL PNG////////////////////////////////////////////////////////

////////////////////////DISEGNO FIORE FATTO A CASO/////////////////////////////////////////////////
func Fiore(t *terrapin.Terrapin, lung float64, liv int) {
    for i := 0; i < 3; i++ {
        if liv == 0 {
            t.Forward(lung)
            return
        }
        Fiore(t, lung, liv - 1)
        t.Left(math.Pi / 3.0)
        Fiore(t, lung, liv - 1)
        t.Right(math.Pi - math.Pi / 3.0)
        Fiore(t, lung, liv - 1)
        t.Left(math.Pi / 3.0)
        Fiore(t, lung, liv - 1)
        t.Right(2.0 * math.Pi / 3.0)
    }

}
////////////////////////DISEGNO FIORE FATTO A CASO/////////////////////////////////////////////////

////////////////////////DISEGNO QUADRATO/////////////////////////////////////////////////
func Quadrato (t *terrapin.Terrapin,volte int,lunghezza float64 ) {
    if volte == 0 {
        t.Forward(lunghezza)
        t.Right(math.Pi / 4.0)
        t.Forward(lunghezza)
        t.Right(math.Pi / 4.0)
        t.Forward(lunghezza)
        t.Right(math.Pi / 4.0)
        t.Forward(lunghezza)
        t.Right(math.Pi / 4.0)
    }else{
        t.Forward(lunghezza)
        t.Right(math.Pi / 4.0)
        t.Forward(lunghezza)
        t.Right(math.Pi / 4.0)
        t.Forward(lunghezza)
        t.Right(math.Pi / 4.0)
        t.Forward(lunghezza)
        t.Right(math.Pi / 4.0)
        Quadrato(t,volte-1,lunghezza/1.2)
    }

}
////////////////////////DISEGNO QUADRATO/////////////////////////////////////////////////

///////////////////////////////////STRA/////////////////////////////////////////////////////////////

func Stra (t *terrapin.Terrapin,volte int,lunghezza float64 ) {
    if volte == 0{
        t.Forward(lunghezza)
        t.Right(math.Pi / 2.0)
        t.Forward(lunghezza)
        t.Right(math.Pi / 2.0)
        t.Forward(lunghezza)
        t.Right(math.Pi / 2.0)
        t.Forward(lunghezza)
        t.Right(math.Pi / 3.0)
    }else{
        t.Forward(lunghezza)
        t.Right(math.Pi / 2.0)
        t.Forward(lunghezza)
        t.Right(math.Pi / 2.0)
        t.Forward(lunghezza)
        t.Right(math.Pi / 2.0)
        t.Forward(lunghezza)
        t.Right(math.Pi / 3.0)
        Stra(t,volte-1,lunghezza/1.2)
    }
}

///////////////////////////////////////////STRA---------------------------------------------------------------

///////////////////////////////DISEGNO FIOCCO DI KOCH//////////////////////////////////////////////////////////
func kochCurve(t *terrapin.Terrapin, lung float64, liv int) {
    if liv == 0 {
        t.Forward(lung)
        return
    }
    kochCurve(t, lung, liv - 1)
    t.Left(math.Pi / 3.0)
    kochCurve(t, lung, liv - 1)
    t.Right(math.Pi - math.Pi / 3.0)
    kochCurve(t, lung, liv - 1)
    t.Left(math.Pi / 3.0)
    kochCurve(t, lung, liv - 1)
}

func kochSnowFlake(t *terrapin.Terrapin, lung float64, liv int) {
    for i := 0; i < 3; i++ {
        kochCurve(t, lung, liv)
        t.Right(2.0 * math.Pi / 3.0)
    }
}
///////////////////////////////DISEGNO FIOCCO DI KOCH//////////////////////////////////////////////////////////

////////////////////////////////SITI VERI/////////////////////////////////////////////////////////////
func pagemain(w http.ResponseWriter, r *http.Request){
    w.Write([]byte( `
        <!doctype html>
        <title>Ciao</title>
        <h1>Come stai</h1>
        <img src="https://www.4gamehz.com/wp-content/uploads/2019/10/league-of-legends-4.jpg">
        `))
    }

    func OutFiore(w http.ResponseWriter, r *http.Request){
        w.Write([]byte( `
            <!doctype html>
            <title>Fiore Fatto a caso scusa kekw</title>
            <h1>TESTO A CASO</h1>
            <img src="/Fiore.png">
            `))
        }

        func OutKoch(w http.ResponseWriter, r *http.Request){
            w.Write([]byte( `
                <!doctype html>
                <title>Koch  Fatto a caso scusa kekw</title>
                <h1>Fiocco di Koch</h1>
                <img src="/Koch.png">
                `))
            }

            func OutQuadrato(w http.ResponseWriter, r *http.Request){
                w.Write([]byte( `
                    <!doctype html>
                    <title>Koch  Fatto a caso scusa kekw</title>
                    <h1>Spirale</h1>
                    <img src="/Quadrato.png">
                    `))
                }

                func OutStra(w http.ResponseWriter, r *http.Request){
                    w.Write([]byte( `
                        <!doctype html>
                        <title>Koch  Fatto a caso scusa kekw</title>
                        <h1>Stra</h1>
                        <img src="/Stra.png">
                        `))
                    }

            ////////////////////////////////SITI VERI/////////////////////////  ////////////////////////////////////

            func main(){
                http.HandleFunc("/a/",pagemain)

                http.HandleFunc("/Fiore.html",OutFiore)
                http.HandleFunc("/Fiore.png",pageImmagine)

                http.HandleFunc("/Koch.html",OutKoch)
                http.HandleFunc("/Koch.png",Disegnokoch)

                //Lo so si chiama quadrato ma disegna una spirale ma non ho voglia di cambiarlo dappertutto just stfu
                http.HandleFunc("/Spirale.html",OutQuadrato)
                http.HandleFunc("/Quadrato.png",DisegnoQuadrato)

                http.HandleFunc("/Stra.html",OutStra)
                http.HandleFunc("/Stra.png",DisegnoStra)

                fmt.Println("Listening on http://localhost:3000/")
                    http.ListenAndServe(":3000",nil)

            }
