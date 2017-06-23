#coding:utf-8
import sqlite3
from lib import logger
import requests
import json
import re
class Core(object):
    def __init__(self):
        self.conn=sqlite3.connect("database/core.db")
        self.conn.text_factory=str
        self.cu=self.conn.cursor()

    def exit(self):
        self.conn.close()

    def init(self):
        try:
            self.cu.execute("DROP TABLE info;")
            self.cu.execute("DROP TABLE sniff")
            self.cu.execute("CREATE TABLE info (SESSION STRING,STATUS STRING,URL STRING,DATA TEXT);")
            self.cu.execute("CREATE TABLE sniff (SESSION STRING,METHOD STRING,URL STRING,COOKIE STRING,DATA TEXT);")
            logger.success("初始化成功")
        except:
            logger.error("初始化错误")
            logger.info("从新建立数据库")
            self.cu.execute("CREATE TABLE info (SESSION STRING,STATUS STRING,URL STRING,DATA TEXT);")
            self.cu.execute("CREATE TABLE sniff (SESSION STRING,METHOD STRING,URL STRING,COOKIE STRING,DATA TEXT);")
            logger.success("建立成功")

    def search_task(self,status,session):
        if status['status'] != u"running":
            self.cu.execute("UPDATE info SET STATUS='" + status['status'] + "' WHERE SESSION='" + session[0] + "';")
            self.conn.commit()

    def status(self,arg):
        if arg=="":
            sessions=self.cu.execute("SELECT SESSION FROM info;")
            for session in self.cu.fetchall():
                try:
                    status=requests.get("http://127.0.0.1:8775/scan/" + session[0] + "/status").json()
                    self.search_task(status,session)
                except:
                    logger.error("查询错误")
                    logger.error("Sqlmap中没有这个%s任务"%session[0])
            return self.cu.execute("SELECT SESSION,STATUS,URL FROM info;")

        else:
            try:
                status = requests.get("http://127.0.0.1:8775/scan/" + arg + "/status").json()
                self.search_task(status,arg)
                return self.cu.execute("SELECT SESSION,URL,STATUS FROM info WHERE SESSION='"+arg+"';")
            except:
                logger.error("Sqlmap中没有这个%s任务"%arg)
                return self.cu.execute("SELECT SESSION,URL,STATUS FROM info WHERE SESSION='" + arg + "';")

    def info(self,arg):
        try:
            data=re.sub('[\'\"]',' ',str(requests.get("http://127.0.0.1:8775/scan/"+arg+"/data").json()['data']))
            if data!="":
                self.cu.execute("UPDATE info SET DATA='" + str(data) + "' WHERE SESSION='" + str(arg) + "';")
                self.conn.commit()
                return data
            else:
                return "没有数据"
        except:
            logger.error("查询错误")

    def save(self,arg):
        pass

    def count(self):
        self.cu.execute("SELECT COUNT(*) FROM info;")
        return self.cu.fetchall()

    def ok_count(self):
        self.cu.execute("select count(*) from info where status='terminated'")
        return self.cu.fetchall()

    @classmethod
    def create(cls,args):
        self=Core()
        print self.cu.execute("SELECT * FROM info WHERE URL='%s'"%(args['URL'])).fetchone()
        if self.cu.execute("SELECT URL FROM info WHERE URL='%s'"%(args['URL'])).fetchone():return
        sess=requests.get("http://127.0.0.1:8775/task/new").json()
        cr_task=requests.post("http://127.0.0.1:8775/scan/"+sess['taskid']+"/start",data=json.dumps({'url':args['URL']}),headers={'Content-type':'application/json'})
        self.cu.execute("INSERT INTO sniff values ('%s','%s','%s','%s','%s');" % (sess['taskid'], args['METHOD'], args['URL'], args['COOKIE'], str(args['DATA'].__dict__.items())[0].strip()))
        self.cu.execute("INSERT INTO info values ('%s','%s','%s','%s');" % (sess['taskid'], sess[u'success'], args['URL'],str(args['DATA'].__dict__.items())[0].strip()))
        self.conn.commit()
        logger.info("\n扫描"+args['URL'])

