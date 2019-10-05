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
}

func (tc *typicode) Decode(result interface{}) error {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		fmt.Println("error requset", err)
		return err
	}
	return json.NewDecoder(resp.Body).Decode(&result)
}

func NewGet() *typicode {
	return &typicode{}
}
