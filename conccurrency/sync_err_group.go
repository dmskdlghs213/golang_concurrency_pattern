package conccurrency

import (
	"fmt"

	"github.com/dmskdlghs213/golang_concurrency_pattern/client"
	"github.com/labstack/gommon/log"
	"golang.org/x/sync/errgroup"
)

func SyncErrGroup(counts []uint16) error {

	var eg errgroup.Group
	for i := 0; i < len(counts); i++ {
		n := i
		fmt.Println("========start==========")
		eg.Go(func() error {
			fmt.Println(n, "回目")

			_, err := client.APICall()
			if err != nil {
				return err
			}

			defer func() {
				if r := recover(); r != nil {
					log.Error(fmt.Sprintf("Recovered from:%v", r))
				}
			}()

			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return err
	}

	return nil
}
