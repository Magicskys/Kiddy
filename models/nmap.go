package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

func Exists_Nmap_Task(name string,host string) bool {
	mgoc,session:=CopySession(name)
	defer session.Close()
	query,err:=mgoc.Find(bson.M{"class":"general","typeof":"nmap","host":host}).Count()
	if err!=nil || query!=0{
		return false
	}
	return true
}

func Insert_Nmap_Out(name string,host string,result string,status string) bool {
	mgoc,session:=CopySession(name)
	defer session.Close()
	_,err:=mgoc.Upsert(bson.M{"class":"general","typeof":"nmap","host":host},bson.M{"class":"general","typeof":"nmap","host":host,"message":result,"datetime":time.Now(),"status":status})
	if err!=nil{
		return false
	}
	return true
}

func Remmove_nmap_taskId(name string,taskId string) bool {
	mgoc,session:=CopySession(name)
	defer session.Close()
	err:=mgoc.RemoveId(bson.ObjectIdHex(taskId))
	if err!=nil{
		return false
	}
	return true
}