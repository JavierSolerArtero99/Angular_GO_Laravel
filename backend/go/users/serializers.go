package users

import (
	"github.com/gin-gonic/gin"

	"App/common"
)

type ProfileSerializer struct {
	C *gin.Context
	users
}

// Declare your response schema here
type ProfileResponse struct {
	ID        uint    `json:"id"`
	Username  string  `json:"username"`
	Bio       string  `json:"bio"`
	Image     *string `json:"image"`
	Following bool    `json:"following"`
	Role      bool    `json:"role"`
}

// Put your response logic including wrap the userModel here.
func (self *ProfileSerializer) Response() ProfileResponse {
	myusers := self.C.MustGet("my_user_model").(users)
	profile := ProfileResponse{
		ID:        self.ID,
		Username:  self.Username,
		Bio:       self.Bio,
		Image:     self.Image,
		Following: myusers.isFollowing(self.users),
		Role:      self.Role,
	}
	return profile
}

type UserSerializer struct {
	c *gin.Context
}

type UserResponse struct {
	ID       uint    `json:"id"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Bio      string  `json:"bio"`
	Image    *string `json:"image"`
	Token    string  `json:"token"`
	Role     bool    `json:"role"`
}

func (self *UserSerializer) Response() UserResponse {
	myusers := self.c.MustGet("my_user_model").(users)
	user := UserResponse{
		ID:       myusers.ID,
		Username: myusers.Username,
		Email:    myusers.Email,
		Bio:      myusers.Bio,
		Image:    myusers.Image,
		Token:    common.GenToken(myusers.ID),
		Role:     myusers.Role,
	}
	return user
}
