package api

import (
	"net/http"
	"server/e"
	"server/models"
	"server/util"

	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string
	Password string
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	data := make(map[string]interface{})
	code := e.SUCCESS
	if models.CheckAuth(username, password) {
		token, err := util.GenerateToken(username, password)
		if err != nil {
			code = e.ERROR_AUTH_TOKEN
		} else {
			data["token"] = token
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
