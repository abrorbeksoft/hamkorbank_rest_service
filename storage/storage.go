package storage

import (
	"context"
	"rest_service/models"
)

type StorageI interface {
	CloseDB()
	Phone() PhoneRepoI
}

type PhoneRepoI interface {
	Get(ctx context.Context, req *models.Phone) (*models.Phone, bool, error)
	GetAll(ctx context.Context, req *models.Phone) ([]models.Phone, error)
}
