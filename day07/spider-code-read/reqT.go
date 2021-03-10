package main

import (
	"fmt"

	"github.com/zhshch2002/goreq"
)

func main() {
	req := goreq.Get("https://www.bilibili.com/video/BV1at411a7RS").AddHeader("cookie", "_uuid=1B9F036F-8652-DCDD-D67E-54603D58A9B904750infoc; buvid3=5D62519D-8AB5-449B-A4CF-72D17C3DFB87155806infoc; sid=9h5nzg2a; LIVE_BUVID=AUTO7815811574205505; CURRENT_FNVAL=16; im_notify_type_403928979=0; rpdid=|(k|~uu|lu||0J'ul)ukk)~kY; _ga=GA1.2.533428114.1584175871; PVID=1; DedeUserID=403928979; DedeUserID__ckMd5=08363945687b3545; SESSDATA=b4f022fe%2C1601298276%2C1cf0c*41; bili_jct=2f00b7d205a97aa2ec1475f93bfcb1a3; bp_t_offset_403928979=375484225910036050")
	// is, _ := req.GetBody()
	fmt.Println(req.Response)
}
