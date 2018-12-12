package icotemplate

import (
	"github.com/CityOfZion/neo-storm/interop/engine"
	"github.com/CityOfZion/neo-storm/interop/output"
	"github.com/CityOfZion/neo-storm/interop/runtime"
	"github.com/CityOfZion/neo-storm/interop/storage"
	"github.com/CityOfZion/neo-storm/interop/transaction"
	"github.com/CityOfZion/neo-storm/interop/util"
)

const (
	factor     int = 100000000
	neodecimal int = 100000000
)

var neoAssetID = []byte{155, 124, 255, 218, 166, 116, 190, 174, 15, 147, 14, 190, 96, 133, 175, 144, 147, 229, 254, 86, 179, 74, 92, 34, 12, 205, 207, 110, 252, 51, 111, 197}

//ICOSettings settings for a ICO
type ICOSettings struct {
	Name        string
	Symbol      string
	Owner       []byte
	Decimals    int
	Factor      int
	NeoDecimal  int
	PreICOCap   int
	BasicRate   int
	Start       int
	End         int
	TotalAmount int
}

//CreateICO create a ico settings
func CreateICO() ICOSettings {
	return ICOSettings{
		Name:        "ztcoin",
		Symbol:      "ZTC",
		Owner:       util.FromAddress("AK2nJJpJr6o664CWJKi1QRXjqeic2zRp8y"),
		Decimals:    8,
		Factor:      factor,
		NeoDecimal:  neodecimal,
		PreICOCap:   30000000 * factor,
		BasicRate:   1000 * factor,
		Start:       1506787200,
		End:         1538323200,
		TotalAmount: 100000000 * factor,
	}
}

//Main entrance
func Main(op string, args []interface{}) interface{} {
	I := CreateICO()
	if runtime.GetTrigger() == runtime.Verification() {
		if len(I.Owner) == 20 ||
			len(I.Owner) == 33 {
			return runtime.CheckWitness(I.Owner)
		}
		return false
	} else if runtime.GetTrigger() == runtime.Application() {
		if op == "name" {
			return I.Name
		}
		if op == "symbol" {
			return I.Symbol
		}
		ctx := storage.GetContext()
		if op == "totalSupply" {
			return I.TotalSupply(ctx)
		}
		if op == "balanceOf" {
			if len(args) < 1 {
				return false
			}
			address := args[0].([]byte)
			return I.BalanceOf(ctx, address)
		}
		if op == "transfer" {
			if len(args) < 3 {
				return false
			}
			from := args[0].([]byte)
			to := args[1].([]byte)
			value := args[2].(int)
			return I.Transfer(ctx, from, to, value)
		}
		if op == "deploy" {
			return I.Deploy(ctx)
		}
		if op == "mintTokens" {
			return I.MintTokens(ctx)
		}
	}
	return nil
}

//Deploy initialize token
func (I ICOSettings) Deploy(ctx storage.Context) bool {
	totalsupply := storage.Get(ctx, "totalSupply")
	if 0 != totalsupply {
		return false
	}
	storage.Put(ctx, I.Owner, I.PreICOCap)
	storage.Put(ctx, "totalSupply", I.PreICOCap)
	return true
}

//MintTokens mint a number of tokens to neo
func (I ICOSettings) MintTokens(ctx storage.Context) bool {
	sender := GetSender()
	if len(sender) == 0 {
		return false
	}
	value := GetContributeValue()
	rate := I.CurrentSwapRate(ctx)
	if rate == 0 {
		return false
	}
	if value == 0 {
		return true
	}
	token := I.CurrentSwapToken(ctx, value, rate)
	if token == 0 {
		return false
	}
	balance := storage.Get(ctx, sender).(int)
	storage.Put(ctx, sender, balance+token)
	totalsupply := storage.Get(ctx, "totalSupply").(int)
	storage.Put(ctx, "totalSupply", totalsupply+token)
	return true
}

//TotalSupply total token supply
func (I ICOSettings) TotalSupply(ctx storage.Context) int {
	totalsupply := storage.Get(ctx, "totalSupply").(int)
	return totalsupply
}

//CurrentSwapRate current exchange rate between neo and token
func (I ICOSettings) CurrentSwapRate(ctx storage.Context) int {
	now := runtime.GetTime()
	if now < I.Start {
		return 0
	} else if I.End < now {
		return 0
	}
	return I.BasicRate
}

//CurrentSwapToken whether over contribute capacity, you can get the token amount
func (I ICOSettings) CurrentSwapToken(ctx storage.Context, value int, rate int) int {
	token := value / I.NeoDecimal * rate
	totalsupply := storage.Get(ctx, "totalSupply").(int)
	if I.TotalAmount <= totalsupply {
		return 0
	}
	if I.TotalAmount-totalsupply < token {
		token = I.TotalAmount - totalsupply
	}
	return token
}

//Transfer someone transfer tokens to other one
func (I ICOSettings) Transfer(ctx storage.Context, from []byte, to []byte, value int) bool {
	if value <= 0 {
		return false
	}
	if !runtime.CheckWitness(from) {
		return false
	}
	if len(to) != 20 {
		return false
	}
	frombalance := storage.Get(ctx, from).(int)
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
	tobalance := storage.Get(ctx, to).(int)
	storage.Put(ctx, to, tobalance+value)
	return true
}

//BalanceOf get the balance of a account with address
func (I ICOSettings) BalanceOf(ctx storage.Context, address []byte) int {
	balance := storage.Get(ctx, address).(int)
	return balance
}

//GetSender check whether asset is neo and get sender script hash
func GetSender() []byte {
	tx := engine.GetScriptContainer()
	reference := transaction.GetReferences(tx)
	for i := 0; i < len(reference); i++ {
		o := reference[i].(output.Output)
		if util.Equals(output.GetAssetID(o), neoAssetID) {
			return output.GetScriptHash(o)
		}
	}
	return []byte{}
}

//GetReceiver get smart contract script hash
func GetReceiver() []byte {
	return engine.GetExecutingScriptHash()
}

//GetContributeValue get all you contribute neo amount
func GetContributeValue() int {
	tx := engine.GetScriptContainer()
	outputs := transaction.GetOutputs(tx)
	var value int
	for i := 0; i < len(outputs); i++ {
		v := outputs[i]
		if util.Equals(output.GetScriptHash(v), GetReceiver()) &&
			util.Equals(output.GetAssetID(v), neoAssetID) {
			value += output.GetValue(v)
		}
	}
	return value
}
