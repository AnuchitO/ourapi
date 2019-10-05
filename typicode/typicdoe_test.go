package typicode

import "testing"

func TestDoGetTypicode(t *testing.T) {
	tc := &typicode{
		url: "http://dummy.com",
	}
	result := struct {
		Name string `json:"name"`
	}{}

	err := tc.Decode(&result)

	if err != nil {
		t.Error("expecte should not error but got", err)
	}
}
