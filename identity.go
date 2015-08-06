package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    //"github.com/go-martini/martini"
)

const AUTH_POINT = "https://staging.identity.api.rackspacecloud.com/v2.0/tokens"

func main() {
    token := "AADR4CYXAu4zFye0t5j5DieNzQDrRYev35c26HTYP00J8kQYgZ8_6BMPXtpzv59rReJO0jO_xq-4OwPbNgmbINT0oaZQLfayb_1jYezECd4Z4Scle_LGXkHD"

    req, err := http.NewRequest("GET", AUTH_POINT+"/"+token, nil)
    req.Header.Set("X-Auth-Token", token+"A")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Request has error")
        panic(err)
    }
    if resp.Status != "200 OK" {
        fmt.Println(resp.Status)
        panic("not autherized.")
    }

    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
    fmt.Println(resp)
    fmt.Println(err)
}

// func main(){
//     m := martini.Classic()
//     m.Use(auth.Token("TokenFoundInCache"))
//     m.Run()
// }