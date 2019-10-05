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

type typicode struct {
	client Doer
}

func (tc *typicode) Decode(result interface{}) error {
	resp, err := tc.client.Do()
	if err != nil {
		return err
	}

	return json.NewDecoder(resp.Body).Decode(&result)
}

func Get(path string) *typicode {
	return &typicode{
		client: &doGet{
			url: "https://jsonplaceholder.typicode.com" + path,
		},
	}
}
