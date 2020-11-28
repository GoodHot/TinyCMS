package common

type ORM struct {
	Type    string    `val:"${props.orm.type}"`
	LevelDB *struct{} `val:"${props.orm.levelDB}"`
	MongoDB *struct{} `val:"${props.orm.mongodb}"`
	ins     DBOperator
}

func (my *ORM) Ins() DBOperator {
	if my.ins != nil {
		return my.ins
	}
	if my.Type == "mongodb" {
		my.ins = &MongoDB{}
	} else {
		my.ins = &ObjectBox{}
	}
	return my.ins
}

type DBOperator interface {
}

type MongoDB struct {
}

type ObjectBox struct {
}
