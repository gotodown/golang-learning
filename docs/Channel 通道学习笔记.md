### Channel 通道学习笔记

##### 通道初始化

有缓冲的与无缓冲的channel初始化的区别：

```go
c1 := make(chan int)  // 无缓冲，  
c2 := make(chan int, 10)  //有缓冲
c3 := make(chan<- int)  // 单向通道， 只能向c3 发送数据， c3 不能向外发送数据 
c1 <- 1 //接收 receive
<-c1 // 发送 send
close(c1)   //关闭通道
```

无缓冲： 不仅仅是向c1 通道传1 ，而是一直要等待别的协程(gorouting) 接收了这个通道的值， 那么c1 <- 1 才会继续下去，要不然，会一直阻塞

有缓冲： c2<- 1 则不会阻塞， 因为缓冲大小是10（其实缓冲大小为0），只有当放入超出缓冲大小的值个数时，才会阻塞（放置11个数，而前面的数没有被取走）

channel 的声明： `var ch chan <元素类型>`

#### `select` 多路复用

Go内置了`select`关键字， 可哟同时响应多个通道的操作

`select` 的使用类似与`switch`语句， 它有一系列的case分支和default分支， 每个case会对应一个通道的通信（接收或者发送）过程，`select`会一直等待， 直到某个`case`的通信操作完成， 才会执行`case`分支对应的语句

```go
select {
    case <- ch1:
    	...
    case data:= <-ch2:
       ...
    case ch3 <- data:
    	...
    default:
    	...
}
```

使用`select`语句能提高代码的可读性

* 可以处理一个或多个channel的发送/接收操作
* 如果多个`case`同时满足， `select`会随机选择一个case执行
* 对于没有case的select{} 会一直等待， 可用于阻塞main函数

#### 通道总结

`Channel` 常见的异常总结，如

![异常](/home/ljd/files/code/golang-learning/docs/imgs/channel异常总结.png "channel 异常总结")



