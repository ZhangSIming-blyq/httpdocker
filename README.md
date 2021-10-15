# 构建本地镜像
# 2. 编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化
# 3. 将镜像推送至 docker 官方镜像仓库
# 4. 通过 docker 命令本地启动 httpserver
# 5. 通过 nsenter 进入容器查看 IP 配置

```bash
$ cat Dockerfile
FROM golang:1.16
WORKDIR /go/src/httpdemo
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o httpdemo server.go

FROM golang:1.16
LABEL author="ZhangSiming",project="httpdemo"
WORKDIR /root/
COPY --from=0 /go/src/httpdemo/httpdemo .

ENTRYPOINT ["/root/httpdemo"]

$ docker build -t httpdemo .
$ docker run --rm -it httpdemo:latest -p 80:8081
$ go run client.go 
Status =  200 OK
StatusCode =  200
Header =  map[Content-Length:[5] Content-Type:[text/plain; charset=utf-8] Date:[Mon, 11 Oct 2021 02:32:31 GMT]]
Body =  &{0xc00005c040 {0 0} false <nil> 0x6318e0 0x631860}
alive

# 进入容器查看
$ docker ps
CONTAINER ID   IMAGE             COMMAND                  CREATED         STATUS         PORTS                                   NAMES
151faffc51d8   httpdemo:latest   "/root/httpdemo"         4 seconds ago   Up 3 seconds   0.0.0.0:80->8081/tcp, :::80->8081/tcp   optimistic_panini
$ sudo docker inspect -f {{.State.Pid}} 151faffc51d8                                      
1718
$ sudo nsenter --target 1718 --mount --uts --ipc --net --pid
[151faffc51d8]:# ls
bin  boot  dev	etc  go  home  lib  lib64  media  mnt  opt  proc  root	run  sbin  srv	sys  tmp  usr  var
[151faffc51d8]:# ls /root
httpdemo
[151faffc51d8]:# hostname -I
172.17.0.3
```