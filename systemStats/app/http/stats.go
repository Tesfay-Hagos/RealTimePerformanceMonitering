package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"systemstats/app/config"
)

func SendStatsToProducer(stats map[string]int) (map[string]int, error) {
	producerUrl := config.NewAppConfig().ProducerUrl

	requestBodyBytes, err := json.Marshal(stats)
	if err != nil {
		return nil, err
	}
	requestBody := bytes.NewReader(requestBodyBytes)
	request, err := http.NewRequest("POST", producerUrl, requestBody)
	if err != nil {
		fmt.Println("Request error", err)
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println("Post request failed", err)
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("unable to get send system stats to producer")
	}
	if err != nil {
		return nil, err
	}
	return nil, nil
}
