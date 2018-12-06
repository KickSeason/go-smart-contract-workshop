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
    * Learn using neo-python's build command with the test argument
    * Test differences between Log and Notify
* runtime smart contract, using runtime.GetTrigger, runtime.Verification, runtime.CheckWitness and runtime.Application: [runtime.go](https://github.com/CityOfZion/neo-storm/blob/master/examples/runtime/runtime.go)
* storage using smart contract: [storage.go](https://github.com/CityOfZion/neo-storm/blob/master/examples/storage/storage.go)
    * Storage is one of the key components of most smart contracts.
    * support bytes and string.
    * ~~Learn about debugstorage on/off/reset.~~
* Check out Dictionary support and neo.Runtime.Serialize
* a domain registration smart contract: [domain.go](https://github.com/KickSeason/neo-go-workshop/blob/master/domain/domain.go)
    * users can query, register, transfer and delete domains
    * important concept: checking of ownership
* neo nep5 smart contract: [nep5.go](https://github.com/CityOfZion/neo-storm/tree/master/examples/token), reference: [Step-by-step guide on issuing your NEP-5 token on NEO’s Private net using Go](https://medium.com/coinmonks/neo-token-contract-nep-5-in-go-f6b0102c59ee)
* ICO template: [ICOTemplate.go](https://github.com/KickSeason/neo-go-workshop/blob/master/ICOTemplate/ICOTemplate.go)

# QuickStart

# Typical method signatures

# Often used imports

# Often used build commands

# Reference
[neo-storm: the neo smart contract framework for golang](https://github.com/CityOfZion/neo-storm)

[examples](https://github.com/CityOfZion/neo-storm/tree/master/examples)
