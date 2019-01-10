package utils

//import (
//	"context"
//	"encoding/json"
//	"fmt"
//	"gopkg.in/mgo.v2/bson"
//	"Kiddy/models"
//	"net/http"
//	"time"
//)
//
//
//
//type AdminList struct {
//	Tasks map[string]string	`json:"tasks"`
//	Tasks_num int64	`json:"tasks_num"`
//	Success	bool `json:"success"`
//
//}
//
//func Init_SqlmapApi_Result()  {
//	if models.Get_Sqlmap_autRefres(){
//		var ctx context.Context
//		ctx, models.SqlmapCancel = context.WithCancel(context.Background())
//		go Select_Sqlmap_Api(ctx)
//	}
//}
//
//func Sqlmap_api_update(){
//	client:=&http.Client{
//		Timeout:3*time.Second,
//	}
//	resp,err:=client.Get("http://127.0.0.1:8775/admin/list")
//	if err!=nil{
//		mgoc,session:=models.CopySession("settings")
//		defer session.Close()
//		err=mgoc.Update(bson.M{"sqlmap.autoRefresh":true},bson.M{"$set":bson.M{"sqlmap.autoRefresh":false}})
//		if err!=nil{
//			fmt.Println("mongodb update setting sqlmap auto Refresh")
//		}
//		return
//	}
//
//	if err!=nil{
//		fmt.Println("mongodb update setting sqlmap auto Refresh")
//	}
//	if resp.StatusCode==http.StatusOK{
//		var result AdminList
//		mgoc,session:=models.CopySession("sqlmap")
//		defer session.Close()
//		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {}
//		resp.Body.Close()
//		for key,value:=range result.Tasks{
//			_,err:=mgoc.Upsert(bson.M{"taskId":key},bson.M{"$set":bson.M{"status":value}})
//			if err!=nil{
//				fmt.Println("Update data error")
//			}
//		}
//	}
//}
//
//
//func Select_Sqlmap_Api(ctx context.Context)  {
//	for{
//		select {
//		case <-ctx.Done():
//			fmt.Println("sqlmap auto refresh done")
//			return
//		default:
//			Sqlmap_api_update()
//		}
//		fmt.Println("auto update sqlmap result ...")
//		time.Sleep(time.Second*time.Duration(30))
//	}
//}