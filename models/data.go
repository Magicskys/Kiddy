package models

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"Kiddy/container/form"
	"Kiddy/setting"
)


var (
	MgoSession *mgo.Session
	SqlmapCancel context.CancelFunc
	SqlmapStartCancel context.CancelFunc
	MonitorCancel context.CancelFunc
	)

func InitData() {
	var err error
	MgoSession, err = mgo.Dial("mongodb://"+setting.DataBaseHost)
	if err != nil {
		fmt.Print("mongodb session Connection failed")
		panic(err)
	}
	//defer MgoSession.Close()
	if err:=MgoSession.Ping();err!=nil{
		fmt.Print("mongodb session Connection failed")
		panic(err)
	}
	// Optional. Switch the session to a monotonic behavior.
	MgoSession.SetMode(mgo.Monotonic, true)
	fmt.Println("mongodb session ok")
}

func init_Sql_Settings(){
	mgoc,session:=CopySession("settings")
	defer session.Close()
	query,err:=mgoc.Find(bson.M{"sqlmap":bson.M{"$type":3}}).Count()
	if err!=nil || query!=0{
		return
	}else{
		err:=mgoc.Insert(bson.M{"sqlmap":bson.M{"options":bson.M{"risk" : 1, "level" : 1, "delay" : 0, "randomAgent" : false, "tech" : "BEUSTQ"},"sqlmapApi":"127.0.0.1:8775","sqlmapLocation":"","autoRefresh" : true,"start":false}})
		if err!=nil{
			panic(errors.New("sqlmap settings insert into mongodb error"))
		}
	}
}

func init_General_Settings(){
	mgoc,session:=CopySession("settings")
	defer session.Close()
	query,err:=mgoc.Find(bson.M{"general":bson.M{"$type":3}}).Count()
	if err!=nil || query!=0{
		return
	}else{
		err:=mgoc.Insert(bson.M{"general":bson.M{"portScan":false,"portRange":[]int{1,65535},"plugin":[]string{},"portSchema":"sS"}})
		if err!=nil{
			panic(errors.New("general settings insert into mongodb fail"))
		}
	}
}

func init_Plugins(){
	mgoc,session:=CopySession("plugins")
	defer session.Close()
	query,err:=mgoc.Find(bson.M{"pinyin":bson.M{"$in":[]string{"portscan","xssinject","sqlinject"}}}).Count()
	if err!=nil || query==3{
		return
	}else{
		err:=mgoc.Insert(bson.M{"pinyin":"sqlinject","title":"SQL注入","danger":"","classification":"基线测试"},bson.M{"pinyin":"xssinject","title":"XSS注入","danger":"","classification":"基线测试"},bson.M{"pinyin":"portscan","title":"端口扫描","danger":"","classification":"基线测试"})
		if err!=nil{
			panic(errors.New("poc plugins insert into mongodb fail"))
		}
		fmt.Println("Initial plugins")
	}
}

func Init_Settings(){
	init_Sql_Settings()
	init_General_Settings()
	init_Plugins()
}

func CopySession(name string) (*mgo.Collection,*mgo.Session) {
	var mgoc *mgo.Collection
	session:=MgoSession.Copy()
	mgoc = session.DB(setting.TABLENAME).C(name)

	return mgoc,session
}

func GetAllInfo(name string) ([]form.Info,error) {
	mgoc,session:=CopySession(name)
	defer session.Close()
	var infos []form.Info
	err:=mgoc.Find(bson.M{"scheme":bson.M{"$type":2}}).All(&infos)
	if err!=nil{
		return nil,errors.New("not found")
	}
	return infos,nil
}

func GetIdInfo(name string,uid string) (form.Info,error) {
	mgoc,session:=CopySession(name)
	defer session.Close()
	var info form.Info
	err:=mgoc.Find(bson.M{"_id":bson.ObjectIdHex(uid)}).One(&info)
	if err!=nil{
		return form.Info{},errors.New("not found info id")
	}
	return info,nil
}


func GetAllResultSql(name string) ([]bson.M,error) {
	mgoc,session:=CopySession(name)
	defer session.Close()
	pipeline:=[]bson.M{
		bson.M{"$match": bson.M{"taskId": bson.M{"$type":2}}},
		//bson.M{"$match": bson.M{"class": bson.M{"$type":2}}},
		bson.M{"$lookup": bson.M{"from": "sqlmap", "localField": "taskId", "foreignField": "taskId", "as": "result"}},
		bson.M{"$lookup":bson.M{"from":"start","localField":"objectId","foreignField":"_id","as":"result2"}},
		bson.M{"$project":bson.M{"_id":0,"objectId":0,"typeof":0,"result":bson.M{"_id":0,"taskId":0},"result2":bson.M{"_id":0,"headers":0,"body":0,"path":0,"port":0,"host":0,"schema":0,"http_version":0}}},
	}
	pipe:=mgoc.Pipe(pipeline)
	resp := []bson.M{}
	err := pipe.All(&resp)
	if err!=nil{
		return nil,errors.New("not found")
	}
	return resp,nil
}

