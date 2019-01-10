package form

type SqlMapOptions struct {
	Level int64 `json:"level"`
	Delay int64 `json:"delay"`
	RandomAgent bool `json:"randomagent" bson:"randomAgent"`
	Tech string `json:"tech"`
	Risk int64 `json:"risk"`
}

type Sqlmap struct {
	Options *SqlMapOptions
	SqlmapApi string `json:"sqlmap_api" bson:"sqlmapApi"`
	SqlmapLocation string `json:"sqlmap_location" bson:"sqlmapLocation"`
	AutoRefresh bool `json:"auto_refresh" bson:"autoRefresh"`
	Start bool `json:"start" bson:"start"`
}

type SqlmapSettings struct {
	Sqlmap *Sqlmap `json:"sqlmap"`
}

type Post_Sqlmap_setting struct {
	Sqlmap_localhost string `form:"sqlmap_localhost" binding:"required"`
	Sqlmap_api string `form:"sqlmap_api" binding:"required"`
	Region uint `form:"region" binding:"required"`
	Level uint `form:"level" binding:"required"`
	Risk uint `form:"risk" binding:"required"`
	Refresh bool `form:"refresh" binding:"exists"`
	Start bool `form:"start" binding:"exists"`
	Tech string `form:"tech[]" binding:"required"`
	User_agent bool `form:"user_agent" binding:"exists"`
}

type General struct {
	Plugin []string `json:"plugin" bson:"plugin"`
	PortScan bool `json:"port_scan" bson:"portScan"`
	PortRange []int `json:"port_range" bson:"portRange"`
	PortSchema string `json:"portschema" bson:"portSchema"`
}

type GeneralSettings struct {
	General *General `json:"general"`
}

type Post_General_setting struct {
	Portrange []int `form:"portrange[]" binding:"exists"`
	Plugin []string `form:"plugin[]" binding:"exists"`
	Porttype bool `form:"porttype" binding:"required"`
	PortSchema string `form:"portschema" binding:"required"`
}