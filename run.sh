#!/bin/sh

#shell脚本语言，chmod u+x run.sh先赋予执行权限，在./run.sh执行脚本
#执行内容

go build -o main ./main.go
./main
