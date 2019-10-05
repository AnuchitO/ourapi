package users

import (
	"encoding/json"
	"net/http"
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
	result []byte
}

func (tc *mockTypicode) Decode(result interface{}) error {
	return json.Unmarshal(tc.result, &result)
}

func TestGetAllUsersHandler(t *testing.T) {
	t.Run("get all users with status OK", func(t *testing.T) {
		c := &mockContext{}
		api := &usersAPI{
			service: &mockTypicode{
				result: []byte(`[{
					"id": 1,
					"name": "Leanne Graham",
					"username": "Bret",
					"email": "Sincere@april.biz",
					"phone": "1-770-736-8031 x56442"
				}]`),
			},
		}

		api.getUsers(c)

		if c.statusCode != http.StatusOK {
			t.Errorf("expect status %d, but got %d", http.StatusOK, c.statusCode)
		}
	})

	t.Run("get all user fail decode ", func(t *testing.T) {
		c := &mockContext{}
		api := &usersAPI{
			service: &mockTypicode{
				result: []byte(`invalid json should error`),
			},
		}

		api.getUsers(c)

		if c.statusCode != http.StatusInternalServerError {
			t.Errorf("expect status %d, but got %d", http.StatusInternalServerError, c.statusCode)
		}
	})
}
