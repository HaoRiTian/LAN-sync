#!/bin/sh
#version=$(git describe --tags $(git rev-list --tags='v[0-9].[0-9]*' --max-count=1))
# 获取最近一次的tag标签作为版本号
version=$(git describe --abbrev=0 --tags)

go generate
# -H windowsgui 隐藏命令行窗口 -w 裁剪gdb调试信息
GOOS=windows GOARCH=amd64 go build -ldflags "-H windowsgui -w -s" -o build/LAN-sync-${version}.exe
