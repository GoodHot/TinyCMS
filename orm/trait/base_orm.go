package trait

import "time"

type BaseORM struct {
	ID       int        `json:"id"`
	Created  *time.Time `json:"created"`
	Modified *time.Time `json:"modified"`
}
