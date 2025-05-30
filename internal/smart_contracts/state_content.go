package smart_contracts

import "reflect"

type ContractContent[T any] struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
	Data   *T     `json:"data"`
}

func (cnt *ContractContent[T]) GetType() reflect.Type {
	if cnt == nil {
		return nil
	}
	return reflect.TypeFor[T]()
}
