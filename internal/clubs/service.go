package clubs

type ClubRepository interface {
	GetAllClubs() ([]Club, error)
}

type ClubService struct {
	repo ClubRepository
}

func NewClubService(r ClubRepository) *ClubService {
	return &ClubService{repo: r}
}

func (s *ClubService) GetAllClubs() ([]Club, error) {
	return s.repo.GetAllClubs()
}
