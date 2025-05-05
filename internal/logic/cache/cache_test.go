package cache

import (
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

var ()

func Test_Set(t *testing.T) {
	key := "test"
	value := "1234"
	err := service.Cache().GetCacheInstance().Set(ctx, key, value, 0)
	if err != nil {
		panic(err)
	}
	//time.Sleep(time.Second * 2)
	get, err := service.Cache().GetCacheInstance().Get(ctx, key)
	if err != nil {
		return
	}
	g.Dump(get)
}

func Test_Get(t *testing.T) {
	get, err := service.Cache().GetCacheInstance().Get(ctx, "test")
	if err != nil {
		return
	}
	g.Dump(get)
}

func Test_Keys(t *testing.T) {
	keys, err := service.Cache().GetCacheInstance().Keys(ctx)
	g.Dump(keys, err)
}

func Test_Del(t *testing.T) {
	//keys := g.Slice{"gf_cms:public:system_setting:admin_emails", "gf_cms:public:system_setting:auto_art_pic"}
	//_, err := service.Cache().GetCacheInstance().Del(ctx, keys)
	//if err != nil {
	//	g.Dump(err)
	//}
}
