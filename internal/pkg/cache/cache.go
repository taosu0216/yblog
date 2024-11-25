package cache

import (
	"log"
	"sync"
	"time"
)

type request struct {
	key    string
	result chan int
}

func Memoize[T comparable, R any](f func(T) R) func(T) R {
	cache := make(map[T]R)
	var mu sync.Mutex

	return func(x T) R {
		mu.Lock()
		if result, found := cache[x]; found {
			mu.Unlock()
			return result
		}
		mu.Unlock()

		result := f(x)

		mu.Lock()
		cache[x] = result
		mu.Unlock()

		return result
	}
}
func IpCounter() func(key string) int {
	cache := make(map[string]int)

	getAndIncrement := make(chan request)

	// 启动一个管理 cache 的 goroutine
	go func() {
		for {
			select {
			case req := <-getAndIncrement:
				cache[req.key]++
				req.result <- cache[req.key]
			case <-time.After(1 * time.Minute):
				cache = make(map[string]int) // 每分钟清空一次缓存
				log.Println("缓存已清空")
			}
		}
	}()

	return func(key string) int {
		if key == "" || key == "172.17.0.1" || key == "103.151.172.38" {
			return 0
		}
		result := make(chan int)
		getAndIncrement <- request{key: key, result: result}
		return <-result
	}

}
func BlackIpList() []string {
	return []string{}
}
