package main

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gotodown/gospider"
	"github.com/zhshch2002/goreq"
)

// BaseDownloader is default downloader of goribot
type BaseDownloader struct {
	Client *http.Client
}

func NewBaseDownloader() *BaseDownloader {
	j, _ := cookiejar.New(nil)
	return &BaseDownloader{Client: &http.Client{Jar: j}}
}

func newSpider(e ...interface{}) *gospider.Spider {

	flo := &goreq.FilterLimiterOpinion{
		Allow: true,
	}
	lm := goreq.LimiterMatcher{
		Glob: "*.bilibili.com",
	}
	flo.LimiterMatcher = lm
	s := &gospider.Spider{
		Name:    "spider",
		Logging: false,
		Client: goreq.NewClient(
			goreq.WithRandomUA(),
			goreq.WithFilterLimiter(false, flo),
			goreq.WithRefererFiller(),
		),
		Status: gospider.NewSpiderStatus(),
	}
	s.SetWaitGroup()
	s.Use(e...)
	return s
}

func main() {
	// 1. 创建Spider
	s := newSpider(
		gospider.WithDepthLimit(3),
		// gospider.WithDeduplicate(),
	)
	do := NewBaseDownloader()
	// 获取视频信息
	var getVideoInfo = func(ctx *gospider.Context) {
		results, err := ctx.Resp.JSON()
		if err != nil {
			fmt.Println(err)
		}

		res := map[string]interface{}{
			"bvid":  results.Get("data.bvid").String(),
			"title": results.Get("data.title").String(),
			"des":   results.Get("data.des").String(),
			"pic":   results.Get("data.pic").String(),   // 封面图
			"tname": results.Get("data.tname").String(), // 分类名
			"owner": map[string]interface{}{ //视频作者
				"name": results.Get("data.owner.name").String(),
				"mid":  results.Get("data.owner.mid").String(),
				"face": results.Get("data.owner.face").String(), // 头像
			},
			"ctime":   results.Get("data.ctime").String(),   // 创建时间
			"pubdate": results.Get("data.pubdate").String(), // 发布时间
			"stat": map[string]interface{}{ // 视频数据
				"view":     results.Get("data.stat.view").Int(),
				"danmaku":  results.Get("data.stat.danmaku").Int(),
				"reply":    results.Get("data.stat.reply").Int(),
				"favorite": results.Get("data.stat.favorite").Int(),
				"coin":     results.Get("data.stat.coin").Int(),
				"share":    results.Get("data.stat.share").Int(),
				"like":     results.Get("data.stat.like").Int(),
				"dislike":  results.Get("data.stat.dislike").Int(),
			},
		}
		ctx.AddItem(res)
	}
	var findVideo gospider.Handler
	findVideo = func(ctx *gospider.Context) {
		u := ctx.Req.URL.String()
		fmt.Println("请求的网址----", u)
		if strings.HasPrefix(u, "https://www.bilibili.com/video/") {
			if strings.Contains(u, "?") {
				u = u[:strings.Index(u, "?")]
			}
			u = u[31:]
			fmt.Println("paodaozheli ------------->", u)
			ctx.AddTask(goreq.Get("https://api.bilibili.com/x/web-interface/view?bvid="+u), getVideoInfo)
		}
		html, err := ctx.Resp.HTML()
		if err != nil {
			fmt.Println(err)
			return
		}

		html.Find("a[href]").Each(func(i int, sel *goquery.Selection) {
			if h, ok := sel.Attr("href"); ok {
				fmt.Println(h)
				ctx.AddTask(goreq.Get(h), findVideo)
			}
		})
	}
	s.OnItem(func(ctx *gospider.Context, i interface{}) interface{} {
		// fmt.Println(i)
		return i
	})
	s.SeedTask(goreq.Get("https://www.bilibili.com/video/BV1at411a7RS").AddHeader("cookie", "_uuid=1B9F036F-8652-DCDD-D67E-54603D58A9B904750infoc; buvid3=5D62519D-8AB5-449B-A4CF-72D17C3DFB87155806infoc; sid=9h5nzg2a; LIVE_BUVID=AUTO7815811574205505; CURRENT_FNVAL=16; im_notify_type_403928979=0; rpdid=|(k|~uu|lu||0J'ul)ukk)~kY; _ga=GA1.2.533428114.1584175871; PVID=1; DedeUserID=403928979; DedeUserID__ckMd5=08363945687b3545; SESSDATA=b4f022fe%2C1601298276%2C1cf0c*41; bili_jct=2f00b7d205a97aa2ec1475f93bfcb1a3; bp_t_offset_403928979=375484225910036050"),
		findVideo)
	// s.Run()
	s.Wait()

}
