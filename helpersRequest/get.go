package helpersRequest

import (
	"errors"
	"io"
	"log"
	"net/http"
)

func (r *Request) Get() ([]byte, error) {
	if r.ContentType == "" {
		r.ContentType = "application/json"
	}

	client := &http.Client{}

	url := params(r.Url, r.Params)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	req.Header.Add("Content-Type", r.ContentType)
	r.customHeaders(req)

	if r.Authorization != "" {
		req.Header.Add("Authorization", r.Authorization)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer resp.Body.Close()

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
