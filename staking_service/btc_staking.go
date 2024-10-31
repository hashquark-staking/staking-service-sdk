package staking_service

func (stakingService *StakingService) GenerateStakingRequest(params GenerateStakingRequestParams) (result *GenerateStakingRequestResp, code uint, msg string, err error) {
	req := stakingService.getBaseRequest()
	req.SetBody(params)
	res, err := req.Send("POST", "/openapi/btc-staking/babylon/stake")
	if err != nil {
		return
	}

	return processResponse[GenerateStakingRequestResp](res)
}

func (stakingService *StakingService) GenerateBroadcastStakingRequest(params GenerateBroadcastStakingRequestParams) (result *BroadcastBabylonStakingTxResp, code uint, msg string, err error) {
	req := stakingService.getBaseRequest()
	req.SetBody(params)
	res, err := req.Send("POST", "/openapi/btc-staking/babylon/transaction/broadcast")
	if err != nil {
		return
	}

	return processResponse[BroadcastBabylonStakingTxResp](res)
}
