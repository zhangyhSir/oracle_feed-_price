package proxy

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"oracle/config"
	"os"
)

func UseProxy(method, coinUrl string) (string, error) {

	if method != "GET" && method != "POST" {
		return "", errors.New("提交方式method只能是'GET'或者'POST'")
	}

	// 设置Clash代理地址
	proxyURL, err := url.Parse(config.ProxyUrl)
	if err != nil {
		fmt.Println("Failed to parse proxy URL:", err)
		os.Exit(1)
	}

	// 创建自定义的Transport
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	// 创建自定义的Client
	client := &http.Client{
		Transport: transport,
	}

	// 创建GET/Post请求
	req, err := http.NewRequest(method, coinUrl, nil)
	if err != nil {
		fmt.Println("Failed to create request:", err)
		os.Exit(1)
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response:", err)
		os.Exit(1)
	}

	return string(body), nil
}
