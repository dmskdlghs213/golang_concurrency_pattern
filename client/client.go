package client

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

func APICall() (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// "hello!!とだけ返すAPIサーバーを想定"
	req, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
		return "", errors.Wrap(err, "request error")
	}

	rsp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return "", errors.Wrap(err, "response error")
	}

	defer func() error {
		err = rsp.Body.Close()
		if err != nil {
			return err
		}
		return nil
	}()

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "", errors.Wrap(err, "response body reading error")
	}

	fmt.Println(string(body))

	return string(body), nil
}
