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
	ID             uint64  `json:"id"`
	BrokerID       uint64  `json:"brokerID"`
	UID            string  `json:"uid"`
	Email          *string `json:"email"`
	Address        *string `json:"address"`
	CommissionRate uint    `json:"commissionRate"`
}

type UserDetails struct {
	User
	Validators ValidatorList `json:"validators"`
}

type Validator struct {
	UID            string `json:"uid"`
	ValidatorIndex uint64 `json:"validatorIndex"`
	Pubkey         string `json:"pubkey"`
	Balance        string `json:"balance"`
	TotalRewards   string `json:"totalRewards"`
	Status         string `json:"status"`
	ActivationTime string `json:"ActivationTime"`
	ExitTime       string `json:"exitTime"`
	Principal      string `json:"principal"`
}

type ValidatorList []Validator

type CreateUserParams struct {
	UID     string `json:"uid"`
	Email   string `json:"email"`
	Address string `json:"address"`
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
	Validators ValidatorList `json:"validators"`
}

type UserValidators struct {
	BrokerUserID uint64        `json:"-"`
	UID          string        `json:"uid"`
	Validators   ValidatorList `json:"validators"`
}

type ValidatorAssignmentParams struct {
	Assignments []Assignment `json:"assignments"`
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
