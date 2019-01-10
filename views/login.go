package views

import (
	"github.com/gin-gonic/gin"
	"Kiddy/container/form"
	"Kiddy/container/utils"
	"net/http"
)

func GetLogin(c *gin.Context){
	c.HTML(http.StatusOK,"login.html",nil)
	return
}


func Login(c *gin.Context){
	var login_from form.Login
	if c.ShouldBind(&login_from)==nil{
		if login_from.Username== "Start" && login_from.Password == "Start" {
			c.SetCookie("PHPSESSION",utils.SetToken(),60*15,"/","",false,true)
			//c.Redirect(http.StatusMovedPermanently,"/index")
			c.JSON(http.StatusOK,gin.H{"code":200})
			return
		} else {
			c.JSON(http.StatusAccepted, gin.H{"code": 202})
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{"code": "400"})
	return
}
