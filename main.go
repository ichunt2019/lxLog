package main

import (
	 . "github.com/ichunt2019/lxLog/log"
	"time"
)



//测试日志实例打点
func main (){
	Init("./logs/","sku","spu")


	Instance("sku").Info("111111111111111111111")
	Instance("sku").Warn("111111111111111111111")
	Instance("sku").Error("111111111111111111111")


	Instance("spu").Info("111111111111111111111")
	Instance("spu").Warn("111111111111111111111")
	Instance("spu").Error("111111111111111111111")

	for {
		Instance("sku").Info("111111111111111111111")
		Instance("sku").Warn("111111111111111111111")
		Instance("sku").Error("111111111111111111111")


		Instance("spu").Info("111111111111111111111")
		Instance("spu").Warn("111111111111111111111")
		Instance("spu").Error("111111111111111111111")
		time.Sleep(100*time.Millisecond)
	}


	time.Sleep(100*time.Second)
}