## [H128](https://www.h128.com) 爬虫

> 该网站壁纸页面从旧到新以自然数递增，因此十分容易构造下载地址，但有些页面是不存在的，不是程序错误；代码仅用于学习交流目的，请勿大规模爬取，以免对该网站造成困扰！


### 示例一：通过命令行下载图片

```go
/*
	根据编号下载图片到当前工作目录
*/

package main

import (
	"flag"
	"github.com/cxfans/walls"
)

// go build -o bin/console examples/console.go
func main() {
	num := flag.Int("n", 20000, "图片编号")
	flag.Parse()
	_ = walls.Crawler(*num, "example.jpg")
}
```

### 示例二：24小时间隔爬取

```go
/*
	H128 网限制终身会员每天最多下载200张图片，
	因此可以通过每天爬取网站，以获取全部资源
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
	stop  = 25636
)

// go build -o bin/cycle examples/cycle.go
func main() {
	// 倒序方式
	for d := stop; d >= start; d -= 200 {
		for n := d; n > d-200; n-- {
			_ = walls.Crawler(n, fmt.Sprintf("e:/recycle/2025/相册/彼岸图网/h128_%d.jpg", n))
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
```
