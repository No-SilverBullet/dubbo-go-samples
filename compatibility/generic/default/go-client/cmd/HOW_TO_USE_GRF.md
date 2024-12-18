### 1.install goref

Make sure your go version is \>= go1.21 

```go
go install github.com/cloudwego/goref/cmd/grf@latest
```

### 2.start docker

In dubbo-go-samples root directory

```sh
make -f build/Makefile docker-up
```

### 3.start server 

for example

```sh
cd ./compatibility/generic/default/go-server/cmd
export DUBBO_GO_CONFIG_PATH="../conf/dubbogo.yml"
go run .
```

### 4.edit client code

[go ref test samples](https://github.com/cloudwego/goref/blob/main/testdata/mockleak/main.go)

for example

```sh
cd ./compatibility/generic/default/go-client/cmd
```

import proof

```go
import _ "net/http/pprof"
```

start to listen 6060 port, add these lines in the begining of client.go/main()

```go
	go func() {
		log.Println(http.ListenAndServe("127.0.0.1:6060", nil))
	}()
```

modify the code to simulate multiple generic-calls, i simulated a four-minute generic-calls with the following code

```go
func main() {
	go func() {
		log.Println(http.ListenAndServe("127.0.0.1:6060", nil))
	}()
	// register POJOs
	hessian.RegisterPOJO(&pkg.User{})
	anchorTime := time.Now()
	for {
		if time.Since(anchorTime) > 4*time.Minute {
			break
		}
		dubboRefConf := newRefConf("org.apache.dubbo.samples.UserProvider", dubbo.DUBBO)
		callGetUser(dubboRefConf)
	}
	//wait gc, 1 minute is enough
	time.Sleep(1 * time.Minute)
	select {}
}
```

Add script gif_test.sh

```shell
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

#wait client run and gc
sleep 300

#使用 grf attach 到目标进程
if [ -n "$PID" ]; then
    echo "> Attaching grf to process $PID..."
    grf attach $PID
else
    echo "Failed to retrieve PID!"
    exit 1
fi

kill -9 $PID

```

wait for terminal output （maybe a few minutes）

```sh
successfully output to `grf.out`
```

### 5.Open grf.out file

install `graphviz`, and then

```go
go tool pprof -http=:5079 ./grf.out
```

