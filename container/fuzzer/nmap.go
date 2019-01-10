package fuzzer

import (
	"bytes"
	"fmt"
	"Kiddy/models"
	"os/exec"
)

func exec_Shell(s string) string {
	cmd:=exec.Command("/bin/bash","-c",s)
	var out bytes.Buffer
	cmd.Stdout=&out
	err:=cmd.Run()
	if err!=nil{
		return err.Error()
	}
	return out.String()
}

func Join_Nmap_Scan(username string,targetUrls []string)  {
	general_settings,err:=models.Get_General_settings_struct("settings")
	if err!=nil{
		return
	}
	for _,targetUrl:=range targetUrls{
		target,err:=models.GetIdInfo(username,targetUrl)
		if err!=nil{
			continue
		}
		if !models.Exists_Nmap_Task("start",target.Host,){
			fmt.Println("任务存在")
			continue
		}
		if !models.Insert_Nmap_Out("start",target.Host,"","start"){
			fmt.Println("初始化任务失败")
			continue
		}
		var cmdString string
		if general_settings.General.PortScan{
			cmdString=fmt.Sprintf("nmap -F -Pn %s -oX -",target.Host)
		}else{
			cmdString=fmt.Sprintf("nmap -p %d-%d -Pn %s -oX -",general_settings.General.PortRange[0],general_settings.General.PortRange[1],target.Host)
		}
		cmdOut:=exec_Shell(cmdString)
		if !models.Insert_Nmap_Out("start",target.Host,cmdOut,"end"){
			fmt.Println("插入失败")
		}
		fmt.Println("NMAP扫描成功")
	}
}

func Kill_Nmap_taskId(targetUrls []string) {
	for _,targetUrl:=range targetUrls{
		if !models.Remmove_nmap_taskId("start",targetUrl){
			fmt.Println("remove mongodb taskid general fail ",targetUrl)
		}
	}
}