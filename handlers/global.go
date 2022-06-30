package handlers

import (
    "LAN-sync/utils"
    "fmt"
)

var UploadPath = ""

func init() {
    path, err := utils.GetUploadPath()
    if err != nil {
        fmt.Printf("%v, init global var failed!\n", err)
        return
    }
    UploadPath = path
}
