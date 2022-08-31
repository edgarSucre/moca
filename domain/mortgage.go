package domain

import (
	"fmt"
	"math"
)

type (

	// validation rules are business rules
	Mortgage struct {
		PropertyValue      float64 `validate:"required,gt=0"`
		DownPayment        float64 `validate:"required,gt=0,ltecsfield=PropertyValue"`
		AnualInterestRate  float64 `validate:"required,gt=0,lt=100"`
		AmortizationPeriod int16   `validate:"required,oneof=5 10 15 20 25 30"`
		PaymentSchedule    int16   `validate:"required,oneof=12 24 26"`
	}
)

var (
	// Minimum down payment based on property value
	minimumDownPaymentPercentage map[bracketRange]float64

	// default insurance rate based on down payment
	insuranceRateIndex map[bracketRange]float64

	ErrInvalidDownPayment        = fmt.Errorf("invalid down payment")
	ErrInvalidAmortizationPeriod = fmt.Errorf("invalid amortization period")
)

// ValidateAmortizationPeriod checks the minimun downPayment against the AmortizationPeriod
func (m Mortgage) ValidateAmortizationPeriod() error {
	i, err := m.requiredInsurance()
	if err != nil {
		return err
	}

	if i > 0 && m.AmortizationPeriod > 25 {
		return ErrInvalidAmortizationPeriod
	}
	return nil
}

// ValidateDownPayment checks the minimun downPayment against the property value
func (m Mortgage) ValidateDownPayment() error {
	for bracked, minimun := range minimumDownPaymentPercentage {
		if bracked.withinRange(m.downPaymentPercentage()) && m.downPaymentPercentage() >= minimun {
			return nil
		}
	}
	return ErrInvalidDownPayment
}

// downPaymentPercentage returns what percetage of the property value the downpayment represents
func (m Mortgage) downPaymentPercentage() float64 {
	raw := m.DownPayment / m.PropertyValue * 100
	return math.Ceil(raw*100) / 100
}

// GetPayment returns the payment per payment schedule
func (m Mortgage) GetPayment() (float64, error) {
	p, err := m.principalAfterIssurance()
	if err != nil {
		return 0, err
	}

	subExp := math.Pow((1 + m.paymentInterestRate()), float64(m.numberOfPayments()))
	raw := (p * m.paymentInterestRate() * subExp) / (subExp - 1)
	return math.Ceil(raw*100) / 100, nil
}

func (m Mortgage) principalAfterIssurance() (float64, error) {
	insurance, err := m.requiredInsurance()
	if err != nil {
		return 0, err
	}
	return m.PropertyValue - m.DownPayment + insurance, nil
}

func (m Mortgage) principalBeforeIssurance() float64 {
	return m.PropertyValue - m.DownPayment
}

// numberOfPayments returns the total number of paymnets
func (m Mortgage) numberOfPayments() int16 {
	return m.AmortizationPeriod * m.PaymentSchedule
}

// requiredInsurance returns the insurance obligation for the mortgage
func (m Mortgage) requiredInsurance() (float64, error) {
	for bracket, rate := range insuranceRateIndex {
		if bracket.withinRange(m.downPaymentPercentage()) {
			return rate * m.principalBeforeIssurance(), nil
		}
	}

	return 0, ErrInvalidDownPayment
}

// paymentInterestRate its the interest rate per payment in a year
func (m Mortgage) paymentInterestRate() float64 {
	return m.AnualInterestRate / 100 / float64(m.PaymentSchedule)
}
