package v1

import (
    "LAN-sync/handlers"
    "LAN-sync/ws"
    "embed"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "io/fs"
    "log"
    "net/http"
    "strconv"
    "strings"
    "time"
)

func RunGin(port int, FS embed.FS) {
    gin.SetMode(gin.ReleaseMode)
    r := gin.Default()
    r.Use(Cors())

    sFS, _ := fs.Sub(FS, "frontend/dist")
    r.StaticFS("/static", http.FS(sFS))
    r.NoRoute(func(ctx *gin.Context) {
        path := ctx.Request.URL.Path
        if strings.HasPrefix(path, "/static") {
            if reader, err := sFS.Open("index.html"); err == nil {
                defer reader.Close()
                if stat, err := reader.Stat(); err == nil {
                    ctx.DataFromReader(http.StatusOK, stat.Size(), "text/html", reader, nil)
                }
            }
        }
        ctx.Status(http.StatusNotFound)
    })

    g := r.Group("/api/v1")
    Router(g)

    wsManager := ws.NewWsManager()
    wsManager.Run()
    r.GET("/ws", func(ctx *gin.Context) {
        wsManager.WsClientHandler(ctx)
    })

    r.GET("/uploads/:path", handlers.UploadsHandler)

    r.Run(":" + strconv.Itoa(port))
}

func Router(r *gin.RouterGroup) {
    r.POST("/texts", handlers.TextHandler)
    r.POST("/files", handlers.FilesHandler)
    r.GET("/addresses", handlers.AddressesHandler)
    r.GET("/qrcodes", handlers.QrCodeHandler)
}

// Cors 解决跨域问题
func Cors() gin.HandlerFunc {
    return cors.New(cors.Config{
        AllowMethods:     []string{"PUT", "PATCH", "POST"},
        AllowHeaders:     []string{"Origin", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        AllowOriginFunc: func(origin string) bool {
            if origin == "http://127.0.0.1:3000" || origin == "http://localhost:3000" {
                return true
            } else {
                log.Printf("%v is now allowed", origin)
                return false
            }
        },
        MaxAge: 12 * time.Hour,
    })

}
