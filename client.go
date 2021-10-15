// http client端

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//使用Get方法获取服务器响应包数据
	// resp, err := http.Get("http://localhost:8081/hello?name=lisa&age=24&class=1411")
	resp, err := http.Get("http://localhost:8081/healthz")

	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	// 获取服务器端读到的数据
	fmt.Println("Status = ", resp.Status)         // 状态
	fmt.Println("StatusCode = ", resp.StatusCode) // 状态码
	fmt.Println("Header = ", resp.Header)         // 响应头部
	fmt.Println("Body = ", resp.Body)             // 响应包体
	//读取body内的内容
	content, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(content))

}
