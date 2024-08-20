package routes

import (
	"mobile-app-backend/dao"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (routes Routes) PostFriendRequest(c *gin.Context) {
	var friendRequest dao.FriendRequest
	if err := c.BindJSON(&friendRequest); err != nil {
		panic(err)
	}

	err := routes.database.SendFriendRequest(friendRequest)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, nil)
}

func (routes Routes) GetFriendRequestStatus(c *gin.Context) {
	fstUser, err := primitive.ObjectIDFromHex(c.Param("fstUser"))
	if err != nil {
		panic(err)
	}

	sndUser, err := primitive.ObjectIDFromHex(c.Param("sndUser"))
	if err != nil {
		panic(err)
	}

	status, err := routes.database.GetFriendRequestStatus(fstUser, sndUser)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, status)
}
