package friends

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FriendHandler struct {
	service *FriendService
}

func NewFriendHandler(s *FriendService) *FriendHandler {
	return &FriendHandler{service: s}
}

func (h *FriendHandler) PostFriendRequest(c *gin.Context) {
	var friend FriendInput
	if err := c.BindJSON(&friend); err != nil {
		panic(err)
	}

	err := h.service.CreateFriend(friend)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, nil)
}

func (h *FriendHandler) GetFriendStatus(c *gin.Context) {
	fstUser, err := primitive.ObjectIDFromHex(c.Param("fstUser"))
	if err != nil {
		panic(err)
	}

	sndUser, err := primitive.ObjectIDFromHex(c.Param("sndUser"))
	if err != nil {
		panic(err)
	}

	status, err := h.service.GetFriendStatus(fstUser, sndUser)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, status)
}

func (h *FriendHandler) DeleteFriend(c *gin.Context) {
	fstUser, err := primitive.ObjectIDFromHex(c.Param("fstUser"))
	if err != nil {
		panic(err)
	}

	sndUser, err := primitive.ObjectIDFromHex(c.Param("sndUser"))
	if err != nil {
		panic(err)
	}

	err = h.service.RemoveFriend(fstUser, sndUser)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, nil)
}

func (h *FriendHandler) GetFriendRequests(c *gin.Context) {
	userID, err := primitive.ObjectIDFromHex(c.Query("userID"))
	if err != nil {
		panic(err)
	}

	requests, err := h.service.GetFriendRequests(userID)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, requests)
}

func (h *FriendHandler) GetFriends(c *gin.Context) {
	userID, err := primitive.ObjectIDFromHex(c.Query("userID"))
	if err != nil {
		panic(err)
	}

	friends, err := h.service.GetUsersFriends(userID)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, friends)
}
