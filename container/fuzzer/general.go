package fuzzer

import (
	"bytes"
	"fmt"
	"Kiddy/container/form"
	"Kiddy/models"
	"os/exec"
	"strings"
)

func exec_General_Shell(s string,scheme string,url string) string {
	e := fmt.Sprintln("exec('''" + strings.Replace(s,"{#url#}",scheme+"://"+url,-1) + "''')")
	e = strings.TrimSpace(e)
	s = fmt.Sprintf("/Users/galan/anaconda2/bin/python -c %#v", e)
	cmd := exec.Command("/bin/bash", "-c", s)
	var out bytes.Buffer
	cmd.Stdout=&out
	//cmd.Stderr = os.Stderr
	err:=cmd.Run()
	if err!=nil{
		return ""
	}
	return out.String()
}

func General_Scan(username string,uid []string,plugins []form.PluginsPoc){
	for _,targetUrl:=range uid{
		for _,value:=range plugins{
			switch value.Pinyin {
			case "portscan":
				Join_Nmap_Scan("start",uid)
				continue
			case "sqlinject":
				if models.SqlmapStartCancel!=nil{
					Join_Sqlmap_Scan("start",uid)
					continue
				}
			case "xssinject":
				continue
			default:
				target,err:=models.GetIdInfo(username,targetUrl)
				if err!=nil{
					continue
				}
				cmdOut:=exec_General_Shell(value.Poc,target.Scheme,target.Host)
				if cmdOut!="" {
					if !models.Insert_General_Out("start",value.Classification, target.URL, cmdOut, "end") {
						fmt.Println("插入失败")
					}
					fmt.Println(value.Title,"扫描成功")
				}
			}
		}
	}
}