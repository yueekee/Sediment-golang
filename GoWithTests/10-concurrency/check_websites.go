package main

import (

)

// 使用一个函数作为参数
type WebsiteChecker func(string) bool

// 单一线程
// func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
// 	results := make(map[string]bool)

// 	for _, url := range urls {
// 		results[url] = wc(url)
// 	}
// 	return results
// }

type result struct {
	string
	bool
}

// 多线程并发
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func (u string)  {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.string] = result.bool
	}

	return results
}