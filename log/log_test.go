package log

import (
	"testing"
	"time"
)



//测试日志实例打点
func TestLogInstance(t *testing.T) {
	Init("./logs/","sku","spu")
	Instance("sku").Info("111111111111111111111")
	Instance("sku").Warn("111111111111111111111")
	Instance("sku").Error("111111111111111111111")


	Instance("spu").Info("111111111111111111111")
	Instance("spu").Warn("111111111111111111111")
	Instance("spu").Error("111111111111111111111")


	time.Sleep(100*time.Second)
}