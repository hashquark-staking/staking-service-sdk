package staking_service

import (
	"fmt"
)

func (stakingService *StakingService) ListDelegatedValidators(chainName string) (result *DelegatedValidatorList, code uint, msg string, err error) {
	req := stakingService.getBaseRequest()
	res, err := req.Send("GET", fmt.Sprintf("/openapi/delegation/%s/validators", chainName))
	if err != nil {
		return
	}

	return processResponse[DelegatedValidatorList](res)
}
func (stakingService *StakingService) GetDelegatedValidatorInfo(chainName, validatorID string) (result *DelegatedValidator, code uint, msg string, err error) {
	req := stakingService.getBaseRequest()
	res, err := req.Send("GET", fmt.Sprintf("/openapi/delegation/%s/validators/%s", chainName, validatorID))
	if err != nil {
		return
	}

	return processResponse[DelegatedValidator](res)
}

// ListDelegatesForValidator 方法用于获取某个验证者的所有委托记录
func (stakingService *StakingService) ListDelegatesForValidator(chainName, validatorID string, pageParams PageParams) (result *PageResult[DelegateTransaction], code uint, msg string, err error) {
	req := stakingService.getBaseRequest()
	req.SetQueryParam("pageNum", fmt.Sprintf("%d", pageParams.PageNum))
	req.SetQueryParam("pageSize", fmt.Sprintf("%d", pageParams.PageSize))
	res, err := req.Send("GET", fmt.Sprintf("/openapi/delegation/%s/validators/%s/delegates", chainName, validatorID))
	if err != nil {
		return
	}

	// 处理响应并返回结果
	return processResponse[PageResult[DelegateTransaction]](res)
}

func (stakingService *StakingService) ListRewardsForValidator(chainName, validatorID string, pageParams PageParams) (result *PageResult[DelegatedValidatorReward], code uint, msg string, err error) {
	req := stakingService.getBaseRequest()
	req.SetQueryParam("pageNum", fmt.Sprintf("%d", pageParams.PageNum))
	req.SetQueryParam("pageSize", fmt.Sprintf("%d", pageParams.PageSize))
	res, err := req.Send("GET", fmt.Sprintf("/openapi/delegation/%s/validators/%s/rewards", chainName, validatorID))
	if err != nil {
		return
	}

	return processResponse[PageResult[DelegatedValidatorReward]](res)
}

func (stakingService *StakingService) ListDelegators(chainName, validatorID string) (result *[]Delegator, code uint, msg string, err error) {
	req := stakingService.getBaseRequest()
	req.SetQueryParam("validatorID", validatorID)
	res, err := req.Send("GET", fmt.Sprintf("/openapi/delegation/%s/delegators", chainName))
	if err != nil {
		return
	}

	return processResponse[[]Delegator](res)
}

func (stakingService *StakingService) GetDelegatorInfo(chainName, validatorID, delegatorID string) (result *Delegator, code uint, msg string, err error) {
	req := stakingService.getBaseRequest()
	req.SetQueryParam("validatorID", validatorID)
	res, err := req.Send("GET", fmt.Sprintf("/openapi/delegation/%s/delegators/%s", chainName, delegatorID))
	if err != nil {
		return
	}

	return processResponse[Delegator](res)
}

func (stakingService *StakingService) GetDelegatorOverview(chainName, validatorID, delegatorID string) (result *DelegatorOverview, code uint, msg string, err error) {
	req := stakingService.getBaseRequest()
	req.SetQueryParam("validatorID", validatorID)
	res, err := req.Send("GET", fmt.Sprintf("/openapi/delegation/%s/delegators/%s/overview", chainName, delegatorID))
	if err != nil {
		return
	}

	return processResponse[DelegatorOverview](res)
}
func (stakingService *StakingService) ListDelegatesForDelegator(chainName, validatorID, delegatorID string, pageParams PageParams) (result *PageResult[DelegateTransaction], code uint, msg string, err error) {
	req := stakingService.getBaseRequest()
	req.SetQueryParam("validatorID", validatorID)
	req.SetQueryParam("pageNum", fmt.Sprintf("%d", pageParams.PageNum))
	req.SetQueryParam("pageSize", fmt.Sprintf("%d", pageParams.PageSize))
	res, err := req.Send("GET", fmt.Sprintf("/openapi/delegation/%s/delegators/%s/delegates", chainName, delegatorID))
	if err != nil {
		return
	}

	return processResponse[PageResult[DelegateTransaction]](res)
}
func (stakingService *StakingService) ListRewardsForDelegator(chainName, validatorID, delegatorID string, pageParams PageParams) (result *PageResult[DelegatorReward], code uint, msg string, err error) {
	req := stakingService.getBaseRequest()
	req.SetQueryParam("validatorID", validatorID)
	req.SetQueryParam("pageNum", fmt.Sprintf("%d", pageParams.PageNum))
	req.SetQueryParam("pageSize", fmt.Sprintf("%d", pageParams.PageSize))
	res, err := req.Send("GET", fmt.Sprintf("/openapi/delegation/%s/delegators/%s/rewards", chainName, delegatorID))
	if err != nil {
		return
	}

	return processResponse[PageResult[DelegatorReward]](res)
}
