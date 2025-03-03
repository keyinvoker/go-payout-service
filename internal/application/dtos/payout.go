package dtos

type CreatePayoutRequest struct {
	LoanID      int     `json:"loan_id" binding:"required"`
	UserID      int     `json:"user_id" binding:"required"`
	Principal   float64 `json:"principal" binding:"required" default:"0"`
	Interest    float64 `json:"interest" binding:"required" default:"0"`
	Fine        float64 `json:"fine" default:"0"`
	Description string  `json:"description"`
}
