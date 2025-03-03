package usecases

import (
	"context"

	"github.com/keyinvoker/go-payout-service/internal/application/dtos"
	"github.com/keyinvoker/go-payout-service/internal/application/services"
	"github.com/keyinvoker/go-payout-service/internal/domain/models"
)

type CreatePayoutUsecase struct {
	payoutService *services.PayoutService
}

func NewCreatePayoutUsecase(payoutService *services.PayoutService) *CreatePayoutUsecase {
	return &CreatePayoutUsecase{payoutService: payoutService}
}

func (u *CreatePayoutUsecase) Execute(
	ctx context.Context,
	data *dtos.CreatePayoutRequest,
) (*models.Payout, error) {

	payout := &models.Payout{
		LoanID:    data.LoanID,
		UserID:    data.UserID,
		Principal: data.Principal,
		Interest:  data.Interest,
		Fine:      data.Fine,
	}

	payout, err := u.payoutService.Create(ctx, payout)
	if err != nil {
		return nil, err
	}

	return payout, nil
}
