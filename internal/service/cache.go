// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
)

type (
	ICache interface {
		// Set 设置缓存
		Set(ctx context.Context, key string, value interface{}, duration time.Duration) (err error)
		// Get 获取缓存
		Get(ctx context.Context, key string) (get *gvar.Var, err error)
	}
)

var (
	localCache ICache
)

func Cache() ICache {
	if localCache == nil {
		panic("implement not found for interface ICache, forgot register?")
	}
	return localCache
}

func RegisterCache(i ICache) {
	localCache = i
}
