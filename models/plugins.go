package models

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
	"Kiddy/container/form"
	"time"
)

func GetAllPlugins(name string) ([]form.Plugins,error) {
	mgoc,session:=CopySession(name)
	defer session.Close()
	var plugins []form.Plugins
	err:=mgoc.Find(bson.M{}).All(&plugins)
	if err!=nil{
		return nil,errors.New("not found")
	}
	return plugins,nil
}

func GetUseAllPlugins(name string) []form.PluginsPoc {
	mgoc,session:=CopySession(name)
	defer session.Close()
	var plugins []form.PluginsPoc
	plugin,err:=Get_General_settings_struct("settings")
	if err!=nil{
		return nil
	}
	for _,i:=range plugin.General.Plugin{
		p:=form.PluginsPoc{}
		_=mgoc.Find(bson.M{"pinyin":i}).One(&p)
		plugins=append(plugins,p)
	}
	return plugins
}

func GetPluginsList(name string,pluginList []string) []string {
	mgoc,session:=CopySession(name)
	defer session.Close()
	var plugins []form.Plugins
	err:=mgoc.Find(bson.M{"pinyin":bson.M{"$in":pluginList}}).All(&plugins)
	if err!=nil{
		return nil
	}
	result:=make([]string,len(plugins))
	for i,value:=range plugins{
		result[i]=value.Pinyin
	}
	return result
}

func RemovePluginId(name string,idList []string) bool {
	mgoc,session:=CopySession(name)
	defer session.Close()
	plugins:=make([]form.Plugins,3)
	err:=mgoc.Find(bson.M{"pinyin":bson.M{"$in":[]string{"portscan","xssinject","sqlinject"}}}).All(&plugins)
	if err!=nil{
		return true
	}
	idList2:=[]bson.ObjectId{}
	status:=false
	for _,id:=range idList{
		t:=true
		for _,pid:=range plugins{
			if id==pid.Id.Hex(){
				t=false
			}
		}
		if t{
			idList2=append(idList2,bson.ObjectIdHex(id))
		}
	}
	_,err=mgoc.RemoveAll(bson.M{"_id":bson.M{"$in":idList2}})
	if err!=nil{
		status=true
	}
	return status
}

func InsertPlugin(name string,postForm *form.PostPlugin) bool {
	mgoc,session:=CopySession(name)
	defer session.Close()
	err:=mgoc.Insert(bson.M{"title":postForm.Title,"classification":postForm.Classification,"danger":postForm.Danger,"poc":postForm.Poc,"pinyin":postForm.Title})
	if err!=nil{
		return false
	}
	return true
}

func Insert_General_Out(name string,typeof string,host string,result string,status string) bool {
	mgoc,session:=CopySession(name)
	defer session.Close()
	_,err:=mgoc.Upsert(bson.M{"class":"general","typeof":typeof,"host":host},bson.M{"class":"general","typeof":typeof,"host":host,"message":result,"datetime":time.Now(),"status":status})
	if err!=nil{
		return false
	}
	return true
}