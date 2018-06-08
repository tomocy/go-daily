package midtest

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func LiveDailyLife() {
	tests := []string{
		"english",
		"math",
		"japanese",
	}
	wg := &sync.WaitGroup{}
	testCnt := len(tests)
	for i := 0; i < testCnt; i++ {
		wg.Add(1)
		go func(indx int) {
			defer wg.Done()

			takeMidtermTest(tests[indx])

		}(i)
	}

	wg.Wait()

	fmt.Println("All tests were over!!!!")
}

func takeMidtermTest(test string) {
	exps := []string{
		"so difficult :(",
		"so so...",
		"too easy!",
	}

	rand.Seed(time.Now().UnixNano())
	exp := exps[rand.Intn(len(exps))]
	b := make([]byte, 0, 10)

	b = append(b, "This "...)
	b = append(b, test...)
	b = append(b, " is "...)
	b = append(b, exp...)

	time.Sleep(2 * time.Second)
	fmt.Println(string(b))
	time.Sleep(3 * time.Second)

}
