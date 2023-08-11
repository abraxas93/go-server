package user

type IUserService interface {
}

type UserService struct {
	repo UserRepository
}

func NewService(r *UserRepository) {

}

func CreateNewUser() {

}
