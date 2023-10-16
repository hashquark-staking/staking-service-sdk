package staking_service

import "fmt"

func (stakingService *StakingService) ListDeposits(pageNum, pageSize uint) (result *PageResult[Deposit], code uint, msg string, err error) {
	req := stakingService.getBaseRequest()
	req.SetQueryParam("pageNum", fmt.Sprintf("%d", pageNum))
	req.SetQueryParam("pageSize", fmt.Sprintf("%d", pageSize))
	res, err := req.Send("GET", "/openapi/deposits")
	if err != nil {
		return
	}

	return processResponse[PageResult[Deposit]](res)
}

func (stakingService *StakingService) GetDepositDetails(txHash string) (result *DepositDetails, code uint, msg string, err error) {
	req := stakingService.getBaseRequest()
	res, err := req.Send("GET", fmt.Sprintf("/openapi/deposits/%s", txHash))
	if err != nil {
		return
	}

	return processResponse[DepositDetails](res)
}

func (stakingService *StakingService) AssignDepositedValidators(txHash string, validatorAssignmentParams ValidatorAssignmentParams) (result *[]UserValidators, code uint, msg string, err error) {
	req := stakingService.getBaseRequest()
	req.SetBody(validatorAssignmentParams)
	res, err := req.Send("POST", fmt.Sprintf("/openapi/deposits/%s/assignment", txHash))
	if err != nil {
		return
	}

	return processResponse[[]UserValidators](res)
}

func (stakingService *StakingService) BatchAssignDepositedValidators(batchValidatorAssignmentParams BatchValidatorAssignmentParams) (result *[]UserValidators, code uint, msg string, err error) {
	req := stakingService.getBaseRequest()
	req.SetBody(batchValidatorAssignmentParams)
	res, err := req.Send("POST", "/openapi/deposits_assignments")
	if err != nil {
		return
	}

	return processResponse[[]UserValidators](res)
}

func (stakingService *StakingService) DepositData(depositDataRequestParams DepositDataRequestParams) (result *DepositDataResponse, code uint, msg string, err error) {
	req := stakingService.getBaseRequest()
	req.SetBody(depositDataRequestParams)
	res, err := req.Send("POST", "/openapi/deposits/deposit_data")
	if err != nil {
		return
	}

	return processResponse[DepositDataResponse](res)
}

func (stakingService *StakingService) ManualDepositBrokerUser(txHash string) (result *string, code uint, msg string, err error) {
	req := stakingService.getBaseRequest()
	res, err := req.Send("GET", fmt.Sprintf("/openapi/deposits/manual_deposit/broker_user/%s", txHash))
	if err != nil {
		return
	}

	return processResponse[string](res)
}

func (stakingService *StakingService) ManualDepositBroker(txHash string) (result *string, code uint, msg string, err error) {
	req := stakingService.getBaseRequest()
	res, err := req.Send("GET", fmt.Sprintf("/openapi/deposits/manual_deposit/broker/%s", txHash))
	if err != nil {
		return
	}

	return processResponse[string](res)
}

func (stakingService *StakingService) SendTransaction(sendTransactionRequestParams SendTransactionRequestParams) (result *SendTransactionResponse, code uint, msg string, err error) {
	req := stakingService.getBaseRequest()
	req.SetBody(sendTransactionRequestParams)
	res, err := req.Send("POST", "/openapi/deposits")
	if err != nil {
		return
	}

	return processResponse[SendTransactionResponse](res)
}
