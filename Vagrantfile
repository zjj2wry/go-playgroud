# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  config.vm.network :forwarded_port, guest: 22, host: 2323
  config.vm.network "private_network", ip: "192.168.50.4"
  config.vm.box = "bento/ubuntu-16.04"
  config.vm.hostname = 'develop'
  config.vm.provision "shell", privileged: false, inline: <<-SHELL
    set -e -x -u
    sudo apt-get update && sudo apt-get install -y bridge-utils
    # Install Golang
    wget --quiet https://storage.googleapis.com/golang/go1.9.1.linux-amd64.tar.gz
    sudo tar -zxf go1.9.1.linux-amd64.tar.gz -C /usr/local/
    echo 'export GOROOT=/usr/local/go' >> /home/vagrant/.bashrc
    echo 'export GOPATH=$HOME/go' >> /home/vagrant/.bashrc
    echo 'export PATH=$PATH:$GOROOT/bin:$GOPATH/bin' >> /home/vagrant/.bashrc
    export GOROOT=/usr/local/go
    export GOPATH=$HOME/go
    export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
    mkdir -p /home/vagrant/go/src
    rm -rf /home/vagrant/go1.9.1.linux-amd64.tar.gz
  SHELL
  config.vm.provider :virtualbox do |v|
      # v.customize ["modifyvm", :id, "--cpus", 4]
      # v.customize ["modifyvm", :id, "--memory", 4096]
      # v.customize ['modifyvm', :id, '--nicpromisc1', 'allow-all']
  end
end
