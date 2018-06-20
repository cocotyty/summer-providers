package summer_providers

import (
	"github.com/bmizerany/mc"
	"github.com/jolestar/go-commons-pool"
	"context"
)

type PoolMemcachedClient struct {
}

func (pc *PoolMemcachedClient) Init() {
	ctx := context.Background()
	factory := pool.NewPooledObjectFactorySimple(func(context.Context) (interface{}, error) {
		return mc.Dial("tcp", "")
	})
	p := pool.NewObjectPoolWithDefaultConfig(ctx, factory)
	p.BorrowObject(ctx)
}
