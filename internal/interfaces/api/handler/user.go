package handler

import (
	"errors"
	"go-example/internal/application/usecase"
	e "go-example/internal/utils/errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	GetUser(c *gin.Context)
	PostUser(c *gin.Context)
	PatchUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type userHandler struct {
	u usecase.UserUseCase
}

func NewUserHandler(u usecase.UserUseCase) UserHandler {
	return &userHandler{u}
}

func (h userHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	log.Print(id)
	if id == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	u, err := h.u.GetUserByID(c, id)
	if err != nil {
		var en *e.ErrNotFound
		if errors.As(err, &en) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		log.Printf("failed to get user: %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": u,
	})
}

func (h userHandler) PostUser(c *gin.Context) {
	var d struct {
		FirstName string    `json:"firstName" binding:"required,min=1,max=30"`
		LastName  string    `json:"lastName" binding:"required,min=1,max=30"`
		Email     string    `json:"email" binding:"required,email"`
		Birthdate time.Time `json:"birthdate" binding:"required"`
	}
	if err := c.BindJSON(&d); err != nil {
		log.Printf("failed to parse json: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	u, err := h.u.CreateUser(c, d.FirstName, d.LastName, d.Email, d.Birthdate)
	if err != nil {
		var ei *e.ErrInvalidData
		if errors.As(err, &ei) {
			log.Printf("invalid data: %v", err)
			c.AbortWithStatus(http.StatusBadRequest)
		}
		log.Printf("failed to get user: %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"user": u,
	})
}

func (h userHandler) PatchUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var d struct {
		FirstName *string    `json:"firstName,omitempty" binding:"omitempty,min=1,max=30"`
		LastName  *string    `json:"lastName,omitempty" binding:"omitempty,min=1,max=30"`
		Email     *string    `json:"email,omitempty" binding:"omitempty,email"`
		Birthdate *time.Time `json:"birthdate,omitempty" binding:"omitempty"`
	}
	if err := c.BindJSON(&d); err != nil {
		log.Printf("failed to parse json: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	u, err := h.u.UpdateUser(c, id, d.FirstName, d.LastName, d.Email, d.Birthdate)
	if err != nil {
		var ei *e.ErrInvalidData
		if errors.As(err, &ei) {
			log.Printf("invalid data: %v", err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		var en *e.ErrNotFound
		if errors.As(err, &en) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		log.Printf("failed to get user: %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": u,
	})
}

func (h userHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err := h.u.DeleteUser(c, id); err != nil {
		var en *e.ErrNotFound
		if errors.As(err, &en) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		log.Printf("failed to get user: %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusNoContent)
}
