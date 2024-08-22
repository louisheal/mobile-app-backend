package users

type UserRepository interface {
	SearchUsers(string) ([]User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(r UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) SearchUsers(username string) ([]User, error) {
	return s.repo.SearchUsers(username)
}
