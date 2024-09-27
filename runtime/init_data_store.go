package runtime

import (
	redisDataStore "github.com/gaullow/goflow/core/redis-datastore"
	"github.com/gaullow/goflow/core/sdk"
)

func initDataStore(redisURI string, password string) (dataStore sdk.DataStore, err error) {
	dataStore, err = redisDataStore.GetRedisDataStore(redisURI, password)
	return dataStore, err
}
