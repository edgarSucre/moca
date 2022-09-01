package usecase

import (
	"fmt"

	"github.com/edgarSucre/moca/internal/domain"
	"github.com/go-playground/validator/v10"
)

type MortgageUseCase struct {
	validator *validator.Validate
}

func New() *MortgageUseCase {
	return &MortgageUseCase{
		validator: validator.New(),
	}
}

func (m *MortgageUseCase) GetMortgagePayment(params domain.Mortgage) (float64, error) {
	if err := m.validateMortgage(params); err != nil {
		return 0, err
	}

	payment, err := params.GetPayment()
	if err != nil {
		return 0, err
	}

	return payment, nil
}

func (m *MortgageUseCase) validateMortgage(params domain.Mortgage) error {
	if err := m.validator.Struct(params); err != nil {
		if orig, ok := err.(validator.ValidationErrors); ok {
			return fmt.Errorf("%s", getMessage(orig))
		}

		return err
	}

	if err := params.ValidateDownPayment(); err != nil {
		return err
	}

	if err := params.ValidateAmortizationPeriod(); err != nil {
		return err
	}

	return nil
}
