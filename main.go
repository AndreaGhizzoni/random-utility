package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/AndreaGhizzoni/Random/randutil"
)

func main() {
	rand.Seed(time.Now().Unix())
	fmt.Println(randutil.RandomInt64(1, 10))
	for i := 0; i < 10; i++ {
		fmt.Println(randutil.RandomFloat64(1, 10))
	}
}
