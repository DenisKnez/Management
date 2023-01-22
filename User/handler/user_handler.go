package handler

import (
	"io/ioutil"
	"net/http"

	"github.com/DenisKnez/management/user/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService service.Service
}

func (handler *UserHandler) UploadFile(c *gin.Context) {
	name := c.PostForm("name")
	formFile, err := c.FormFile("data")
	if err != nil {
		c.Error(err)
		return
	}

	file, err := formFile.Open()
	if err != nil {
		c.Error(err)
		return
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		c.Error(err)
		return
	}

	err = handler.UserService.UploadFile(c.Request.Context(), service.File{
		Name: name,
		Data: data,
	})
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

type CreateUserRequest struct {
	Name string `json:"name"`
}

func (handler *UserHandler) CreateUser(c *gin.Context) {
	req := CreateUserRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, "invalid request body provided")
		return
	}

	err = handler.UserService.CreateUser(c.Request.Context(), service.User{
		Name: req.Name,
	})
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, "failed to create user")
		return
	}

	c.JSON(http.StatusCreated, nil)
}

type UpdateUserRequest struct {
	Name string `json:"name"`
}

func (handler *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	req := UpdateUserRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, "invalid request body provided")
		return
	}

	err = handler.UserService.UpdateUser(c.Request.Context(), id, service.User{
		Name: req.Name,
	})
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, "failed to update user")
		return
	}

	c.JSON(http.StatusOK, nil)
}
