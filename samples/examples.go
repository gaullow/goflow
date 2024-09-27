package main

import (
	"fmt"
	"github.com/gaullow/goflow/samples/condition"
	"github.com/gaullow/goflow/samples/loop"
	"github.com/gaullow/goflow/samples/myflow"
	"github.com/gaullow/goflow/samples/parallel"
	"github.com/gaullow/goflow/samples/serial"
	"github.com/gaullow/goflow/samples/single"
	goflow "github.com/gaullow/goflow/v1"
)

func main() {
	fs := &goflow.FlowService{
		Port: 8080,
		//RedisURL: "localhost:6379",
		RedisURL:          "52.77.42.5:6379",
		RedisPassword:     "game",
		WorkerConcurrency: 5,
		EnableMonitoring:  true,
		DebugEnabled:      true,
	}
	fs.Register("single", single.DefineWorkflow)
	fs.Register("serial", serial.DefineWorkflow)
	fs.Register("parallel", parallel.DefineWorkflow)
	fs.Register("condition", condition.DefineWorkflow)
	fs.Register("loop", loop.DefineWorkflow)
	fs.Register("myflow", myflow.DefineWorkflow)
	fmt.Println(fs.Start())
}
