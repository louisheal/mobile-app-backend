package friends

import (
	"mobile-app-backend/internal/users"
)

type FriendRepository interface {
	CreateFriend(FriendInput) error
	FriendExists(users.UserID, users.UserID) (bool, error)
	DeleteFriend(users.UserID, users.UserID) error
	GetUsersFriends(users.UserID) ([]FriendInput, error)
	GetUser(users.UserID) (users.User, error)
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

	if err := s.repo.DeleteFriend(sndUser, fstUser); err != nil {
		return err
	}

	return nil
}

func (s *FriendService) GetFriendRequests(userID users.UserID) ([]users.User, error) {
	friends, err := s.repo.GetUsersFriends(userID)
	if err != nil {
		return []users.User{}, err
	}

	result := []users.User{}
	for _, friend := range friends {
		exists, err := s.repo.FriendExists(userID, friend.Sender)
		if err != nil {
			return []users.User{}, err
		}
		if !exists {
			friend, err := s.repo.GetUser(friend.Sender)
			if err != nil {
				return []users.User{}, err
			}
			result = append(result, friend)
		}
	}

	return result, nil
}
