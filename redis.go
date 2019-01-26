package summer_providers

import (
	"github.com/cocotyty/summer"
	. "gopkg.in/redis.v4"
)

var _ = summer.Provider(&Redis{})
var _ = summer.Destroy(&Redis{})

type Redis struct {
	Address  string `sm:"#.(.).address"`
	PoolSize int    `sm:"#.(.).poolSize"`
	Password string `sm:"#.(.).password"`
	DB       int    `sm:"#.(.).db"`
	client   *Client
}

func (rp *Redis) Init() {
	rp.client = NewClient(&Options{
		Addr:     rp.Address,
		PoolSize: rp.PoolSize,
		Password: rp.Password,
		DB:       rp.DB,
	})
}
func (rp *Redis) Provide() interface{} {
	return rp.client
}

func (rp *Redis) Destroy() {
	rp.client.Close()
}
