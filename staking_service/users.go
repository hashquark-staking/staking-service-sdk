package staking_service

import (
	"fmt"
)

func (stakingService *StakingService) ListUsers(pageNum, pageSize uint) (result *PageResult[User], code uint, msg string, err error) {
	req := stakingService.getBaseRequest()
	req.SetQueryParam("pageNum", fmt.Sprintf("%d", pageNum))
	req.SetQueryParam("pageSize", fmt.Sprintf("%d", pageSize))
	res, err := req.Send("GET", "/openapi/users")
	if err != nil {
		return
	}

	return processResponse[PageResult[User]](res)
}

func (stakingService *StakingService) GetUserDetails(uid string) (result *UserDetails, code uint, msg string, err error) {
	req := stakingService.getBaseRequest()
	res, err := req.Send("GET", fmt.Sprintf("/openapi/users/%s", uid))
	if err != nil {
		return
	}

	return processResponse[UserDetails](res)
}

func (stakingService *StakingService) CreateUser(uid string, depositaddresses []string) (result *User, code uint, msg string, err error) {
	createUserParams := new(CreateUserParams)
	createUserParams.UID = uid
	createUserParams.DepositAddresses = depositaddresses

	// paramBytes, err := json.Marshal(createUserParams)
	// if err != nil {
	// 	return
	// }
	req := stakingService.getBaseRequest()
	req.SetBody(createUserParams)
	res, err := req.Send("POST", "/openapi/users")
	if err != nil {
		return
	}

	return processResponse[User](res)
}

func (stakingService *StakingService) AddDepositAddress(uid string, depositaddresses []string) (result *string, code uint, msg string, err error) {
	addDepositAddressRequest := new(AddDepositAddressRequest)
	addDepositAddressRequest.DepositAddresses = depositaddresses

	// paramBytes, err := json.Marshal(createUserParams)
	// if err != nil {
	// 	return
	// }
	req := stakingService.getBaseRequest()
	req.SetBody(addDepositAddressRequest)
	res, err := req.Send("POST", fmt.Sprintf("/openapi/users/%s/add_deposit_address", uid))
	if err != nil {
		return
	}

	return processResponse[string](res)
}

func (stakingService *StakingService) DeleteUser(uid string) (result *string, code uint, msg string, err error) {
	req := stakingService.getBaseRequest()
	res, err := req.Send("DELETE", fmt.Sprintf("/openapi/users/%s", uid))
	if err != nil {
		return
	}

	return processResponse[string](res)
}
