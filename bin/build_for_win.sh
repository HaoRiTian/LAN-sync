go generate
# -H windowsgui 隐藏命令行窗口 -w 裁剪gdb调试信息
GOOS=windows GOARCH=amd64 go build -ldflags "-H windowsgui -w -s" -o build/LAN-sync.exe
