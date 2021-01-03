package trait

import "time"

type BaseORM struct {
	ID       int        `json:"id"`
	Created  *time.Time `json:"created"`
	Modified *time.Time `json:"modified"`
	Deleted  bool       `json:"deleted"`
}

type Page struct {
	Page  int         `json:"page"`  // current page number
	Size  int         `json:"size"`  // pre page size
	Count int         `json:"count"` // total data count
	Total int         `json:"pages"` // total page count
	List  interface{} `json:"list"`  // page data result
}

type Query struct {
	Page   int                    `json:"page"`    // current page number
	Size   int                    `json:"size"`    // page size
	LastID int                    `json:"last_id"` // pre page last result id
	Param  map[string]interface{} `json:"param"`   // query param
	Order  map[string]string      `json:"order"`   // order field eg: id->desc, time->asc
}
