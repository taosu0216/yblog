package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func main() {
	// 并发访问的次数
	concurrency := 300

	// 创建一个等待组，用于等待所有协程完成
	var wg sync.WaitGroup
	wg.Add(concurrency)

	// 创建一个通道，用于收集响应
	responses := make(chan string, concurrency)

	// 并发发送请求
	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()

			data := map[string]string{
				"username": "root",
				"password": "root",
			}

			jsonData, err := json.Marshal(data)
			if err != nil {
				panic(err)
			}

			body := bytes.NewBuffer(jsonData)

			req, err := http.NewRequest("POST", "http://127.0.0.1:8000/api/user/login", body)
			if err != nil {
				panic(err)
			}

			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}

			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			respBody, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}

			// 将响应发送到通道
			responses <- string(respBody)
		}()
	}

	// 等待所有协程完成
	wg.Wait()
	close(responses)

	// 收集并打印所有响应
	for resp := range responses {
		fmt.Println(resp)
	}
}
