package handlers

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/skip2/go-qrcode"
    "net/http"
)

// QrCodeHandler TODO：创建二维码交给前端最好
func QrCodeHandler(c *gin.Context) {
    var err error
    if content := c.Query("content"); content != "" {
        png, err := qrcode.Encode(content, qrcode.Medium, 256)
        if err == nil {
            c.Data(http.StatusOK, "image/png", png)
        }
    }
    if err != nil {
        fmt.Printf("%v\n", err)
    }
    c.Status(http.StatusBadRequest)
}
