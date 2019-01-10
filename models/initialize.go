package models

import (
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)



type AdminList struct {
	Tasks map[string]string	`json:"tasks"`
	Tasks_num int64	`json:"tasks_num"`
	Success	bool `json:"success"`

}

func Init_SqlmapApi_Result()  {
	refres,sqlmap_api:=Get_Sqlmap_autRefres();if refres{
		var ctx context.Context
		ctx, SqlmapCancel = context.WithCancel(context.Background())
		go Select_Sqlmap_Api(ctx,sqlmap_api)
	}
}

func Init_Sqlmap_Start() {
	start,sqlmap_location,sqlmap_api:=Get_Sqlmap_Start();if start && sqlmap_location!="" && sqlmap_api!=""{
		var ctx context.Context
		ctx, SqlmapCancel = context.WithCancel(context.Background())
		go Select_Sqlmap_Start(ctx,sqlmap_location,sqlmap_api)
	}else{
		Update_Sqlmap_settings_Initial_Start("settings")
	}
}

func Sqlmap_api_update(sqlmap_api string){
	client:=&http.Client{
		Timeout:3*time.Second,
	}
	resp,err:=client.Get("http://"+sqlmap_api+"/admin/list")
	if err!=nil{
		mgoc,session:=CopySession("settings")
		defer session.Close()
		err=mgoc.Update(bson.M{"sqlmap.autoRefresh":true},bson.M{"$set":bson.M{"sqlmap.autoRefresh":false}})
		if err!=nil{
			fmt.Println("mongodb update setting sqlmap auto Refresh")
		}
		return
	}

	if err!=nil{
		fmt.Println("mongodb update setting sqlmap auto Refresh")
	}
	if resp.StatusCode==http.StatusOK{
		var result AdminList
		mgoc,session:=CopySession("sqlmap")
		defer session.Close()
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {}
		_=resp.Body.Close()
		for key,value:=range result.Tasks{
			_,err:=mgoc.Upsert(bson.M{"taskId":key},bson.M{"$set":bson.M{"status":value}})
			if err!=nil{
				fmt.Println("Update data error")
			}
		}
	}
}

func Select_Sqlmap_Api(ctx context.Context,sqlmap_api string)  {
	for{
		select {
		case <-ctx.Done():
			fmt.Println("sqlmap auto refresh done")
			return
		default:
			Sqlmap_api_update(sqlmap_api)
		}
		fmt.Println("auto update sqlmap result ...")
		time.Sleep(time.Second*time.Duration(30))
	}
}

func exec_Shell_awit(location string,apiHost string) {
	host:=strings.Split(apiHost,":")
	cmd:=exec.Command("bash","-c",fmt.Sprintf("python sqlmapapi.py -s -H %s -p %s",host[0],host[1]))
	cmd.Dir=location
	cmd.Stdout=os.Stdout
	_=cmd.Start()
}

func Sqlmap_Start(location string,apiHost string) {
	exec_Shell_awit(location,apiHost)
}

func Select_Sqlmap_Start(ctx context.Context,location string,apiHost string) {
	fmt.Println("启动SQL注入组件")
	Sqlmap_Start(location,apiHost)
	for{
		select {
		case <-ctx.Done():
			fmt.Println("关闭SQL注入组件")
			return
		}
	}
}

func exec_Shell_Monitor_awit() {
	cmd:=exec.Command("python","./container/monitor.py")
	//cmd.Dir=location
	cmd.Stdout=os.Stdout
	cmd.Stderr = os.Stderr
	_=cmd.Start()
}

func Monitor_Start(){
	exec_Shell_Monitor_awit()
}

func Select_Monitor_Start(ctx context.Context) {
	fmt.Println("启动被动式监听组件")
	Monitor_Start()
	for{
		select {
		case <-ctx.Done():
			fmt.Println("关闭被动式监听组件")
			return
		}
	}
}