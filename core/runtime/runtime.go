package runtime

import (
	"github.com/gaullow/goflow/core/sdk/executor"
)

type Runtime interface {
	Init() error
	CreateExecutor(*Request) (executor.Executor, error)
}