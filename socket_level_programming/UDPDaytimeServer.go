/* UDPDaytimeServer
 */
package main

import (
    "fmt"
    "net"
    "os"
    "time"
)

func main() {

    service := "localhost:1200"
    udpAddr, err := net.ResolveUDPAddr("udp4", service)
    checkError(err)

    for {
        conn, err := net.ListenUDP("udp", udpAddr)
        checkError(err)
        fmt.Println("Server start a connection.")
        go handleClient(conn)
    }
}

func handleClient(conn *net.UDPConn) {
    
    defer conn.Close()
    
    var buf [512]byte

    _, addr, err := conn.ReadFromUDP(buf[0:])
    if err != nil {
        return
    }

    daytime := time.Now().String() + "/n"

    conn.WriteToUDP([]byte(daytime), addr)
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
        os.Exit(1)
    }
}