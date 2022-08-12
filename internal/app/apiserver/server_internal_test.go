package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mouldykitz/duty_checklist/store/sqlstore/teststore"
	"github.com/stretchr/testify/assert"
)

func TestServer_HandleUsersCreate(t *testing.T) {
	s := newServer(teststore.New())
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":     "user@example.com",
				"passsword": "password",
			},
			expectedCode: http.StatusCreated,
		},
	}

	// Цикл, который запустает test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/users", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}

	// rec := httptest.NewRecorder()
	// req, _ := http.NewRequest(http.MethodPost, "/users", nil)
	// s := newServer(teststore.New())
	// s.ServeHTTP(rec, req)
	// assert.Equal(t, rec.Code, http.StatusOK)
}
