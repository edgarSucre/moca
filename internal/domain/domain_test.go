package domain_test

import (
	"errors"
	"testing"

	"github.com/edgarSucre/moca/internal/domain"
)

func TestValidateAmortizationPeriod(t *testing.T) {
	testcases := []struct {
		name          string
		input         domain.Mortgage
		checkResponse func(*testing.T, error)
	}{
		{
			name: "No Insurance",
			input: domain.Mortgage{
				PropertyValue:      100000,
				DownPayment:        20000,
				AnualInterestRate:  3.5,
				AmortizationPeriod: 30,
				PaymentSchedule:    12,
			},
			checkResponse: func(t *testing.T, err error) {
				if err != nil {
					t.Errorf("Test Failed, Expected: nil, Got: %v", err)
				}
			},
		},

		{
			name: "With Insurance - Invalid Downpayment",
			input: domain.Mortgage{
				PropertyValue:      100000,
				DownPayment:        -20000,
				AnualInterestRate:  3.5,
				AmortizationPeriod: 30,
				PaymentSchedule:    12,
			},
			checkResponse: func(t *testing.T, err error) {
				if !errors.Is(err, domain.ErrInvalidDownPayment) {
					t.Errorf("Test Failed, Expected: domain.ErrInvalidDownPayment, Got: %v", err)
				}
			},
		},

		{
			name: "With Insurance - Invalid Amortization",
			input: domain.Mortgage{
				PropertyValue:      100000,
				DownPayment:        10000,
				AnualInterestRate:  3.5,
				AmortizationPeriod: 30,
				PaymentSchedule:    12,
			},
			checkResponse: func(t *testing.T, err error) {
				if !errors.Is(err, domain.ErrInvalidAmortizationPeriod) {
					t.Errorf("Test Failed, Expected: domain.ErrInvalidAmortizationPeriod, Got: %v", err)
				}
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.input.ValidateAmortizationPeriod()
			tc.checkResponse(t, err)
		})
	}
}

func TestValidateDownPayment(t *testing.T) {
	testcases := []struct {
		name          string
		input         domain.Mortgage
		checkResponse func(*testing.T, error)
	}{
		{
			name: "OK - 5 %",
			input: domain.Mortgage{
				PropertyValue:      100000,
				DownPayment:        5000,
				AnualInterestRate:  3.5,
				AmortizationPeriod: 30,
				PaymentSchedule:    12,
			},
			checkResponse: func(t *testing.T, err error) {
				if err != nil {
					t.Errorf("Test Failed, Expected: nil, Got: %v", err)
				}
			},
		},

		{
			name: "OK - 10 %",
			input: domain.Mortgage{
				PropertyValue:      600000,
				DownPayment:        60000,
				AnualInterestRate:  3.5,
				AmortizationPeriod: 30,
				PaymentSchedule:    12,
			},
			checkResponse: func(t *testing.T, err error) {
				if err != nil {
					t.Errorf("Test Failed, Expected: nil, Got: %v", err)
				}
			},
		},

		{
			name: "OK - 20 %",
			input: domain.Mortgage{
				PropertyValue:      1000000,
				DownPayment:        200000,
				AnualInterestRate:  3.5,
				AmortizationPeriod: 30,
				PaymentSchedule:    12,
			},
			checkResponse: func(t *testing.T, err error) {
				if err != nil {
					t.Errorf("Test Failed, Expected: nil, Got: %v", err)
				}
			},
		},

		{
			name: "Low Downpayment %",
			input: domain.Mortgage{
				PropertyValue:      1000000,
				DownPayment:        20000,
				AnualInterestRate:  3.5,
				AmortizationPeriod: 30,
				PaymentSchedule:    12,
			},
			checkResponse: func(t *testing.T, err error) {
				if !errors.Is(err, domain.ErrInvalidDownPayment) {
					t.Errorf("Test Failed, Expected: domain.ErrInvalidDownPayment, Got: %v", err)
				}
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.input.ValidateDownPayment()
			tc.checkResponse(t, err)
		})
	}
}
