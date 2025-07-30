package ethereum

import (
	"context"
	simplestorage "goledger-challenge-besu/abi"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func MustConnectToNode(url string) *ethclient.Client {
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatalf("Erro ao conectar no nó: %v", err)
	}

	return client
}

func MustCreateContract(address string, client *ethclient.Client) *simplestorage.SimpleStorage {
	contract, err := simplestorage.NewSimpleStorage(common.HexToAddress(address), client)
	if err != nil {
		log.Fatalf("Erro ao instanciar contrato: %v", err)
	}

	return contract
}

func MustCreateSigner(hexKey string, client *ethclient.Client) *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA(hexKey)
	if err != nil {
		log.Fatalf("Erro ao carregar chave privada: %v", err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Erro ao obter ChainID: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Erro ao criar transactor: %v", err)
	}

	return auth
}
