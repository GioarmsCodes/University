package main

import (
  "encoding/json"
  "fmt"
  "os"
  "io/ioutil"
)


func main() {
  s, err := readFile(os.Args[1])
  if err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
  for _, x := range s {
    fmt.Printf("%-20s\t%.2f\t%6.2f", x.Name, average(x), average110(x))
    for _, sc := range x.Score {
      fmt.Printf("\t%2d", sc)
    }
    fmt.Println()
  }

  fmt.Printf("Now I copy the data to the file named %s\n", os.Args[2])
  b, _ := json.Marshal(s)
  ioutil.WriteFile(os.Args[2], b, 0644)  // 0644 the permission flag of the file

  fmt.Printf("Now I read back the json file")
  b, _ = ioutil.ReadFile(os.Args[2])
  var t []student
  json.Unmarshal(b, &t)
  fmt.Println(t)
}
