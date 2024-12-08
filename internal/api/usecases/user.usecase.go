package usecases

import (
	"errors"

	"github.com/funukonta/management_toko/internal/api/dtos"
	"github.com/funukonta/management_toko/internal/api/models"
	"github.com/funukonta/management_toko/internal/api/repositories"
	"github.com/funukonta/management_toko/pkg/common"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UsecaseUser struct {
	repo repositories.RepoUser
}

func NewUsecaseUser(repo repositories.RepoUser) *UsecaseUser {
	return &UsecaseUser{repo: repo}
}

func (uc *UsecaseUser) Login(req dtos.LoginReq) (*common.RespData, error) {
	user, err := uc.repo.GetUser(models.UserCondition{
		Where: models.UserWhere{
			Username: req.Username,
		},
	})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("wrong username/password")
		}
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("wrong username/password")
	}

	resp := &common.RespData{
		Message: "Login Success",
		Data:    map[string]any{"token": "exampleToken!"},
	}

	return resp, nil
}
