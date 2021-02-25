package main

import (
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


	time.Sleep(100*time.Second)

}
