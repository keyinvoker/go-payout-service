package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/keyinvoker/go-payout-service/internal/domain/constants"
	"github.com/keyinvoker/go-payout-service/internal/domain/exceptions"
	"github.com/keyinvoker/go-payout-service/internal/domain/models"

	"gorm.io/gorm"
)

type IPayoutRepository interface {
	Create(ctx context.Context, entity *models.Payout) error
	Update(ctx context.Context, entity *models.Payout) error
	Delete(ctx context.Context, id int) error
	Count(ctx context.Context) (int64, error)
	GetByID(ctx context.Context, id int) (*models.Payout, error)
	GetAll(ctx context.Context, filters map[string]interface{}) ([]*models.Payout, error)
	GetPayoutsByStatus(ctx context.Context, status constants.PayoutStatus) ([]*models.Payout, error)
}

type PayoutRepository struct {
	db *gorm.DB
}

func NewPayoutRepository(db *gorm.DB) *PayoutRepository {
	return &PayoutRepository{db: db}
}

func (r *PayoutRepository) Create(ctx context.Context, entity *models.Payout) error {
	result := r.db.WithContext(ctx).Create(entity)
	if result.Error != nil {
		return fmt.Errorf("failed to create payout: %w", exceptions.ErrDatabase{Op: "create", Err: result.Error, Resource: "payout"})
	}
	return nil
}

func (r *PayoutRepository) Update(ctx context.Context, entity *models.Payout) error {
	result := r.db.WithContext(ctx).Save(entity)
	if result.Error != nil {
		return fmt.Errorf("failed to update payout: %w", exceptions.ErrDatabase{Op: "update", Err: result.Error, Resource: "payout"})
	}
	return nil
}

func (r *PayoutRepository) Delete(ctx context.Context, id int) error {
	var payout models.Payout
	result := r.db.WithContext(ctx).First(&payout, id)
	if result.Error != nil {
		return fmt.Errorf("failed to find payout for deletion: %w", exceptions.ErrDatabase{Op: "delete", Err: result.Error, Resource: "payout"})
	}

	result = r.db.WithContext(ctx).Delete(&payout)
	if result.Error != nil {
		return fmt.Errorf("failed to delete payout: %w", exceptions.ErrDatabase{Op: "delete", Err: result.Error, Resource: "payout"})
	}
	return nil
}

func (r *PayoutRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	result := r.db.WithContext(ctx).Model(&models.Payout{}).Count(&count)
	if result.Error != nil {
		return 0, fmt.Errorf("failed to count payouts: %w", exceptions.ErrDatabase{Op: "count", Err: result.Error, Resource: "payout"})
	}
	return count, nil
}

func (r *PayoutRepository) GetByID(ctx context.Context, id int) (*models.Payout, error) {
	var payout models.Payout
	result := r.db.WithContext(ctx).First(&payout, "id = ?", id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, exceptions.ErrNotFound{Resource: "payout", ID: id}
		}
		return nil, exceptions.ErrDatabase{Op: "get", Err: result.Error, Resource: "payout"}
	}

	return &payout, nil
}

func (r *PayoutRepository) GetAll(ctx context.Context, filters map[string]interface{}) ([]*models.Payout, error) {
	var payouts []*models.Payout
	query := r.db.WithContext(ctx).Model(&models.Payout{})

	if filters != nil {
		query = query.Where(filters)
	}

	result := query.Find(&payouts)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to get payout list: %w", exceptions.ErrDatabase{Op: "get", Err: result.Error, Resource: "payout"})
	}
	return payouts, nil
}

func (r *PayoutRepository) GetPayoutsByStatus(ctx context.Context, status constants.PayoutStatus) ([]*models.Payout, error) {
	var payouts []*models.Payout
	result := r.db.WithContext(ctx).Where("payout_status = ?", status).Find(&payouts)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to get payouts by status: %w", exceptions.ErrDatabase{Op: "get", Err: result.Error, Resource: "payout"})
	}
	return payouts, nil
}
