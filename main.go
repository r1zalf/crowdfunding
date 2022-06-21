package main

import (
	"bwastartup/auth"
	"bwastartup/handler"
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(localhost:3306)/bwa_startup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// AUTH
	authService := auth.NewService()

	// USER

	userRepository := user.NewRepository(db)

	userService := user.NewService(userRepository)

	userHandler := handler.NewUserHandler(authService, userService)

	router := gin.Default()

	api := router.Group("/api/v1")

	api.POST("/users/register", userHandler.Register)
	api.POST("/users/login", userHandler.Login)

	api.GET("/test", authMiddleware(authService, userService), func(ctx *gin.Context) {
		fmt.Println("test")
	})

	router.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			webResponseFail := helper.WebResposne{
				Meta: helper.Meta{
					Code:    http.StatusUnauthorized,
					Message: "Unauthorized",
				},
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, webResponseFail)
			return
		}

		tokenString := ""

		tokenAuth := strings.Split(authHeader, " ")

		if len(tokenAuth) == 2 {
			tokenString = tokenAuth[len(tokenAuth)-1]

		}

		tokenValidate, err := authService.ValidateToken(tokenString)
		if err != nil {
			webResponseFail := helper.WebResposne{
				Meta: helper.Meta{
					Code:    http.StatusUnauthorized,
					Message: "Unauthorized",
				},
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, webResponseFail)
			return
		}

		claim, ok := tokenValidate.Claims.(jwt.MapClaims)
		if !ok || !tokenValidate.Valid {
			webResponseFail := helper.WebResposne{
				Meta: helper.Meta{
					Code:    http.StatusUnauthorized,
					Message: "Unauthorized",
				},
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, webResponseFail)
			return
		}

		userId := int64(claim["user_id"].(float64))

		userEntity, err := userService.GetUserById(userId)
		if err != nil {
			webResponseFail := helper.WebResposne{
				Meta: helper.Meta{
					Code:    http.StatusUnauthorized,
					Message: "Unauthorized",
				},
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, webResponseFail)
			return
		}

		c.Set("currentUser", userEntity)
	}
}
