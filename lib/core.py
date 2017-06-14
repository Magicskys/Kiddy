#coding:utf-8
import sqlite3
from lib import logger
import requests

class Core():
    def __init__(self):
        self.conn=sqlite3.connect("database/core.db")
        self.conn.text_factory=str
        self.cu=self.conn.cursor()

    def exit(self):
        self.conn.close()

    def init(self):
        try:
            self.cu.execute("DROP TABLE info;")
            self.cu.execute("CREATE TABLE info (SESSION STRING,URL STRING,STATUS STRING,DATA STRING);")
            logger.success("初始化成功")
        except:
            logger.error("初始化错误")
            logger.info("从新建立数据库")
            self.cu.execute("CREATE TABLE info (SESSION STRING,URL STRING,STATUS STRING,DATA STRING);")
            logger.success("建立成功")

    def status(self,arg):
        if arg=="":
            sessions=self.cu.execute("SELECT SESSION FROM info;")
            for session in sessions:
                try:
                    status=requests.get("127.0.0.1:8775/scan/" + session + "/status").json()['status']
                    if status!="running":
                        self.cu.execute("update info set status="+status['status']+"data="+status['data']+" where session='"+session+"';")
                    return self.cu.execute("SELECT SESSION,URL,STATUS FROM info;")
                except:
                    logger.error("查询错误")
                    return self.cu.execute("SELECT SESSION,STATUS,URL FROM info;")
                    # logger.error("Sqlmap中没有这个%s任务"%session)
        else:
            try:
                status = requests.get("127.0.0.1:8775/scan/" + arg + "/status").json()['status']
                if status != "running":
                    self.cu.execute("UPDATE info SET status=" + status['status'] + "data=" + status['data'] + " where session='" + arg + "';")
                return self.cu.execute("SELECT SESSION,URL,STATUS FROM info WHERE SESSION='"+arg+"';")
            except:
                logger.error("Sqlmap中没有这个%s任务"%arg)
                self.cu.execute("SELECT SESSION,URL,STATUS FROM info WHERE SESSION='" + arg + "';")
                return self.cu.fetchall()


    def info(self,arg):
        if arg=="":
            return self.cu.execute("SELECT SESSION,URL,STATUS FROM info where status='terminal';")
        else:
            return self.cu.execute("SELECT * FROM info WHERE SESSION='" + arg + "';")

    def save(self):
        pass

    def count(self):
        self.cu.execute("SELECT COUNT(*) FROM info;")
        return self.cu.fetchall()

    def ok_count(self):
        self.cu.execute("select count(*) from info where status='terminal'")
        return self.cu.fetchall()

    def start(self):
        