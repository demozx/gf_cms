package cache

import (
	"context"
	"gf_cms/internal/logic/util"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
	"time"

	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
)

const (
	//缓存驱动方式
	cacheDriverMemory = "memory" // 内存
	cacheDriverRedis  = "redis"  //redis
)

var (
	ctx      = gctx.New()
	insCache = sCache{}
	cache    = gcache.New()
)

type sCache struct{}

func init() {
	service.RegisterCache(New())
	initDriver()
}

func New() *sCache {
	return &sCache{}
}

func Cache() *sCache {
	return &insCache
}

func initDriver() {
	cacheDriverConfig := util.Util().GetConfig("server.cacheDriver")
	switch cacheDriverConfig {
	case "":
	case cacheDriverMemory:
		cache.SetAdapter(gcache.NewAdapterMemory())
	case cacheDriverRedis:
		redisConfig, err := g.Cfg().GetWithEnv(ctx, "redis.default")
		if err != nil {
			panic(err)
		}
		yamlConfig, err := gredis.ConfigFromMap(redisConfig.Map())
		if err != nil {
			panic(err)
		}
		gredis.Instance()
		redis, err := gredis.New(yamlConfig)
		if err != nil {
			panic(err)
		}
		cache.SetAdapter(gcache.NewAdapterRedis(redis))
		g.Dump("缓存驱动：redis")
	default:
		panic("缓存驱动方式配置错误")
	}
	return
}

// Set 设置缓存
func (s *sCache) Set(ctx context.Context, key string, value interface{}, duration time.Duration) (err error) {
	err = cache.Set(ctx, key, value, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get 获取缓存
func (s *sCache) Get(ctx context.Context, key string) (get *gvar.Var, err error) {
	get, err = cache.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	return
}
