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
	Quantity          int64  `json:"quantity"`
	WithdrawalAddress string `json:"withdrawalAddress"`
}

type SendTransactionRequestParams struct {
	SignedTransaction string `json:"signedTransaction"`
}

type SendTransactionResponse struct {
	TxHash string `json:"txHash"`
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
	Period   string            `json:"period"`
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

type PoolRewardInfo struct {
	AvgYearApr       string `json:"avg_year_apr"`
	BeginSlot        int    `json:"begin_slot"`
	CumulativeReward string `json:"cumulative_reward"`
	DailyReward      string `json:"daily_reward"`
	Date             string `json:"date"`
	EndSlot          int    `json:"end_slot"`
}

type PoolRewardInfoList []PoolRewardInfo

type PooledStakingInfo struct {
	Amount           string `json:"amount"`
	CumulativeReward string `json:"cumulative_reward"`
	Point            string `json:"point"`
	UserAddress      string `json:"user_address"`
}

type PooledStakingInfoList []PooledStakingInfo

type UserPooledStakingInfo struct {
	DepositedEth string `json:"deposited_eth"`
	TotalEth     string `json:"total_eth"`
}

type PooledWithdrawRequestInfo struct {
	RequestID      int    `json:"request_id"`
	User           string `json:"user"`
	Point          string `json:"point"`
	Amount         string `json:"amount"`
	Principal      string `json:"principal"`
	WithdrawTxhash string `json:"withdraw_txhash"`
	Description    string `json:"description"`
}

type WithdrawPossibleBlock struct {
	CanWithdrawBlock string `json:"can_withdraw_block"`
}

type TransactionList []*TransactionInfo

type TransactionInfo struct {
	Amount         float32 `json:"amount"`
	Type           uint64  `json:"type"`
	TxTime         uint64  `json:"txTime"`
	Status         string  `json:"status"`
	TxHash         string  `json:"txHash"`
	WithdrawalHash string  `json:"withdrawalHash"`
}

type DelegatedValidator struct {
	ChainName    string `json:"chainName"`
	ValidatorID  string `json:"validatorID"`
	Commission   uint   `json:"commission"`
	Principal    string `json:"principal"`
	Status       string `json:"status"`
	StatusOrigin int    `json:"statusOrigin"`
	APR          int64  `json:"apr"`
	AppearTime   uint64 `json:"appearTime"`
	DelegatorNum int64  `json:"delegatorNum"`
}

type DelegatedValidatorList []DelegatedValidator

type Delegator struct {
	ChainName   string `json:"chainName"`
	ValidatorID string `json:"validatorID"`
	DelegatorID string `json:"delegatorID"`
	Principal   string `json:"principal"`
}

type DelegatorOverview struct {
	StakedAmount                  string `json:"stakedAmount"`
	UnlockAmount                  string `json:"unlockAmount"`
	ReadyToClaim                  string `json:"readyToClaim"`
	HistoricalRewardsForStaker    string `json:"historicalRewardsForStaker"`
	HistoricalClaimedRewards      string `json:"historicalClaimedRewards"`
	HistoricalRewardsForValidator string `json:"historicalRewardsForValidator"`
	HistoricalReferralCommission  string `json:"historicalReferralCommission"`
	SettledReferralCommission     string `json:"settledReferralCommission"`
	UnsettledReferralCommission   string `json:"unsettledReferralCommission"`
}

type DelegatedValidatorReward struct {
	ChainName   string `json:"chainName"`
	ValidatorID string `json:"validatorID"`
	// Identifier为收益周期，即如果是每个Epoch统计收益即为Epoch，如果是某个BlockHeight分发收益则记录BlockHeight，如果是按时间周期，则记录时间或者日期。
	Identifier string `json:"identifier"`
	// 确定的Block Height或者收益周期的block起始值，用逗号分隔
	Block      string `json:"block"`
	Amount     string `json:"amount"`
	RewardTime uint64 `json:"rewardTime"`
	Principal  string `json:"principal"`
	Commission uint   `json:"commission"` // Commission Rate, 统一为 百分比 * 100，例如: 12.34% 为 1234
}

type DelegatedValidatorRewardList []DelegatedValidatorReward

type DelegatorReward struct {
	DelegatedValidatorReward
	DelegatorID string `json:"delegatorID"`
}

type DelegatorRewardList []DelegatorReward

type DelegateTransaction struct {
	ChainName   string `json:"chainName"`
	ValidatorID string `json:"validatorID"`
	DelegatorID string `json:"delegatorID"`
	Amount      string `json:"amount"`
	// 交易中的operationType在数据库中存为数字，1为Delegate，2为Undelegate，3为Claim Reward，注意在任何场景中，给用户打钱的交易都记为Claim，即假如取款不分两部，则取款交易记为3
	OperationType       string `json:"operationType"`
	OperationTypeOrigin int    `json:"operationTypeOrigin"`
	Block               string `json:"block"` // 在aptos中block和txHash都是version
	TxHash              string `json:"txHash"`
	TxTime              int64  `json:"txTime"`
}

type DelegateTransactionList []DelegateTransaction
