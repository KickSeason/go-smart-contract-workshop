package ICOTemplate

import "github.com/CityOfZion/neo-storm/interop/runtime"

func Main(op string, args []interface{}) {
    runtime.Notify("Hello world!")
}