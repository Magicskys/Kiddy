#coding:utf-8
from mitmproxy import proxy, options
from mitmproxy.tools.dump import DumpMaster
import pymongo

from pymongo import MongoClient
conn = MongoClient('0.0.0.0', 27017)
db = conn.scan
my_set = db.start

class AddHeader:
    def request(self, flow):
        # flow.request.data.headers
        # flow.request.path
        # flow.request.url
        # flow.request.method
        # flow.request.cookies
        # flow.request.http_version
        # flow.request.port
        # flow.request.host
        # flow.request.scheme
        if 'google.com' in flow.request.host:
            return

        if flow.request.path!="" or flow.request.method=='POST':
            if my_set.find_one({"method":flow.request.method,"url":flow.request.url}) is None:
                headers=';'.join(['%s=%s' % (i, j) for i, j in flow.request.data.headers.items()])
                my_set.insert_one({"scheme":flow.request.scheme,"host":flow.request.host,"port":flow.request.port,"http_version":flow.request.http_version,"method":flow.request.method,"url":flow.request.url,
                               "path":flow.request.path,"headers":headers,"body":flow.request.data.content.decode("utf-8")})

opts = options.Options(listen_host='0.0.0.0',listen_port=8080)
opts.add_option("body_size_limit", int, 0, "")
opts.add_option("keep_host_header", bool, True, "")
pconf = proxy.config.ProxyConfig(opts)

m = DumpMaster(None)
m.server = proxy.server.ProxyServer(pconf)
m.addons.add(AddHeader())

try:
    m.run()
except KeyboardInterrupt:
    m.shutdown()