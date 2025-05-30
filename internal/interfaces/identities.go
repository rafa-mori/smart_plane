package interfaces

import "github.com/hyperledger/fabric-contract-api-go/contractapi"

type IIdentityBase interface {
	CreateUser(ctx contractapi.TransactionContextInterface, id string, role string) error
	GetUser(ctx contractapi.TransactionContextInterface, id string) (IIdentityBase, error)
	ValidateUser(ctx contractapi.TransactionContextInterface, id string) error
	UserExists(ctx contractapi.TransactionContextInterface, id string) (bool, error)
}