func GetAllResultGeneral(name string) ([]form.GeneralResult,error) {
	var general_result []form.GeneralResult
	mgoc,session:=CopySession(name)
	defer session.Close()
	err:=mgoc.Find(bson.M{"class": bson.M{"$type":2}}).All(&general_result)
	if err!=nil{
		return nil,errors.New("select result general fail")
	}
	return general_result,nil
}

func Get_Sqlmap_autRefres() (bool,string) {
	var sqlmap_settings form.SqlmapSettings
	mgoc,session:=CopySession("settings")
	defer session.Close()
	err:=mgoc.Find(bson.M{"sqlmap":bson.M{"$type":3}}).One(&sqlmap_settings)
	if err!=nil{
		return false,""
	}
	if sqlmap_settings.Sqlmap.AutoRefresh==true{
		return true,sqlmap_settings.Sqlmap.SqlmapApi
	}
	return false,""
}

func Get_Sqlmap_Start() (bool,string,string) {
	var sqlmap_settings form.SqlmapSettings
	mgoc,session:=CopySession("settings")
	defer session.Close()
	err:=mgoc.Find(bson.M{"sqlmap":bson.M{"$type":3}}).One(&sqlmap_settings)
	if err!=nil{
		return false,"",""
	}
	if sqlmap_settings.Sqlmap.Start==true{
		return true,sqlmap_settings.Sqlmap.SqlmapLocation,sqlmap_settings.Sqlmap.SqlmapApi
	}
	return false,"",""
}

func Get_Sqlmap_settings() ([]byte,string) {
	var sqlmap_settings form.SqlmapSettings
	mgoc,session:=CopySession("settings")
	defer session.Close()
	err:=mgoc.Find(bson.M{"sqlmap":bson.M{"$type":3}}).One(&sqlmap_settings)
	if err!=nil{
		return nil,""
	}
	response,_:=json.Marshal(sqlmap_settings.Sqlmap.Options)
	return response,sqlmap_settings.Sqlmap.SqlmapApi
}

func Get_Sqlmap_settings_struct(name string) (form.SqlmapSettings,error) {
	var sqlmap_settings form.SqlmapSettings
	mgoc,session:=CopySession(name)
	defer session.Close()
	err:=mgoc.Find(bson.M{"sqlmap":bson.M{"$type":3}}).One(&sqlmap_settings)
	if err!=nil{
		return form.SqlmapSettings{},errors.New("get sqlmap settings for mongo error")
	}
	return sqlmap_settings,nil
}

func Get_General_settings_struct(name string) (form.GeneralSettings,error) {
	var general_settings form.GeneralSettings
	mgoc,session:=CopySession(name)
	defer session.Close()
	err:=mgoc.Find(bson.M{"general":bson.M{"$type":3}}).One(&general_settings)
	if err!=nil{
		return form.GeneralSettings{},errors.New("get general settings for mongo error")
	}
	return general_settings,nil
}

func Update_Sqlmap_settings_Initial_Start(name string){
	mgoc,session:=CopySession(name)
	defer session.Close()
	err:=mgoc.Update(bson.M{"sqlmap":bson.M{"$type":3}},bson.M{"$set":bson.M{"sqlmap.start":false,"sqlmap.autoRefresh":false}})
	if err!=nil{
		fmt.Println("Initial update settings sqlmap start fail")
	}
}

func Update_Sqlmap_settings(name string,sqlmap *form.Post_Sqlmap_setting) bool {
	mgoc,session:=CopySession(name)
	defer session.Close()
	if SqlmapCancel==nil && sqlmap.Refresh {
		var ctx context.Context
		ctx, SqlmapCancel = context.WithCancel(context.Background())
		go Select_Sqlmap_Api(ctx,sqlmap.Sqlmap_api)
	}else if SqlmapStartCancel==nil && sqlmap.Start{
		var ctx context.Context
		ctx, SqlmapStartCancel = context.WithCancel(context.Background())
		go Select_Sqlmap_Start(ctx,sqlmap.Sqlmap_localhost,sqlmap.Sqlmap_api)
	}else if SqlmapCancel!=nil && !sqlmap.Refresh {
		SqlmapCancel()
		SqlmapCancel=nil
	}else if SqlmapStartCancel!=nil && !sqlmap.Start{
		SqlmapStartCancel()
		SqlmapStartCancel=nil
	}
	err:=mgoc.Update(bson.M{"sqlmap":bson.M{"$type":3}},bson.M{"sqlmap":bson.M{"sqlmapApi":sqlmap.Sqlmap_api,"sqlmapLocation":sqlmap.Sqlmap_localhost,"autoRefresh":sqlmap.Refresh,"start":sqlmap.Start,"options":bson.M{"risk":sqlmap.Risk,"level":sqlmap.Level,"delay":sqlmap.Region,"randomAgent":sqlmap.User_agent,"tech":sqlmap.Tech}}})
	if err!=nil{
		return false
	}
	return true
}

func Update_General_settings(name string,general *form.Post_General_setting) bool {
	plugins:=GetPluginsList("plugins",general.Plugin)
	mgoc,session:=CopySession(name)
	defer session.Close()
	err:=mgoc.Update(bson.M{"general":bson.M{"$type":3}},bson.M{"general":bson.M{"portScan":general.Porttype,"portRange":general.Portrange,"plugin":plugins,"portSchema":general.PortSchema}})
	if err!=nil{
		return false
	}
	return true
}