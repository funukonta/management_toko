package repositories

import (
	"github.com/funukonta/management_toko/internal/api/models"
	"gorm.io/gorm"
)

type RepoUser interface {
	GetUser(models.UserCondition) (*models.Users, error)
	GetUserList()
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
	populateUserWithWhere(q, cond.Where)

	err := q.First(user).Error

	return user, err

}

func populateUserWithWhere(db *gorm.DB, where models.UserWhere) {
	if where.ID != "" {
		db.Where("id = ?", where.ID)
	}

	if where.Username != "" {
		db.Where("username = ?", where.Username)
	}
}

func (r *repoUser) GetUserList() {}
