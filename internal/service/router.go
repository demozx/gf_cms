package service

// 路由
type sRouter struct{}

var (
	insRouter = sRouter{}
)

func Router() *sRouter {
	return &insRouter
}

func (*sRouter) Handle() {

}
