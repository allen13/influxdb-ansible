ansible influxdb cluster playbook
=================================

####configuration overview

The goal of this playbook is to make use of InfluxDB as a system metric store.
Features:

* Can automatically cluster n-number of influxdb nodes
* Uses telegraf to gather metrics using UDP
* Sets up a grafana dashboard
* sets up downsampling for all metrics

####notes

* This playbook is specific to Centos7 and there is no intent to generalize it to
other distros. I am willing to accept pull requests that make this work with
apt and yum provided it's done cleanly.

####requirements

* [Virtualbox](https://www.virtualbox.org/wiki/Downloads)
* [Vagrant](http://www.vagrantup.com/downloads) on your local machine
* Ansible via pip install ansible

####running in vagrant

    vagrant up
    bin/provision_development

####destroying

    bin/destroy

####Influxdb Admin Site

    http://http://10.0.168.101:8083

####Grafana Dashboard

    http://http://10.0.168.101:3006

####custom ansible modules

* **influxdb_user** - create users for Influxdb
* **influxdb_retention_policy** - create retention policies for Influxdb
* **influxdb_downsample** - given measurements of [system_load1, mem_available]
   and retention policy 30s:30d,1m:52w continuous queries will be generated
   that downsample the metrics accordingly. The naming scheme
   will generate continuous queries system_load1_30s with assigned retention policy rp_30d and
   system_load1_1m with assigned retention policy rp_52w.

####custom filter plugins

* **to_url_with_port** - assist in building the list of influxdb servers for telegraf
  to_url_with_port(groups,'influxdb','udp://','private_ip','8089')|join(', ')
