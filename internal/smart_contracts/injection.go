package smart_contracts

import (
	"crypto/rsa"
	"fmt"
	"reflect"

	l "github.com/faelmori/logz"
	"github.com/google/uuid"
	gl "github.com/rafa-mori/smart_plane/logger"
	t "github.com/rafa-mori/smart_plane/types"
)

type BaseContractInfoAPI[T any] struct {
	Logger            l.Logger `json:"-"`
	*BaseContractInfo `json:"baseContractInfo,omitempty"`

	PublicKey  *rsa.PublicKey  `json:"publicKey,omitempty"`
	PrivateKey *rsa.PrivateKey `json:"signature,omitempty"`

	AccountID    uuid.UUID `json:"accountId,omitempty"`    // ID of the account that owns the contract
	AccessToken  string    `json:"accessToken,omitempty"`  // Optional, used for API access control
	RefreshToken string    `json:"refreshToken,omitempty"` // Optional, used for API access control

	// Requests will handle validation and other contract-related internal methods.
	// Internally we ONLY accept a map of this type:
	//
	// map[method]*t.ValidationFunc[*T] `json:"requests"`
	//
	// Type definition:
	// type ValidationFunc[T any] struct {
	//     Priority int
	//     Func     func(value *T, args ...any) interfaces.IValidationResult
	//     Result   interfaces.IValidationResult
	// }
	Requests map[string]any `json:"requests,omitempty"`

	Error error `json:"error,omitempty"` // Used to store any error that occurs during processing
}

func NewBaseContractInfoAPI[T any](contractInfo *BaseContractInfo, accountID uuid.UUID, publicKey *rsa.PublicKey, privateKey *rsa.PrivateKey, logger l.Logger) (*BaseContractInfoAPI[T], error) {
	if logger == nil {
		logger = l.GetLogger("BaseContractInfoAPI")
	}
	if publicKey == nil {
		gl.Log("error", "Public key cannot be nil")
		return nil, fmt.Errorf("public key cannot be nil")
	}
	if privateKey == nil {
		gl.Log("error", "Private key cannot be nil")
		return nil, fmt.Errorf("private key cannot be nil")
	}
	if contractInfo == nil {
		gl.Log("error", "Contract info cannot be nil")
		return nil, fmt.Errorf("contract info cannot be nil")
	}
	if accountID == uuid.Nil {
		gl.Log("error", "Account ID cannot be nil")
		return nil, fmt.Errorf("account ID cannot be nil")
	}

	return &BaseContractInfoAPI[T]{
		Logger:           logger,
		BaseContractInfo: contractInfo,
		PublicKey:        publicKey,
		PrivateKey:       privateKey,
		AccountID:        accountID,
		Error:            nil,
	}, nil
}

func (cnt *BaseContractInfoAPI[T]) RegisterRequest(method string, request t.ValidationFunc[*T]) {
	if cnt.Requests == nil {
		cnt.Requests = make(map[string]any)
	}
	cnt.Requests[method] = request
}

func (cnt *BaseContractInfoAPI[T]) GetRequest(method string) (t.ValidationFunc[*T], bool) {
	if cnt.Requests == nil {
		return t.ValidationFunc[*T]{}, false
	}

	if req, ok := cnt.Requests[method]; ok {
		if validationFunc, ok := req.(t.ValidationFunc[*T]); ok {
			return validationFunc, true
		}
	}
	return t.ValidationFunc[*T]{}, false
}

func (cnt *BaseContractInfoAPI[T]) GetType() reflect.Type {
	if cnt == nil {
		return nil
	}
	return reflect.TypeFor[T]()
}

func (cnt *BaseContractInfoAPI[T]) GetError() error {
	if cnt == nil {
		return fmt.Errorf("BaseContractInfoAPI is nil")
	}
	if cnt.BaseContractInfo == nil {
		return fmt.Errorf("BaseContractInfo is nil")
	}
	if cnt.PublicKey == nil {
		return fmt.Errorf("PublicKey is nil")
	}
	if cnt.PrivateKey == nil {
		return fmt.Errorf("PrivateKey is nil")
	}
	return nil
}

func (cnt *BaseContractInfoAPI[T]) SetError(err error) {
	if cnt == nil {
		return
	}
	if cnt.Error != nil {
		gl.LogObjLogger(&cnt, "error", "Error already set, overwriting with new error")
	}
	cnt.Error = err
	gl.LogObjLogger(&cnt, "error", fmt.Sprintf("Set error in BaseContractInfoAPI: %s", err))
}
