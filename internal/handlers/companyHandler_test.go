package handlers

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	mock "github.com/Safwanseban/voixme-project/internal/mockdata"
	"github.com/Safwanseban/voixme-project/internal/types"
	"github.com/Safwanseban/voixme-project/internal/usecases"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/valyala/fasthttp"

	"github.com/stretchr/testify/require"
)

func TestFetch(t *testing.T) {

	type args struct {
		query string
	}
	tests := []struct {
		name             string
		args             args
		expectedStruct   []types.OfferCompany
		assertFunc       func(*http.Response)
		wantCall         bool
		expectedHttpCode int
		expectedError    error
	}{

		{
			name: "error invalid query",
			args: args{
				query: "",
			},
			expectedStruct: []types.OfferCompany{
				{
					OfferID:     1,
					ClientID:    1,
					Country:     "india",
					Image:       "test_image",
					ImageWidth:  100,
					ImageHeight: 100,
				},
			},
			wantCall: false,
			assertFunc: func(w *http.Response) {
				require.Equal(t, http.StatusBadRequest, w.StatusCode)
				data, err := ioutil.ReadAll(w.Body)
				require.NoError(t, err)
				require.Equal(t, string(data), `{"message":"provide valid country parameter"}`)
			},
			expectedHttpCode: http.StatusBadRequest,
			expectedError:    errors.New("provide valid country parameter"),
		},
		{

			name: "error db",
			args: args{
				query: "?country=India",
			},
			wantCall: true,
			expectedStruct: []types.OfferCompany{
				{
					OfferID:     1,
					ClientID:    1,
					Country:     "india",
					Image:       "test_image",
					ImageWidth:  100,
					ImageHeight: 100,
				},
			},
			assertFunc: func(w *http.Response) {
				require.Equal(t, http.StatusInternalServerError, w.StatusCode)
				data, err := ioutil.ReadAll(w.Body)
				require.NoError(t, err)
				require.Equal(t, string(data), `{"message":"error fetching the records"}`)
			},
			expectedHttpCode: http.StatusOK,
			expectedError:    errors.New("error fetching the records"),
		},
		{

			name: "success",
			args: args{
				query: "?country=India",
			},
			wantCall: true,
			expectedStruct: []types.OfferCompany{
				{
					OfferID:     1,
					ClientID:    1,
					Country:     "india",
					Image:       "test_image",
					ImageWidth:  100,
					ImageHeight: 100,
				},
			},
			assertFunc: func(w *http.Response) {
				require.Equal(t, http.StatusOK, w.StatusCode)
				data, err := ioutil.ReadAll(w.Body)
				require.NoError(t, err)

				content, err := ioutil.ReadFile("testData/fetch_success.json")
				require.NoError(t, err)
				require.JSONEq(t, string(data), string(content))
			},
			expectedHttpCode: http.StatusOK,
			expectedError:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var w *http.Response
			if tt.wantCall {
				ctrl := gomock.NewController(t)
				defer ctrl.Finish()
				usecase := mock.NewMockUsecasesCompany(ctrl)
				usecase.EXPECT().ShowOfferCompany(gomock.Any()).
					Times(1).Return(tt.expectedStruct, tt.expectedError)

				w = testSetup(http.MethodGet, "/"+tt.args.query, usecase)
			} else {
				w = testSetup(http.MethodGet, "/"+tt.args.query, nil)
			}
			tt.assertFunc(w)
		})
	}
}

func testSetup(method string, url string,
	usecases usecases.UsecasesCompany,

) *http.Response {
	router := fiber.New()
	router.AcquireCtx(&fasthttp.RequestCtx{})
	NewServer(router, usecases)
	req := httptest.NewRequest(method, url, nil)
	resp, _ := router.Test(req)
	return resp

}
