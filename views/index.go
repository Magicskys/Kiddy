package views

import (
	"context"
	"github.com/gin-gonic/gin"
	"Kiddy/container/form"
	"Kiddy/container/fuzzer"
	"Kiddy/container/utils"
	"Kiddy/models"
	"net/http"
)


func GetIndex(c *gin.Context){
	c.HTML(http.StatusOK,"index.html",nil)
	return
}

func GetInfo(c *gin.Context)  {
	infos,err:=models.GetAllInfo("start")
	if err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"code":404})
		return
	}
	c.JSON(http.StatusOK,gin.H{"code":200,"data":infos,"username":"start"})
	return
}

func GetResultSql(c *gin.Context)  {
	result,err:=models.GetAllResultSql("start")
	if err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"code":404})
		return
	}
	c.JSON(http.StatusOK,gin.H{"code":200,"data":result,"username":"start"})
	return
}

func GetResultGeneral(c *gin.Context){
	result,err:=models.GetAllResultGeneral("start")
	if err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"code":404})
		return
	}
	c.JSON(http.StatusOK,gin.H{"code":200,"data":result,"username":"start"})
	return
}

func GetSettingsql(c *gin.Context){
	result,err:=models.Get_Sqlmap_settings_struct("settings")
	if err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"code":404})
		return
	}
	c.JSON(http.StatusOK,gin.H{"code":200,"data":result,"username":"start"})
	return
}

func PostTaskSql(c *gin.Context)  {
	uid:=c.PostFormArray("uid[]")
	go fuzzer.Join_Sqlmap_Scan("start",uid)
	c.JSON(http.StatusOK,gin.H{"code":200,"data":uid})
}

func PostSettingsql(c *gin.Context)  {
	var post_form form.Post_Sqlmap_setting
	err:=c.ShouldBind(&post_form);if err!=nil{
		c.JSON(http.StatusNotFound,gin.H{})
		return
	}
	if len(post_form.Sqlmap_localhost)==0 || !utils.Exists(post_form.Sqlmap_localhost){
		c.JSON(http.StatusOK,gin.H{"code":400,"message":"update file fail , the file doest not exist"})
		return
	}
	result:=models.Update_Sqlmap_settings("settings",&post_form)
	if result==false{
		c.JSON(http.StatusNotFound,gin.H{"code":404,"message":"update fail"})
		return
	}
	c.JSON(http.StatusOK,gin.H{"code":200})
	return
}

func GetSettinggeneral(c *gin.Context)  {
	result,err:=models.Get_General_settings_struct("settings")
	if err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"code":404})
		return
	}
	plugins,err:=models.GetAllPlugins("plugins")
	if err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"code":404})
		return
	}
	c.JSON(http.StatusOK,gin.H{"code":200,"data":result,"username":"start","plugins":plugins})
	return
}

func PostSettinggeneral(c *gin.Context)  {
	var post_form form.Post_General_setting
	err:=c.ShouldBind(&post_form);if err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"code":400})
		return
	}
	portSchema:=[]string{"sS", "sT", "sU", "sF", "sX", "sN", "sW", "sV","sP","sA"}
	for _,j:=range portSchema{
		if post_form.PortSchema==j{
			break
		}else{
			post_form.PortSchema="sS"
		}
	}
	result:=models.Update_General_settings("settings",&post_form)
	if result==false{
		c.JSON(http.StatusNotFound,gin.H{"code":400,"message":"update fail"})
		return
	}
	c.JSON(http.StatusOK,gin.H{"code":200})
	return
}

func PostTaskGeneral(c *gin.Context)  {
	uid:=c.PostFormArray("uid[]")
	plugins:=models.GetUseAllPlugins("plugins")
	if plugins==nil{
		c.JSON(http.StatusNotFound,gin.H{})
		return
	}
	go fuzzer.General_Scan("start",uid,plugins)
	c.JSON(http.StatusOK,gin.H{"code":200,"data":uid})
	return
}

func GetPocList(c *gin.Context)  {
	plugins,err:=models.GetAllPlugins("plugins")
	if err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"code":404})
		return
	}
	c.JSON(http.StatusOK,gin.H{"code":200,"plugins":plugins})
	return
}

func PostPoc(c *gin.Context) {
	var post_form form.PostPlugin
	err:=c.ShouldBind(&post_form);if err!=nil{
		c.JSON(http.StatusNotFound,gin.H{})
		return
	}
	danger:=true
	for _,i:=range []string{"low","medium","high"}{
		if post_form.Danger==i{
			danger=false
			break
		}
	}
	if danger{
		c.JSON(http.StatusAccepted,gin.H{"code":202,"message":"danger"})
		return
	}
	result:=models.InsertPlugin("plugins",&post_form)
	if !result{
		c.JSON(http.StatusNotFound,gin.H{"code":404,"message":"insert fail"})
		return
	}
	c.JSON(http.StatusOK,gin.H{"code":200})
	return
}

func PostMonitor(c *gin.Context) {
	monitor,err:=c.GetPostForm("monitor")
	if !err{
		c.JSON(http.StatusNotFound,gin.H{})
		return
	}
	switch monitor {
	case "true":
		if models.MonitorCancel==nil{
			var ctx context.Context
			ctx, models.MonitorCancel = context.WithCancel(context.Background())
			go models.Select_Monitor_Start(ctx)
		}
	case "false":
		if models.MonitorCancel!=nil{
			models.MonitorCancel()
			models.MonitorCancel=nil
		}
	default:
		if models.MonitorCancel==nil{
			c.JSON(http.StatusOK,gin.H{"code":201,"message":false})
		}else{
			c.JSON(http.StatusOK,gin.H{"code":201,"message":true})
		}
		return
	}
	c.JSON(http.StatusOK,gin.H{"code":200})
	return
}