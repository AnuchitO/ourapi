package typicode

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Decoder interface {
	Decode(result interface{}) error
}

type typicode struct {
	url string
}

func (tc *typicode) Decode(result interface{}) error {
	resp, err := http.Get(tc.url)
	if err != nil {
		fmt.Println("error requset", err)
		return err
	}
	return json.NewDecoder(resp.Body).Decode(&result)
}

func Get(path string) *typicode {
	return &typicode{
		url: "https://jsonplaceholder.typicode.com" + path,
	}
}
