package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HealthHandler struct {
	db *gorm.DB
}

func NewHealthHandler(db *gorm.DB) *HealthHandler {
	return &HealthHandler{db: db}
}

func (h *HealthHandler) CheckHealth(ctx *gin.Context) {
	sqlDB, err := h.db.DB()
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "Failed to get database connection"},
		)
		return
	}

	err = sqlDB.Ping()
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "Failed to ping database"},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{"status": "OK", "message": "Database connection is healthy!"},
	)
}
