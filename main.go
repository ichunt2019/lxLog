package main

import (
	 . "github.com/ichunt2019/lxLog/log"
	"time"
)



//测试日志实例打点
func main (){
	Init("./logs/","sku","spu")


	Instance("sku").Info("111111111111111111111")
	Instance("sku").Warn("22222222222222222")
	Instance("sku").Error("333333333333333333")


	Instance("spu").Info("777777777777777777")
	Instance("spu").Warn("888888888888888888888")
	Instance("spu").Error("999999999999999999")

	for {
		Instance("sku").Info("44444444444444444")
		Instance("sku").Warn("555555555555555555")
		Instance("sku").Error("666666666666666666")


		Instance("spu").Info("787878787878787878787878")
		Instance("spu").Warn("89898989898989898989898989898989")
		Instance("spu").Error("90909090909090909090909090909090")
		time.Sleep(time.Second*3)
	}


	for{
		time.Sleep(time.Second)
	}
}