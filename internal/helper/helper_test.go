package helper

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// Kullanacağımız test struct
type User struct {
	Name  string `json:"name"  validate:"required,min=3"`
	Email string `json:"email" validate:"required,email"`
}

func TestRandomNumbers(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{"Length 4", 4},
		{"Length 9", 9},
		{"Length greater than 9", 15}, // 9'a düşecek
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			num, err := RandomNumbers(tt.length)
			assert.NoError(t, err)
			assert.NotZero(t, num)
			// Length kontrolü
			assert.LessOrEqual(t, len(string(rune(num))), tt.length)
		})
	}
}

func TestParseValidated(t *testing.T) {
	app := fiber.New()

	app.Post("/test", func(c *fiber.Ctx) error {
		var user User
		if err := ParseValidated(c, &user); err != nil {
			return err
		}
		return c.JSON(user)
	})

	tests := []struct {
		name           string
		body           string
		expectedStatus int
		expectedError  string
	}{
		{
			name:           "Invalid JSON",
			body:           `{"name": "John"`, // eksik kapatma
			expectedStatus: 400,
			expectedError:  "Invalid request body",
		},
		{
			name:           "Missing fields",
			body:           `{}`,
			expectedStatus: 400,
			expectedError:  "Name is required", // içerik kontrolü
		},
		{
			name:           "Invalid email and short name",
			body:           `{"name":"Jo","email":"not-email"}`,
			expectedStatus: 400,
			expectedError:  "Name must be at least 3 characters", // bir kısmı
		},
		{
			name:           "Valid request",
			body:           `{"name":"John","email":"john@example.com"}`,
			expectedStatus: 200,
			expectedError:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(
				"POST",
				"/test",
				bytes.NewBuffer([]byte(tt.body)),
			)
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			buf := new(bytes.Buffer)
			buf.ReadFrom(resp.Body)
			bodyStr := buf.String()

			if tt.expectedError != "" {
				assert.Contains(t, bodyStr, tt.expectedError)
			}
		})
	}
}
