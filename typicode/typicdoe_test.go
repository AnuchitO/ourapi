package typicode

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

type mockDoGet struct {
}

func (m *mockDoGet) Do() (resp *http.Response, err error) {
	return &http.Response{
		Body: ioutil.NopCloser(strings.NewReader(`{"name": "AnuchitO"}`)),
	}, nil
}
func TestDoGetTypicode(t *testing.T) {
	tc := &typicode{
		client: &mockDoGet{},
	}
	result := struct {
		Name string `json:"name"`
	}{}

	err := tc.Decode(&result)

	if err != nil {
		t.Error("expecte should not error but got", err)
	}

	if result.Name != "AnuchitO" {
		t.Errorf("expected %s but got %s", "AnuchitO", result.Name)
	}
}
