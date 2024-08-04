package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func ReadApi(url string, sptr interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = errors.New(" error in error")
		return err
	}
	sbyte, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(sbyte, sptr)
	if err != nil {
		return err
	}
	return nil
}
