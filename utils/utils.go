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
    if !PathExists(upPath) {
        err = os.Mkdir(upPath, os.ModePerm)
    }
    return upPath, err
}

func PathExists(path string) bool {
    _, err := os.Stat(path)
    return If(err == nil, true, !os.IsNotExist(err)).(bool)
}

func If(cond bool, a interface{}, b interface{}) interface{} {
    if cond {
        return a
    }
    return b
}
