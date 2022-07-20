package snowflake

import (
	"fmt"
	"math/rand"
	"time"
)

func Snowflake(datacenterID, machineID uint64) uint64 {

	timestamp := time.Now().UnixMilli()

	fmt.Printf("%b\n", datacenterID)

	return uint64(timestamp) + datacenterID + machineID + rand.Uint64()
}
