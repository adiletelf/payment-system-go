package handler

import (
	"net/http"

	"github.com/adiletelf/payment-system-go/internal/model"
	"github.com/adiletelf/payment-system-go/internal/token"
	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *BaseHandler) Register(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	a := model.Admin{}

	a.Username = input.Username
	a.Password = input.Password

	_, err := h.ar.Save(&a)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "registration success"})
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *BaseHandler) Login(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	a := model.Admin{}
	a.Username = input.Username
	a.Password = input.Password

	accessToken, refreshToken, err := h.ar.LoginCheck(a.Username, a.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accessToken": accessToken,
		"refreshToken": refreshToken,
	})
}

func (h *BaseHandler) CurrentAdmin(c *gin.Context) {
	admin_id, err := token.ExtractTokenID(c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	a, err := h.ar.GetAdminById(admin_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": a})
}

type RefreshInput struct {
	RefreshToken string `json:"refreshToken"`
}

func (h *BaseHandler) Refresh(c *gin.Context) {
	var input RefreshInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
