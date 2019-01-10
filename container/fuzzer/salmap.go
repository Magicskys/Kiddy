package fuzzer

import (
	"fmt"
	"github.com/gin-gonic/gin/json"
	"Kiddy/models"
	"net/http"
	"strings"
)

type NewTask struct {
	Taskid string `json:"taskid"`
	Success bool `json:"success"`
}

func Join_Sqlmap_Scan(username string,targetUrls []string){
	sqlmap_settings,sqlmap_api:=models.Get_Sqlmap_settings()
	for _,targetUrl:=range targetUrls{
		target,err:=models.GetIdInfo(username,targetUrl)
		if err!=nil{
			continue
		}
		resp,err:=http.Get("http://"+sqlmap_api+"/task/new")
		if err!=nil{
			continue
		}
		var newTask NewTask
		defer resp.Body.Close()
		err=json.NewDecoder(resp.Body).Decode(&newTask);if err!=nil{
			continue
		}
		if newTask.Success==true{
			resp,err=http.Post(fmt.Sprintf("http://"+sqlmap_api+"/option/%s/set",newTask.Taskid),"application/json",strings.NewReader(string(sqlmap_settings)))
			if err!=nil{
				continue
			}
			postValue := map[string]string{}
			postValue["url"]=target.URL
			postValue["Cookie"]=target.Cookies
			//postValue["Referer"]=target
			postValue["data"]=target.Body
			response,_:=json.Marshal(postValue)
			resp,_=http.Post(fmt.Sprintf("http://"+sqlmap_api+"/scan/%s/start",newTask.Taskid),"application/json",strings.NewReader(string(response)))
			models.Insert_Sqlmap_Task(username,targetUrl,newTask.Taskid)
			models.Sqlmap_api_update(sqlmap_api)
		}
	}
}


func Stop_taskId(targetUrls []string) {
	_,sqlmap_api:=models.Get_Sqlmap_settings()
	for _,targetUrl:=range targetUrls{
		_,err:=http.Get(fmt.Sprintf("http://"+sqlmap_api+"/scan/%s/stop",targetUrl))
		if err!=nil{
			fmt.Println("stop taskid sqlmap fail ",targetUrl)
		}
	}
}

func Kill_taskId(targetUrls []string) {
	_,sqlmap_api:=models.Get_Sqlmap_settings()
	for _,targetUrl:=range targetUrls{
		_,err:=http.Get(fmt.Sprintf("http://"+sqlmap_api+"/scan/%s/kill",targetUrl))
		if err!=nil{
			fmt.Println("kill taskid sqlmap fail ",targetUrl)
		}
		_,err=http.Get(fmt.Sprintf("http://"+sqlmap_api+"/task/%s/delete",targetUrl))
		if err!=nil{
			fmt.Println("delete taskid sqlmap fail ",targetUrl)
		}else{
			if !models.Remmove_name_taskId("start",targetUrl){
				fmt.Println("remove mongodb taskid sqlmap fail ",targetUrl)
			}
		}

	}
}