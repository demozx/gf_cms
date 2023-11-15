// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	jwt "github.com/gogf/gf-jwt/v2"
)

type (
	IAuth interface {
		JWTAuth() *jwt.GfJWTMiddleware
		// PayloadFunc is a callback function that will be called during login.
		// Using this function it is possible to add additional payload data to the webtoken.
		// The data is then made available during requests via c.Get("JWT_PAYLOAD").
		// Note that the payload is not encrypted.
		// The attributes mentioned on jwt.io can't be used as keys for the map.
		// Optional, by default no additional data will be set.
		PayloadFunc(data interface{}) jwt.MapClaims
		// IdentityHandler get the identity from JWT and set the identity for every request
		// Using this function, by r.GetParam("id") get identity
		IdentityHandler(ctx context.Context) interface{}
		// Unauthorized is used to define customized Unauthorized callback function.
		Unauthorized(ctx context.Context, code int, message string)
		// Authenticator is used to validate login parameters.
		// It must return user data as user identifier, it will be stored in Claim Array.
		// if your identityKey is 'id', your user data must have 'id'
		// CheckByRoleId error (e) to determine the appropriate error message.
		Authenticator(ctx context.Context) (interface{}, error)
	}
)

var (
	localAuth IAuth
)

func Auth() IAuth {
	if localAuth == nil {
		panic("implement not found for interface IAuth, forgot register?")
	}
	return localAuth
}

func RegisterAuth(i IAuth) {
	localAuth = i
}
