#!/usr/bin/env python
#coding:utf-8
import cmd
import os
from lib import logger
from lib.core import Core
from lib.sniff import sniff_main,Sniff
from mitmproxy.proxy import ProxyServer, ProxyConfig
from mitmproxy import flow, controller, options
import subprocess
import threading

class Interface(cmd.Cmd,Core):
    def __init__(self):
        cmd.Cmd.__init__(self)
        Core.__init__(self)
        self.prompt = "[*]探测装置->"
        self.name="sniff"
        self.start_status=False


    def do_help(self, arg):
        commands={
            'help':"帮助菜单",
            'start':'启动探测',
            'stop':'停止探测',
            'inti':'初始化数据库  ',
            'status':'查看任务状态(可加session号查询)',
            'info':'查看任务的扫描结果',
            'exit':'退出'
        }

        print "\nCore Commands\n=============\n"
        print "%-30s%s" % ("命令", "描述")
        print "%-30s%s" % ("-------", "-----------")
        for command in commands:
            print "%-30s%s" % (command, commands[command])
        print

    def do_exit(self,arg):
        import sys
        self.exit()
        try:
            if self.th_sniff.isAlive():
                sys.exit()
        except:
            pass
        exit()

    def do_start(self,arg):
        if not self.start_status:
            logger.info("OK")
            self.prompt="\033[1;35m [*]探测装置-> \033[0m"
            self.start_status=True
            opts = options.Options(
                upstream_server="http://localhost:8080", cadir="~/.mitmproxy/")
            config = ProxyConfig(opts)
            state = flow.State()
            server = ProxyServer(config)
            self.m = Sniff(opts, server, state)
            self.th_sniff=threading.Thread(target=self.m.run,name='sniff')
            self.th_sniff.setDaemon(True)
            self.th_sniff.start()
        else:
            logger.error("Sniff已开启")

    def do_status(self,arg):
        print "%-20s%-15s%s" % ("ID", "状态", "地址")
        print "%-20s%-13s%s" % ("-------", "------", "---------")
        if not arg:
            if self.status(arg):
                for session, status, url in self.status(arg):
                    print "%-20s%-13s%s" % (session, status, url)
                print
                print "共有%d个任务,%d个任务扫描完成" %(self.count()[0][0],self.ok_count()[0][0]  )
            else:
                logger.info("没有任务")
        else:
            for session, url, status in self.status(arg):
                print "%-20s%-13s%s" % (session, status, url)

    def do_info(self,arg):
        if not arg:
            logger.error("请跟上参数SESSION")
        else:
            logger.info("%s的数据为"%arg)
            print self.info(arg)

    def do_save(self,arg):
        self.save(arg)

    def do_init(self, arg):
        logger.info("初始化数据库")
        self.init()

    def do_thread(self,arg):
        print "线程名：",self.th_sniff.getName()
        print "线程活动状态：",self.th_sniff.isAlive()

    def do_stop(self,arg):
        try:
            if self.start_status:
                Sniff.sniff_stop(self.m)
                self.prompt = "[*]探测装置->"
                logger.info("暂停")
                self.start_status=False
            else:
                logger.error('Sniff未开启')
        except Exception,e:
            print e
            logger.error("Sniff暂停失败")
