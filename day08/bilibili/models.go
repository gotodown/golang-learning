package main

// 将数据进行序列化， 并保存至数据库
// Bilibili watch
type Bilibili struct {
	ID      int64     `json:"id" gorm:"column:id; PRIMARY_KEY"`
	Code    int64     `json:"code"`
	Data    VideoData `json:"data" gorm:"foreignkey:Aid"`
	Message string    `json:"message"`
	TTL     int64     `json:"ttl"`
}

// StatInfo 互动统计信息
type StatInfo struct {
	Aid        int64  `json:"aid"`
	ArgueMsg   string `json:"argue_msg"`
	Coin       int64  `json:"coin"`
	Danmaku    int64  `json:"danmaku"`
	Dislike    int64  `json:"dislike"`
	Evaluation string `json:"evaluation"`
	Favorite   int64  `json:"favorite"`
	HisRank    int64  `json:"his_rank"`
	Like       int64  `json:"like"`
	NowRank    int64  `json:"now_rank"`
	Reply      int64  `json:"reply"`
	Share      int64  `json:"share"`
	View       int64  `json:"view"`
}

// VideoData
type VideoData struct {
	Aid       int64  `json:"aid" gorm:"foreignkey:Aid"`
	Bvid      string `json:"bvid"`
	Cid       int64  `json:"cid"`
	Copyright int64  `json:"copyright"`
	Ctime     int64  `json:"ctime"`
	Desc      string `json:"desc"`
	// Dimension DimensionInfo `json:"dimension"`
	Duration  int64  `json:"duration"`
	Dynamic   string `json:"dynamic"`
	MissionID int64  `json:"mission_id"`
	NoCache   bool   `json:"no_cache"`
	// Owner     UserInfo `json:"owner"`
	// Pages     []VideoPage   `json:"pages"`
	Pic     string `json:"pic"`
	Pubdate int64  `json:"pubdate"`
	// Rights    VideoRights   `json:"rights"`
	// Stat      StatInfo      `json:"stat"`
	State int64 `json:"state"`
	// Subtitle  Subtitle      `json:"subtitle"`
	Tid   int64  `json:"tid"`
	Title string `json:"title"`
	Tname string `json:"tname"`
	// UserGarb  UserGarb      `json:"user_garb"`
	Videos int64 `json:"videos"`
}

type Subtitle struct {
	AllowSubmit bool          `json:"allow_submit"`
	List        []interface{} `json:"list"`
}

type VideoRights struct {
	Autoplay      int64 `json:"autoplay"`
	Bp            int64 `json:"bp"`
	CleanMode     int64 `json:"clean_mode"`
	Download      int64 `json:"download"`
	Elec          int64 `json:"elec"`
	Hd5           int64 `json:"hd5"`
	IsCooperation int64 `json:"is_cooperation"`
	IsSteinGate   int64 `json:"is_stein_gate"`
	Movie         int64 `json:"movie"`
	NoBackground  int64 `json:"no_background"`
	NoReprint     int64 `json:"no_reprint"`
	Pay           int64 `json:"pay"`
	UgcPay        int64 `json:"ugc_pay"`
	UgcPayPreview int64 `json:"ugc_pay_preview"`
}

// VideoPage 翻页信息
type VideoPage struct {
	Cid       int64         `json:"cid"`
	Dimension DimensionInfo `json:"dimension"`
	Duration  int64         `json:"duration"`
	From      string        `json:"from"`
	Page      int64         `json:"page"`
	Part      string        `json:"part"`
	Vid       string        `json:"vid"`
	Weblink   string        `json:"weblink"`
}

type UserInfo struct {
	Face string `json:"face"`
	Mid  int64  `json:"mid"`
	Name string `json:"name"`
}

type DimensionInfo struct {
	Height int64 `json:"height"`
	Rotate int64 `json:"rotate"`
	Width  int64 `json:"width"`
}

type UserGarb struct {
	URLImageAniCut string `json:"url_image_ani_cut"`
}
