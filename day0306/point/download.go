package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	for i := 0; i < 5000; i++ {
		url := fmt.Sprintf("http://c.biancheng.net/view/%d.html", i)
		filename := fmt.Sprintf("%d.html", i)
		DownloadFile(url, filename)
	}

}

// func saveFile(filename string, content io.)

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
