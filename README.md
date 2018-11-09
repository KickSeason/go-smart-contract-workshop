# neo-go-workshop
neo workshop for golang developers
# Requisite
* operating system ubuntu descktop 18.04 lts
* Git
* golang v1.10, dep 
* docker ce, docker-compose
# Install neo-storm
cd $GOPATH/src
git clone https://github.com/CityOfZion/neo-storm.git
cd neo-storm
dep ensure -update
make install
> **NOTE**: if permission problems occur, please check the permission of the directory. Be careful with the GOPATH env between root and ordinary users.
# Steps in workshop
* setup your private neo blockchain using [neo-local](https://github.com/cityofzion/neo-local)
* First smart contract, definately "Hello World": helloworld.go
* runtime smart contract: runtime.go
* storage using smart contract: storage.go
* a domain registration smart contract: domain.go
* neo nep5 smart contract: nep5.go
* ICO template: ICOTemplate.go

# QuickStart

# Typical method signatures

# Often used imports

# Often used build commands

# Reference

