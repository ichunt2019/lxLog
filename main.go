package lxLog

import (
	log "github.com/ichunt2019/lxLog/lib"
	"time"
)

func main(){
	log.Init()
	Instance("sku").Info("111111111111111111111")
	Instance("sku").Warn("111111111111111111111")
	Instance("sku").Error("111111111111111111111")


	Instance("spu").Info("111111111111111111111")
	Instance("spu").Warn("111111111111111111111")
	Instance("spu").Error("111111111111111111111")


	time.Sleep(100*time.Second)
}
