package typicode

import (
	"encoding/json"
	"net/http"
)

type Decoder interface {
	Decode(result interface{}) error
}

type Doer interface {
	Do() (resp *http.Response, err error)
}

type doGet struct {
	url string
}

func (do *doGet) Do() (resp *http.Response, err error) {
	return http.Get(do.url)
}

type getTypicode struct {
	client Doer
}

func (tc *getTypicode) Decode(result interface{}) error {
	resp, err := tc.client.Do()
	if err != nil {
		return err
	}

	return json.NewDecoder(resp.Body).Decode(&result)
}

func Get(path string) *getTypicode {
	return &getTypicode{
		client: &doGet{
			url: "https://jsonplaceholder.typicode.com" + path,
		},
	}
}
