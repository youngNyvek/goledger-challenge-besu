package service

import (
	"context"
	"fmt"
	simplestorage "goledger-challenge-besu/abi"
	"goledger-challenge-besu/internal/repository/model"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ContractLogRepository interface {
	Save(*model.ContractLog) error
	FindLatest() (*model.ContractLog, error)
}

type ContractService struct {
	contract *simplestorage.SimpleStorage
	client   *ethclient.Client
	signer   *bind.TransactOpts
	repo     ContractLogRepository
}

func NewContractService(contract *simplestorage.SimpleStorage, client *ethclient.Client, signer *bind.TransactOpts, repo ContractLogRepository) *ContractService {
	return &ContractService{contract: contract, client: client, signer: signer, repo: repo}
}

func (s *ContractService) SetValue(ctx context.Context, value int64) (txHash string, blockNumber string, err error) {
	tx, err := s.contract.Set(s.signer, big.NewInt(value))
	if err != nil { return "", "", err }

	receipt, err := bind.WaitMined(ctx, s.client, tx)
	if err != nil { return "", "", err }

	return fmt.Sprintf("0x%x", tx.Hash()), receipt.BlockNumber.String(), nil
}

func (s *ContractService) GetValue(ctx context.Context) (string, error) {
	val, err := s.contract.Get(&bind.CallOpts{Context: ctx})
	if err != nil { return "", err }
	return val.String(), nil
}

func (s *ContractService) Sync(ctx context.Context) (string, error) {
	val, err := s.GetValue(ctx)
	if err != nil { return "", err }
	if err := s.repo.Save(&model.ContractLog{Value: val}); err != nil { return "", err }
	return val, nil
}

func (s *ContractService) Check(ctx context.Context) (bool, error) {
	onChain, err := s.GetValue(ctx)
	if err != nil { return false, err }

	log, err := s.repo.FindLatest()
	if err != nil { return false, err }

	return onChain == log.Value, err
}