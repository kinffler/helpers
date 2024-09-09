package helpers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

func (r *Request) Put() ([]byte, error) {
	if r.ContentType == "" {
		r.ContentType = "application/json"
	}

	payload, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	url := params(r.Url, r.Params)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", r.ContentType)
	r.customHeaders(req)

	if r.Authorization != "" {
		req.Header.Add("Authorization", r.Authorization)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	decoded, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	r.StatusCode = resp.StatusCode

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		message := string(decoded)
		log.Println(message)
		return nil, errors.New(resp.Status)
	}

	return decoded, nil
}
