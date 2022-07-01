package handlers

import (
    "LAN-sync/utils"
    "fmt"
    "github.com/gin-gonic/gin"
    "log"
    "net/http"
    "path"
    "path/filepath"
    "strings"
)

func FilesHandler(c *gin.Context) {
    file, err := c.FormFile("raw")
    if err != nil {
        log.Fatal(err)
    }
    fileName := geneAvailFileName(file.Filename)
    saveFileErr := c.SaveUploadedFile(file, filepath.Join(UploadPath, fileName))
    if saveFileErr != nil {
        log.Fatal(saveFileErr)
    }
    c.JSON(http.StatusOK, gin.H{"url": "/uploads/" + fileName})
}

// 可能出现文件重名，重名则在文件名后加 “(%d)”
func geneAvailFileName(fileName string) string {
    fileSuffix := path.Ext(fileName)
    filenameOnly := strings.TrimSuffix(fileName, fileSuffix)
    for i := 1; utils.FileOrPathIsExists(filepath.Join(UploadPath + fileName)); i++ {
        fileName = filenameOnly + fmt.Sprintf("(%d)", i) + fileSuffix
    }
    return fileName
}
