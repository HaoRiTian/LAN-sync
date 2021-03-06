package handlers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "path/filepath"
)

func UploadsHandler(c *gin.Context) {
    if path := c.Param("path"); path != "" {
        target := filepath.Join(UploadPath, path)
        c.Header("Content-Description", "File Transfer")
        c.Header("Content-Transfer-Encoding", "binary")
        c.Header("Content-Disposition", "attachment; filename="+path)
        c.Header("Content-Type", "application/octet-stream")
        c.File(target)
    } else {
        c.Status(http.StatusNotFound)
    }
}
