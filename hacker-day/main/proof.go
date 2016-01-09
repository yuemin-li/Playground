package main

import (
    "fmt"
    "net/http"
)

func main(){
    count := []int{1,2,3,4,5,6,7,8,9,10}
    
    for i := range(count) {
        fmt.Printf("Start a gorotine: %d\n", i)
        go request(i)
    }
    
    var input string
    fmt.Scanln(&input)
    fmt.Println("done")

}

func request(count int){
        req, err := http.NewRequest("GET", "http://localhost:3000", nil)
        if err != nil{
            panic(err)
        }
        client := &http.Client{}
        
        resp, err := client.Do(req)
        if err != nil{
            panic(err)
        }
        
        defer resp.Body.Close()
        
        fmt.Println("resp.Status:", resp.Status)
        fmt.Printf("request No. %d finished.", count)
}