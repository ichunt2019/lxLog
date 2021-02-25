# lxLog

使用方法


```
package main

import (
	"errors"
	"fmt"
	xlog "github.com/ichunt2019/lxLog/log"
	"time"
)

func main(){
	xlog.Init("./logs/","sku","spu")


	xlog.Instance("sku").Info("111111111111111111111")
	xlog.Instance("sku").Warn("111111111111111111111")
	xlog.Instance("sku").Error("111111111111111111111")

	xlog.Instance("spu").Info("111111111111111111111")
	xlog.Instance("spu").Warn("111111111111111111111")
	xlog.Instance("spu").Error("111111111111111111111")

	err := test()
	if err != nil{
		xlog.Instance("sku").Info(fmt.Sprintf("%s",err))
	}

	time.Sleep(100*time.Second)

}



func test() (err error){
	xlog.Instance("sku").Info("666")
	return errors.New("错误")
}


```

记录日志


```
[WARN][2021/02/25 18:48:11][main.go:13] 111111111111111111111
[ERROR][2021/02/25 18:48:11][main.go:14] 111111111111111111111
[WARN][2021/02/25 18:49:56][main.go:15] 111111111111111111111
[ERROR][2021/02/25 18:49:56][main.go:16] 111111111111111111111
[WARN][2021/02/25 18:50:41][main.go:15] 111111111111111111111
[ERROR][2021/02/25 18:50:41][main.go:16] 111111111111111111111

```


