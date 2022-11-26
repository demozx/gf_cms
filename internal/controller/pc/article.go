package pc

import (
	"context"
	"gf_cms/api/pc"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	Article = cArticle{}
)

type cArticle struct{}

func (c *cArticle) List(ctx context.Context, req *pc.ArticleListReq) (res *pc.ArticleListRes, err error) {
	g.Dump("pc.article_list", req.Id)
	return
}

func (c *cArticle) Detail(ctx context.Context, req *pc.ArticleDetailReq) (res *pc.ArticleDetailRes, err error) {
	g.Dump("pc.article_detail")
	return
}
