/*
	获取最新壁纸
*/

package main

import (
	"fmt"
	"github.com/cxfans/walls"
	"time" 
)

// go build -o bin/cycle examples/range.go
// go run examples/range.go
func main() {
	// 2020.7.5
	start := 30995
	stop, _ := walls.Latest()
	for d := start; d <= stop; d++ {
		_ = walls.Crawler(d, fmt.Sprintf("C:/Merls/彼岸图网/h128_%d.jpg", d))
		fmt.Println("DOWNLOAD: ", d)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Today's mission has been completed")
	colored := fmt.Sprintf("%c[1;00;31m%d%c[0m\n", 0x1B, stop+1, 0x1B)
	fmt.Println("Next time will start from", colored)
}
