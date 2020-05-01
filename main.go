package main

import (
	"log"
	"net/http"
	"os"
)

// 路由
func router(w http.ResponseWriter, r *http.Request) {

	// 允许访问所有域
	w.Header().Set(`Access-Control-Allow-Origin`, `*`)

	// 路由规则
	switch r.URL.Path {
	case `/yunhei`:
		outputJSON(w, getWebContent(`https://gitee.com/ranbom/SakuraBan/raw/master/a.json`))
	default:
		w.WriteHeader(404)
		outputText(w, `没有这个功能`)
	}
}

// 向客户端或者浏览器输出json数据
func outputJSON(w http.ResponseWriter, content string) {
	w.Header().Set(`content-type`, `application/json`)
	outputText(w, content)
}

// 向客户端或者浏览器输出文本内容
func outputText(w http.ResponseWriter, content string) {
	w.Write([]byte(content))
}

// 使用Get方式获取指定网址内容
func getWebContent(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(`请求网址：` + url + ` 失败`)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		buf := make([]byte, 1024)
		var context string
		for {
			n, _ := resp.Body.Read(buf)
			if 0 == n {
				break
			}
			context += string(buf[:n])
		}
		return context
	}
	return ``
}

// 主函数：一般情况下不用动
func main() {
	// 获取日志文件句柄
	logFile, err := os.OpenFile(os.Args[0]+`.log`, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	// 设置存储位置
	log.SetOutput(logFile)
	// 设置访问的路由
	http.HandleFunc(`/`, router)
	// 设置监听的端口
	err = http.ListenAndServe(`:233`, nil)
	if err != nil {
		log.Fatal(err)
	}
}
