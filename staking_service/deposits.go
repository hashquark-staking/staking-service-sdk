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
	res, err := req.Send("POST", fmt.Sprintf("/openapi/deposits/validators_assign"))
	if err != nil {
		return
	}

	return processResponse[[]UserValidators](res)
}
