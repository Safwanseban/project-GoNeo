package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Safwanseban/voixme-project/internal/types"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
)

func TestFetch(t *testing.T) {

	type args struct {
		query string
	}
	tests := []struct {
		name             string
		args             args
		expectedStruct   *types.OfferCompany
		expectedHttpCode int
		expectedError    error
	}{

		{
			name: "error fetching data",
			args: args{
				query: "",
			},
			expectedStruct: &types.OfferCompany{
				OfferID:     1,
				ClientID:    1,
				Country:     "india",
				Image:       "test_image",
				ImageWidth:  100,
				ImageHeight: 100,
			},
			expectedHttpCode: http.StatusBadRequest,
			expectedError:    errors.New("provide valid country parameter"),
		},
		{

			name: "success",
			args: args{
				query: "?country=India",
			},
			expectedStruct: &types.OfferCompany{
				OfferID:     1,
				ClientID:    1,
				Country:     "india",
				Image:       "test_image",
				ImageWidth:  100,
				ImageHeight: 100,
			},
			expectedHttpCode: http.StatusOK,
			expectedError:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var company *types.OfferCompany
			w := testSetup(http.MethodGet, "/"+tt.args.query, nil, nil)
			data, err := io.ReadAll(w.Body)
			require.NoError(t, err)
			err = json.Unmarshal(data, company)
			require.NoError(t, err)
			require.Equal(t, w.StatusCode, tt.expectedHttpCode)
		})
	}
}

func testSetup(method string, url string,
	body *bytes.Buffer, context map[string]interface{},

) *http.Response {
	router := fiber.New()
	newserver := &Server{
		App: router,
	}
	router.Get("/", newserver.Fetch)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(new(types.OfferCompany))
	}))
	client := server.Client()
	req, _ := http.NewRequest(method, server.URL, body)
	resp, _ := client.Do(req)

	return resp

}
