class VagrantPlugins::ProviderVirtualBox::Action::Network
  def dhcp_server_matches_config?(dhcp_server, config)
    true
  end
end

Vagrant.configure("2") do |config|
  (1..3).each do |i|
    config.vm.define "node#{i}" do |node|
      node.vm.box = "ubuntu/bionic64"
      node.vm.network "private_network", ip: "192.168.56.#{10 + i}"
      node.vm.provision "shell", inline: <<-SHELL
        apt-get update
        apt-get install -y software-properties-common
        add-apt-repository -y ppa:projectatomic/ppa
        apt-get update
        apt-get install -y podman
        apt-get install -y golang
        mkdir -p /home/vagrant/kmeans-backend
        chown vagrant:vagrant /home/vagrant/kmeans-backend
      SHELL
      node.vm.synced_folder "./kmeans-backend", "/home/vagrant/kmeans-backend"
      node.vm.network "forwarded_port", guest: 8080, host: 8080 + i
    end
  end
end

  