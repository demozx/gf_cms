package cache

import (
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
	"time"
)

var ()

func Test_setCache(t *testing.T) {
	key := "test"
	value := "1234"
	err := service.Cache().Set(ctx, key, value, time.Second)
	if err != nil {
		panic(err)
	}
	//time.Sleep(time.Second * 2)
	get, err := service.Cache().Get(ctx, key)
	if err != nil {
		return
	}
	g.Dump(get)
}
