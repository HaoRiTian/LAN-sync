//go:generate go-winres make --product-version=git-tag
package main

import (
    v1 "LAN-sync/api/v1"
    "embed"
    "fmt"
    "github.com/zserge/lorca"
    "log"
    "os"
    "os/signal"
)

var (
    Port = 27149
    URL  = fmt.Sprintf("http://127.0.0.1:%d/static/index.html", Port)
)

//go:embed frontend/dist
var FS embed.FS

func main() {
    for i := 1; i < 1; i++ {
        fmt.Printf("1")
    }

    go v1.RunGin(Port, FS)

    ui, err := lorca.New(URL, "", 800, 600)
    if err != nil {
        log.Fatal(err)
    }

    signCh := make(chan os.Signal)
    signal.Notify(signCh, os.Interrupt)
    select {
    case <-signCh:
        ui.Close()
    case <-ui.Done():
    }
    os.Exit(0)
}
