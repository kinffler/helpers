package helpers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func (r *Request) Post() ([]byte, error) {
	if r.ContentType == "" {
		r.ContentType = "application/json"
	}

	payload, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	var req *http.Request
	if r.ContentType == "multipart/form-data" {
		bodyBuffer := &bytes.Buffer{}
		bodyWriter := multipart.NewWriter(bodyBuffer)

		if r.File != "" {
			file, err := os.Open(r.File)
			if err != nil {
				return nil, err
			}
			defer file.Close()

			if r.FileKey == "" {
				r.FileKey = "file"
			}

			fileWriter, err := bodyWriter.CreateFormFile(r.FileKey, filepath.Base(r.File))
			if err != nil {
				return nil, err
			}
			_, err = io.Copy(fileWriter, file)
			if err != nil {
				return nil, err
			}
		}

		for key, value := range r.Params {
			_ = bodyWriter.WriteField(key, fmt.Sprintf("%v", value))
		}

		err = bodyWriter.Close()
		if err != nil {
			return nil, err
		}

		req, err = http.NewRequest("POST", r.Url, bodyBuffer)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", bodyWriter.FormDataContentType())

	} else {
		req, err = http.NewRequest("POST", r.Url, bytes.NewBuffer(payload))
		if err != nil {
			return nil, err
		}

		req.Header.Add("Content-Type", r.ContentType)
		r.customHeaders(req)
	}

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
