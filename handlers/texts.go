package handlers

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "net/http"
    "os"
    "path/filepath"
    "time"
)

func TextHandler(c *gin.Context) {
    var text struct {
        Raw string `json:"raw"`
    }
    var err error
    if err = c.ShouldBindJSON(&text); err == nil {
        filename := uuid.New().String() + "-" + time.Now().Format("20060102-15-04-05") + ".txt"
        file := filepath.Join(UploadPath, filename)
        err = os.WriteFile(file, []byte(text.Raw), os.ModePerm)
        if err != nil {
            fmt.Printf("%v\n", err)
        }
    }
    if err == nil {
        c.JSON(http.StatusOK, gin.H{})
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    }
}
