package repository

import (
	"raytheon/datamodels"

	"github.com/jinzhu/gorm"
)

type Query func(user datamodels.User) bool

//UserRepository user repository interface
type UserRepository interface {
	FetchOne(query Query) (user datamodels.User, found bool)
	FetchAll(query Query, limit int) (results []datamodels.User)
	InsertOrUpdate(user datamodels.User) (updateUser datamodels.User, err error)
}

type UserDBRepository struct {
	db *gorm.DB
}

func (r *UserDBRepository) FetchOne(query Query) (user datamodels.User, found bool) {
	panic("implement me")
}

func (r *UserDBRepository) FetchAll(query Query, limit int) (results []datamodels.User) {
	panic("implement me")
}

func (r *UserDBRepository) InsertOrUpdate(user datamodels.User) (updateUser datamodels.User, err error) {
	panic("implement me")
}
