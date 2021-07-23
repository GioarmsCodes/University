package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "strings"
)

type Response struct {
    Sol    []string    `json:"sol_keys"`
}

func main() {
    response, err := http.Get("https://api.nasa.gov/insight_weather/?api_key=RWcQarejc1zbbky57xOc1nPb4fFCwQCIMbip7XO5&feedtype=json&ver=1.0")

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

    fulljson := string(responseData)

    var responseObject Response
    json.Unmarshal(responseData, &responseObject)

    //fmt.Println(responseObject.Sol)

    stringhe := strings.Split(fulljson,"\n")
    for _,x := range stringhe[2:9]{

        fmt.Println(x)

}
}
