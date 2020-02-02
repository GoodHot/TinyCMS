package config

import (
	"fmt"
	"github.com/GoodHot/TinyCMS/common/strs"
	"time"
)

type db struct {
	TablePrefix     string        `json:"table_prefix"`      // table name prefix , like t_user. default "t_"
	Dialect         string        `json:"dialect"`           // database type, support MySQL,PostgreSQL,Sqlite3, SQL Server
	URL             string        `json:"url"`               // connection database url
	User            string        `json:"user"`              // connection database username
	Passwd          string        `json:"passwd"`            // connection database password
	MaxIdleConns    int           `json:"max_idle_conns"`    // database max idle conn
	MaxOpenConns    int           `json:"max_open_conns"`    // database max open conn
	ConnMaxLifetime time.Duration `json:"conn_max_lifetime"` // database max life time
}

type cache struct {
	Type    string `json:"type"`     // cache type, support redis,memory cache. [rerdis|memory]
	URL     string `json:"url"`      // connection cache server url
	Passwd  string `json:"passwd"`   // connection cache server password
	DBIndex uint   `json:"db_index"` // if type is redis, this value is redis db index
}

type log struct {
}

type controller struct {
	AdminPrefix string `json:"admin_prefix"`
	WebPrefix   string `json:"web_prefix"`
	APIPrefix   string `json:"api_prefix"`
}

type Config struct {
	Env        string      `json:"env"`        // application run environment. [prod, test, dev]
	AppDir     string      `json:"app_dir"`    // application in system directory path
	ConfigDir  string      `json:"config_dir"` // application config directory
	DB         *db         `json:"db"`         // database config
	Cache      *cache      `json:"cache"`      // cache config
	Log        *log        `json:"log"`        // system logger cconfig
	Controller *controller `json:"controller"`
}

func (s *Config) GetDirPath(path string) string {
	return s.AppDir + path
}

func (s *Config) Load() {
	fmt.Println("load config " + s.GetDirPath(strs.Fmt("%s/config_%s.json", s.ConfigDir, s.Env)))
}
