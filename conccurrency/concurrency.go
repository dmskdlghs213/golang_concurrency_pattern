package conccurrency

import (
	"fmt"
	"math/rand"
	"time"
)

func ConccurrencyCall() (string, error) {

	rand.Seed(time.Now().Unix())
	randomCount := rand.Intn(60)

	// 実行回数はランダム
	counts := make([]uint16, 0, randomCount)
	for i := 1; i <= randomCount; i++ {
		count := uint16(i)
		counts = append(counts, count)
	}

	fmt.Println(counts)

	var msgResult string
	switch {
	case len(counts) < 20:
		fmt.Println("20回以内です")
		msg, err := Channel(counts)
		if err != nil {
			return "", err
		}
		msgResult = msg
	case len(counts) < 40:
		fmt.Println("40回以内です")
	case len(counts) < 60:
		fmt.Println("60回以内です")
	}

	return msgResult, nil
}
