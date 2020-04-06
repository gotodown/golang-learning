package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
使用goroutine和channel实现一个计算int64随机数各位数和的程序。
开启一个goroutine循环生成int64类型的随机数，发送到jobChan
开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
主goroutine从resultChan取出结果并打印到终端输出
为了保证业务代码的执行性能将之前写的日志库改写为异步记录日志方式。
*/
type job struct {
	value int64
}

type result struct {
	job *job
	sum int64
}

var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)
var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go generator(jobChan)
	wg.Add(24)
	for i := 0; i < 24; i++ {
		go results(jobChan, resultChan)
	}
	for res := range resultChan {
		fmt.Println(res.job.value, res.sum)
	}
}

// 循环生成int64类型的随机数，发送到jobChan
func generator(ml chan<- *job) {
	defer wg.Done()
	for {
		value := rand.Int63() / 1000000000000
		newJob := &job{
			value: value,
		}
		ml <- newJob
		time.Sleep(time.Second * 1)
	}
}

// 计算数字的各位数的和
func results(ml <-chan *job, res chan<- *result) {
	defer wg.Done()
	for {
		job := <-ml
		num := job.value
		sum := Sum(num)
		newResult := &result{
			job: job,
			sum: sum,
		}
		res <- newResult
	}
	wg.Wait()
}

func Sum(num int64) int64 {
	var sum int64
	sum = 0
	if num%10 == num {
		return sum + num
	} else {
		sum := num % 10
		num = num / 10
		return sum + Sum(num)
	}
}
