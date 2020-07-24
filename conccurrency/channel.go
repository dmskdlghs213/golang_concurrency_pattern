package conccurrency

import (
	"fmt"
	"sync"
	"time"

	"github.com/dmskdlghs213/golang_concurrency_pattern/client"
	"github.com/labstack/gommon/log"
)

func Channel(counts []uint16) (string, error) {

	msgCh := make(chan string, 1)
	errCh := make(chan error, 1)

	// 実行数を制御
	var mu sync.Mutex
	for i := 0; i < len(counts); i++ {
		go func() {
			msg, err := client.APIClient()
			if err != nil {
				errCh <- err
			}

			msgCh <- msg
			errCh <- nil
			mu.Lock()

			defer func() {
				mu.Unlock()
				if r := recover(); r != nil {
					log.Error(fmt.Sprintf("Recovered from:%v", r))
				}
			}()
		}()
	}

	// メインルーチンを止めてサブルーチンの終了を待つ
	time.Sleep(1 * time.Second)

	if errCh != nil {
		return "", <-errCh
	}

	var msg string
	msg = <-msgCh

	return msg, nil
}
