/*
实现断点续传功能
// done
*/
package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

// 思路：
//   边复制，边记录复制的总量
//  方法： 通过seeker 随机的读写
func main() {

	src := "/root/tmp/mysql/mysql-community-8.0.18-1.el7.src.rpm"
	des := "/root/gdriver/code/mysql-community-8.0.18-1.el7.src.rpm"
	tfile := src + "-tmp.txt" // 记录访问游标的位置
	fr, _ := os.Open(src)
	fw, _ := os.OpenFile(des, os.O_CREATE|os.O_WRONLY, 0644)
	ft, _ := os.OpenFile(tfile, os.O_CREATE|os.O_WRONLY, 0644)

	// 关闭
	defer fr.Close()
	defer fw.Close()
	// 读取ft 文件中的数据
	// ft.Seek(0, io.SeekStart)
	bs := make([]byte, 100, 100)
	n1, err := ft.Read(bs)
	if err != nil {
		return
	}
	fmt.Println("读取了配置文件， 长度有。。。", n1)
	countStr := string(bs[:n1])
	// 字符串toint
	count, _ := strconv.ParseInt(countStr, 10, 64)

	// 设置读写的偏移量
	fr.Seek(count, io.SeekStart)
	fw.Seek(count, io.SeekStart)

	data := make([]byte, 1024, 1024)
	var n2 int
	var n3 int
	total := int(count)
	for {
		// 读取数据
		n2, err = fr.Read(data)
		if err == io.EOF {
			fmt.Println("文件读取完毕")
			ft.Close()
			os.Remove(tfile)
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		n3, _ = fw.Write(data[:n2])
		total += n3
		ft.Seek(0, io.SeekStart)
		ft.WriteString(strconv.Itoa(total))

		if total > 8000 {
			panic("假装断电了。。。，假装的。。。")
		}

	}

}
