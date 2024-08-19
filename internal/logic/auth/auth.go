package auth

import (
	"context"
	"gf_cms/internal/logic/admin"
	"gf_cms/internal/logic/util"
	"gf_cms/internal/model"
	"gf_cms/internal/service"
	"time"

	jwt "github.com/gogf/gf-jwt/v2"
	"github.com/gogf/gf/v2/frame/g"
)

var authService *jwt.GfJWTMiddleware

type (
	sAuth struct{}
)

var (
	insAuth = sAuth{}
)

func Auth() *sAuth {
	return &insAuth
}

func (*sAuth) JWTAuth() *jwt.GfJWTMiddleware {
	return authService
}

func init() {
	service.RegisterAuth(New())

	auth := jwt.New(&jwt.GfJWTMiddleware{
		Realm:           util.Util().ProjectName() + "_backend",
		Key:             []byte(util.Util().JwtKey()),
		Timeout:         time.Minute * 5,
		MaxRefresh:      time.Minute * 5,
		IdentityKey:     "id",
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
		Authenticator:   Auth().Authenticator,
		Unauthorized:    Auth().Unauthorized,
		PayloadFunc:     Auth().PayloadFunc,
		IdentityHandler: Auth().IdentityHandler,
	})
	authService = auth
}

func New() *sAuth {
	return &sAuth{}
}

// PayloadFunc is a callback function that will be called during login.
// Using this function it is possible to add additional payload data to the webtoken.
// The data is then made available during requests via c.Get("JWT_PAYLOAD").
// Note that the payload is not encrypted.
// The attributes mentioned on jwt.io can't be used as keys for the map.
// Optional, by default no additional data will be set.
func (*sAuth) PayloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	params := data.(map[string]interface{})
	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}

// IdentityHandler get the identity from JWT and set the identity for every request
// Using this function, by r.GetParam("id") get identity
func (*sAuth) IdentityHandler(ctx context.Context) interface{} {
	claims := jwt.ExtractClaims(ctx)
	return claims[authService.IdentityKey]
}

// Unauthorized is used to define customized Unauthorized callback function.
func (*sAuth) Unauthorized(ctx context.Context, code int, message string) {
	r := g.RequestFromCtx(ctx)
	r.Response.WriteJson(g.Map{
		"code":    code,
		"message": message,
	})
	r.ExitAll()
}

// Authenticator is used to validate login parameters.
// It must return user data as user identifier, it will be stored in Claim Array.
// if your identityKey is 'id', your user data must have 'id'
// CheckByRoleId error (e) to determine the appropriate error message.
func (*sAuth) Authenticator(ctx context.Context) (interface{}, error) {
	var (
		r  = g.RequestFromCtx(ctx)
		in model.AdminLoginInput
	)
	if err := r.Parse(&in); err != nil {
		return "", err
	}

	if user := admin.Admin().GetUserByUserNamePassword(ctx, in); user != nil {
		return user, nil
	}

	return nil, jwt.ErrFailedAuthentication
}
