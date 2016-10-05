package main

import (
  "fmt"
  "os"
  "os/signal"
  "time"
)

func main() {
  cocoers := []string{"Berni", "Huni", "Cata", "Tommy", "Horia", "Zoli", "Sandor", "Anca", "Dani", "Mihai"}
  
  c := make(chan os.Signal, 1)
  signal.Notify(c, os.Interrupt)
  go func(){
    for _ = range c {
        fmt.Println("\n\n...goodbye :)")
        os.Exit(1) //<- it's just wrong
    }
  }()
  
  for true {
    for _, cocoer := range cocoers {
      fmt.Printf("%s says...", cocoer)
    }
    time.Sleep(30 * time.Millisecond)
  }
}
