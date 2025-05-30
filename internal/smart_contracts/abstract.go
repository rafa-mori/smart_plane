package smart_contracts

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type BaseContract[T any] struct {
	contractapi.Contract
}

func (bc *BaseContract[T]) Put(ctx contractapi.TransactionContextInterface, id string, data T) (bool, error) {
	stateJSON, err := readState(ctx, id)
	if err != nil {
		return false, fmt.Errorf("erro ao ler estado: %v", err)
	}
	if stateJSON != nil {
		// If the state already exists, return true and an error
		return true, fmt.Errorf("data already registered")
	}
	txJSON, err := json.Marshal(data)
	if err != nil {
		return false, fmt.Errorf("erro ao serializar dados: %v", err)
	}
	if err := ctx.GetStub().PutState(id, txJSON); err != nil {
		return false, fmt.Errorf("erro ao gravar estado: %v", err)
	} else {
		return true, nil
	}
}

func (bc *BaseContract[T]) Get(ctx contractapi.TransactionContextInterface, id string) (T, error) {
	if txJSON, err := ctx.GetStub().GetState(id); err != nil {
		var zero T
		return zero, fmt.Errorf("erro ao obter item %s: %v", id, err)
	} else if txJSON == nil {
		var zero T
		return zero, fmt.Errorf("item %s não encontrado", id)
	} else {
		var data T
		err = json.Unmarshal(txJSON, &data)
		if err != nil {
			return data, err
		}
		return data, nil
	}
}

func (bc *BaseContract[T]) Delete(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	if id == "" {
		return false, fmt.Errorf("ID cannot be empty")
	}
	if !bc.Exists(ctx, id) {
		// If the item does not exist, return true and an error
		// True because the data does not exist, so no deletion is needed
		return true, fmt.Errorf("item %s não encontrado", id)
	}
	if err := ctx.GetStub().DelState(id); err != nil {
		// If there was an error during deletion, return false and the error
		return false, fmt.Errorf("erro ao deletar item %s: %v", id, err)
	} else {
		// If the deletion was successful, return true and no error
		return true, nil
	}
}

func (bc *BaseContract[T]) Exists(ctx contractapi.TransactionContextInterface, id string) bool {
	if stateJSON, err := readState(ctx, id); err != nil {
		fmt.Printf("erro ao ler estado: %v", err)
		return false
	} else if stateJSON == nil {
		// If the state does not exist, return false
		return false
	} else {
		// If the state exists, return true
		return true
	}
}

func (bc *BaseContract[T]) History(ctx contractapi.TransactionContextInterface, id string) ([]T, error) {
	resultsIterator, err := ctx.GetStub().GetHistoryForKey(id)
	if err != nil {
		return nil, fmt.Errorf("erro ao obter histórico do item %s: %v", id, err)
	}
	defer func(resultsIterator shim.HistoryQueryIteratorInterface) {
		_ = resultsIterator.Close()
	}(resultsIterator)
	var history []T
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("erro ao iterar no histórico: %v", err)
		}
		var data T
		if err = json.Unmarshal(queryResponse.Value, &data); err != nil {
			return nil, fmt.Errorf("erro ao deserializar histórico: %v", err)
		}
		history = append(history, data)
	}
	return history, nil
}

func readState(ctx contractapi.TransactionContextInterface, id string) ([]byte, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %w", err)
	}

	if assetJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", id)
	}

	return assetJSON, nil
}
