from ansible import errors, runner
import json

def to_servers_with_port(host_vars, groups, target, attribute, port):
    if type(host_vars) != runner.HostVars:
        raise errors.AnsibleFilterError("|failed expects a HostVars")

    if type(groups) != dict:
        raise errors.AnsibleFilterError("|failed expects a Dictionary")

    hosts_in_group = []
    for host in groups[target]:
        hosts_in_group.append(host_vars[host])

    hosts_with_port = map((lambda host: host[attribute] + ':' + port), hosts_in_group)
    return hosts_with_port

class FilterModule (object):
    def filters(self):
        return {"to_servers_with_port": to_servers_with_port}
