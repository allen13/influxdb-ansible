ansible influxdb cluster playbook
=================================

####configuration overview

The goal of this playbook is to make use of [InfluxDB](https://influxdb.com) as a system metric store.
Features:

* Can automatically cluster n-number of influxdb nodes
* Uses [Telegraf](https://github.com/influxdb/telegraf) to gather metrics using UDP
* Sets up [Kapacitor](https://github.com/influxdb/kapacitor) for processing, monitoring, and alerting
* Sets up a grafana dashboard
* Sets up downsampling for all metrics

####notes

* All of the packages in this repo are set to state=latest so that I don't have
to continuously maintain the version pin.
If you intend to use this in production you should pin all of the versions and
do testing between upgrades.
Additionally the influxdb modules make a lot of assumptions about the way influx
does things and are vulnerable to being broken between versions.

* This playbook is specific to Centos7 and there is no intent to generalize it to
other distros, but it should provide a good enough example to get going in any environment.

####requirements

* [Virtualbox](https://www.virtualbox.org/wiki/Downloads)
* [Vagrant](http://www.vagrantup.com/downloads)
* [Ansible](http://www.ansible.com) via pip install ansible

####running in vagrant

    vagrant up
    bin/provision_development

####destroying

    bin/destroy

####influxdb admin site

    http://10.0.168.101:8083

####grafana dashboard

    http://10.0.168.101:3006

####custom ansible modules

* **influxdb_user** - create users for Influxdb
* **influxdb_retention_policy** - create retention policies for Influxdb
* **influxdb_downsample** - given measurements [{'measurement':'system', 'fields':['load1','load5','load15']]
   and retention policy 30s:30d,1m:52w continuous queries will be generated
   that downsample the metrics accordingly. The naming scheme
   will generate the continuous query system_30s with assigned retention policy rp_30d.

####custom filter plugins

* **to_url_with_port** - assist in building the list of influxdb servers for telegraf.
	* Example: hostvars|to_url_with_port(groups,'influxdb','udp://','private_ip','8089')|join(', ')
