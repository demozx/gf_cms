package consts

const (
	// AdminSessionKeyPrefix 后台用户session前缀
	AdminSessionKeyPrefix = "admin_session"

	// ChannelModelArticle 文章模型
	ChannelModelArticle = "article"
	// ChannelModelArticleDesc 文章模型描述
	ChannelModelArticleDesc = "文章"
	// ChannelModelImage 图集模型
	ChannelModelImage = "image"
	// ChannelModelImageDesc 图集模型描述
	ChannelModelImageDesc = "图集"
	// ChannelModelSinglePage 单页模型
	ChannelModelSinglePage = "single_page"
	// ChannelModelSinglePageDesc 单页模型描述
	ChannelModelSinglePageDesc = "单页"

	// PcHomeAdChannelId pc首页广告分类id
	PcHomeAdChannelId = 1
	// GoodsChannelTid 产品栏目顶级id
	GoodsChannelTid = 4
	// AbortChannelTid 关于我们栏目顶级id
	AbortChannelTid = 8
	// NewsChannelTid 新闻栏目顶级id
	NewsChannelTid = 1

	// SwaggerUIPageContent 自定义SwaggerUI
	SwaggerUIPageContent = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1" />
  <meta name="description" content="SwaggerUI"/>
  <title>SwaggerUI</title>
  <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@latest/swagger-ui.css" />
</head>
<body>
<div id="swagger-ui"></div>
<script src="https://unpkg.com/swagger-ui-dist@latest/swagger-ui-bundle.js" crossorigin></script>
<script>
	window.onload = () => {
		window.ui = SwaggerUIBundle({
			url:    '/api.json',
			dom_id: '#swagger-ui',
		});
	};
</script>
</body>
</html>
`
)
