package di

type User struct {
	ID int64
}

type IRepo interface {
	Get(int64) (*User, error)
}

type Repo struct{}

func (r *Repo) Get(id int64) (*User, error) {
	return &User{ID: id}, nil
}

func NewRepo() IRepo {
	return &Repo{}
}

type Service struct {
	R IRepo
}

func NewService(r IRepo) *Service {
	return &Service{R: r}
}
