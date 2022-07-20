package snowflake

import (
	"fmt"
	"math/rand"
)

func GenerateKey(datacenterID, machineID uint64) uint64 {
	fmt.Println(datacenterID, machineID)
	return datacenterID + machineID + rand.Uint64()
}
