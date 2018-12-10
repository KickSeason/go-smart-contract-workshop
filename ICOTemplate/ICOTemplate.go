package ICOTemplate

import (
	"github.com/CityOfZion/neo-storm/interop/engine"
	"github.com/CityOfZion/neo-storm/interop/runtime"
	"github.com/CityOfZion/neo-storm/interop/storage"
	"github.com/CityOfZion/neo-storm/interop/transaction"
	"github.com/CityOfZion/neo-storm/interop/util"
)

const (
	factor  uint = 100000000
	decimal uint = 100000000
)

type Token struct {
	Name        string
	Symbol      string
	Owner       []byte
	Decimals    uint
	AssetID     []byte
	TotalAmount uint
}

type ICOSettings struct {
	Factor     uint
	NeoDecimal uint
	PreICOCap  uint
	BasicRate  uint
	Start      int
	End        int
	Token      Token
}

func Main(op string, args []interface{}) interface{} {
	runtime.Notify("Hello world!")
	return nil
}

func (I ICOSettings) Deploy(ctx storage.Context) bool {
	totalsupply := storage.Get(ctx, "totalSupply")
	if 0 != totalsupply {
		return false
	}
	storage.Put(ctx, I.Token.Owner, I.PreICOCap)
	storage.Put(ctx, "totalSupply", I.PreICOCap)
	return true
}

func (I ICOSettings) MintToken(ctx storage.Context) bool {
	return true
}

func (I ICOSettings) TotalSupply(ctx storage.Context) uint {
	totalsupply := storage.Get(ctx, "totalSupply").(uint)
	return totalsupply
}

func (I ICOSettings) CurrentSwapRate(ctx storage.Context) uint {
	return 0
}

func (I ICOSettings) CurrentSwapToken(ctx storage.Context) uint {
	return 0
}
func (T Token) Transfer(ctx storage.Context, from []byte, to []byte, value uint) bool {
	if value <= 0 {
		return false
	}
	if !runtime.CheckWitness(from) {
		return false
	}
	if len(to) != 20 {
		return false
	}
	frombalance := storage.Get(ctx, from).(uint)
	if frombalance < value {
		return false
	}
	if util.Equals(from, to) {
		return true
	}
	if frombalance == value {
		storage.Delete(ctx, from)
	} else {
		storage.Put(ctx, from, frombalance-value)
	}
	tobalance := storage.Get(ctx, to).(uint)
	storage.Put(ctx, to, tobalance+value)
	return true
}

func (T Token) BalanceOf(ctx storage.Context, address []byte) uint {
	balance := storage.Get(ctx, address).(uint)
	return balance
}

func GetSender() []byte {
	tx := engine.GetScriptContainer()
	reference := transaction.GetReferences(tx)
	return []byte{}
}
func GetReceiver() []byte {
	return []byte{}
}
func GetContributeValue() uint64 {
	return 0
}
