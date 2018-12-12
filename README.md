# neo-go-workshop
neo smart contract workshop for golang developers
# Requisite
* operating system ubuntu descktop 18.04 lts
* Git
* golang v1.10, dep 
* docker-ce, docker-compose

# Install neo-storm
```
cd $GOPATH/src
git clone https://github.com/CityOfZion/neo-storm.git
cd neo-storm
dep ensure -update
make install
```
> **NOTE**: if permission problems occur, please check the permission of the directory. Be careful with the GOPATH env between root and ordinary users.

# Develop Enviroment
 now neo-storm not support test and deploy contract, so we can use neo-python
 
* setup your private neo blockchain using [neo-local](https://github.com/cityofzion/neo-local)
* install neo-python: https://neo-python.readthedocs.io/en/latest/install.html , and refering this for interacting with smart contract:https://neo-python.readthedocs.io/en/latest/neo/SmartContract/smartcontracts.html
* initialize your contract
```
neo-storm init -n mycontract
```
* deploy your contract using np-promt after you install neo-python, for example:
```
import contract ~/mycontract/main.avm 0710 f0 True False False
```
* test your contract
```
testinvoke 0xb1f0d096431eac7c90eba08a6b9dae692e01681c  b'register' [b'neo.org', b'\xef\x41\x9e\x5c\x66\x05\xf0\xad\xc6\xb8\x6f\x47\xff\x61\x19\xf9\xc2\x5d\x0e\x0f']
```
# Steps in workshop
* First smart contract, definitely "Hello World": [helloworld.go](https://github.com/KickSeason/neo-go-workshop/tree/master/helloworld)
    * Learn using neo-python's build command with the test argument
    * Test differences between Log and Notify
* runtime smart contract, using runtime.GetTrigger, runtime.Verification, runtime.CheckWitness and runtime.Application: [runtime.go](https://github.com/CityOfZion/neo-storm/blob/master/examples/runtime/runtime.go)
   * as now neo-storm not support deploy and test, you can use neo-python
* storage using smart contract: [storage.go](https://github.com/CityOfZion/neo-storm/blob/master/examples/storage/storage.go)
    * Storage is one of the key components of most smart contracts.
    * support bytes and string.
    * ~~Learn about debugstorage on/off/reset.~~
* Check out Dictionary support and neo.Runtime.Serialize
* a domain registration smart contract: [domain.go](https://github.com/KickSeason/neo-go-workshop/blob/master/domain/domain.go)
    * users can query, register, transfer and delete domains
    * important concept: checking of ownership
* neo nep5 smart contract: [nep5.go](https://github.com/CityOfZion/neo-storm/tree/master/examples/token)
   * reference: [Step-by-step guide on issuing your NEP-5 token on NEO’s Private net using Go](https://medium.com/coinmonks/neo-token-contract-nep-5-in-go-f6b0102c59ee)
* ICO template: [ICOTemplate.go](https://github.com/KickSeason/neo-go-workshop/blob/master/ICOTemplate/ICOTemplate.go)

# Warning
support build in functions so far: "__len__", "__append__", "__SHA256__", "__SHA1__", "__Hash256__", "__Hash160__", "__FromAddress__", "__Equals__".
> __FromAddress__ only support basic type 'string', variable not supported presently. 

# Reference
[neo-storm: the neo smart contract framework for golang](https://github.com/CityOfZion/neo-storm)

[examples](https://github.com/CityOfZion/neo-storm/tree/master/examples)

# TODO
workshop PPT
