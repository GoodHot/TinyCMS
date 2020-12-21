package trait

type Admin struct {
	ID       int
	Username string
	Email    string
	Password string
}

type AdminORM interface {
	Initial() error
	GetByUsername(username string) (*Admin, error)
	GetByEmail(email string) (*Admin, error)
}
