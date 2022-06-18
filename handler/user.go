package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Register(c *gin.Context)
}

type userHandler struct {
	service user.Service
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

	newUser, err := h.service.Register(input)
	if true {
		webResponseFail := helper.WebResposne{
			Meta: helper.Meta{
				Code:    http.StatusBadRequest,
				Message: "Register failed" + err.Error(),
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
		Data: user.FormaterUser(newUser, ""),
	}

	c.JSON(http.StatusOK, webResponse)
}

func NewUserHandler(service user.Service) UserHandler {

	return &userHandler{service}

}
