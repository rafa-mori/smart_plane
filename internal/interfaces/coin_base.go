package interfaces

import "github.com/hyperledger/fabric-contract-api-go/contractapi"

type ICoinBase interface {
	PutCoinBase(ctx contractapi.TransactionContextInterface, id string, amount int, from string, to string) error
	GetCoinBase(ctx contractapi.TransactionContextInterface, id string) (ICoinBase, error)
	DeleteCoinBase(ctx contractapi.TransactionContextInterface, id string) error
	CoinBaseExists(ctx contractapi.TransactionContextInterface, id string) (bool, error)
	HistoryCoinBase(ctx contractapi.TransactionContextInterface, id string) ([]ICoinBase, error)
}
