package handlers

import (
	"git/config"
	"git/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var user models.User

func RegisterHandler(c *gin.Context) {

	// var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Dados inválidos"})
		return
	}
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao registrar usuário"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"mensagem": "Usuário registrado"})
}

func LoginHandler(c *gin.Context) {
	var loginData models.User

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Dados inválidos"})
		return
	}

	if err := config.DB.Where("email = ? AND password = ?", loginData.Email, loginData.Password).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "Credenciais inválidas"})
		return
	}

	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Save()

	c.JSON(http.StatusOK, gin.H{"mensagem": "Login realizado com sucesso"})
}
func LogoutHandler(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("user_id")
	session.Save()

	c.JSON(http.StatusOK, gin.H{"mensagem": "Logout realizado com sucesso"})
}
func DashboardHandler(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user_id")

	// var user User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Usuário não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensagem": "Bem-vindo", "email": user.Email})
}
