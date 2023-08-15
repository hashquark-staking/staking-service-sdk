package staking_service

import (
	"fmt"
)

func (stakingService *StakingService) ValidatorList(validatorListRequestParam ValidatorListRequestParam) (result *PageResult[ValidatorDetail], code uint, msg string, err error) {
	req := stakingService.getBaseRequest()
	req.SetQueryParam("pageNum", fmt.Sprintf("%d", validatorListRequestParam.PageNum))
	req.SetQueryParam("pageSize", fmt.Sprintf("%d", validatorListRequestParam.PageSize))
	req.SetQueryParam("uid", fmt.Sprintf("%s", validatorListRequestParam.Uid))
	res, err := req.Send("GET", fmt.Sprintf("/openapi/validators"))
	if err != nil {
		return
	}
	return processResponse[PageResult[ValidatorDetail]](res)
}

func (stakingService *StakingService) ValidatorExit(pubkey string, validatorExitParam ValidatorExitParam) (result *string, code uint, msg string, err error) {
	req := stakingService.getBaseRequest()
	req.SetBody(validatorExitParam)
	res, err := req.Send("POST", fmt.Sprintf("/openapi/validators/%s/exit", pubkey))
	if err != nil {
		return
	}

	return processResponse[string](res)
}

type ValidatorListRequestParam struct {
	Uid string `form:"uid"`
	PageParams
}

type PageParams struct {
	PageNum  int `json:"pageNum" form:"pageNum"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

type ValidatorExitParam struct {
	Epoch     uint64 `json:"epoch"`
	Broadcast uint64 `json:"broadcast"`
}
