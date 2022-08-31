package usecase_test

import (
	"errors"
	"testing"

	"github.com/edgarSucre/moca/domain"
	"github.com/edgarSucre/moca/usecase"
	"github.com/stretchr/testify/require"
)

func TestGetPayment(t *testing.T) {
	tescases := []struct {
		name        string
		input       domain.Mortgage
		checkResult func(*testing.T, float64, error)
	}{
		{
			name: "OK",
			input: domain.Mortgage{
				PropertyValue:      100000,
				DownPayment:        20000,
				AnualInterestRate:  3.5,
				AmortizationPeriod: 30,
				PaymentSchedule:    12,
			},
			checkResult: func(t *testing.T, resp float64, err error) {
				require.NoError(t, err)
				require.Equal(t, 359.24, resp)
			},
		},

		{
			name:  "Validation Error - Required",
			input: domain.Mortgage{},
			checkResult: func(t *testing.T, resp float64, err error) {
				require.ErrorContains(t, err, "PropertyValue is required")
				require.ErrorContains(t, err, "DownPayment is required")
				require.ErrorContains(t, err, "AnualInterestRate is required")
				require.ErrorContains(t, err, "AmortizationPeriod is required")
				require.ErrorContains(t, err, "PaymentSchedule is required")
				require.Zero(t, resp)
			},
		},

		{
			name: "Validation Error - GT",
			input: domain.Mortgage{
				PropertyValue:      -1,
				DownPayment:        -1,
				AnualInterestRate:  -1,
				AmortizationPeriod: 30,
				PaymentSchedule:    12,
			},
			checkResult: func(t *testing.T, resp float64, err error) {
				require.ErrorContains(t, err, "PropertyValue must be greater than 0")
				require.ErrorContains(t, err, "DownPayment must be greater than 0")
				require.ErrorContains(t, err, "AnualInterestRate must be greater than 0")
				require.Zero(t, resp)
			},
		},

		{
			name: "Validation Error - Downpayment",
			input: domain.Mortgage{
				PropertyValue:      100000,
				DownPayment:        100001,
				AnualInterestRate:  3.5,
				AmortizationPeriod: 30,
				PaymentSchedule:    12,
			},
			checkResult: func(t *testing.T, resp float64, err error) {
				require.ErrorContains(t, err, "DownPayment must be smaller than PropertyValue")
				require.Zero(t, resp)
			},
		},

		{
			name: "Validation Error - Invalid Options",
			input: domain.Mortgage{
				PropertyValue:      100000,
				DownPayment:        100001,
				AnualInterestRate:  3.5,
				AmortizationPeriod: 26,
				PaymentSchedule:    30,
			},
			checkResult: func(t *testing.T, resp float64, err error) {
				require.ErrorContains(t, err, "AmortizationPeriod must be one of 5 10 15 20 25 30")
				require.ErrorContains(t, err, "PaymentSchedule must be one of 12 24 26")
				require.Zero(t, resp)
			},
		},

		{
			name: "Invalid Downpayment",
			input: domain.Mortgage{
				PropertyValue:      100000,
				DownPayment:        10,
				AnualInterestRate:  3.5,
				AmortizationPeriod: 30,
				PaymentSchedule:    12,
			},
			checkResult: func(t *testing.T, resp float64, err error) {
				if !errors.Is(err, domain.ErrInvalidDownPayment) {
					t.Errorf("Test Failed, Expected: domain.ErrInvalidDownPayment, Got: %v", err)
				}
				require.Zero(t, resp)
			},
		},

		{
			name: "Invalid Amortization",
			input: domain.Mortgage{
				PropertyValue:      100000,
				DownPayment:        10000,
				AnualInterestRate:  3.5,
				AmortizationPeriod: 30,
				PaymentSchedule:    12,
			},
			checkResult: func(t *testing.T, resp float64, err error) {
				if !errors.Is(err, domain.ErrInvalidAmortizationPeriod) {
					t.Errorf("Test Failed, Expected: domain.ErrInvalidAmortizationPeriod, Got: %v", err)
				}
				require.Zero(t, resp)
			},
		},
	}

	for _, tc := range tescases {
		t.Run(tc.name, func(t *testing.T) {
			uc := usecase.New()
			resp, err := uc.GetMortgagePayment(tc.input)
			tc.checkResult(t, resp, err)
		})
	}
}
