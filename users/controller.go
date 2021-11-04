package users

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prakharmaurya/gin_test/conn"
	"gopkg.in/mgo.v2/bson"
)

const UserCollection = "user"

var (
	errNotExist        = errors.New("users are not exist")
	errInvalidID       = errors.New("invalid ID")
	errInvalidBody     = errors.New("invalid request body")
	errInsertionFailed = errors.New("error in the user insertion")
	errUpdationFailed  = errors.New("error in the user updation")
	errDeletionFailed  = errors.New("error in the user deletion")
)

// GetAllUser Endpoint
func GetAllUser(c *gin.Context) {
	// Get DB from Mongo Config
	db := conn.GetMongoDB()
	users := make([]User, 0, 10)
	err := db.C(UserCollection).Find(nil).All(&users)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errNotExist.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "users": &users})
}

// GetUser Endpoint
func GetUser(c *gin.Context) {
	var id bson.ObjectId = bson.ObjectIdHex(c.Param("id")) // Get Param
	user, err := UserInfo(id, UserCollection)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidID.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "user": &user})
}

// CreateUser Endpoint
func CreateUser(c *gin.Context) {
	// Get DB from Mongo Config
	db := conn.GetMongoDB()
	user := User{}
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidBody.Error()})
		return
	}
	user.ID = bson.NewObjectId()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	err = db.C(UserCollection).Insert(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInsertionFailed.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "user": &user})
}

// UpdateUser Endpoint
func UpdateUser(c *gin.Context) {
	// Get DB from Mongo Config
	db := conn.GetMongoDB()
	var id bson.ObjectId = bson.ObjectIdHex(c.Param("id")) // Get Param
	existingUser, err := UserInfo(id, UserCollection)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidID.Error()})
		return
	}
	// user := user.User{}
	err = c.Bind(&existingUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidBody.Error()})
		return
	}
	existingUser.ID = id
	existingUser.UpdatedAt = time.Now()
	err = db.C(UserCollection).Update(bson.M{"_id": &id}, existingUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errUpdationFailed.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "user": &existingUser})
}

// DeleteUser Endpoint
func DeleteUser(c *gin.Context) {
	// Get DB from Mongo Config
	db := conn.GetMongoDB()
	var id bson.ObjectId = bson.ObjectIdHex(c.Param("id")) // Get Param
	err := db.C(UserCollection).Remove(bson.M{"_id": &id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errDeletionFailed.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "User deleted successfully"})
}
