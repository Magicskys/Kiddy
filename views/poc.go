package views

import (
	"github.com/gin-gonic/gin"
	"Kiddy/container/form"
	"Kiddy/models"
	"net/http"
)


func Post_Action_POC(c *gin.Context) {
	var action_form form.Action_Sqlmap
	err:=c.ShouldBind(&action_form);if err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"code":404})
		return
	}
	switch action_form.Action {
	case "kill":
		if models.RemovePluginId("plugins",action_form.Id){
			c.JSON(http.StatusNotFound,gin.H{"code":404})
			return
		}
	}
	c.JSON(http.StatusOK,gin.H{"code":200})
	return
}