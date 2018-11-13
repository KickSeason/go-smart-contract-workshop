package domain

import (
	"github.com/CityOfZion/neo-storm/interop/runtime"
	"github.com/CityOfZion/neo-storm/interop/storage"
)

func Main(op string, args []interface{}) interface{} {
	if len(args) == 0 {
		runtime.Log("args lenght 0")
		return false
	}
	if "query" == op {
		if checkArgs(args, 1) {
			dname := args[0].([]byte)
			return QueryDomain(dname)
		}
	} else if "delete" == op {
		if checkArgs(args, 1) {
			dname := args[0].([]byte)
			return DeleteDomain(dname)
		}
	} else if "register" == op {
		if checkArgs(args, 2) {
			dname := args[0].([]byte)
			scriptHash := args[1].([]byte)
			return RegisterDomain(dname, scriptHash)
		}
	} else if "transfer" == op {
		if checkArgs(args, 2) {
			dname := args[0].([]byte)
			toScriptHash := args[1].([]byte)
			return TransferDomain(dname, toScriptHash)
		}
	}
	return false
}
func checkArgs(args []interface{}, cnt int) bool {
	if len(args) < cnt {
		return false
	}
	return true
}
func QueryDomain(dname []byte) interface{} {
	runtime.Notify("QueryDomain: ", dname)

	ctx := storage.GetContext()
	owner := storage.Get(ctx, dname).([]byte)
	if len(owner) == 0 {
		runtime.Notify("Domain is not yet registered.")
		return false
	}
	return owner
}

func RegisterDomain(dname []byte, scriptHash []byte) bool {
	runtime.Notify("RegisterDomain: ", dname, scriptHash)

	if !runtime.CheckWitness(scriptHash) {
		runtime.Notify("Owner arguement is not the same as the sender.")
		return false
	}
	ctx := storage.GetContext()
	ownerExist := storage.Get(ctx, dname).([]byte)
	if len(ownerExist) == 0 {
		runtime.Notify("Domain has already been registered already.")
		return false
	}
	storage.Put(ctx, dname, scriptHash)
	return true
}

func TransferDomain(dname []byte, toScriptHash []byte) bool {
	runtime.Notify("TransferDomain: ", dname, toScriptHash)

	ctx := storage.GetContext()
	owner := storage.Get(ctx, dname).([]byte)
	if len(owner) == 0 {
		runtime.Notify("Domain is not yet registered.")
		return false
	}
	if !runtime.CheckWitness(owner) {
		runtime.Notify("Sender is not the owner.")
		return false
	}
	if len(toScriptHash) != 20 {
		runtime.Notify("Invalid new owner address scripthash. Must be exactly 20 characters.")
		return false
	}
	storage.Put(ctx, dname, toScriptHash)
	return true
}

func DeleteDomain(dname []byte) bool {
	runtime.Notify("DeleteDomain: ", dname)

	ctx := storage.GetContext()
	owner := storage.Get(ctx, dname).([]byte)
	if len(owner) == 0 {
		runtime.Notify("Domain is not yet registered.")
		return false
	}
	if !runtime.CheckWitness(owner) {
		runtime.Notify("Sender is not the owner.")
		return false
	}
	storage.Delete(ctx, dname)
	return true
}
