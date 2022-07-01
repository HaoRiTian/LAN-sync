package utils

import (
    "os"
    "path/filepath"
)

func GetUploadPath() (string, error) {
    var err error
    exe, _ := os.Executable()
    exeDir := filepath.Dir(exe)
    upPath := filepath.Join(exeDir, "uploads")
    if !FileOrPathIsExists(upPath) {
        err = os.Mkdir(upPath, os.ModePerm)
    }
    return upPath, err
}

func FileOrPathIsExists(path string) bool {
    _, err := os.Stat(path)
    return err == nil || os.IsExist(err)
}

func If(cond bool, a interface{}, b interface{}) interface{} {
    if cond {
        return a
    }
    return b
}
