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
