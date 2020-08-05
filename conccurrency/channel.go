package conccurrency

import (
	"fmt"
	"time"

	"github.com/dmskdlghs213/golang_concurrency_pattern/client"
	"github.com/labstack/gommon/log"
)

func UseChannel(counts []uint16) (string, error) {

	msgCh := make(chan string, 1)
	errCh := make(chan error, 1)

	// 実行数を制御
	for i := 0; i < len(counts); i++ {

		fmt.Println("========start==========")
		go func(i int) {
			fmt.Println(i, "回目")

			msg, err := client.APICall()
			if err != nil {
				errCh <- err
			}

			msgCh <- msg
			errCh <- nil

			defer func() {
				if r := recover(); r != nil {
					log.Error(fmt.Sprintf("Recovered from:%v", r))
				}
			}()
		}(i)
	}

	// メインルーチンを止めてサブルーチンの終了を待つ
	time.Sleep(2 * time.Second)

	if errCh != nil {
		return "", <-errCh
	}

	msg := <-msgCh

	return msg, nil
}
