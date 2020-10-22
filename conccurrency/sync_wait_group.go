package conccurrency

import (
	"fmt"
	"sync"

	"github.com/dmskdlghs213/golang_concurrency_pattern/client"
	"github.com/labstack/gommon/log"
)

func SyncWaitGroup(counts []uint16) error {

	var wg sync.WaitGroup
	for i := 0; i < len(counts); i++ {
		wg.Add(1)

		fmt.Println("========start==========")
		go func(i int) {
			fmt.Println(i, "回目")

			_, err := client.APICall()
			if err != nil {
				log.Error(err)
			}

			defer func() {
				wg.Done()
				if r := recover(); r != nil {
					log.Error(fmt.Sprintf("Recovered from:%v", r))
				}
			}()
		}(i)
	}
	wg.Wait()

	return nil
}
