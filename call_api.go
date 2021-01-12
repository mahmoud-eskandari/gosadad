package gosadad

import (
	"io"
	"io/ioutil"
	"net/http"
)

//Call HTTP
func call(url string, body io.Reader) ([]byte, error) {
	client := &http.Client{}
	httpReq, err := http.NewRequest("POST", url, body)

	if err != nil {
		return nil, err
	}
	httpReq.Header.Add("Content-Type", "application/json")

	res, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer func() { _ = res.Body.Close() }()

	b, err := ioutil.ReadAll(res.Body)
	return b, err
}
