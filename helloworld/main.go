package helloworld

import "github.com/CityOfZion/neo-storm/interop/runtime"

func Main(op string, args []interface{}) {
	runtime.Log("Hello world!")
}
