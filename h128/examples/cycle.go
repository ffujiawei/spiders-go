/*
	网站限制终身会员每天最多下载200张图片，
	因此可以通过每天爬取网站，以获取全部资源。
*/

package main

import (
	"fmt"
	"github.com/cxfans/walls"
	"os"
	"time"
)

const (
	start = 469
	// 2020.5.18
	stop  = 17340
)

// go build -o bin/cycle examples/cycle.go
// go run examples/cycle.go
func main() {
	// 倒序方式
	for d := stop; d >= start; d -= 200 {
		for n := d; n > d-200; n-- {
			_ = walls.Crawler(n, fmt.Sprintf("imgs/h128_%d.jpg", n))
			fmt.Println("DOWNLOAD: ", n)
			time.Sleep(1 * time.Second)
		}
		fmt.Println("Today's mission has been completed.")
		if d >= start {
			// 休眠24小时
			time.Sleep(24 * time.Hour)
		} else {
			os.Exit(0)
		}
	}
}
