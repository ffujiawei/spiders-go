/*
	H128 4K 壁纸（https://www.h128.com）爬虫
*/

package walls

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

// n：图片编号；
// p：保存位置及命名
func Crawler(n int, p string) (err error) {
	queryUrl := "https://www.h128.com/tools/newdown.ashx"

	// 查询图片所需的表格数据
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("site_id", "1")
	_ = writer.WriteField("id", strconv.Itoa(n))
	_ = writer.WriteField("channel_id", "4")
	_ = writer.Close()

	client := &http.Client{}
	req, _ := http.NewRequest("POST", queryUrl, payload)
	req.Header.Add("cookie", "ASP.NET_SessionId=kpl1ony2oawu3jppzlmxvmgb; dt_cookie_url_referrer=https%3a%2f%2fwww.h128.com%2flist%2f0%2f0%2f2%2f0%2f0%2fd.html; dt_cookie_user_name_remember=DTcms=18268237856; dt_cookie_user_pwd_remember=DTcms=79EFED596C479D818E16191895B4EDC4")
	req.Header.Add("User-Agent", "Mozilla/5.0(WindowsNT10.0;Win64;x64)AppleWebKit/537.36(KHTML,likeGecko)Chrome/78.0.3904.97Safari/537.36")
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, _ := client.Do(req)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// 解析 JSON 数据
	var data struct {
		Status int    `json:"status"`
		Url    string `json:"file_path"`
		Msg    string `json:"msg"`
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	// 文件不存在或已经被删除
	if data.Status == 0 {
		fmt.Println(data.Msg)
		return nil
	}

	// 下载图片
	downUrl := data.Url
	req, _ = http.NewRequest("GET", downUrl, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0(WindowsNT10.0;Win64;x64)AppleWebKit/537.36(KHTML,likeGecko)Chrome/78.0.3904.97Safari/537.36")
	res, _ = client.Do(req)
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	// 路径不存在居然不报错
	img, err := os.Create(p)
	if err != nil {
		return err
	}
	_, _ = img.Write(body)
	return nil
}

// 自动获取最新图片的编号
func Latest() (l int, err error) {
	newUrl := "https://www.h128.com/list/0/0/2/0/0/t.html"
	resp, _ := http.Get(newUrl)
	bytes, _ := ioutil.ReadAll(resp.Body)
	reg, _ := regexp.Compile(`/show-([0-9]+).html`)
	n := reg.FindStringSubmatch(string(bytes))[1]
	l, err = strconv.Atoi(n)
	if err != nil {
		return 0, err
	}
	return l, nil
}
