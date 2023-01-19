package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"rest_service/models"
	"rest_service/storage"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) storage.PhoneRepoI {
	return &userRepo{
		db: db,
	}
}

func (u userRepo) Get(ctx context.Context, req *models.Phone) (*models.Phone, bool, error) {
	var resp = models.Phone{}
	row := u.db.QueryRow(ctx, "select id, phone, created_at::varchar from phones where id = $1", req.ID)

	err := row.Scan(&resp.ID, &resp.Phone, &resp.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, true, err
	}

	if err != nil {
		return nil, false, err
	}

	return &resp, false, nil
}

func (u userRepo) GetAll(ctx context.Context, req *models.Phone) ([]models.Phone, error) {
	resp := make([]models.Phone, 0)

	rows, err := u.db.Query(ctx, "select id, phone, created_at::varchar from phones")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		temp := models.Phone{}
		err = rows.Scan(&temp.ID, &temp.Phone, &temp.CreatedAt)
		fmt.Println(temp)
		if err != nil {
			return nil, err
		}
		resp = append(resp, temp)
	}

	return resp, nil
}
