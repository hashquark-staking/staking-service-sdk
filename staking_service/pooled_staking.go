package staking_service

import (
	"fmt"
	"strconv"
)

func (stakingService *StakingService) GetPoolRewards(date *string, limit *uint) (result *PoolRewardInfoList, code uint, msg string, err error) {
	req := stakingService.getBaseRequest()
	if date != nil {
		req.SetQueryParam("date", *date)
	}
	if limit != nil {
		req.SetQueryParam("limit", strconv.FormatUint(uint64(*limit), 10))
	}
	res, err := req.Send("GET", "/openapi/pooled_staking/rewards")
	if err != nil {
		return
	}

	return processResponse[PoolRewardInfoList](res)
}

func (stakingService *StakingService) GetUserPooledStakingInfo(address string) (result *UserPooledStakingInfo, code uint, msg string, err error) {
	req := stakingService.getBaseRequest()
	res, err := req.Send("GET", fmt.Sprintf("/openapi/pooled_staking/users/%s/staking_info", address))
	if err != nil {
		return
	}

	return processResponse[UserPooledStakingInfo](res)
}

func (stakingService *StakingService) GetPooledStakingInfo() (result *PooledStakingInfoList, code uint, msg string, err error) {
	req := stakingService.getBaseRequest()
	res, err := req.Send("GET", "/openapi/pooled_staking/staking_info")
	if err != nil {
		return
	}

	return processResponse[PooledStakingInfoList](res)
}

func (stakingService *StakingService) GetWithdrawRequestInfo(requestID uint64) (result *PooledWithdrawRequestInfo, code uint, msg string, err error) {
	req := stakingService.getBaseRequest()
	res, err := req.Send("GET", fmt.Sprintf("/openapi/pooled_staking/withdraw_request/%d", requestID))
	if err != nil {
		return
	}

	return processResponse[PooledWithdrawRequestInfo](res)
}

func (stakingService *StakingService) GetUserWithdrawPossible(userAddress string) (result *WithdrawPossibleBlock, code uint, msg string, err error) {
	req := stakingService.getBaseRequest()
	res, err := req.Send("GET", fmt.Sprintf("/openapi/pooled_staking/users/%s/withdraw_possible", userAddress))
	if err != nil {
		return
	}

	return processResponse[WithdrawPossibleBlock](res)
}

func (stakingService *StakingService) GetTransactionHistory(userAddress string, txType, pageNum, pageSize int64) (result *PageResult[TransactionList], code uint, msg string, err error) {
	req := stakingService.getBaseRequest()
	req.SetQueryParam("type", strconv.FormatInt(txType, 10))
	req.SetQueryParam("pageNum", strconv.FormatInt(pageNum, 10))
	req.SetQueryParam("pageSize", strconv.FormatInt(pageSize, 10))

	res, err := req.Send("GET", fmt.Sprintf("/openapi/pooled_staking/users/%s/history", userAddress))
	if err != nil {
		return
	}

	return processResponse[PageResult[TransactionList]](res)
}
