package main

import (
	"bwastartup/auth"
	"bwastartup/handler"
	"bwastartup/user"

	"github.com/gin-gonic/gin"
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

	router.Run()
}
