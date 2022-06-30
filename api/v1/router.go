package v1

import (
    "LAN-sync/handlers"
    "embed"
    "github.com/gin-gonic/gin"
    "io/fs"
    "net/http"
)

func RunGin(FS embed.FS) {
    gin.SetMode(gin.ReleaseMode)
    r := gin.Default()

    sFS, _ := fs.Sub(FS, "frontend/dist")
    r.StaticFS("/static", http.FS(sFS))

    g := r.Group("/api/v1")
    Router(g)

    r.Run(":8080")
}

func Router(r *gin.RouterGroup) {
    r.POST("/texts", handlers.TextHandler)
    r.POST("/files", handlers.FilesHandler)
    r.POST("/addresses", handlers.AddressesHandler)
    r.POST("/uploads", handlers.UploadsHandler)
    r.POST("/qrcode", handlers.QrCodeHandler)
}
