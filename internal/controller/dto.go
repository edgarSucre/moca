package controller

import "github.com/edgarSucre/moca/internal/domain"

type (
	MortgageRequest struct {
		PropertyValue      float64 `json:"propertyValue"`
		DownPayment        float64 `json:"downPayment"`
		AnualInterestRate  float64 `json:"anualInterestRate"`
		AmortizationPeriod int16   `json:"amortizationPeriod"`
		PaymentSchedule    int16   `json:"paymentSchedule"`
	}

	MortgageResponse struct {
		Payment float64 `json:"payment"`
	}

	ErrorResponse struct {
		Error string `json:"error"`
	}
)

func (mr MortgageRequest) ToDomain() domain.Mortgage {
	return domain.Mortgage{
		PropertyValue:      mr.PropertyValue,
		DownPayment:        mr.DownPayment,
		AnualInterestRate:  mr.AnualInterestRate,
		AmortizationPeriod: mr.AmortizationPeriod,
		PaymentSchedule:    mr.PaymentSchedule,
	}
}
