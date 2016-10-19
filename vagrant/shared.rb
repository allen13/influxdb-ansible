def create_vm(config, options = {})
  dirname = File.dirname(__FILE__)

  name = options.fetch(:name, "node")
  id = options.fetch(:id, 1)
  vm_name = "%s-%02d" % [name, id]

  memory = options.fetch(:memory, 1024)
  cpus = options.fetch(:cpus, 1)

  config.vm.define vm_name do |config|
    config.vm.box = "centos/7"
    config.vm.hostname = vm_name

    private_ip = "192.0.2.11#{id}"
    config.vm.network :private_network, ip: private_ip, netmask: "255.255.255.128"

    public_ip = "192.0.2.21#{id}"
    config.vm.network :private_network, ip: public_ip, netmask: "255.255.255.128"

    config.vm.provider :virtualbox do |vb|
      vb.memory = memory
      vb.cpus = cpus
    end
  end
end
