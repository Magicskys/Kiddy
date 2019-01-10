package views

import (
	"github.com/gin-gonic/gin"
	"Kiddy/container/form"
	"Kiddy/container/fuzzer"
	"net/http"
)

func Post_Action_General(c *gin.Context) {
	var action_form form.Action_Sqlmap
	err:=c.ShouldBind(&action_form);if err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"code":404})
	}
	switch action_form.Action {
	case "kill":
		fuzzer.Kill_Nmap_taskId(action_form.Id)
	}
	c.JSON(http.StatusOK,gin.H{"code":200})
}