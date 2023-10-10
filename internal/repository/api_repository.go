// internal/repository/api_repository.go
package repository

import (
	"context"

	"ddvinyaninov/assets-tracker-api/internal/domain"

	"github.com/jmoiron/sqlx"
)

type apiRepository struct {
	db *sqlx.DB
}

func NewApiRepository(db *sqlx.DB) domain.ApiRepository {
	return &apiRepository{db: db}
}

func (r apiRepository) Select(ctx context.Context) ([]*domain.Api, error) {
	assets := []*domain.Api{}
	if err := r.db.Select(&assets, `SELECT 
				'Feature' as type,
				json_build_object(
					'id', a.id,
					'name', a.name,
					'description', a.description,
					'status', a.status,
					'health', a.health
				) as properties,
				json_build_object(
					'type', 'Point',
					'coordinates', json_build_array(
						g.long+offsetx*extract(second from current_timestamp),
						g.lat+offsety*extract(second from current_timestamp))
				) as geometry
				FROM asset a 
				INNER JOIN geo g ON a.id = g.aid`); err != nil {
		return nil, err
	}

	return assets, nil
}

func (r apiRepository) Insert(ctx context.Context, body string) (*domain.Api, error) {
	lastInsertId := 0
	if err := r.db.
		QueryRow("INSERT INTO asset(name) VALUES($1) RETURNING id", body).
		Scan(&lastInsertId); err != nil {
		return nil, err
	}

	var api domain.Api
	if err := r.db.Get(&api, "SELECT * FROM asset WHERE id = $1", lastInsertId); err != nil {
		return nil, err
	}

	return &api, nil
}
