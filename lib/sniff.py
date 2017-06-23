from mitmproxy import flow, controller, options
from mitmproxy.proxy import ProxyServer, ProxyConfig
from core import Core
import re

def proxy_address(flow):
    return ("remote_ip", 18002)

class Sniff(flow.FlowMaster):
    def run(self):
        try:
            flow.FlowMaster.run(self)
        except KeyboardInterrupt:
            self.shutdown()

    def sniff_stop(self):
        self.shutdown()

    @controller.handler
    def request(self, f):
        address = proxy_address(flow)
        # if f.live:
            # print("inside")
            # f.live.change_upstream_proxy_server(address)
        url=f.request.url
        if re.search(r'\?\S+=',url):
            data={'METHOD':f.request.method,'URL':url,'COOKIE':f.request.cookies.to_dict(),'DATA':f.request.data}
            Core().create(data)

    # @controller.handler
    # def response(self, f):
    #     print("response", f)

    # @controller.handler
    # def error(self, f):
    #     print("error", f)

    # @controller.handler
    # def log(self, l):
    #     print("log", l.msg)

def sniff_main():
    opts = options.Options(
        upstream_server="http://localhost:8080", cadir="~/.mitmproxy/")
    config = ProxyConfig(opts)
    state = flow.State()
    server = ProxyServer(config)
    m = Sniff(opts, server, state)
    m.run()