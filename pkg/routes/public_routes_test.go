package routes

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestPublicRoutes(t *testing.T) {

	tests := []struct {
		description   string
		route         string
		expectedError bool
		expectedCode  int
	}{
		{
			description:   "get all campains ",
			route:         "/api/v1/campaigns/",
			expectedError: false,
			expectedCode:  200,
		},
		{
			description:   "get campain by id",
			route:         "/api/v1/campaigns/" + uuid.New().String(),
			expectedError: false,
			expectedCode:  404,
		},
		{
			description:   "get campains by invalid ID (non UUID)",
			route:         "/api/v1/campaigns/123456",
			expectedError: false,
			expectedCode:  404, //ganti jadi 500 nanti
		},
	}

	app := fiber.New()

	// define routes
	PublicRoutes(app)

	for _, test := range tests {
		//
		req := httptest.NewRequest("GET", test.route, nil)
		req.Header.Set("Content-Type", "application/json")

		// Perform the request plain with the app.
		resp, err := app.Test(req, -1) // the -1 disables request latency

		assert.Equalf(t, test.expectedError, err != nil, test.description)

		if test.expectedError {
			continue
		}

		assert.Equal(t, test.expectedCode, resp.StatusCode, test.description)
	}
}
