#!/usr/bin/env python
#coding:utf-8
import cmd
import os
from lib import logger
from lib.core import Core

class Interface(cmd.Cmd,Core):
    def __init__(self):
        cmd.Cmd.__init__(self)
        Core.__init__(self)
        self.prompt = "[*]探测装置->"


    def do_help(self, arg):
        commands={
            'help':"帮助菜单",
            'start':'启动探测',
            'inti':'初始化数据库',
            'status':'查看任务状态(可加session号查询)',
            'info':'查看成功的任务(可加session号查询)',
            'exit':'退出'
        }

        print "\nCore Commands\n=============\n"
        print "%-30s%s" % ("命令", "描述")
        print "%-30s%s" % ("-------", "-----------")
        for command in commands:
            print "%-30s%s" % (command, commands[command])
        print

    def do_exit(self,arg):
        self.exit()
        exit()

    def do_start(self,arg):
        logger.info("OK")

    def do_status(self,arg):
        print "%-20s%-15s%s" % ("ID", "状态", "地址")
        print "%-20s%-13s%s" % ("-------", "------", "---------")
        if not arg:
            for session, url, status in self.status(arg):
                print "%-20s%-13s%s" % (session, url, status)
            print
            print "共有%d个任务,%d个任务成功" %(self.count()[0][0],self.ok_count()[0][0])
        else:
            self.status(arg)

    def do_info(self,arg):
        print "%-20s%-15s%-15s%s" % ("ID", "状态", "地址","数据")
        print "%-20s%-13s%-15s%s" % ("-------", "------","-----","------")
        if not arg:
            for session, url, status in self.info(arg):
                print "%-20s%-13s%s" % (session, url, status)
            print
            print "共有%d个任务成功"%self.ok_count()[0][0]
        else:
            for session, url, status,data in self.info(arg):
                print "%-20s%-13s%-15s%s" % (session, url, status,data)
            print

    def do_save(self,arg):
        pass

    def do_init(self, arg):
        logger.info("初始化数据库")
        self.init()

