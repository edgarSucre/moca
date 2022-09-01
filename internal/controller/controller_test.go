package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/edgarSucre/moca/internal/controller"
	"github.com/edgarSucre/moca/internal/domain"
	mock "github.com/edgarSucre/moca/internal/mocks"
	"github.com/golang/mock/gomock"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestGetMortgagePayment(t *testing.T) {
	testcases := []struct {
		name      string
		input     any
		stub      func(*mock.MockUsecase, domain.Mortgage)
		checkResp func(*testing.T, httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			input: controller.MortgageRequest{
				PropertyValue:      100000,
				DownPayment:        20000,
				AnualInterestRate:  3.5,
				AmortizationPeriod: 30,
				PaymentSchedule:    12,
			},
			stub: func(mu *mock.MockUsecase, params domain.Mortgage) {
				mu.EXPECT().
					GetMortgagePayment(gomock.Eq(params)).
					Times(1).
					Return(200.00, nil)
			},
			checkResp: func(t *testing.T, r httptest.ResponseRecorder) {

				require.Equal(t, r.Code, http.StatusOK)

				resp := controller.MortgageResponse{}
				decoder := json.NewDecoder(r.Body)
				err := decoder.Decode(&resp)

				require.NoError(t, err)
				require.Equal(t, resp.Payment, 200.00)
			},
		},

		{
			name: "Invalid Downpayment",
			input: controller.MortgageRequest{
				PropertyValue:      100000,
				DownPayment:        200000,
				AnualInterestRate:  3.5,
				AmortizationPeriod: 30,
				PaymentSchedule:    12,
			},
			stub: func(mu *mock.MockUsecase, params domain.Mortgage) {
				mu.EXPECT().
					GetMortgagePayment(gomock.Eq(params)).
					Times(1).
					Return(0.00, domain.ErrInvalidDownPayment)
			},
			checkResp: func(t *testing.T, r httptest.ResponseRecorder) {

				require.Equal(t, r.Code, http.StatusBadRequest)

				resp := controller.ErrorResponse{}
				decoder := json.NewDecoder(r.Body)
				err := decoder.Decode(&resp)

				require.NoError(t, err)
				require.Equal(t, resp.Error, "invalid down payment")
			},
		},

		{
			name:  "Bad Request",
			input: "this should fail",
			stub: func(mu *mock.MockUsecase, params domain.Mortgage) {
				mu.EXPECT().
					GetMortgagePayment(gomock.Eq(params)).
					Times(0)
			},
			checkResp: func(t *testing.T, r httptest.ResponseRecorder) {

				require.Equal(t, r.Code, http.StatusBadRequest)

				resp := controller.ErrorResponse{}
				decoder := json.NewDecoder(r.Body)
				err := decoder.Decode(&resp)

				require.NoError(t, err)
				require.True(t, strings.Contains(resp.Error, "cannot unmarshal"))
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			uc := mock.NewMockUsecase(ctrl)

			input, ok := tc.input.(controller.MortgageRequest)
			if ok {
				tc.stub(uc, input.ToDomain())
			} else {
				tc.stub(uc, domain.Mortgage{})
			}

			body := bytes.NewBuffer([]byte{})
			encoder := json.NewEncoder(body)
			err := encoder.Encode(tc.input)
			require.NoError(t, err)

			logger := log.WithFields(log.Fields{
				"layer": "test",
			})

			handler := controller.New(uc, logger)
			request := httptest.NewRequest(http.MethodPost, "/api/payment", body)
			recorder := httptest.NewRecorder()

			handler.ServeHTTP(recorder, request)
			tc.checkResp(t, *recorder)
		})
	}

}

func TestSpaHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock.NewMockUsecase(ctrl)
	logger := log.WithFields(log.Fields{
		"layer": "test",
	})

	handler := controller.New(uc, logger)
	request := httptest.NewRequest(http.MethodGet, "/index.html", nil)
	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, request)

	//server does not start properly from test env
	require.Equal(t, 301, recorder.Code)
}
