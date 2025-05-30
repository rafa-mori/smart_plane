package contracts

import "github.com/hyperledger/fabric-contract-api-go/contractapi"

type IBaseContract[T any] interface {
	Put(ctx contractapi.TransactionContextInterface, id string, data T) error
	Get(ctx contractapi.TransactionContextInterface, id string) (T, error)
	Delete(ctx contractapi.TransactionContextInterface, id string) error
	Exists(ctx contractapi.TransactionContextInterface, id string) (bool, error)
	History(ctx contractapi.TransactionContextInterface, id string) ([]T, error)
}
