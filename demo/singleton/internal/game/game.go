package game

import (
	"fmt"
	"math/rand"
	"time"
)

func Startup() {
	rand.Seed(time.Now().Unix())
	fmt.Println("Game Server Startup")
}
