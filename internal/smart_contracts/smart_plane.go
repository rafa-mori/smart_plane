package smart_contracts

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	ds "github.com/rafa-mori/smart_documents/data_structures"
	sd "github.com/rafa-mori/smart_documents/document_base"
)

type BlockchainManager struct {
	contracts map[string]contractapi.ContractInterface
}

func NewBlockchainManager() *BlockchainManager {
	return &BlockchainManager{
		contracts: map[string]contractapi.ContractInterface{
			"ApprovalContract":  &sd.ApprovalContract{},
			"SignatureContract": &sd.SignatureContract{},
			"TrafficContract":   &sd.TrafficContract{},
		},
	}
}

func (bm *BlockchainManager) RegisterDocument(contractName, id, content string) error {
	contract, exists := bm.contracts[contractName]
	if !exists {
		return fmt.Errorf("contrato %s não encontrado", contractName)
	}

	switch c := contract.(type) {
	case *sd.ApprovalContract:
		return c.RegisterDocument(nil, id, content)
	case *sd.TrafficContract:
		return c.RegisterTrafficDocument(nil, id, content)
	default:
		return fmt.Errorf("contrato %s não suporta registro de documentos", contractName)
	}
}

func (bm *BlockchainManager) GetDocumentHistory(contractName, id string) ([]string, error) {
	contract, exists := bm.contracts[contractName]
	if !exists {
		return nil, fmt.Errorf("contrato %s não encontrado", contractName)
	}

	switch c := contract.(type) {
	case *sd.ApprovalContract:
		return c.GetDocumentHistory(nil, id)
	case *sd.SignatureContract:
		return c.GetDocumentHistory(nil, id)
	case *sd.TrafficContract:
		return c.GetDocumentHistory(nil, id)
	default:
		return nil, fmt.Errorf("contrato %s não suporta consulta de histórico", contractName)
	}
}

func (bm *BlockchainManager) DeleteDocumentState(contractName, id string) error {
	contract, exists := bm.contracts[contractName]
	if !exists {
		return fmt.Errorf("contrato %s não encontrado", contractName)
	}

	switch c := contract.(type) {
	case *sd.ApprovalContract:
		return c.DeleteDocumentState(nil, id)
	case *sd.SignatureContract:
		return c.DeleteDocumentState(nil, id)
	case *sd.TrafficContract:
		return c.DeleteDocumentState(nil, id)
	default:
		return fmt.Errorf("contrato %s não suporta exclusão de estado", contractName)
	}
}

func (bm *BlockchainManager) ApproveDocument(contractName, id string) error {
	contract, exists := bm.contracts[contractName]
	if !exists {
		return fmt.Errorf("contrato %s não encontrado", contractName)
	}

	switch c := contract.(type) {
	case *sd.ApprovalContract:
		return c.ApproveDocument(nil, id)
	default:
		return fmt.Errorf("contrato %s não suporta aprovação de documentos", contractName)
	}
}

func (bm *BlockchainManager) SignDocument(contractName, id, signature string) error {
	contract, exists := bm.contracts[contractName]
	if !exists {
		return fmt.Errorf("contrato %s não encontrado", contractName)
	}

	switch c := contract.(type) {
	case *sd.SignatureContract:
		return c.SignDocument(nil, id, signature)
	default:
		return fmt.Errorf("contrato %s não suporta assinatura de documentos", contractName)
	}
}

func (bm *BlockchainManager) GetDocumentState(contractName, id string) (*ds.Document, error) {
	contract, exists := bm.contracts[contractName]
	if !exists {
		return nil, fmt.Errorf("contrato %s não encontrado", contractName)
	}

	switch c := contract.(type) {
	case *sd.ApprovalContract:
		return c.GetDocumentState(nil, id)
	case *sd.SignatureContract:
		return c.GetDocumentState(nil, id)
	case *sd.TrafficContract:
		return c.GetDocumentState(nil, id)
	default:
		return nil, fmt.Errorf("contrato %s não suporta consulta de estado", contractName)
	}
}
