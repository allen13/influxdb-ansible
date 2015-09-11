require 'rbconfig'

def create_vm(config, options = {})
  dirname = File.dirname(__FILE__)
  config.vm.synced_folder "#{dirname}/..", '/vagrant', disabled: false
  config.ssh.password = 'vagrant'

  name = options.fetch(:name, "node")
  id = options.fetch(:id, 1)
  vm_name = "%s-%02d" % [name, id]

  memory = options.fetch(:memory, 1024)
  cpus = options.fetch(:cpus, 1)

  config.vm.define vm_name do |config|
    config.vm.box = "bento/centos-7.1"
    config.vm.hostname = vm_name

    public_ip = "10.0.168.10#{id}"
    config.vm.network :private_network, ip: public_ip, netmask: "255.255.255.0"

    private_ip = "10.0.169.10#{id}"
    config.vm.network :private_network, ip: private_ip, netmask: "255.255.255.0"

    config.vm.provider :virtualbox do |vb|
      vb.memory = memory
      vb.cpus = cpus
    end
  end
end
