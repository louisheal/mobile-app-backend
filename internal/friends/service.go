package friends

import "go.mongodb.org/mongo-driver/bson/primitive"

type FriendRepository interface {
	CreateFriend(FriendInput) error
	FriendExists(primitive.ObjectID, primitive.ObjectID) (bool, error)
}

type FriendService struct {
	repo FriendRepository
}

func NewFriendService(r FriendRepository) *FriendService {
	return &FriendService{repo: r}
}

func (s *FriendService) CreateFriend(friend FriendInput) error {
	return s.repo.CreateFriend(friend)
}

func (s *FriendService) GetFriendStatus(fstUser primitive.ObjectID, sndUser primitive.ObjectID) (FriendStatus, error) {
	fstExists, err := s.repo.FriendExists(fstUser, sndUser)
	if err != nil {
		return None, err
	}

	sndExists, err := s.repo.FriendExists(sndUser, fstUser)
	if err != nil {
		return None, err
	}

	var status FriendStatus
	switch {
	case fstExists && sndExists:
		status = Accepted
	case sndExists:
		status = Accept
	case fstExists:
		status = Pending
	default:
		status = Send
	}

	return status, nil
}
