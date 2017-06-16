from mitmproxy.script import concurrent
import mitmproxy

@concurrent
def request(flow):
    mitmproxy.ctx.log("handle request: %s%s" % (flow.request.host, flow.request.path))
