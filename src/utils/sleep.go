package utils

import (
	"math/rand"
	"time"
)

func DoSleepN(n int) {
	time.Sleep(time.Duration(n) * time.Second)
}

func DoSleepRange(min int, max int) {
	var randN int = rand.Intn(max-min) + min

	rand.Seed(time.Now().UnixMicro())
	time.Sleep(time.Duration(randN) * time.Second)
}
