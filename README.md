# neo-go-workshop
neo workshop for golang developers
# Requisite
* operating system ubuntu descktop 18.04 lts
* Git
* golang v1.10, dep 
* docker ce, docker-compose
# Install neo-storm
```
cd $GOPATH/src
git clone https://github.com/CityOfZion/neo-storm.git
cd neo-storm
dep ensure -update
make install
```
> **NOTE**: if permission problems occur, please check the permission of the directory. Be careful with the GOPATH env between root and ordinary users.
# Steps in workshop
* setup your private neo blockchain using [neo-local](https://github.com/cityofzion/neo-local)
* First smart contract, definitely "Hello World": [helloworld.go](https://github.com/KickSeason/neo-go-workshop/tree/master/helloworld)
* runtime smart contract: [runtime.go](https://github.com/CityOfZion/neo-storm/blob/master/examples/runtime/runtime.go)
* storage using smart contract: [storage.go](https://github.com/CityOfZion/neo-storm/blob/master/examples/storage/storage.go)
* a domain registration smart contract: [domain.go](https://github.com/KickSeason/neo-go-workshop/blob/master/domain/domain.go)
* neo nep5 smart contract: nep5.go, reference: [Step-by-step guide on issuing your NEP-5 token on NEO’s Private net using Go](https://medium.com/coinmonks/neo-token-contract-nep-5-in-go-f6b0102c59ee)
* ICO template: ICOTemplate.go

# QuickStart

# Typical method signatures

# Often used imports

# Often used build commands

# Warning
support build in functions so far: "__len__", "__append__", "__SHA256__", "__SHA1__", "__Hash256__", "__Hash160__", "__FromAddress__", "__Equals__".
> __FromAddress__ only support basic type 'string', variable not supported presently. 
# Reference
[neo-storm: the neo smart contract framework for golang](https://github.com/CityOfZion/neo-storm)

[examples](https://github.com/CityOfZion/neo-storm/tree/master/examples)
