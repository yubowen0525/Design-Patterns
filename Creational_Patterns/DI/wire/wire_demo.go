package wi

import "github.com/google/wire"

type User struct {
	ID int64
}

type Repo interface {
	Get(int64) (*User, error)
}

type repo struct{}

func (r *repo) Get(id int64) (*User, error) {
	return &User{ID: id}, nil
}

func NewRepo() Repo {
	return &repo{}
}

var (
	Set = wire.NewSet(
		NewRepo,
	)
)

type Service struct {
	R Repo
}

func NewService(r Repo) (*Service, error) {
	return &Service{R: r}, nil
}

func InitApp() (*Service, error) {
	panic(wire.Build(Set, NewService))
}
