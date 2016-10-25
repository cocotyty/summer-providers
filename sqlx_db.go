package summer_providers

import (
	"github.com/cocotyty/summer"
	"github.com/jmoiron/sqlx"
	"time"
)

var _ = summer.Provider(&SqlXDB{})
var _ = summer.Destroy(&SqlXDB{})

type SqlXDB struct {
	DriverName  string `sm:"#.(.).driver"`
	Address     string `sm:"#.(.).address"`
	MaxOpenConn int    `sm:"#.(.).maxOpenConn"`
	MaxIdleConn int    `sm:"#.(.).maxIdleConn"`
	MaxLifetime int    `sm:"#.(.).maxLifetime"`
	db          *sqlx.DB
}

func (s *SqlXDB) Init() {
	db := sqlx.MustOpen(s.DriverName, s.Address)
	db.SetConnMaxLifetime(time.Duration(s.MaxLifetime) * time.Second)
	db.SetMaxIdleConns(s.MaxIdleConn)
	db.SetMaxOpenConns(s.MaxOpenConn)
	s.db = db
}

func (s *SqlXDB) Provide() interface{} {
	return s.db
}
func (s *SqlXDB) Destroy() {
	s.db.Close()
}
