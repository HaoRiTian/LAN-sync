#!/bin/sh

Script_Dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
Work_Dir="$( cd ${Script_Dir}/.. && pwd )"
cd ${Work_Dir}

#version=$(git describe --tags $(git rev-list --tags='v[0-9].[0-9]*' --max-count=1))
# 获取最近一次的tag标签作为版本号，且不会获取到奇奇怪怪诸如 0.1.0-5-gfef09d5 这样的标签值
version=$(git describe --abbrev=0 --tags)

go generate
# -H windowsgui 隐藏命令行窗口 -w 裁剪gdb调试信息
GOOS=windows GOARCH=amd64 go build -ldflags "-H windowsgui -w -s" -o build/LAN-sync-${version}.exe
