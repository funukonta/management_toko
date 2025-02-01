package repositories

import (
	"errors"
	"fmt"
	"log"

	"github.com/funukonta/management_toko/internal/api/models"
	"gorm.io/gorm"
)

type RepoUser interface {
	GetUser(models.UserCondition) (*models.Users, error)
	GetUserList()
	CreateUser(*models.Users) error
}

type repoUser struct {
	db *gorm.DB
}

func NewRepoUser(db *gorm.DB) RepoUser {
	return &repoUser{db: db}
}

func (r *repoUser) GetUser(cond models.UserCondition) (*models.Users, error) {
	user := &models.Users{}
	q := r.db.Model(user)
	cond.Where.ImpelementWhere(q)

	err := q.First(user).Error

	return user, err

}

func (r *repoUser) GetUserList() {}

func (r *repoUser) CreateUser(u *models.Users) error {
	q := r.db.Create(u)

	err := q.Error
	if err != nil {
		log.Println("Error on user", err.Error())
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return err
		}

		return fmt.Errorf("USER")
	}

	return nil

}
