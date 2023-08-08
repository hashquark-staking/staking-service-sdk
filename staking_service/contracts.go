package staking_service

func (stakingService *StakingService) GetUnusedKeys() (result *uint64, code uint, msg string, err error) {
	req := stakingService.getBaseRequest()
	res, err := req.Send("GET", "/openapi/contracts/key_registry/unused_keys")
	if err != nil {
		return
	}

	return processResponse[uint64](res)
}
