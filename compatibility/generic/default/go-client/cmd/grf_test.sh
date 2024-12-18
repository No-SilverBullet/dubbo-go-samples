#!/bin/bash

# 设置变量
GO_FILE="client.go"

OUTPUT_BINARY="client_binary"
LOG_FILE="app.log"
PID_FILE="app.pid"
CONFIG_FILE="../conf/dubbogo.yml"
export PATH=$PATH:$(go env GOPATH)/bin
export DUBBO_GO_CONFIG_PATH=$CONFIG_FILE

# 编译 Go 文件
echo "> Compiling $GO_FILE..."
go build -o $OUTPUT_BINARY $GO_FILE
if [ $? -ne 0 ]; then
    echo "Compilation failed!"
    exit 1
fi
echo "> Compilation successful: $OUTPUT_BINARY"

# 将 PID 保存到文件中
nohup ./$OUTPUT_BINARY >$LOG_FILE 2>&1 &
echo $! >$PID_FILE
PID=$(cat $PID_FILE)
echo "> Application started with PID: $PID"

#等待客户端模拟以及GC完成
sleep 300

#使用 grf attach 到目标进程
if [ -n "$PID" ]; then
    echo "> Attaching grf to process $PID..."
    grf attach $PID
    kill -9 $PID
    echo "> Killed process with PID $PID."
else
    echo "Failed to retrieve PID!"
    exit 1
fi
