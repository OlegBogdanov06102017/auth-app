package repository

type Authorization interface {
}

type List interface {
}

type Repository struct {
	Authorization
	List
}

func NewRepository() *Repository {
	return &Repository{}
}
