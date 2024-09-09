package helpersRequest

import "net/http"

func (r *Request) customHeaders(req *http.Request) {
	for key, value := range r.Header {
		req.Header.Add(key, value)
	}
}
