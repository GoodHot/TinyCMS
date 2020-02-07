package orm

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type ORM struct {
	Dialect          string `val:"${db.dialect}"`
	URL              string `val:"${db.url}"`
	TablePrefix      string `val:"${db.table_prefix}"`
	PhysicalDeletion bool   `val:${db.physical_deletion}`
	DB               *gorm.DB
}

func (s *ORM) Init(model ...interface{}) error {
	if s.Dialect != "mysql" && s.Dialect != "postgres" && s.Dialect != "sqlite3" && s.Dialect != "mssql" {
		return errors.New("unsupport database dialect:" + s.Dialect)
	}
	var err error
	s.DB, err = gorm.Open(s.Dialect, s.URL)
	if err != nil {
		return err
	}
	s.DB.SingularTable(true)
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return s.TablePrefix + defaultTableName;
	}
	s.DB.AutoMigrate(model...)
	return nil
}

// Transaction
func (s *ORM) Tx(callback func(db *gorm.DB) error) error {
	tx := s.DB.Begin()
	err := callback(tx)
	if err != nil {
		tx.Rollback()
		return err
	} else {
		tx.Commit()
	}
	return nil
}
