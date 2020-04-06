### goroutine 学习

#### goroutine 的本质

goroutine的调度模型： `GMP`（goroutine machine process）

`GPM` 是Go语言运行时（runtime）层面的实现， 是go语言自己实现的一套调度系统， 区别于操作系统调度OS线程

* `G(goroutine)`， 里面除了存放当前`goroutine` 信息外，还有与所在P的绑定等信息
* `P(Processor)` 管理着一组`goroutine`，P里面会存储当前goroutine运行的上下文环境（函数指针，堆栈地址及地址边界），P会对自己管理的goroutine队列做一些调度（比如把占用CPU时间较长的goroutine暂停、运行后续的goroutine等等）当自己的队列消费完了就去全局队列里取，如果全局队列里也消费完了会去其他P的队列里抢任务。
* `M（machine）`是Go运行时（runtime）对操作系统内核线程的虚拟， M与内核线程一般是一一映射的关系， 一个groutine最终是要放到M上执行的；

P与M一般也是一一对应的。他们关系是： P管理着一组G挂载在M上运行。当一个G长久阻塞在一个M上时，runtime会新建一个M，阻塞G所在的P会把其他的G 挂载在新建的M上。当旧的G阻塞完成或者认为其已经死掉时 回收旧的M。

`M:N` 把m个goroutine分配给n个操作系统线程

#### goroutine与操作系统线程（OS线程）的区别

`goroutine`是用户态线程，比内核态的线程更轻量，初始化只占用2kb的栈空间。

`goroutine`是由Go语言的运行时（runtime）调度完成，而线程是由操作系统调度完成

#### runtime.GOMAXPROCS

GO 1.5 之后默认就是操作系统的逻辑核心数，默认跑满CPU

`runtime.GOMAXPROCS(1)`: 只占用一个核



#### sync 同步

##### sync.WaitGroup

在 goroutine中，可使用`sync.WaitGroup` 来实现并发任务的同步

`sync.WaitGroup` 有一下几个方法

|             方法名              |        功能         |
| :-----------------------------: | :-----------------: |
| (wg * WaitGroup) Add(delta int) |    计数器+delta     |
|     (wg *WaitGroup) Done()      |      计数器-1       |
|     (wg *WaitGroup) Wait()      | 阻塞直到计数器变为0 |

需要注意`sync.WaitGroup`是一个结构体，传递的时候要传递指针。

##### sync.Once

`sync.Once`是goroutine 确保某些操作在高并发场景下只执行一次的解决方案

`sync.Once` 只有一个`Do`方法，具体如下

```go
func (o *Once) Do(f func()) {}
```

