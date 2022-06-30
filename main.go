package main

import (
    v1 "LAN-sync/api/v1"
    "embed"
    "github.com/zserge/lorca"
    "log"
    "os"
    "os/signal"
)

//go:embed frontend/dist
var FS embed.FS

func main() {
    go v1.RunGin(FS)

    ui, err := lorca.New("", "", 480, 320, "disable-infobars")
    if err != nil {
        log.Fatal(err)
    }
    defer ui.Close()

    signCh := make(chan os.Signal)
    signal.Notify(signCh, os.Interrupt)
    select {
    case <-signCh:
    case <-ui.Done():
    }
    os.Exit(0)
}
