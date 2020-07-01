package orm

import (
	"errors"
	"github.com/GoodHot/TinyCMS/brain"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"strconv"
	"time"
)

type Model struct {
	ID        int        `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

type ORM struct {
	Log                  brain.Logger `ioc:"auto"`
	Dialect              string       `val:"${db.dialect}"`
	URL                  string       `val:"${db.url}"`
	TablePrefix          string       `val:"${db.table_prefix}"`
	SQLLog               bool         `val:"${db.sql_log}"`
	PhysicalDeletion     bool         `val:"${db.physical_deletion}"`
	ArticleSubTableCount int          `val:"${db.article_sub_table_count}"`
	DB                   *gorm.DB
}

func (s *ORM) Startup() error {
	if s.Dialect != "mysql" && s.Dialect != "postgres" && s.Dialect != "sqlite3" && s.Dialect != "mssql" {
		s.Log.Error("unsupport database dialect:%s", s.Dialect)
		return errors.New("unsupport database dialect:" + s.Dialect)
	}
	s.Log.Info("ORM Connect database %s[%s]. table prefix is [%s], physical deletion is [%v]", s.Dialect, s.URL, s.TablePrefix, s.PhysicalDeletion)
	db, err := gorm.Open(s.Dialect, s.URL)
	if err != nil {
		return err
	}
	s.DB = db
	s.DB.LogMode(s.SQLLog)
	s.DB.SingularTable(true)
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return s.TablePrefix + defaultTableName
	}
	return nil
}

func (s *ORM) AutoMigrate(contentModel interface{}, model ...interface{}) error {
	err := s.DB.AutoMigrate(model...).Error
	if err != nil {
		return err
	}
	for i := 0; i < s.ArticleSubTableCount; i++ {
		tableName := s.ArticleContentTableName(i)
		if s.DB.HasTable(tableName) {
			continue
		}
		err := s.DB.Table(tableName).CreateTable(contentModel).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *ORM) ArticleContentTableName(index int) string {
	return s.TablePrefix + "article_content_" + strconv.Itoa(index)
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

type Page struct {
	List       interface{} `json:"list"`
	PageNum    int         `json:"page_num"`
	PageSize   int         `json:"page_size"`
	TotalPage  int         `json:"total_page"`
	TotalCount int         `json:"total_count"`
}

type ORMPage struct {
	PageNum  int
	PageSize int
	OrderBy  string
	Where    ORMWhere
	Result   interface{}
}

func (s *ORM) Page(page ORMPage) *Page {
	// set order by
	var db *gorm.DB
	if page.OrderBy != "" {
		db = s.DB.Order(page.OrderBy)
	} else {
		db = s.DB.Order("id desc")
	}
	// set sql where
	if page.Where.Has {
		db = db.Where(page.Where.Cond, page.Where.Param...)
	}
	// find count
	count := 0
	db.Model(page.Result).Count(&count)
	// set page skip and limit
	offset := (page.PageNum - 1) * page.PageSize
	db = db.Limit(page.PageSize).Offset(offset)
	// find result
	db.Find(page.Result)
	// get total page
	totalPage := count / page.PageSize
	if (count % page.PageSize) > 0 {
		totalPage++
	}
	// result
	return &Page{
		List:       page.Result,
		PageNum:    page.PageNum,
		PageSize:   page.PageSize,
		TotalPage:  totalPage,
		TotalCount: count,
	}
}

type ORMWhere struct {
	Has   bool
	Cond  string
	Param []interface{}
}

func Where(where string, params ...interface{}) ORMWhere {
	return ORMWhere{
		Has:   true,
		Cond:  where,
		Param: params,
	}
}
