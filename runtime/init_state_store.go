package runtime

import (
	redisStateStore "github.com/gaullow/goflow/core/redis-statestore"
	"github.com/gaullow/goflow/core/sdk"
)

func initStateStore(redisURI string, password string) (stateStore sdk.StateStore, err error) {
	stateStore, err = redisStateStore.GetRedisStateStore(redisURI, password)
	return stateStore, err
}
