#!/usr/bin/env python
#coding:utf-8
import cmd
import os

class Interface(cmd.Cmd):
    def __init__(self):
        cmd.Cmd.__init__(self)
        self.prompt = "[*]探测装置->"


    def do_help(self, arg):
        commands={
            'help':"帮助菜单",
            'start':'启动探测',
            'exit':'退出'
        }

        print "\nCore Commands\n=============\n"
        print "%-30s%s" % ("命令", "描述")
        print "%-30s%s" % ("-------", "-----------")
        for command in commands:
            print "%-30s%s" % (command, commands[command])
        print

    def do_exit(self,arg):
        exit()


    def do_start(self,arg):
        print "OK"


    def do_status(self,arg):
        pass

    def do_info(self,arg):
        pass

    def do_save(self,arg):
        pass

    def do_init(self,arg):
        pass