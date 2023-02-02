package code

import (
	"fmt"
	"time"
)

func Fetch(url, path string) {
	start := time.Now()
	fmt.Println("Start ==> ", url)
	lea := NewCrawler()
	lea.Fetch(url, path)
	fmt.Printf("耗时: %v\n", time.Now().Sub(start))
}
