package identities

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Estrutura para identidade de usuário na rede Fabric
type IdentityBase struct {
	ID       string `json:"id"`
	Role     string `json:"role"`
	Verified bool   `json:"verified"`
}

// Criar um novo usuário na blockchain
func (ic *IdentityBase) CreateUser(ctx contractapi.TransactionContextInterface, id string, role string) error {
	exists, err := ic.UserExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("usuário %s já existe", id)
	}

	user := IdentityBase{
		ID:       id,
		Role:     role,
		Verified: false,
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, userJSON)
}

// Consultar um usuário na blockchain
func (ic *IdentityBase) GetUser(ctx contractapi.TransactionContextInterface, id string) (*IdentityBase, error) {
	userJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("falha ao obter usuário: %v", err)
	}
	if userJSON == nil {
		return nil, fmt.Errorf("usuário %s não encontrado", id)
	}

	var user IdentityBase
	err = json.Unmarshal(userJSON, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Validar a identidade de um usuário
func (ic *IdentityBase) ValidateUser(ctx contractapi.TransactionContextInterface, id string) error {
	user, err := ic.GetUser(ctx, id)
	if err != nil {
		return err
	}

	user.Verified = true
	userJSON, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, userJSON)
}

// Verificar se o usuário já existe na blockchain
func (ic *IdentityBase) UserExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	userJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("falha ao verificar usuário: %v", err)
	}

	return userJSON != nil, nil
}
