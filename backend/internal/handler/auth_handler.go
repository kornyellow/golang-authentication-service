package handler

import (
	"errors"
	"net/http"

	"backend/internal/domain"
	"backend/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req domain.RegisterRequest

	// 1. Bind JSON and Validate
	if err := c.ShouldBindJSON(&req); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]string, len(ve))
			for i, fe := range ve {
				out[i] = fe.Field() + ": " + msgForTag(fe)
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": out})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. Let Service handle Registration
	if err := h.service.Register(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 3. Success
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req domain.LoginRequest

	// 1. Bind JSON and Validate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. Let Service handle Login
	token, err := h.service.Login(req)
	if err != nil {
		// If Login fails
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// 3. Success and return Token
	c.JSON(http.StatusOK, gin.H{
		"message":  "Login successful",
		"token":    token,
		"username": req.Username,
	})
}

func (h *AuthHandler) GetProfile(c *gin.Context) {
	username, _ := c.Get("username")

	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the secret zone!",
		"user":    username,
	})
}

func msgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "min":
		return "Must be longer than " + fe.Param() + " characters"
	case "eqfield":
		return "Passwords do not match"
	}
	return fe.Error()
}
