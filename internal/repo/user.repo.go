package repo

// type UserRepo struct {}

// func NewUserRepo() *UserRepo {
// 	return &UserRepo{}
// }

// func (ur *UserRepo) GetInfoUser() string {
// 	return "tung-lee"
// }

// interface version

type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct{}



func NewUserRepository() IUserRepository {
	return &userRepository{}
}

// GetUserByEmail implements IUserRepository.
func (ur *userRepository) GetUserByEmail(email string) bool {
	return true
}
