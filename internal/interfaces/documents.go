package interfaces

import "github.com/hyperledger/fabric-contract-api-go/contractapi"

type IDocumentBase interface {
	CreateDocument(ctx contractapi.TransactionContextInterface, id string, hash string, owner string, timestamp string) error
	GetDocument(ctx contractapi.TransactionContextInterface, id string) (IDocumentBase, error)
	UpdateDocument(ctx contractapi.TransactionContextInterface, id string, newOwner string) error
	DeleteDocument(ctx contractapi.TransactionContextInterface, id string) error
	ValidateDocument(ctx contractapi.TransactionContextInterface, id string) error
	DocumentExists(ctx contractapi.TransactionContextInterface, id string) (bool, error)
}
