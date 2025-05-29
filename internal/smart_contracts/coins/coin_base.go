package financials

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type CoinBase struct {
	contractapi.Contract

	ID     string `json:"id"`
	Amount int    `json:"amount"`
	From   string `json:"from"`
	To     string `json:"to"`
}

func (tc *CoinBase) PutCoinBase(ctx contractapi.TransactionContextInterface, id string, amount int, from string, to string) error {
	exists, err := tc.CoinBaseExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("transação %s já existe", id)
	}

	tx := CoinBase{
		ID:     id,
		Amount: amount,
		From:   from,
		To:     to,
	}

	txJSON, err := json.Marshal(tx)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, txJSON)
}

func (tc *CoinBase) GetCoinBase(ctx contractapi.TransactionContextInterface, id string) (*CoinBase, error) {
	txJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("falha ao obter transação: %v", err)
	}
	if txJSON == nil {
		return nil, fmt.Errorf("transação %s não encontrada", id)
	}

	var tx CoinBase
	err = json.Unmarshal(txJSON, &tx)
	if err != nil {
		return nil, err
	}

	return &tx, nil
}

func (tc *CoinBase) DeleteCoinBase(ctx contractapi.TransactionContextInterface, id string) error {
	exists, err := tc.CoinBaseExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("transação %s não encontrada", id)
	}

	return ctx.GetStub().DelState(id)
}

func (tc *CoinBase) CoinBaseExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	txJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("falha ao verificar transação: %v", err)
	}

	return txJSON != nil, nil
}

func (tc *CoinBase) CoinBaseHistory(ctx contractapi.TransactionContextInterface, id string) ([]*CoinBase, error) {
	resultsIterator, err := ctx.GetStub().GetHistoryForKey(id)
	if err != nil {
		return nil, fmt.Errorf("falha ao obter histórico da transação: %v", err)
	}
	defer resultsIterator.Close()

	var history []*CoinBase
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("falha ao iterar no histórico: %v", err)
		}

		var tx CoinBase
		err = json.Unmarshal(queryResponse.Value, &tx)
		if err != nil {
			return nil, fmt.Errorf("falha ao deserializar transação: %v", err)
		}

		history = append(history, &tx)
	}

	return history, nil
}
