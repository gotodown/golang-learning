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
var once sync.Once
func (o *Once) Do(f func()) {}
```

`sync.Once` 内部包含一个互斥锁和一个布尔值， 互斥锁保证布尔值和数据的安全， 而布尔值用来记录初始化是否完成。 技能保证 初始化操作是并发安全的并且初始化操作不会被执行多次

#### 并发安全和锁

当多个`goroutine`同时操作一个资源时，对现有的资源会有个竞争，即`竞态问题`（数据竞态），经典的是`生产-消费者`,

```go
var x int64
var wg sync.WaitGroup
func add() {
    for i:=0; i< 500; i++{
        x += 1
    }
    wg.Done()
}
func main(){
    wg.Add(2)
    go add()
    go add()
    wg.Wait()
    fmt.Println(x)
}
```

上面的代码中的`goroutine`去累加变量x的值， `goroutine`在访问和修改`x`变量的时候就会存在数据竞争，导致最后的结果与期待的不符

已经锁定的Mutex并不与特定的goroutine相关联，这样可以利用一个goroutine对其进行加锁，再利用其他的goroutine对其解锁

##### 读写互斥锁

读写锁在Go语言中使用`sync`包中的`RWMute`类型

读写锁分为两种： 读锁和写锁，当一个goroutine获取读锁之后， 其他的`goroutine`如果是获取读锁则能继续获取读锁，如果是获取写锁就会等待;当一个`goroutine`获取写锁之后， 其他的`goroutine`获取读或写锁都会等待， 读写锁适合在读多写少的场景

```go
import sync

rwlock sync.RWMutex

//加锁
rwlock.RLock()
//释放锁
rwlock.RUnlock()

```



#### sync.Map

Go 语言中内置的map不是并发安全的， 示例

```go
var m = make(map[string]int)
func get(key sting) int {
    return m[key] = value
}

func main(){
    wg := sync.WaitGroup{}
    for i:=0; i<20; i++ {
        wg.Add(1)
        go func(n int) {
            key := strconv.Itoa(n)
            set(key, n)
            fmt.Printf("k:=%v, v:= %v", key, get(key))
            wg.Done()
        }(i)
    }
    wg.Wait()
}
```

以上代码在开启少量goroutine时 可能能正常运行，当并发多了之后即报`fatal error: concurrent map writes`

使用`sync.Map`来做并发， 开箱即用， 不需要使用类似与make函数进行初始化即可直接使用，同时内置了`Store`， `Load` ，`LoadOrStore`、`Delete`、 `Range`等操作方法

```go
var m = sync.Map{}

func main() {
    wg := sync.WaitGroup{}
    for i:=0;i <20 i++ {
        wg.Add(1)
        go func(n int){
        key := strconv.Itoa(n)
		m.Store(key, n)
		value, _ := m.Load(key)
		fmt.Printf("k=:%v,v:=%v\n", key, value)
        wg.Done()
    	}(i)
  	}
    wg.Wait()
}

```







