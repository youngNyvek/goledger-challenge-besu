package repository

import (
	"database/sql"
	"goledger-challenge-besu/internal/repository/model"

	_ "github.com/lib/pq"
)

type PostgresContractLogRepository struct { db *sql.DB }

func NewPostgresContractLogRepository(db *sql.DB) *PostgresContractLogRepository {
	return &PostgresContractLogRepository{db: db}
}

func (r *PostgresContractLogRepository) Save(contractLog *model.ContractLog) error {
	_, err := r.db.Exec(
		"INSERT INTO contract_logs (value) VALUES ($1)",
		contractLog.Value,
	)
	return err
}

func (r *PostgresContractLogRepository) FindLatest() (*model.ContractLog, error) {
	row := r.db.QueryRow(`SELECT id, value, created_at FROM contract_logs ORDER BY id DESC LIMIT 1`)
	var log model.ContractLog
	err := row.Scan(&log.ID, &log.Value, &log.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &log, nil
}