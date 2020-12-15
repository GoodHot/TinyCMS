package trait

type Admin struct {
	ID       int
	Username string
	Email    string
	Password string
}

type AdminORM interface {
	GetByUsername(username string) (*Admin, error)
	GetByEmail(email string) (*Admin, error)
}
