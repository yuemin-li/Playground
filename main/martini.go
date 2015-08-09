package main

import (
    "github.com/go-martini/martini"
    "math/rand"
    "time"
    "fmt"
)

func main() {
    rand.Seed(10)
    m := martini.Classic()
    m.Get("/", func() string {
        tmp := rand.Intn(100)
        fmt.Printf("Sleep for  %d Second", tmp)
        time.Sleep(time.Duration(tmp)*time.Second)
        return "Hello world!"
    })
    m.Run()
}