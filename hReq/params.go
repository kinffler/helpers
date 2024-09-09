package helpers

import (
	"fmt"
	"log"
	"net/url"
)

func params(baseURL string, params map[string]interface{}) string {
	if len(params) == 0 {
		return baseURL
	}

	u, err := url.Parse(baseURL)
	if err != nil {
		log.Println("Error parsing URL:", err)
		return baseURL
	}

	q := u.Query()
	for key, value := range params {
		switch v := value.(type) {
		case int:
			q.Set(key, fmt.Sprintf("%d", v))
		case uint:
			q.Set(key, fmt.Sprintf("%d", v))
		default:
			q.Set(key, fmt.Sprintf("%v", v)) // String
		}
	}
	u.RawQuery = q.Encode()

	return u.String()
}
