package userUsecase

import (
	"context"
	"errors"
	"tansan/modules/user"
	"tansan/modules/user/userRepository"
	"tansan/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

type (
	UserUsecaseService interface {
		CreateUser(pctx context.Context, req *user.CreateUserReq) (string, error)
	}

	userUsecase struct {
		userRepository userRepository.UserRepositoryService
	}
)

func NewUserUsecase(userRepository userRepository.UserRepositoryService) UserUsecaseService {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (u *userUsecase) CreateUser(pctx context.Context, req *user.CreateUserReq) (string, error){
	if (!u.userRepository.IsUniqueUser(pctx, req.Email, req.Username)){
		return "", errors.New("error: email or username already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("error: hashing password")
	}

	// Insert one user
	userId, err := u.userRepository.InsertOneUser(pctx, &user.User{
        Email: req.Email,
        Password: string(hashedPassword),
        Username: req.Username,
        CreateAt: utils.LocalTime(),
        UpdateAt: utils.LocalTime(),
        UserRoles: []user.UserRole{
            {
                RoleTitle: "user",
                RoleCode: 0,
            },
        },
    })

	return userId.Hex(), nil
}