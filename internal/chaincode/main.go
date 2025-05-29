package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Document representa um registro na blockchain
type Document struct {
	ID        string `json:"id"`
	Hash      string `json:"hash"`
	Owner     string `json:"owner"`
	Timestamp string `json:"timestamp"`
	Validated bool   `json:"validated"`
}

// SmartContract implementa o Chaincode
type SmartContract struct {
	contractapi.Contract
}

// Criar um novo documento
func (s *SmartContract) CreateDocument(ctx contractapi.TransactionContextInterface, id string, hash string, owner string, timestamp string) error {
	exists, err := s.DocumentExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("documento %s já existe", id)
	}

	document := Document{
		ID:        id,
		Hash:      hash,
		Owner:     owner,
		Timestamp: timestamp,
		Validated: false,
	}

	documentJSON, err := json.Marshal(document)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, documentJSON)
}

// Consultar um documento pelo ID
func (s *SmartContract) GetDocument(ctx contractapi.TransactionContextInterface, id string) (*Document, error) {
	documentJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("falha ao obter documento: %v", err)
	}
	if documentJSON == nil {
		return nil, fmt.Errorf("documento %s não encontrado", id)
	}

	var document Document
	err = json.Unmarshal(documentJSON, &document)
	if err != nil {
		return nil, err
	}

	return &document, nil
}

// Atualizar um documento (caso necessário)
func (s *SmartContract) UpdateDocument(ctx contractapi.TransactionContextInterface, id string, newOwner string) error {
	document, err := s.GetDocument(ctx, id)
	if err != nil {
		return err
	}

	document.Owner = newOwner
	documentJSON, err := json.Marshal(document)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, documentJSON)
}

// Validar um documento
func (s *SmartContract) ValidateDocument(ctx contractapi.TransactionContextInterface, id string) error {
	document, err := s.GetDocument(ctx, id)
	if err != nil {
		return err
	}

	document.Validated = true
	documentJSON, err := json.Marshal(document)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, documentJSON)
}

// Verificar se um documento já existe
func (s *SmartContract) DocumentExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	documentJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("falha ao verificar documento: %v", err)
	}

	return documentJSON != nil, nil
}

// Função principal
func main() {
	chaincode, err := contractapi.NewChaincode(new(SmartContract))
	if err != nil {
		fmt.Printf("Erro ao criar Chaincode: %v", err)
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Erro ao iniciar Chaincode: %v", err)
	}
}
