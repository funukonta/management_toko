package usecases

import (
	"errors"

	"github.com/funukonta/management_toko/internal/api/constants"
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

func (uc *UsecaseUser) Register(req dtos.RegisterReq) (*common.RespData, error) {

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := models.Users{
		Username: req.Username,
		Name:     req.Name,
		Status:   constants.USER_STATUS_ACTIVE,
		Password: string(hashedPass),
		Role:     req.Role,
	}

	if req.Role == "" {
		user.Role = constants.USER_ROLE_USER
	}

	err = uc.repo.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	// to hide the password
	user.Password = "secret!"

	resp := &common.RespData{
		Message: "Register Success",
		Data:    user,
	}

	return resp, nil
}
