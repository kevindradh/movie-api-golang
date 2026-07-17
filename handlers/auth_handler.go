package handlers

import (
	"MovieAPI/config"
	"MovieAPI/models"
	"MovieAPI/response"
	"MovieAPI/utils"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// Register godoc
// @Summary 		Register new user
// @Description 	Create a new account to access protected features
// @Tags 			auth
// @Accept 			json
// @Produce 		json
// @Param 			user body models.RegisterInput true "Registration data"
// @Success 		201 {object} response.SuccessResponse
// @Failure 		400 {object} response.ErrorResponse
// @Failure 		409 {object} response.ErrorResponse
// @Router 			/auth/register [post]
func Register(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var input models.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, "Invalid registration data", err)
		return
	}

	// Check email already registered
	count, err := config.UserCollection.CountDocuments(ctx, bson.M{"email": input.Email})
	if err != nil {
		response.InternalError(c, "Failed to check email", err)
		return
	}
	if count > 0 {
		c.JSON(409, response.ErrorResponse{
			Status:  "error",
			Message: "Email already registered",
		})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		response.InternalError(c, "Failed to hash password", err)
		return
	}

	newUser := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
		Role:     "user",
	}

	result, err := config.UserCollection.InsertOne(ctx, newUser)
	if err != nil {
		response.InternalError(c, "Failed to insert user", err)
		return
	}

	newUser.ID = result.InsertedID.(primitive.ObjectID)
	newUser.Password = ""

	response.Created(c, "Registration success", newUser)
}

// Login godoc
// @Summary 		Login
// @Description 	Login with email & password, it will return JWT token
// @Tags 			auth
// @Accept 			json
// @Produce 		json
// @Param 			credentials body models.LoginInput true "Email & password"
// @Success 		200 {object} response.SuccessResponse
// @Failure 		400 {object} response.ErrorResponse
// @Failure 		401 {object} response.ErrorResponse
// @Router 			/auth/login [post]
func Login(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var input models.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, "Invalid login data", err)
		return
	}

	var user models.User
	err := config.UserCollection.FindOne(ctx, bson.M{"email": input.Email}).Decode(&user)
	if err != nil {
		response.Unauthorized(c, "Incorrect email or password")
		return
	}

	// Compare password with hashed password in database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		response.Unauthorized(c, "Incorrect email or password")
		return
	}

	token, err := utils.GenerateToken(user.ID.Hex(), user.Email, user.Role)
	if err != nil {
		response.InternalError(c, "Failed to generate token", err)
		return
	}

	response.OK(c, "Login success", gin.H{
		"token": token,
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		},
	})
}
