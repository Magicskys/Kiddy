#!/usr/bin/env python
#coding:utf-8
from lib.console import Interface

def main():
    log="""
    -----------------------
    < 小小怪被动式扫描装置启动 >
    -----------------------
    
        ..||||||||∞ |||||
        ╭||||━━　　━━ ||||╮
        ╰|||　　 ~　　　|||╯
        　||╰╭--╮ˋ╭--╮╯||
        　||　╰/ /　 || ОО 
        
        节操粉碎中 请稍后
         ━━━━━━━━━━━
         ▉▉▉▉▉▉▉▉ 99.9%
         ━━━━━━━━━━━
         
    """

    interface=Interface()
    print log
    while True:
        try:
            interface.cmdloop()
        except KeyboardInterrupt:
            print "输入exit退出"

if __name__ == '__main__':
    main()
