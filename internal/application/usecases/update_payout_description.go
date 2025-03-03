package usecases

import (
	"context"

	"github.com/keyinvoker/go-payout-service/internal/application/dtos"
	"github.com/keyinvoker/go-payout-service/internal/application/services"
	"github.com/keyinvoker/go-payout-service/internal/domain/models"
)

type UpdatePayoutDescriptionUsecase struct {
	payoutService *services.PayoutService
}

func NewUpdatePayoutDescriptionUsecase(payoutService *services.PayoutService) *UpdatePayoutDescriptionUsecase {
	return &UpdatePayoutDescriptionUsecase{payoutService: payoutService}
}

func (u *UpdatePayoutDescriptionUsecase) Execute(
	ctx context.Context,
	data *dtos.UpdatePayoutDescriptionRequest,
) (*models.Payout, error) {

	payout, err := u.payoutService.UpdateDescription(ctx, data.ID, data.Description)
	if err != nil {
		return nil, err
	}

	return payout, nil
}
