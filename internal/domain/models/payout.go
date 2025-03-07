package models

import (
	"time"

	"github.com/keyinvoker/go-payout-service/internal/domain/constants"
)

type Payout struct {
	BaseModel
	LoanID       int                    `gorm:"not null;index:payouts_loan_id_key"`
	UserID       int                    `gorm:"not null;index:payouts_user_id_key"`
	PayoutStatus constants.PayoutStatus `gorm:"not null;type:payout_status;default:PENDING"`
	PayoutDate   *time.Time
	Total        float64 `gorm:"not null;default:0"`
	Principal    float64 `gorm:"not null;default:0"`
	Interest     float64 `gorm:"not null;default:0"`
	Fine         float64 `gorm:"not null;default:0"`
	Description  *string `gorm:"type:text"`
}
