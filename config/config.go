package config

import "time"

type db struct {
	TablePrefix     string        `json:"table_prefix"`      // table name prefix , like t_user
	Dialect         string        `json:"dialect"`           // database type, support MySQL,PostgreSQL,Sqlite3, SQL Server
	URL             string        `json:"url"`               // connection database url
	User            string        `json:"user"`              // connection database username
	Passwd          string        `json:"passwd"`            // connection database password
	MaxIdleConns    int           `json:"max_idle_conns"`    // database max idle conn
	MaxOpenConns    int           `json:"max_open_conns"`    // database max open conn
	ConnMaxLifetime time.Duration `json:"conn_max_lifetime"` // database max life time
}

type cache struct {
	Type string
}

type Config struct {
	DB    *db    `json:"db"`    // database config
	Cache *cache `json:"cache"` // cache config
}
