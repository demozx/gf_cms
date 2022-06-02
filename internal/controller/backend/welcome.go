package backend

import (
	"context"
	"gf_cms/api/backend"
)

var (
	Welcome = cWelcome{}
)

type cWelcome struct{}

func (c *cWelcome) Index(ctx context.Context, req *backend.WelcomeReq) (res *backend.WelcomeRes, err error) {

	return
}
