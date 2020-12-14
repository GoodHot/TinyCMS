package orm

type Admin struct {
	ID       int
	Username string
	Email    string
	Password string
}

type AdminORM interface {
}
