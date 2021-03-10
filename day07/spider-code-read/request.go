package main

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"

	"github.com/PuerkitoBio/goquery"
	"github.com/zhshch2002/goribot"
)

// GetClient 返回一个http客户端
func GetClient() *http.Client {
	j, _ := cookiejar.New(nil)
	return &http.Client{
		Jar: j,
	}
}

// Get requests
func Get(url string) *http.Request {
	req, _ := http.NewRequest("GET", url, nil)
	return req
}

func main() {
	url := "https://www.bilibili.com/video/BV1at411a7RS"
	client := GetClient()
	req := Get(url)
	req.Header.Set("cookie", "_uuid=1B9F036F-8652-DCDD-D67E-54603D58A9B904750infoc; buvid3=5D62519D-8AB5-449B-A4CF-72D17C3DFB87155806infoc; sid=9h5nzg2a; LIVE_BUVID=AUTO7815811574205505; CURRENT_FNVAL=16; im_notify_type_403928979=0; rpdid=|(k|~uu|lu||0J'ul)ukk)~kY; _ga=GA1.2.533428114.1584175871; PVID=1; DedeUserID=403928979; DedeUserID__ckMd5=08363945687b3545; SESSDATA=b4f022fe%2C1601298276%2C1cf0c*41; bili_jct=2f00b7d205a97aa2ec1475f93bfcb1a3; bp_t_offset_403928979=375484225910036050")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	html, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	downloader := goribot.NewBaseDownloader()
	html.Find("a[href]").Each(func(i int, sel *goquery.Selection) {
		if h, ok := sel.Attr("href"); ok {
			fmt.Println(h)
			req1 := goribot.GetReq(h)
			req1.Header.Set("cookie", "_uuid=1B9F036F-8652-DCDD-D67E-54603D58A9B904750infoc; buvid3=5D62519D-8AB5-449B-A4CF-72D17C3DFB87155806infoc; sid=9h5nzg2a; LIVE_BUVID=AUTO7815811574205505; CURRENT_FNVAL=16; im_notify_type_403928979=0; rpdid=|(k|~uu|lu||0J'ul)ukk)~kY; _ga=GA1.2.533428114.1584175871; PVID=1; DedeUserID=403928979; DedeUserID__ckMd5=08363945687b3545; SESSDATA=b4f022fe%2C1601298276%2C1cf0c*41; bili_jct=2f00b7d205a97aa2ec1475f93bfcb1a3; bp_t_offset_403928979=375484225910036050")
			resp1, err := downloader.Do(req1)
			// html1, err := goquery.NewDocumentFromReader(resp1.Body)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("请求的host---->", resp1.Request.URL)
			return
		}
	})

}
