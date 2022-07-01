package handlers

import (
    "LAN-sync/utils"
    "fmt"
)

var UploadPath = ""

func init() {
    path, err := utils.GetUploadPath()
    if err != nil {
        fmt.Printf("%v, init upload path failed!\n", err)
        return
    }
    UploadPath = path
}
