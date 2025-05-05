package cache

import (
	"gf_cms/internal/consts"
	"gf_cms/internal/logic/util"
	"gf_cms/internal/service"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
	"sync"
)

var (
	ctx               = gctx.New()
	insCache          = sCache{}
	cacheDriver       string
	cache             *gcache.Cache
	cacheInstanceOnce sync.Once
)

type sCache struct{}

func init() {
	service.RegisterCache(New())
}

func New() *sCache {
	return &sCache{}
}

func Cache() *sCache {
	return &insCache
}

func (s *sCache) GetCacheInstance() *gcache.Cache {
	cacheInstanceOnce.Do(func() {
		//g.Dump("初始化Cache")
		cache = gcache.New()
		// 初始化适配器
		service.Cache().InitDriver()
	})
	//g.Dump("已经初始化过Cache")
	return cache
}

func (s *sCache) InitDriver() {
	cacheDriverConfig := util.Util().GetConfig("server.cacheDriver")
	switch cacheDriverConfig {
	case "":
		fallthrough
	case consts.CacheDriverMemory:
		// 直接使用已初始化的 cache 实例
		cache.SetAdapter(gcache.NewAdapterMemory())
		cacheDriver = consts.CacheDriverMemory
		g.Dump("缓存驱动：memory")
	case consts.CacheDriverRedis:
		redisConfig, err := g.Cfg().GetWithEnv(ctx, "redis.default")
		if err != nil {
			panic(err)
		}
		yamlConfig, err := gredis.ConfigFromMap(redisConfig.Map())
		if err != nil {
			panic(err)
		}
		redis, err := gredis.New(yamlConfig)
		if err != nil {
			panic(err)
		}
		// 直接使用已初始化的 cache 实例
		cache.SetAdapter(gcache.NewAdapterRedis(redis))
		cacheDriver = consts.CacheDriverRedis
		g.Dump("缓存驱动：redis")
	default:
		panic("缓存驱动方式配置错误")
	}
}

// GetCacheDriver 获取驱动方式
func (s *sCache) GetCacheDriver() string {
	return cacheDriver
}
