
def to_url_with_port(host_vars, groups, target, endpoint_type ,attribute, port):
    hosts_in_group = []
    for host in groups[target]:
        hosts_in_group.append(host_vars[host])

    hosts_with_port = map((lambda host: '"' + endpoint_type + host[attribute] + ':' + port + '"'), hosts_in_group)
    return hosts_with_port

class FilterModule (object):
    def filters(self):
        return {"to_url_with_port": to_url_with_port}
