package main

import(
    "fmt"
    "github.com/yuemin-li/Playground"
)

func main() {
    token := "AABiY7MxzAhs6wVP7dghthtPbqTIZP5NrrrdupAY8qotm6KW8WaiYmidvxvKgJvRk0gi9lTR-6EJaEyviUYWtylKE0RDe3Qe2kxnww9uPg9x7jCVVrzXNr0j"
    url := "https://staging.identity.api.rackspacecloud.com/v2.0/tokens"
    
    status_code := identity.Auth(token, url)

    fmt.Println("Status Code: ", status_code)
}