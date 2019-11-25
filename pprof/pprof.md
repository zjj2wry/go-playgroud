## golang pprof basic

- [golang pprof](https://golang.org/pkg/net/http/pprof/)
- [gops](https://github.com/google/gops)

## 常见问题

- 分配的很多的 cpu，但是压测的时候使用率上不去
    - 数据库连接池设置的过小
- goroutine 泄漏导致内存飙升

## example code
```golang
package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go log.Println(http.ListenAndServe("localhost:6060", nil))
}
```

## memory
```bash
go tools pprof -http=:6060 http://localhost:6060/debug/pprof/heap
```

直接进入 web 查看
```bash
go tool pprof -http=:6061 http://localhost:6060/debug/pprof/heap
```

指定 path 查看
```bash
go tool pprof -http=:6061 /Users/zhengjiajin/pprof/pprof.alloc_objects.alloc_space.inuse_objects.inuse_space.009.pb.gz
```

top: 查看内存使用
```
flat: 当前使用的内存
sum: 累加使用的内存
unuse: 已分配未释放的内存
allocate: 已分配的内存
```

list [regrex]: 查看具体的函数执行

## cpu

获取 cpu 需要指定时间
```bash
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30
```
使用 web 查看
```
go tool pprof -http=:6061 /Users/zhengjiajin/pprof/pprof.samples.cpu.002.pb.gz
```

可以通过火焰图看哪个部分 cpu 占用的百分比更高
