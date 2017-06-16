#!/usr/bin/env python
#coding:utf-8
from termcolor import colored,cprint


def error(string):
    cprint("[-]" + string, "red", attrs=['reverse', 'blink'])
def success(string):
    print colored("[+]"+string, "red","on_cyan")

def info(string):
    print colored("[*]"+string, "blue")