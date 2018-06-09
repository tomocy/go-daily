package morning

import (
	"fmt"
	"math/rand"
	"time"
)

func LiveDailyLife() {
	waked := make(chan bool)
	go func(waked chan bool) {
		t := time.NewTicker(1 * time.Second)
		defer t.Stop()

		fmt.Print("zzz")
		for {
			select {
			case <-t.C:
				if 本当に起きますか() {
					waked <- true
					break
				}

				fmt.Print(".")
			}
		}
	}(waked)

	<-waked

	fmt.Println()
	fmt.Println("Good morning")
	// A day starts
}

func 本当に起きますか() bool {
	prob := 3
	divisor := 10
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(divisor)+1 <= prob
}
