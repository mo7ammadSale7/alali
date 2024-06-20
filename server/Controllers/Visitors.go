package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/Application"
	"server/Models"
)

func CreateUser(c *gin.Context) {
	req := Application.NewRequest(c).Auth()
	if !req.IsAuth {
		req.Response(http.StatusUnauthorized, "You're not Authorized!", []interface{}{})
		return
	}

	user := Models.User{
		Username: "Mo7ammad",
		Email:    "mo7mmad57@gmail.com",
		Password: "Mo7@100842",
	}
	req.DB.Create(&user)
	req.Response(http.StatusCreated, "User Created successfully", user)
}
