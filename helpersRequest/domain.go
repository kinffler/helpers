package helpersRequest

type Request struct {
	Url           string
	ContentType   string
	Authorization string
	Header        map[string]string
	Body          interface{}
	Params        map[string]interface{}
	StatusCode    int
	FileKey       string
	File          string
}
