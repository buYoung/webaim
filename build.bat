@echo off
set GOARCH=amd64
set GOHOSTARCH=amd64
set GOHOSTOS=windows
set GOOS=windows
set GOGCCFLAGS=-m64 -mthreads -fmessage-length=0
go build -ldflags="-s -w -extldflags="-static"" github.com/buYoung/webaim