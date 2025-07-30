package model

import "time"

type ContractLog struct {
	ID        int       `db:"id"`
	Value     string    `db:"value"`
	CreatedAt time.Time `db:"created_at"`
}
