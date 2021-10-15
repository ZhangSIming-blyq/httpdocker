FROM golang:1.16
WORKDIR /go/src/httpdemo
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o httpdemo server.go

FROM golang:1.16
LABEL author="ZhangSiming",project="httpdemo"
WORKDIR /root/
COPY --from=0 /go/src/httpdemo/httpdemo .

ENTRYPOINT ["/root/httpdemo"]
