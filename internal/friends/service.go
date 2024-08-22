package friends

import (
	"mobile-app-backend/internal/users"
)

type FriendRepository interface {
	CreateFriend(FriendInput) error
	FriendExists(users.UserID, users.UserID) (bool, error)
	DeleteFriend(users.UserID, users.UserID) error
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

func (s *FriendService) GetFriendStatus(fstUser users.UserID, sndUser users.UserID) (FriendStatus, error) {
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

func (s *FriendService) RemoveFriend(fstUser users.UserID, sndUser users.UserID) error {
	if err := s.repo.DeleteFriend(fstUser, sndUser); err != nil {
		return err
	}

	return s.repo.DeleteFriend(sndUser, fstUser)
}
