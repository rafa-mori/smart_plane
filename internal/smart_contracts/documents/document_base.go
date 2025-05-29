package documents

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Estrutura do documento registrado na blockchain
type DocumentBase struct {
	contractapi.Contract

	ID        string `json:"id"`
	Hash      string `json:"hash"`
	Owner     string `json:"owner"`
	Timestamp string `json:"timestamp"`
	Validated bool   `json:"validated"`
}

// Criar um novo documento na blockchain
func (dc *DocumentBase) CreateDocument(ctx contractapi.TransactionContextInterface, id string, hash string, owner string, timestamp string) error {
	exists, err := dc.DocumentExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("documento %s já existe", id)
	}

	document := DocumentBase{
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

// Consultar um documento na blockchain
func (dc *DocumentBase) GetDocument(ctx contractapi.TransactionContextInterface, id string) (*DocumentBase, error) {
	documentJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("falha ao obter documento: %v", err)
	}
	if documentJSON == nil {
		return nil, fmt.Errorf("documento %s não encontrado", id)
	}

	var document DocumentBase
	err = json.Unmarshal(documentJSON, &document)
	if err != nil {
		return nil, err
	}

	return &document, nil
}

// Atualizar um documento (caso necessário)
func (dc *DocumentBase) UpdateDocument(ctx contractapi.TransactionContextInterface, id string, newOwner string) error {
	document, err := dc.GetDocument(ctx, id)
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

// Deletar um documento do state, não do histórico
func (dc *DocumentBase) DeleteDocument(ctx contractapi.TransactionContextInterface, id string) error {
	exists, err := dc.DocumentExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("documento %s não encontrado", id)
	}

	return ctx.GetStub().DelState(id)
}

// Validar um documento registrado
func (dc *DocumentBase) ValidateDocument(ctx contractapi.TransactionContextInterface, id string) error {
	document, err := dc.GetDocument(ctx, id)
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

// Verifica se o documento já existe
func (dc *DocumentBase) DocumentExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	documentJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("falha ao verificar documento: %v", err)
	}

	return documentJSON != nil, nil
}
