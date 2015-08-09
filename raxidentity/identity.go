package identity

import (
    "fmt"
    "net/http"
)

func Auth(token, endpoint string) int {
    req, err := http.NewRequest("GET", endpoint+"/"+token, nil)
    req.Header.Set("X-Auth-Token", token)
    
    client := &http.Client{}
    resp, err := client.Do(req)
    RespCode := resp.StatusCode
    
    if err != nil {
        fmt.Println("Error communicating with identity.")
        panic(err)
    }
    defer resp.Body.Close() 
    return RespCode
    
}