package usecases

import (
	"github.com/keyinvoker/go-payout-service/internal/domain/models"
	"github.com/keyinvoker/go-payout-service/internal/domain/repositories"
)

type InquirePayoutUsecase struct {
	payoutRepo repositories.BaseDatabaseRepository[*models.Payout]
}
