package public

import (
	"encoding/json"
	"net/http"
	u "net/url"
)

func getRequest(url string) (*http.Request, error) {
	req := request()

	URL, err := u.Parse(url)
	if err != nil {
		return nil, err
	}

	req.Method = "GET"
	req.URL = URL

	return req, nil
}

func request() *http.Request {
	req := &http.Request{
		Header: http.Header{},
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/vnd.drivr.v3+json")
	// TODO: configurable token
	req.Header.Set("Authorization", `Token token="4993a3bb8d4a40f08065e640621fee63"`)

	return req
}

func decode(resp *http.Response, target interface{}) error {
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}
