#!/bin/bash

# 判断当前目录下是否存在 tmp 目录
if [ ! -d "./tmp" ]; then
    echo "tmp 目录不存在，正在创建..."
    mkdir ./tmp
fi
go run main.go 22 > ./tmp/main.ll
clang -Woverride-module -o ./tmp/main.o ./tmp/main.ll
./tmp/main.o
echo $?