package myAxios

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func GetRequest(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, errors.New("failed to fetch data from the URL")
	}
	defer response.Body.Close()

	if response.StatusCode > 299 {
		return nil, errors.New("failed to fetch Pokemon: " + fmt.Sprint(response.StatusCode))
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("failed to read the response body - " + err.Error())
	}

	return body, nil
}
