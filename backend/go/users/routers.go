package users

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"App/common"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func UsersRegister(router *gin.RouterGroup) {
	router.GET("/redis/", Redis)
	router.POST("/", UsersRegistration)
	router.POST("/login", UsersLogin)
	router.POST("/logout", UsersLogout)
}

func Redis(c *gin.Context) {
	// username := c.Param("username")
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("PETICION A LOS STATS DE LOS USUARIOS")
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	result, err := client.Get("current_users").Result()
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{"current_users": result})
}

func UserRegister(router *gin.RouterGroup) {
	router.GET("/", UserRetrieve)
	router.PUT("/", UserUpdate)
}

func ProfileRegister(router *gin.RouterGroup) {
	router.GET("/:username", ProfileRetrieve)
	router.POST("/:username/follow", ProfileFollow)
	router.DELETE("/:username/follow", ProfileUnfollow)
}

func ProfileRetrieve(c *gin.Context) {
	username := c.Param("username")
	userModel, err := FindOneUser(&users{Username: username})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("profile", errors.New("Invalid username")))
		return
	}
	profileSerializer := ProfileSerializer{c, userModel}
	c.JSON(http.StatusOK, gin.H{"profile": profileSerializer.Response()})
}

func ProfileFollow(c *gin.Context) {
	username := c.Param("username")
	userModel, err := FindOneUser(&users{Username: username})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("profile", errors.New("Invalid username")))
		return
	}
	myusers := c.MustGet("my_user_model").(users)
	err = myusers.following(userModel)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	serializer := ProfileSerializer{c, userModel}
	c.JSON(http.StatusOK, gin.H{"profile": serializer.Response()})
}

func ProfileUnfollow(c *gin.Context) {
	username := c.Param("username")
	userModel, err := FindOneUser(&users{Username: username})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("profile", errors.New("Invalid username")))
		return
	}
	myusers := c.MustGet("my_user_model").(users)

	err = myusers.unFollowing(userModel)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	serializer := ProfileSerializer{c, userModel}
	c.JSON(http.StatusOK, gin.H{"profile": serializer.Response()})
}

func UsersRegistration(c *gin.Context) {
	userModelValidator := NewusersValidator()
	if err := userModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("validation", err))
		return
	}

	if err := SaveOne(&userModelValidator.userModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.Set("my_user_model", userModelValidator.userModel)
	serializer := UserSerializer{c}
	c.JSON(http.StatusCreated, gin.H{"user": serializer.Response()})
}

func UsersLogin(c *gin.Context) {
	loginValidator := NewLoginValidator()
	if err := loginValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	userModel, err := FindOneUser(&users{Email: loginValidator.userModel.Email})

	if err != nil {
		c.JSON(http.StatusForbidden, common.NewError("login", errors.New("Not Registered email or invalid password")))
		return
	}

	if userModel.checkPassword(loginValidator.User.Password) != nil {
		c.JSON(http.StatusForbidden, common.NewError("login", errors.New("Not Registered email or invalid password")))
		return
	}

	UpdateContextusers(c, userModel.ID)
	serializer := UserSerializer{c}

	// Let's encrypt

	id := serializer.Response().ID
	username := serializer.Response().Username
	password := loginValidator.User.Password

	// Encrypting id
	encrypted := strconv.Itoa(int(id)*len(username)) + "@"

	// Encrypting username
	for _, char := range username {
		encrypted += strconv.Itoa(int(char)*len(username)) + "%"
	}

	encrypted = encrypted[:len(encrypted)-1] + "#"

	// Encrypting password
	for _, char := range password {
		encrypted += strconv.Itoa(int(char)*int(char)*int(id)) + "&"
	}

	encrypted = encrypted[:len(encrypted)-1] + "!"

	// Encrypting ip
	for _, char := range strings.Split(c.ClientIP(), ".") {
		num, err := strconv.Atoi(char)
		if err != nil {
			return
		}
		encrypted += strconv.Itoa(num*num) + "$"
	}

	encrypted = encrypted[:len(encrypted)-1]

	// Save new tempkey on db

	err = UpdateTempkey(id, encrypted)
	if err != nil {
		return
	}

	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	json, err := json.Marshal(serializer.Response())
	if err != nil {
		fmt.Println(err)
	}

	// Save user loged on redis
	err = client.Set(serializer.Response().Username, json, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	// Add new user connected
	value, err := client.Get("current_users").Result()

	num := 0
	if len(value) > 0 {
		num, err = strconv.Atoi(value)
		if err != nil {
			return
		}

	}

	err = client.Set("current_users", num+1, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	// Add new user connected
	value, err = client.Get("total_users").Result()

	num = 0
	if len(value) > 0 {
		num, err = strconv.Atoi(value)
		if err != nil {
			return
		}

	}

	err = client.Set("total_users", num+1, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})
}

func UsersLogout(c *gin.Context) {
	loginValidator := NewLoginValidator()
	if err := loginValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	userModel, err := FindOneUser(&users{Email: loginValidator.userModel.Email})
	UpdateContextusers(c, userModel.ID)
	serializer := UserSerializer{c}

	err = UpdateTempkey(serializer.Response().ID, "")
	if err != nil {
		return
	}

	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	value, err := client.Get("current_users").Result()

	num := 1
	if len(value) > 0 {
		num, err = strconv.Atoi(value)
		if err != nil {
			return
		}

	}

	err = client.Set("current_users", num-1, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	err = client.Del(serializer.Response().Username).Err()
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{"user": "success"})
}

func UserRetrieve(c *gin.Context) {
	serializer := UserSerializer{c}
	c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})
}

func UserUpdate(c *gin.Context) {
	myusers := c.MustGet("my_user_model").(users)
	userModelValidator := NewusersValidatorFillWith(myusers)
	if err := userModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	userModelValidator.userModel.ID = myusers.ID
	if err := myusers.Update(userModelValidator.userModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	UpdateContextusers(c, myusers.ID)
	serializer := UserSerializer{c}
	c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})
}
