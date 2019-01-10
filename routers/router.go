package routers

import (
	"github.com/gin-gonic/gin"
	"Kiddy/container/middleware"
	"Kiddy/models"
	"Kiddy/views"
)

func InitRouters() *gin.Engine {
	r:=gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static","./static/")
	r.GET("/",views.GetLogin)
	r.POST("/login",views.Login)
	r.GET("/index",views.GetIndex)
	v1:=r.Group("/v1",middleware.MyMiddelware())
	v1.GET("info",views.GetInfo)

	result:=v1.Group("/result")
	result.GET("sql",views.GetResultSql)
	result.GET("general",views.GetResultGeneral)

	poc:=v1.Group("/poc")
	poc.GET("list",views.GetPocList)

	sett:=v1.Group("/setting")
	sett.GET("sql",views.GetSettingsql)
	sett.POST("sql",views.PostSettingsql)
	sett.GET("general",views.GetSettinggeneral)
	sett.POST("general",views.PostSettinggeneral)
	sett.POST("poc",views.PostPoc)
	sett.POST("monitor",views.PostMonitor)

	action:=v1.Group("action")
	action.POST("sql",views.Post_Action_Sqlmap)
	action.POST("general",views.Post_Action_General)
	action.POST("poc",views.Post_Action_POC)

	task:=v1.Group("/task")
	task.POST("sql",views.PostTaskSql)
	task.POST("general",views.PostTaskGeneral)

	models.InitData()
	models.Init_Sqlmap_Start()
	models.Init_SqlmapApi_Result()
	models.Init_Settings()
	return r
}