package service

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"math"
	"net"
	"time"
)

// Util

type sUtil struct{}

var (
	insUtil           = sUtil{}
	insViewBindFun    = sViewBindFun{}
	Ctx               context.Context
	ProjectName       *gvar.Var
	BackendPrefix     *gvar.Var
	SystemRoot        *gvar.Var
	BackendGroup      string
	BackendApiGroup   string
	PublicCachePreFix string
)

func init() {
	Ctx = gctx.New()
	//项目ProjectName
	ProjectName, _ = g.Cfg().Get(Ctx, "server.projectName", "gf_cms")
	//后台入口前缀
	BackendPrefix, _ = g.Config().Get(Ctx, "server.backendPrefix", "backend")
	//BackendGroup 后台view分组
	BackendGroup = "/" + BackendPrefix.String()
	//BackendApiGroup 后台api分组
	BackendApiGroup = "/" + BackendPrefix.String() + "_api"
	//公共缓存前缀
	PublicCachePreFix = ProjectName.String() + ":public"
	//项目目录
	SystemRoot, _ = g.Cfg().Get(Ctx, "server.systemRoot", "")
}

func Util() *sUtil {
	return &insUtil
}

// ProjectName 获取ProjectName
func (*sUtil) ProjectName() string {
	return ProjectName.String()
}

// BackendPrefix 后台入口前缀
func (*sUtil) BackendPrefix() string {
	return BackendPrefix.String()
}

//BackendGroup 后台view分组
func (*sUtil) BackendGroup() string {
	return "/" + Util().BackendPrefix()
}

//BackendApiGroup 后台api分组
func (*sUtil) BackendApiGroup() string {
	return "/" + Util().BackendPrefix() + "_api"
}

// GetConfig 获取配置文件的配置信息
func (*sUtil) GetConfig(node string) string {
	config, _ := g.Cfg().Get(Ctx, node)
	return config.String()
}

// ClearPublicCache 清除公共缓存
func (*sUtil) ClearPublicCache() (*gvar.Var, error) {
	cacheKey := PublicCachePreFix + ":*"
	keys, err := g.Redis().Do(Ctx, "KEYS", cacheKey)
	if err != nil {
		return nil, err
	}
	for _, key := range keys.Array() {
		_, err := g.Redis().Do(Ctx, "DEL", key)
		if err != nil {
			return nil, err
		}
	}
	return nil, err
}

// GetLocalIP 获取ip
func (*sUtil) GetLocalIP() (ip string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		return ipAddr.IP.String(), nil
	}
	return
}

// FriendyTimeFormat 计算时间差，并以"XXd XXh XXm XXs"返回
func (*sUtil) FriendyTimeFormat(TimeCreate time.Time, TimeEnd time.Time) string {
	SubTime := int(TimeEnd.Sub(TimeCreate).Seconds())
	// 秒
	if SubTime < 60 {
		return fmt.Sprintf("%d秒", SubTime)
	}
	// 分钟
	if SubTime < 60*60 {
		minute := int(math.Floor(float64(SubTime / 60)))
		second := SubTime % 60
		return fmt.Sprintf("%d分%d秒", minute, second)
	}
	// 小时
	if SubTime < 60*60*24 {
		hour := int(math.Floor(float64(SubTime / (60 * 60))))
		tail := SubTime % (60 * 60)
		minute := int(math.Floor(float64(tail / 60)))
		second := tail % 60
		return fmt.Sprintf("%d小时%d分%d秒", hour, minute, second)
	}
	// 天
	day := int(math.Floor(float64(SubTime / (60 * 60 * 24))))
	tail := SubTime % (60 * 60 * 24)
	hour := int(math.Floor(float64(tail / (60 * 60))))
	tail = SubTime % (60 * 60)
	minute := int(math.Floor(float64(tail / 60)))
	second := tail % 60
	return fmt.Sprintf("%d天%d小时%d分%d秒", day, hour, minute, second)
}
