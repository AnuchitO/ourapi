package users

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/labstack/echo"
)

type mockContext struct {
	echo.Context
	statusCode int
}

func (m *mockContext) JSON(code int, i interface{}) error {
	m.statusCode = code
	return nil
}

func mockGet(url string) (*http.Response, error) {
	body := ioutil.NopCloser(strings.NewReader(`[{
		"id": 1,
		"name": "Leanne Graham",
		"username": "Bret",
		"email": "Sincere@april.biz",
		"phone": "1-770-736-8031 x56442"
	}]`))

	resp := &http.Response{
		Body: body,
	}
	return resp, nil
}
func TestGetAllUsersHandler(t *testing.T) {
	c := &mockContext{}
	api := &usersAPI{
		get: mockGet,
	}

	api.getUsers(c)

	if c.statusCode != http.StatusOK {
		t.Errorf("expect status %d, but got %d", http.StatusOK, c.statusCode)
	}
}
