/*
	获取最新壁纸
*/

package main

import (
	"fmt"
	"github.com/cxfans/netbian"
	"time"
)

// go build -o bin/cycle examples/range.go
// go run examples/range.go
func main() {
	// 2020.07.05
	start := 26073
	stop, _ := netbian.Latest()
	for d := start; d <= stop; d++ {
		_ = netbian.Crawler(d, fmt.Sprintf("C:/Merls/彼岸图网/netbian_%d.jpg", d))
		fmt.Println("DOWNLOAD: ", d)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Today's mission has been completed")
	// 颜色增强显示
	colorful := fmt.Sprintf("%c[1;00;31m%d%c[0m\n", 0x1B, stop+1, 0x1B)
	fmt.Println("Next time will start from", colorful)
}
