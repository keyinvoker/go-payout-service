package models

import (
	"time"

	"github.com/keyinvoker/go-payout-service/internal/domain/constants"
	"gorm.io/gorm"
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

func (Payout) BeforeMigrate(db interface{}) error {
	if db, ok := db.(*gorm.DB); ok {
		return db.Exec(`
			DO $$ BEGIN
				CREATE TYPE
					payout_status AS ENUM (
						'PENDING',
						'CALCULATION_FAILED',
						'READY_TO_PAYOUT',
						'ON_PROCESS',
						'PAYOUT_FAILED',
						'PAID_OUT'
					);
			EXCEPTION
				WHEN duplicate_object THEN null;
			END $$;
		`).Error
	}
	return nil
}
