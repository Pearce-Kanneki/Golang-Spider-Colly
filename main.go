package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

// 查看網站源碼
func main() {
	fmt.Println(">> Golang Colly <<")
	basic()
	// onHtml()

}

func basic() {
	c := colly.NewCollector() // 在colly中使用Collector物件來做事情

	// 當Visit訪問網頁後,網頁響應(Response)時候執行的事情
	c.OnResponse(func(r *colly.Response) {

		fmt.Println(string(r.Body)) // 返回Response物件r.Body(格式: []Byte)
	})

	c.OnRequest(func(r *colly.Request) {

		r.Headers.Set(
			"User-Agent",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36",
		)
	})

	c.Visit("https://google.com/") // Visit要放在最後
}

func onHtml() {
	c := colly.NewCollector()

	// 尋找特定的Web元素

	// 抓標籤 Tag
	c.OnHTML("title", func(h *colly.HTMLElement) {
		fmt.Println(h.Text)
	})

	// 抓屬性質 AttrVal
	c.OnHTML("meta[name='test']", func(h *colly.HTMLElement) {
		fmt.Println(h.Text)
	})

	// 抓class 名稱
	c.OnHTML(".class-name", func(h *colly.HTMLElement) {
		fmt.Println(h.Text)
	})

	// 抓唯一識別ID
	c.OnHTML("#id-name", func(h *colly.HTMLElement) {
		fmt.Println(h.Text)
	})

	c.OnRequest(func(r *colly.Request) {

		r.Headers.Set(
			"User-Agent",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36",
		)
	})

	c.Visit("https://google.com/") // Visit要放在最後

}
