package helpers

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SwitchInput(address string, input string, output string) error {
	resp, err := http.Get(fmt.Spritnf("http://%s/Port/Set/%s/%s", address, input, output))

	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		responseBody = ioutil.ReadAll(resp.Body)
		return errors.New(fmt.Sprintf("Pulse eight returned error code: %s and error %s", resp.StatusCode, responseBody))
	}

	return nil
}