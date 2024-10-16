package service

import (
	"backend-go/internal/repo"
	"backend-go/pkg/response"
)

// type UserService struct {
// 	userRepo *repo.UserRepo
// }

// func NewUserService() *UserService {
// 	return &UserService{
// 		userRepo: repo.NewUserRepo(),
// 	}
// }

// func (us *UserService) GetInfoUser() string {
// 	return us.userRepo.GetInfoUser()
// }

// interface version

type IUserService interface {
	Register(email string, purpose string) int
	// ...
}

type userService struct {
	userRepo repo.IUserRepository
	// ...
}

func NewUserService(userRepo repo.IUserRepository) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (us *userService) Register(email string, purpose string) int {
	// 1. Check
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrUserExists
	}
	
	return response.Success
}
