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
		CreateUser(pctx context.Context, req *user.CreateUserReq) (*user.UserProfile, error)
		FindOneUserProfile(pctx context.Context, userId string) (*user.UserProfile, error)
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

func (u *userUsecase) CreateUser(pctx context.Context, req *user.CreateUserReq) (*user.UserProfile, error){
	if (!u.userRepository.IsUniqueUser(pctx, req.Email, req.Username)){
		return nil, errors.New("error: email or username already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("error: hashing password")
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

	return u.FindOneUserProfile(pctx, userId.Hex())
}

func (u *userUsecase) FindOneUserProfile(pctx context.Context, userId string) (*user.UserProfile, error){
	result, err := u.userRepository.FindOneUserProfile(pctx, userId)
	if err != nil {
		return nil, err
	}

	return &user.UserProfile{
		Id: result.Id.Hex(),
		Email: result.Email,
		Username: result.Username,
		CreateAt: result.CreateAt.Format("2006-01-02 15:04:05"),
		UpdateAt: result.UpdateAt.Format("2006-01-02 15:04:05"),
	}, nil
}