// internal/domain/api.go
package domain

import (
	"context"
	"database/sql/driver"
	"encoding/json"
)

type ApiID int

type Api struct {
	Type       string     `db:"type" json:"type"`
	Properties Properties `db:"properties" json:"properties"`
	Geometry   Geometry   `db:"geometry" json:"geometry"`
}

type Properties struct {
	Id          int    `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
	Status      string `db:"status" json:"status"`
	Health      int    `db:"health" json:"health"`
}

func (a *Properties) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		//return Properties.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}

type Geometry struct {
	Type        string    `db:"type" json:"type"`
	Coordinates []float64 `db:"coordinates" json:"coordinates"`
}

func (a Geometry) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *Geometry) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		//return Geometry.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

type ApiRepository interface {
	Select(ctx context.Context) ([]*Api, error)
	Insert(ctx context.Context, body string) (*Api, error)
}

type ApiUsecase interface {
	List(ctx context.Context) ([]*Api, error)
	Create(ctx context.Context, body string) (*Api, error)
}
