package models

import (
	"time"

	"github.com/keyinvoker/go-payout-service/internal/domain/constants"
)

type Payout struct {
	BaseModel
	TransactionCode       string                 `gorm:"not null;uniqueIndex:unique_transaction_code"`
	PlpID                 int                    `gorm:"not null;index:payouts_plp_id_key"`
	LoanID                int                    `gorm:"not null;index:payouts_loan_id_key"`
	UserID                int                    `gorm:"not null;index:payouts_user_id_key"`
	PayoutStatus          constants.PayoutStatus `gorm:"not null;default:pending;type:payout_status"`
	PayoutDate            *time.Time
	PayoutAmount          float64 `gorm:"not null;default:0"`
	PrincipalAmount       float64 `gorm:"not null;default:0"`
	InterestAmount        float64 `gorm:"not null;default:0"`
	InterestTaxPercentage float64 `gorm:"not null;default:0"`
	InterestTaxAmount     float64 `gorm:"not null;default:0"`
	FineAmount            float64 `gorm:"not null;default:0"`
	FineTaxPercentage     float64 `gorm:"not null;default:0"`
	FineTaxAmount         float64 `gorm:"not null;default:0"`
	FailureReason         *string `gorm:"type:text"`
}
