package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	// 存放图片链接
	chanImgUrls chan string
	// 存放页面链接
	chanPageUrls chan string
	// 存放任务
	chanTask  chan string
	waitGroup sync.WaitGroup
	imgFlag   int64
	pageFlag  int64
	downFlag  int64
	wgPage    int64
	wgImg     int64
)

const (
	baseURL = "https://e1.wkcsncjdbd.club/pw/"
	task    = 2
)

func main() {
	// 初始化数据管道
	chanImgUrls = make(chan string, 100000)
	chanTask = make(chan string, 147)
	chanPageUrls = make(chan string, 100000)
	pageFlag = 0
	imgFlag = 0
	downFlag = 0
	wgImg = 0
	wgPage = 0
	//爬虫协程： 不断地往管道中添加图片链接
	for i := 1; i < 4; i++ {
		waitGroup.Add(1)
		wgPage++
		go SpiderPageUrls("https://e1.wkcsncjdbd.club/pw/thread.php?fid=14&page=" + strconv.Itoa(i))
	}

	for i := 1; i < task; i++ {
		waitGroup.Add(1)
		fmt.Println("到处理图片了")
		wgImg++
		go SpiderImgUrls()
	}

	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go DownloadImg()
	}
	waitGroup.Wait()
	fmt.Println("============", imgFlag, pageFlag, downFlag, wgImg, wgPage, "============")
}

func getPageStr(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err, "http.Get(url)")
		return ""
	}
	defer resp.Body.Close()
	fByte, err := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(fByte))
	return string(fByte)
}

// SpiderPrettyLinks 爬取页面上的全部图集的链接
func SpiderPrettyLinks(url string) (urls []string) {
	// getPageStr 下载url的页面数据
	pageStr := getPageStr(url)
	reImg := `<a href="(html_data.*?html)" id`
	re := regexp.MustCompile(reImg)
	results := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("SpiderPrettyLinks -- 共找到%d条结果\n", len(results))

	for _, r := range results {
		url := baseURL + r[1]
		pageFlag++
		urls = append(urls, url)
	}
	fmt.Println(len(urls))
	return urls
}

// GetFilenameFromUrl 从url中提取文件名称
func GetFilenameFromUrl(url string, dirPath string) (filename string) {
	lastIndex := strings.LastIndex(url, "/")
	filename = url[lastIndex+1:]
	timePrefix := strconv.Itoa(int(time.Now().Unix()))
	filename = timePrefix + "_" + filename
	filename = dirPath + filename
	fmt.Println("************" + filename)
	return filename
}

func DownloadFile(url string, filename string) (ok bool) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err, "http.Get(url)")
		return
	}
	defer resp.Body.Close()

	fByte, err := ioutil.ReadAll(resp.Body)
	err = ioutil.WriteFile(filename, fByte, 0644)
	fmt.Println(err, "http.Get(url)")
	if err != nil {
		return false
	} else {
		return true
	}
}

func SpiderPrettyImgUrls(url string) (urls []string) {
	// getPageStr 下载url的页面数据
	pageStr := getPageStr(url)
	reImg := `src="(https.*?\.jpg)" border`
	re := regexp.MustCompile(reImg)
	results := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("SpiderPrettyImgUrls -- 共找到%d条结果\n", len(results))

	for _, r := range results {
		url := r[1]
		imgFlag++
		urls = append(urls, url)
	}
	fmt.Println(len(urls))
	return urls
}

func SpiderImgUrls() {
	for url := range chanPageUrls {
		urls := SpiderPrettyImgUrls(url)
		// 将所有图片超链接丢入数据管道
		for _, url := range urls {
			fmt.Println("丢到通道--", url)
			chanImgUrls <- url
		}

		// 通知当前协程任务完成
		chanTask <- url
	}
	waitGroup.Done()
}

// 爬取页面下的所有图片链接，并丢入全局待下载数据管道
func SpiderPageUrls(url string) {
	urls := SpiderPrettyLinks(url)
	// 将所有图片超链接丢入数据管道
	for _, url := range urls {
		chanPageUrls <- url
	}

	// 通知当前协程任务完成
	waitGroup.Done()
}

// 同步下载图片链接管道中的所有图片
func DownloadImg() {
	fmt.Println("下载任务开启。。。")
	time.Sleep(time.Second * 5)
	for url := range chanImgUrls {
		fmt.Println(url, " =======================")
		// filename := GetFilenameFromUrl(url, "/home/ljd/workstation/WorkStation/code/pyCode/telegram/scrapy/photo/")
		filename := GetFilenameFromUrl(url, "/usr/data/code/tmp/scrapy/photo/")
		ok := DownloadFile(url, filename)
		if ok {
			fmt.Printf("%s下载成功！\n", filename)
		} else {
			fmt.Printf("%s 下载失败！！！\n", filename)
		}
		downFlag++
	}
	fmt.Println("下载任务结束")
	waitGroup.Done()
}
