package summer_providers

import (
	"github.com/bmizerany/mc"
	"github.com/jolestar/go-commons-pool"
)

type PoolMemcachedClient struct {
}

func (pc *PoolMemcachedClient) Init() {
	p := pool.NewObjectPoolWithDefaultConfig(pool.NewPooledObjectFactorySimple(func() (interface{}, error) {
		return mc.Dial("tcp", "")
	}))
	p.BorrowObject()
}
