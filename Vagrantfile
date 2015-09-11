# -*- mode: ruby -*-
# vi: set ft=ruby :

require 'fileutils'
require_relative './vagrant/shared.rb'

Vagrant.configure("2") do |config|
  create_vm(config, name: "influxdb", id: 1)
  create_vm(config, name: "influxdb", id: 2)
  create_vm(config, name: "influxdb", id: 3)
end
