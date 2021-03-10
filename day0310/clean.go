package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

const (
	dayTime   = 60 * 30
	monthTime = dayTime * 30
)

func getAllDirs(path string) []string {
	if !isExist(path) {
		fmt.Printf("getAllDirs--%s:No file or directory\n", path)
		return []string{}
	}
	rd, _ := ioutil.ReadDir(path)
	dirs := make([]string, 0)
	for _, fi := range rd {
		if fi.IsDir() {
			dirs = append(dirs, fi.Name())
		}
	}
	return dirs
}

// 递归删除空目录
func deleteEmptyDirs(path string) error {
	if !isExist(path) {
		return fmt.Errorf("deleteEmptyDirs--%s:No file or directory\n", path)
	}
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return nil
		}

		if f.IsDir() {
			dir, _ := ioutil.ReadDir(path)
			if len(dir) == 0 {
				err := os.Remove(path)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
	return err
}

func deleteAllFiles(path string, deleteTime int64) error {
	now := time.Now().Unix()
	if !isExist(path) {
		return fmt.Errorf("deleteAllFiles--%s:No file or directory", path)
	}
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		lastModTime := f.ModTime().Unix()
		if now-lastModTime > deleteTime*dayTime {
			os.Remove(path)
		}
		return nil
		// if f.IsDir() {}
	})
	return err
}

// CleanInstanceLogs is
func CleanInstanceLogs(path string, deleteTime int64) error {
	// 判断路径是否存在， 如果不存在，返回错误
	if !isExist(path) {
		return fmt.Errorf("CleanInstanceLogs--%s:No file or directory", path)
	}
	// 当实例存在gluedb.pid时， 代表实例还在活跃， 则只清理实例路径下归档目录的过时日志
	if isExist(path + "/gluedb.pid") {
		dirs := getAllDirs(path)
		for _, d := range dirs {
			err := deleteAllFiles(path+"/"+d, deleteTime)
			if err != nil {
				return err
			}
		}

	} else {
		err := deleteAllFiles(path, deleteTime)
		return err
	}
	return nil
}

// Clean is
func Clean(path string, delTime int64) error {
	// 清理当前目录下的所有空目录
	if !isExist(path) {
		return fmt.Errorf("Clean --%s:No file or directory", path)
	}
	err := deleteEmptyDirs(path)
	if err != nil {
		fmt.Println("删除空目录失败", err)
	}
	// 获取所有的实例名
	instanceNames := getAllDirs(path)
	// 逐个实例清理
	for _, insName := range instanceNames {
		// 实例日志路径$mountPath/insName/logs/
		logsPath := path + "/" + insName + "/" + "logs" + "/"
		err := CleanInstanceLogs(logsPath, delTime)
		if err != nil {
			return err
		}
	}
	return nil
}

// Exists 判断所给路径文件/文件夹是否存在
func isExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func main() {
	path := "/tmp/logs"
	// s := getAllDirs(path)
	delTime := int64(1)

	fmt.Println(Clean(path, delTime))
}
