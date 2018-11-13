package helloworld

import (
	"github.com/CityOfZion/neo-storm/interop/runtime"
	"github.com/CityOfZion/neo-storm/interop/util"
)

func Main(op string, args []interface{}) {
	addr := "Adawkc6QsS7KHFNm6YMLhPC2UhMu53FuG8"
	bs := util.FromAddress(addr)
	runtime.CheckWitness(bs)
}
