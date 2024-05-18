package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var rate rateService
var Rate = &rate

type rateService struct {
	apiUrl  string
	curRate string
	curErr  error
}

type apiResponse struct {
	Ask string `json:"ask"`
}

func (r *rateService) fetchRate() (string, error) {
	if len(r.apiUrl) < 1 {
		return "", fmt.Errorf("no apiUrl")
	}

	res, err := http.Get(r.apiUrl)
	if err != nil {
		return "", err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var data []apiResponse

	if err = json.Unmarshal(resBody, &data); err != nil {
		return "", err
	}

	if len(data) < 1 {
		return "", fmt.Errorf("got empty response from API")
	}

	return data[0].Ask, nil
}

func (r *rateService) updateRate() (string, error) {
	r.curRate, r.curErr = r.fetchRate()

	if r.curErr != nil {
		log.Printf("Update rate error: %s", r.curErr)
	} else {
		log.Printf("Updated rate: %s", r.curRate)
	}

	return r.curRate, r.curErr
}

func (r *rateService) Init(url string) error {
	r.apiUrl = url
	_, err := r.updateRate()
	return err
}

func (r *rateService) GetCurrentRate() (string, error) {
	return r.curRate, r.curErr
}

func (r *rateService) GetUpdatedRate() (string, error) {
	return r.updateRate()
}
