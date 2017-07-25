package summer_providers

import (
	"database/sql"
	"github.com/cocotyty/summer"
	"time"
)

var _ = summer.Provider(&SqlDB{})
var _ = summer.Destroy(&SqlDB{})

type SqlDB struct {
	DriverName  string `sm:"#.(.).driver"`
	Address     string `sm:"#.(.).address"`
	MaxOpenConn int    `sm:"#.(.).maxOpenConn"`
	MaxIdleConn int    `sm:"#.(.).maxIdleConn"`
	MaxLifetime int    `sm:"#.(.).maxLifetime"`
	db          *sql.DB
}

func (s *SqlDB) Init() {
	db, err := sql.Open(s.DriverName, s.Address)
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Duration(s.MaxLifetime) * time.Second)
	db.SetMaxIdleConns(s.MaxIdleConn)
	db.SetMaxOpenConns(s.MaxOpenConn)
	s.db = db
}

func (s *SqlDB) Provide() interface{} {
	return s.db
}
func (s *SqlDB) Destroy() {
	s.db.Close()
}
