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
}

type ValidatorList []Validator

type CreateUserParams struct {
	UID     string `json:"uid"`
	Email   string `json:"email"`
	Address string `json:"address"`
}
