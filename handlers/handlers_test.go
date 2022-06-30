package handlers

import (
    "fmt"
    "net"
    "testing"
)

func TestAddressesHandler(t *testing.T) {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        fmt.Printf("%v\n", err)
        return
    }

    var ips []string
    for _, addr := range addrs {
        if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
            if ipNet.IP.To4() != nil {
                ips = append(ips, ipNet.IP.String())
            }
        }
    }
    fmt.Println(ips)
}

func BenchmarkAddressesHandler(b *testing.B) {
    for i := 0; i < b.N; i++ {
        addrs, err := net.InterfaceAddrs()
        if err != nil {
            fmt.Printf("%v\n", err)
            return
        }

        var ips []string
        for _, addr := range addrs {
            if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
                if ipNet.IP.To4() != nil {
                    ips = append(ips, ipNet.IP.String())
                }
            }
        }
        //fmt.Println(ips)
    }
}
