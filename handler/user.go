package handler

import (
	"bwastartup/auth"
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type userHandler struct {
	authService auth.Service
	userService user.Service
}

// Login implements UserHandler
func (h *userHandler) Login(c *gin.Context) {
	var input user.UserLoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		webResponseFail := helper.WebResposne{
			Meta: helper.Meta{
				Code:    http.StatusBadRequest,
				Message: "Login failed " + err.Error(),
			},
		}
		c.JSON(http.StatusBadRequest, webResponseFail)
		return
	}

	userEntity, err := h.userService.Login(input)
	if err != nil {
		webResponseFail := helper.WebResposne{
			Meta: helper.Meta{
				Code:    http.StatusBadRequest,
				Message: "Login failed " + err.Error(),
			},
		}
		c.JSON(http.StatusBadRequest, webResponseFail)
		return
	}

	token, err := h.authService.GenerateToken(userEntity.Id)

	if err != nil {
		webResponseFail := helper.WebResposne{
			Meta: helper.Meta{
				Code:    http.StatusBadRequest,
				Message: "Login failed " + err.Error(),
			},
		}
		c.JSON(http.StatusBadRequest, webResponseFail)
		return
	}

	webResponse := helper.WebResposne{
		Meta: helper.Meta{
			Code:    http.StatusOK,
			Message: "",
		},
		Data: user.FormaterUser(userEntity, token),
	}

	c.JSON(http.StatusOK, webResponse)
}

// Register implements UserHandler
func (h *userHandler) Register(c *gin.Context) {

	var input user.UserRegisterInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		webResponseFail := helper.WebResposne{
			Meta: helper.Meta{
				Code:    http.StatusBadRequest,
				Message: "Register failed" + err.Error(),
			},
		}
		c.JSON(http.StatusBadRequest, webResponseFail)
		return
	}

	userEntity, err := h.userService.Register(input)
	if err != nil {
		webResponseFail := helper.WebResposne{
			Meta: helper.Meta{
				Code:    http.StatusBadRequest,
				Message: "Register failed" + err.Error(),
			},
		}
		c.JSON(http.StatusBadRequest, webResponseFail)
		return
	}

	token, err := h.authService.GenerateToken(userEntity.Id)

	if err != nil {
		webResponseFail := helper.WebResposne{
			Meta: helper.Meta{
				Code:    http.StatusBadRequest,
				Message: "Login failed " + err.Error(),
			},
		}
		c.JSON(http.StatusBadRequest, webResponseFail)
		return
	}

	webResponse := helper.WebResposne{
		Meta: helper.Meta{
			Code:    http.StatusOK,
			Message: "",
		},
		Data: user.FormaterUser(userEntity, token),
	}

	c.JSON(http.StatusOK, webResponse)
}

func NewUserHandler(authService auth.Service, service user.Service) UserHandler {

	return &userHandler{authService, service}

}
