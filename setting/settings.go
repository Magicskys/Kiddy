package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg  *ini.File
	RunMode string
	HTTPPort       int
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	LocateLocation string

	PageSize  int
	JwtSecret string

	DataBaseHost string
	TABLENAME string
	)


func init() {
	var err error
	Cfg,err=ini.Load("conf/app.ini")
	if err!=nil{
		log.Fatalln("读取配置文件出错 conf/app.ini")
	}
	LoadBase()
	LoadServer()
	LoadDataBase()
	LoadApp()
}

func LoadBase(){
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("release")
}

func LoadServer(){
	sec,err:=Cfg.GetSection("server")
	if err!=nil{
		log.Fatalln("读取配置项server出错",err.Error())
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
	LocateLocation = sec.Key("LOCATELOCATION").MustString("Asia/Shanghai")
}

func LoadApp(){
	sec,err:=Cfg.GetSection("app")
	if err!=nil{
		log.Fatalln("读取配置项app出错",err.Error())
	}
	JwtSecret=sec.Key("JWT_SECRET").MustString("!@$^1246%*18745")
	PageSize=sec.Key("PAGE_SIZE").MustInt(10)
}

func LoadDataBase()  {
	sec,err:=Cfg.GetSection("database")
	if err!=nil{
		log.Fatalln("读取配置项database出错",err.Error())
	}
	DataBaseHost=sec.Key("HOST").MustString("127.0.0.1:8775")
	TABLENAME=sec.Key("TABLENAME").MustString("scan")
}