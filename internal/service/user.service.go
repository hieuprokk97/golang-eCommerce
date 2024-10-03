package service

import "github.com/hieuprokk97/golang-eCommerce.git/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

// user repo
func (us *UserService) GetInfoUserService() string {
	return repo.NewUserRepo().GetInfoUser()
}
