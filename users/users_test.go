package users

import (
	"encoding/json"
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

type mockTypicode struct {
}

func (tc *mockTypicode) Decode(result interface{}) error {
	body := strings.NewReader(`[{
		"id": 1,
		"name": "Leanne Graham",
		"username": "Bret",
		"email": "Sincere@april.biz",
		"phone": "1-770-736-8031 x56442"
	}]`)
	return json.NewDecoder(body).Decode(&result)
}

func TestGetAllUsersHandler(t *testing.T) {
	c := &mockContext{}
	api := &usersAPI{
		service: &mockTypicode{},
	}

	api.getUsers(c)

	if c.statusCode != http.StatusOK {
		t.Errorf("expect status %d, but got %d", http.StatusOK, c.statusCode)
	}
}
