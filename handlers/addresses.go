package handlers

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "net"
    "net/http"
)

func AddressesHandler(c *gin.Context) {
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
    c.JSON(http.StatusOK, gin.H{
        "addresses": ips,
    })
}
