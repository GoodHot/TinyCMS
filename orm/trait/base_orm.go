package trait

import "time"

type BaseORM struct {
	ID       int
	Created  *time.Time
	Modified *time.Time
}
