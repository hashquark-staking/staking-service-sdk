package staking_service

type Response[T any] struct {
	Code uint   `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

type PageResult[T any] struct {
	PageNum  int   `json:"pageNum"`
	PageSize int   `json:"pageSize"`
	Total    int64 `json:"total"`
	List     []T   `json:"list"`
}

type User struct {
	ID       uint64 `json:"id"`
	BrokerID uint64 `json:"brokerID"`
	UID      string `json:"uid"`
}

type UserDetails struct {
	User
	DepositAddresses []string            `json:"depositAddresses"`
	Validators       ValidatorDetailList `json:"validators"`
}

type Validator struct {
	UID            string `json:"uid"`
	ValidatorIndex uint64 `json:"validatorIndex"`
	Pubkey         string `json:"pubkey"`
}

type ValidatorDetail struct {
	Validator
	Balance        string `json:"balance"`
	TotalRewards   string `json:"totalRewards"`
	Status         string `json:"status"`
	ActivationTime string `json:"ActivationTime"`
	ExitTime       string `json:"exitTime"`
	Principal      string `json:"principal"`
}

type ValidatorDetailList []ValidatorDetail
type ValidatorList []Validator

type CreateUserParams struct {
	UID              string   `json:"uid"`
	DepositAddresses []string `json:"depositAddresses"`
}

type AddDepositAddressRequest struct {
	DepositAddresses []string `json:"deposit_addresses"`
}

type Deposit struct {
	ID              uint64 `json:"id"`
	BrokerID        uint64 `json:"brokerId"`
	FromAddress     string `json:"fromAddress"`
	ContractAddress string `json:"contractAddress"`
	TxHash          string `json:"txHash"`
	ChainName       string `json:"chainName"`
	BlockHeight     uint64 `json:"blockHeight"`
	Slot            uint64 `json:"slot"`
	BlockTime       uint64 `json:"blockTime"`
}

type DepositDetails struct {
	Deposit
	Validators ValidatorDetailList `json:"validators"`
}

type UserValidators struct {
	UID        string        `json:"uid"`
	Validators ValidatorList `json:"validators"`
}

type ValidatorAssignmentParams struct {
	Assignments []Assignment `json:"assignments"`
}

type BatchValidatorAssignmentParams struct {
	Assignments []Assignment `json:"assignments"`
	TxHashs     []string     `json:"tx_hashs"`
}

type DepositDataRequestParams struct {
	Uid               string `json:"uid"`
	Quantity          uint64 `json:"quantity"`
	WithdrawalAddress string `json:"withdrawalAddress"`
}

type Assignment struct {
	UID            string `json:"uid"`
	ValidatorCount int    `json:"validatorCount"`
}

type Reward struct {
	Date              string `json:"date"`
	TotalCount        int64  `json:"totalCount"`
	ActiveCount       int64  `json:"activeCount"`
	ExitCount         int64  `json:"exitCount"`
	Balance           string `json:"balance"`
	Principal         string `json:"principal"`
	Apr               string `json:"apr"`
	CumulativeRewards string `json:"cumulativeRewards"`
	CumulativeStaking string `json:"cumulativeStaking"`
	CumulativeGasFee  string `json:"cumulativeGasFee"`
	Rewards           string `json:"rewards"`
	Staking           string `json:"staking"`
	GasFee            string `json:"gasFee"`
	Withdrawal        string `json:"withdrawal"`
	Rate              string `json:"rate"`
}

type RewardList []Reward

type DepositDataResponse struct {
	Network  string            `json:"network"`
	Protocol string            `json:"protocol"`
	Ethereum *EthereumResponse `json:"ethereum"`
}

type EthereumResponse struct {
	ContractAddress     string             `json:"contractAddress"`
	EstimatedGas        uint64             `json:"estimatedGas"`
	UnsignedTransaction string             `json:"unsignedTransaction"`
	DepositData         []*DepositDataInfo `json:"depositData"`
}

type DepositDataInfo struct {
	DepositDataRoot      string `json:"depositDataRoot"`
	Pubkey               string `json:"pubkey"`
	Signature            string `json:"signature"`
	WithdrawalCredential string `json:"withdrawalCredential"`
}
