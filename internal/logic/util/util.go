package util

import (
	"context"
	"fmt"
	"gf_cms/internal/service"
	"math"
	"net"
	"os"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

// Util

type sUtil struct{}

var (
	insUtil           = sUtil{}
	Ctx               context.Context
	ProjectName       *gvar.Var
	BackendPrefix     *gvar.Var
	SystemRoot        string
	BackendGroup      string
	BackendApiGroup   string
	PublicCachePreFix string
	ServerRoot        *gvar.Var
)

func init() {
	service.RegisterUtil(New())
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
	SystemRoot, _ = os.Getwd()
	//服务目录
	ServerRoot, _ = g.Cfg().Get(Ctx, "server.serverRoot")
}

func New() *sUtil {
	return &sUtil{}
}

func Util() *sUtil {
	return &insUtil
}

// ProjectName 获取ProjectName
func (*sUtil) ProjectName() string {
	return ProjectName.String()
}

// SystemRoot 获取SystemRoot
func (*sUtil) SystemRoot() string {
	return SystemRoot
}

// BackendPrefix 后台入口前缀
func (*sUtil) BackendPrefix() string {
	return BackendPrefix.String()
}

// BackendApiPrefix 后台入口前缀
func (*sUtil) BackendApiPrefix() string {
	return service.Util().BackendPrefix() + "_api"
}

//BackendGroup 后台view分组
func (*sUtil) BackendGroup() string {
	return "/" + Util().BackendPrefix()
}

//BackendApiGroup 后台api分组
func (*sUtil) BackendApiGroup() string {
	return "/" + Util().BackendPrefix() + "_api"
}

//ServerRoot 服务目录
func (s *sUtil) ServerRoot() string {
	return ServerRoot.String()
}

// GetConfig 获取配置文件的配置信息
func (*sUtil) GetConfig(node string) string {
	config, _ := g.Cfg().Get(Ctx, node)
	return config.String()
}

// GetSetting 获取设置
func (*sUtil) GetSetting(name string) string {
	cacheKey := PublicCachePreFix + ":system_setting:" + name
	exists, err := g.Redis().Do(Ctx, "EXISTS", cacheKey)
	if err != nil {
		panic(err)
	}
	//存在缓存key
	if exists.Bool() {
		value, err := g.Redis().Do(Ctx, "GET", cacheKey)
		if err != nil {
			panic(err)
		}
		return value.String()
	}
	//不存在缓存key，从数据库读取
	val, _ := g.Model("system_setting").Where("name", name).Value("value")
	g.Redis().Do(Ctx, "SET", cacheKey, val.String())
	return val.String()
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

// ClearSystemSettingCache 清除后台设置缓存
func (*sUtil) ClearSystemSettingCache() (*gvar.Var, error) {
	cacheKey := PublicCachePreFix + ":system_setting:*"
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

// ImageOrDefaultUrl 返回图片或默认图url
func (*sUtil) ImageOrDefaultUrl(imgUrl string) string {
	if imgUrl == "" {
		return "/resource/images/no_pic.jpg"
	}
	return imgUrl
}

// PackedAloneRun 判断是否打包后单独运行的
func (*sUtil) PackedAloneRun() bool {
	//通过判断根目录的resource文件夹是否存在
	_, err := os.Stat(Util().SystemRoot() + "/resource/")
	if err != nil {
		if os.IsExist(err) {
			return false
		}
		return true
	}
	return false
}
