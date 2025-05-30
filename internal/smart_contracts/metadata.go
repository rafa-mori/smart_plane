package smart_contracts

type BaseContractInfo struct {
	ContractID          string `json:"contractId"`
	ContractName        string `json:"contractName"`
	ContractDescription string `json:"contractDescription"`
	ContractVersion     string `json:"contractVersion"`
	ContractNamespace   string `json:"contractNamespace"`
	ContractType        string `json:"contractType"`
	ContractOwner       string `json:"contractOwner"`
	ContractStatus      string `json:"contractStatus"`
	ContractStartDate   string `json:"contractStartDate"`
	ContractEndDate     string `json:"contractEndDate"`
}

func NewBaseContractInfo(
	contractID, contractName, contractDescription, contractVersion,
	contractNamespace, contractType, contractOwner, contractStatus,
	contractStartDate, contractEndDate string,
) *BaseContractInfo {
	return &BaseContractInfo{
		ContractID:          contractID,
		ContractName:        contractName,
		ContractDescription: contractDescription,
		ContractVersion:     contractVersion,
		ContractNamespace:   contractNamespace,
		ContractType:        contractType,
		ContractOwner:       contractOwner,
		ContractStatus:      contractStatus,
		ContractStartDate:   contractStartDate,
		ContractEndDate:     contractEndDate,
	}
}
