/* LookupPort
 */

package main

import (
    "net"
    "os"
    "fmt"
    "strings"
)

func main() {
    if len(os.Args) != 3 {
        fmt.Fprintf(os.Stderr,
            "Usage: %s network-type service\n",
            os.Args[0])
        os.Exit(1)
    }
    networkType := os.Args[1]
    service := strings.ToLower(os.Args[2])

    port, err := net.LookupPort(networkType, service)
    if err != nil {
        fmt.Println("Error: ", err.Error())
        os.Exit(1)
    }

    fmt.Println("Service port ", port)
    os.Exit(0)
}
