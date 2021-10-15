// 编写一个http server
// 1. 接收客户端 request，并将 request 中带的 header 写入 response header
// 2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
// 3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
// 4. 当访问 localhost/healthz 时，应返回200

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	//注册回调函数
	http.HandleFunc("/hello", hellohello)
	http.HandleFunc("/healthz", demohealth)
	log.Println("start")
	// 开始监听8081的请求
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func hellohello(w http.ResponseWriter, r *http.Request) {
	//ParseForm解析URL中的查询字符串，并将解析结果更新到r.Form字段
	r.ParseForm()
	//遍历打印解析结果
	for key, value := range r.Form {
		// 响应头加入请求头的header
		w.Header().Set(key, fmt.Sprintf(strings.Join(value, ",")))
	}
	// 获取环境变量的VERSION
	os_version := os.Getenv("VERSION")
	w.Header().Set("version", os_version)
	// 获取客户端ip
	remote_ip := GetIP(r)
	fmt.Println(remote_ip, 200)

	w.WriteHeader(200)
	w.Write([]byte("hello"))
}

func demohealth(w http.ResponseWriter, r *http.Request) {
	// health接口
	w.WriteHeader(200)
	w.Write([]byte("alive"))
}

func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}
